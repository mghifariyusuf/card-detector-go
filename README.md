# Card Network Detector

This is a Go implementation of a program that detects the **card network** based on the card number's prefix and length.

## Problem Statement

Given a card number as a string, determine which network it belongs to by checking:

- The **prefix** (IIN – Issuer Identification Number)
- The **length** of the card number

### Examples

#### Example 1
```
Input: "341234567890123"
Output: "American Express"
```

#### Example 2
```
Input: "5034567890123456"
Output: "Maestro"
```

## How to Run

1. Make sure you have Go 1.18+ installed.
2. Clone this repository:
    ```
    git clone git@github.com:mghifariyusuf/card-detector-go.git
    cd card-detector-go
    ```
3. Run the program:
    ```
    go run main.go
    ```

## Running Tests

This project includes unit tests for various card number scenarios.

To run tests:
```
go test ./...
```

## File Structure

- `main.go` – Contains the main function with test cases and example output.
- `/detector/detector.go` – Implements the card detection logic using rules for prefix and length.
- `/detector/detector_test.go` – Unit tests for verifying detection logic across supported card networks.

## Detection Logic Overview

Each card network is defined by:

- One or more **prefix rules** (either exact match or ranges)
- One or more **length rules** (either fixed or ranged)

The detection checks if the card number matches any of the rules defined for each network.

Supported networks:

| Network          | Prefixes                             | Lengths    |
|------------------|--------------------------------------|------------|
| American Express | 34, 37                               | 15         |
| Diners Club      | 38, 39                               | 14         |
| Visa             | 4                                    | 13, 16, 19 |
| MasterCard       | 51–55                                | 16         |
| Discover         | 6011, 622126–622925, 644–649, 65     | 16, 19     |
| Maestro          | 50, 56–59                            | 12–19      |

## License

This project is licensed under the MIT License.