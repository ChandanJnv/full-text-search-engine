```markdown
# Simple Full-Text Search Engine

This repository contains a simple full-text search engine implemented in Go.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ChandanJnv/full-text-search-engine.git
   cd full-text-search-engine
   ```

2. Download the data dump from [here](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz) and extract it. Place the extracted file inside the repository directory.

3. Install the required Go modules:
   ```sh
   go mod tidy
   ```

### Usage

Run the search engine with the following command:
   ```sh
   go run main.go -p <file_path> -q <search_query>
   ```
Note: Command line arguments `file_path` or `search_query` are optional.

### Project Structure

- `main.go`: The main file that contains the indexing and searching logic.
- `utils/`: Directory containing utility functions.
- `go.mod` and `go.sum`: Go module files.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes or improvements.

## License

This project is licensed under the MIT License.
```

Feel free to modify any sections as needed to better suit your project.