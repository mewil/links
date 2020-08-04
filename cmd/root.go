/*
Copyright © 2020 Michael Wilson

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/badoux/goscraper"
	"github.com/spf13/cobra"
)

var (
	labels      []string
	customTitle string
	repo        = os.Getenv("LINKS_REPO")
)

var rootCmd = &cobra.Command{
	Use:   "links",
	Short: "A CLI tool for creating and organizing websites as GitHub issues",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		title, description, err := fetchNameAndDescription(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := createIssue(url, title, description); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&labels, "label", "l", []string{}, "labels to add to the GitHub issue")
	rootCmd.PersistentFlags().StringVarP(&customTitle, "title", "t", "", "custom title for the GitHub issue")
}

func fetchNameAndDescription(url string) (string, string, error) {
	s, err := goscraper.Scrape(os.Args[1], 5)
	if err != nil {
		return "", "", err
	}
	return s.Preview.Title, s.Preview.Description, nil
}

func createIssue(url, title, description string) error {
	binary, err := exec.LookPath("gh")
	if err != nil {
		return err
	}
	if customTitle != "" {
		title = customTitle
	}
	body := fmt.Sprintf("%s\n\nSource: %s\n", description, url)
	args := []string{"gh", "issue", "create", "--repo", repo, "--title", title, "--body", body}
	for _, label := range labels {
		args = append(args, "--label", label)
	}
	return syscall.Exec(binary, args, os.Environ())
}
