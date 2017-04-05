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
it by a name called [mtool](https://github.com/openebs/elves/tree/master/mtool). 
