
# Journal CLI

A simple command-line application written in Go for managing a journal. The application currently supports adding new entries and deleting all entries. Future support for deleting a single entry is planned.

## Features

- Add a new journal entry.
- Delete all journal entries.
- (Planned) Delete a specific journal entry.

## Installation

1. Ensure you have [Go installed](https://golang.org/doc/install) on your machine.
2. Clone the repository:
    ```sh
    git clone https://github.com/26thavenue/journal-cli.git
    ```
3. Navigate to the project directory:
    ```sh
    cd journal-cli
    ```
4. Build the application:
    ```sh
    go build -o journal
    ```

## Usage

### Adding an Entry

To add a new entry to your journal, use the `-add` flag followed by the task description.

```sh
go run main.go -add "Your task description here"
```

### Deleting All Entries

To delete all entries from your journal, use the `-delAll` flag.

```sh
go run main.go -delAll
```

### (Planned) Deleting a Specific Entry

Support for deleting a specific entry by index will be added in the future. The planned usage is:

```sh
go run main.go -delOne 1
```

## Example

```sh
# Add a new entry
./journal -add "Finish writing the README"

# Delete all entries
./journal -delAll
```



## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

