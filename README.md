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

## Simple yet effective mtest

> Simplicity is the ultimate sophistication - Leonardo da Vinci

To achieve this, we shall try to adhere to some design concepts. We are calling
it by a name called `mtool`. `mtool` currently consists of design thoughts only.
In future, we might put `mtool` concepts in `openebs` repository. The trials
of `mtool` concepts will be done in `elves` repository.

> NOTE: The current code base available in this repo might be removed altogether
to accomodate `mtool`'s concepts.

## mtool design concepts

### Intent

Express the intent of the test case in a yaml file. [details](design/intent.md)

### CLI

- Use a tool that can compose these test cases & expose them as a CLI
  - This CLI can be packaged & released
  - e.g. `mtest` can be a packaged distribution which derives all its design principles from `mtool`.

### Over HTTP

- Expose the CLI over HTTP

### Over SLACK

- SLACK plugin that accepts above CLI commands
