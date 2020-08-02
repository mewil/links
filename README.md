# Links

A CLI tool for creating and organizing websites as GitHub issues. Links will scrape the title and description from the provided URL, and create a new GitHub issue with the information in the repo specified by the environment variable `LINKS_REPO`. You can specify labels to add to the issue with the `--label` or `-l`.

## Usage

1. Install the [GitHub CLI](https://github.com/cli/cli) dependency (MacOS: `brew install github/gh/gh`).
2. Make sure to create any GitHub issue labels you want to use (`gh` will not create labels that don't exist)
3. Install links: `go install github.com/mewil/links`
4. Add a new link: `LINKS_REPO=mewil/links links https://mewil.io --label my-label`
5. [Voila!](https://github.com/mewil/links/issues/4)
