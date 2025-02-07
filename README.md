# protofmt

**protofmt** is a simple tool for formatting Protocol Buffers (`.proto`) files by aligning field types, field names, and the equal signs. This ensures consistent and readable formatting in your protobuf definitions.

## Features

- **Column Alignment:** Automatically aligns field types, names, and the equals sign in protobuf messages.
- **In-Place Modification:** Optionally overwrite the original file using the `-o` (or `--override`) flag.
- **Simple Integration:** Designed to be used as a standalone binary.

## Installation

You can install the latest version of **protofmt** as a binary using Go:

```bash
go install github.com/hungcq/protofmt@latest
```

Make sure your $GOPATH/bin (or Go's installation bin directory) is in your system's PATH.

## Usage

### Basic Formatting
To format a .proto file and print the formatted output to stdout:

```bash
protofmt path/to/file.proto
```

### In-Place Modification
To format a .proto file and overwrite the original file with the formatted version, use the -o or --override flag:

```bash
protofmt -o path/to/file.proto
```

### Pre-commit Hook
To use the tool as a [pre-commit](https://pre-commit.com/) hook, add this to your .pre-commit-config.yaml:
```yaml
default_install_hook_types: [pre-commit, prepare-commit-msg]
repos:
  - repo: https://github.com/hungcq/protofmt.git
    rev: latest
    hooks:
      - id: protofmt
```

## Example
Given a my_msg.proto file with fields like:

```protobuf
syntax = "proto3";

message ParentMsg {
  int64 a = 1;
  string long_field_name = 2;
  bool b = 3;

  oneof event {
    TestMsg test_msg = 10;
    SubMsgWithVeryLongName sub_msg_with_very_long_name = 11;
  }
}

message TestMsg {
  int64 c = 1;
  string long_field_name = 2;
  bool d = 3;
  float test = 4;
}

message SubMsgWithVeryLongName {
  int64 x = 1;
  bool another_field_name = 2;
}
```

Running:
```bash
protofmt -o my_msg.proto
```
will reformat the file so that all field types, names, and the equal signs are aligned properly:

```protobuf
syntax = "proto3";

message ParentMsg {
  int64  a               = 1;
  string long_field_name = 2;
  bool   b               = 3;

  oneof event {
    TestMsg                test_msg                    = 10;
    SubMsgWithVeryLongName sub_msg_with_very_long_name = 11;
  }
}

message TestMsg {
  int64  c               = 1;
  string long_field_name = 2;
  bool   d               = 3;
  float  test            = 4;
}

message SubMsgWithVeryLongName {
  int64 x                  = 1;
  bool  another_field_name = 2;
}
```

## Contributing
Contributions are welcome!
Feel free to open issues or submit pull requests on [GitHub](https://github.com/hungcq/protofmt).

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
