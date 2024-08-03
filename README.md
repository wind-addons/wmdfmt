# mdfmt

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
go install github.com/wind-addons/mdfmt@latest
```

## 💡 Usage

### Help

```text
NAME:
   mdfmt - Format Markdown files

USAGE:
   mdfmt [global options] command [command options]

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
mdfmt README.md

# Format from pipe
echo "# Hello" | mdfmt

# Format multiple files
mdfmt README.md CHANGELOG.md

# Format all markdown files in a directory
mdfmt .

# Format all markdown files in a directory and subdirectories
mdfmt -r .

# Format all markdown files in a directory and subdirectories in-place
mdfmt -r -i .

# Format all markdown files in a directory and subdirectories in-place and ignore CHANGELOG.md
mdfmt -r -i --ignore CHANGELOG.md .
```

## 📜 License

MIT
