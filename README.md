# wmdfmt

This is a simple markdown formatter that will format markdown files with [lute](https://github.com/88250/lute).

<details>
<summary>📖 Table of Contents</summary>

- [🚚 Installation](#-installation)
- [💡 Usage](#-usage)
  - [Help](#help)
  - [Example](#example)
- [📜 License](#-license)

</details>

## 🚚 Installation

```bash
go install github.com/wind-addons/wmdfmt@latest
```

## 💡 Usage

### Help

```text
NAME:
   wmdfmt - Format Markdown files

USAGE:
   wmdfmt [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --in-place, -i   Format file(s) in-place (default: false)
   --recursive, -r  Recursively format files in subdirectories (default: false)
   --ignore value   Ignore files matching this pattern
   --help, -h       show help
```

### Example

```bash
# Format a single file
wmdfmt README.md

# Format from pipe
echo "# Hello" | wmdfmt

# Format multiple files
wmdfmt README.md CHANGELOG.md

# Format all markdown files in a directory
wmdfmt .

# Format all markdown files in a directory and subdirectories
wmdfmt -r .

# Format all markdown files in a directory and subdirectories in-place
wmdfmt -r -i .

# Format all markdown files in a directory and subdirectories in-place and ignore CHANGELOG.md
wmdfmt -r -i --ignore CHANGELOG.md .
```

## 📜 License

MIT
