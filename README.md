# Dynamic Plugin Loader Example

## Overview

This Go program serves as a dynamic plugin loader, loading a shared object file (`plug.so`) at runtime. The program assumes that the plugin defines specific symbols (`V` and `F`) and demonstrates how to interact with them.

## Features

- **Dynamic Loading:** The program dynamically loads a shared object file (`plug.so`) using the Go `plugin` package.
- **Symbol Lookup:** It looks up symbols (`V` and `F`) in the loaded plugin.
- **Command-Line Interaction:** The program takes a command-line argument and sets the value of a variable (`V`) in the plugin.
- **Function Invocation:** It calls a function (`F`) from the plugin with the hostname as an argument.

## Prerequisites

Ensure that you have Go installed on your machine. You can download it from [the official website](https://golang.org/dl/).

## Usage

1. **Build the Program:**

    ```bash
    make build
    ```

2. **Build the Plugin (`plug.so`):**

   The plugin should export the symbols `V` (pointer to a string) and `F` (function taking a string argument). Make sure the plugin is built before running the main program.
    
    ```bash
    go build -buildmode=plugin -o build/plug.so  plugin/*
    ```

3. **Run the Program:**

    ```bash
    ./app <your_argument>
    ```

   Replace `<your_argument>` with the desired command-line argument.

## Configuration

No specific configuration is required for this program. Ensure that the shared object file (`plug.so`) is present in the same directory as the executable.

## Contributing

If you'd like to contribute to the project, feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
