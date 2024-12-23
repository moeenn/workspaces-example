# Go workspace example
Imaging a very large golang project. Implementing a particular new feature requires changing code in multiple repositories. Coordinating the changes can be a big hassle. Go `workspaces` simplifies this workflow.


## Structure
- The shared code goes in `lib/`.
- All services go inside `service/`.
- All sub-packages have their own git repositories and `go.mod` files.
- The root of the whole project will contain one `go.work` file.


## Using `go.work` file
- Reference to all sub-packages can be added to `go.work` file using the following command.

```bash
$ go work use ./path/to/sub-package
```

- In Local development, go toolchain will automatically pick up the `go.work` file from the parent folders, and allow building the binary by automatically linking the sub-packages. The sub-packages will not be downloaded (from Github or where-ever).
- In CI/CD pipelines, each repo changes are pushed separately. This means in that environment, actual released versions of sub-packages will be downloaded from Github (in this example).
- Sub-package executables can be executed from the root of the project using the following command.

```bash
# quick execution
$ go run ./path/to/executable/sub-package

# build for production
$ go build -o ./build/binary-name ./path/to/executable/sub-package
```
