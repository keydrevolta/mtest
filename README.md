# mtest

> A code that cannot be tested is flawed – Anonymous

Mtest will cater to various testing needs of openebs projects. Mtest can be 
easily guided to try out a particular customer use case. On a simpler scheme 
of things Mtest will be used during CI runs.

## mtest objectives

An operator should be able to create or update a test case with ease. She
should be able to infer the results of above executed test case as well. 

Operator need not write logic / programs to do so. Operator might use some scripting 
or binary commands to achieve her objective.

## mtest design concepts

### Building Block

- Use a tool that can express the intent of the test case in a yaml file.
- Use profiles that can feed common/default data to a test case:
  - These profiles can in turn be expressed in yaml file(s)
  - These profiles should be able to express the entire infrastructure's defaults
- In case some amount of programming is required:
  - It should be delegated to existing libraries
  - Alternatively, new utility libraries can be built to solve this purpose

### CLI

- Use a tool that can compose these test cases & expose them as a CLI
  - This CLI can be packaged & distributed with proper release tags
  - Refer [ahoy](https://github.com/ahoy-cli) for details

### Over HTTP

- Use a tool that can expose the CLI over HTTP
  - Refer [webhook](https://github.com/adnanh/webhook) for details
