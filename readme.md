# Anagram Checker

Anagram Checker is a simple Go program that checks if two given words are anagrams of each other. It also validates that the inputs are words consisting of only letters.

## Table of Contents

- [Anagram Checker](#anagram-checker)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Testing](#testing)
  - [Contributing](#contributing)
  - [License](#license)

## Prerequisites

Ensure you have the following installed on your local development environment:

- [Go](https://golang.org/doc/install) (version 1.14 or later)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Baraq23/anagram-checker.git
cd anagram-checker
```

2. Build the project:

```bash
go build
```

## Usage

- To run the Anagram Checker program, use the following command:

```bash
go run . <word1> <word2>
```

- Replace <word1> and <word2> with the words you want to check. For example:

```bash
go run . listen silent
```

- This will output:

```bash
listen is an anagram to silent
```


## Testing

- To run the tests for the Anagram Checker program, use the following command:

```bash
go test -v
```

- This will run the tests and provide detailed output about the test results.


## Contributing

- Contributions are welcome! Please follow these steps:

    Fork the repository.
    Create a new branch (git checkout -b feature-branch).
    Commit your changes (git commit -am 'Add new feature').
    Push to the branch (git push origin feature-branch).
    Create a new Pull Request.


## License

This project is licensed under the MIT License. See the LICENSE file for details.