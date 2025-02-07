#!/bin/bash

# Download protofmt if it does not exist.
if ! command -v protofmt &> /dev/null
then
    echo "protofmt not found. Installing..."
    go install github.com/hungcq/protofmt@latest
fi

# Get a list of changed/added YAML files that are staged for commit
changed_files=$(git diff --cached --name-only | grep '\.proto')

if [[ -z "$changed_files" ]]; then
    echo "No proto files to process."
    exit 0
fi

# Find and format all .proto files (modifying them in place).
for file in $changed_files; do
  echo "Formatting $file in place..."
  protofmt -o "$file"
done

echo "All .proto files have been formatted."
exit 0