Joe Matta
Distributed Systems
Task 2
2/29/2020

This program multiplies two square matrices that have the same dimensions.
Each matrix is sent to a single server and multiplied by a process on the server.
The product is sent back to client and displayed.

Instructions:

- The shared folder must be placed in the following directory on the server and client machines:

    C:\Go\src

    This folder contains:

    interface.go
    shared_structs.go

1) run server.go:
    $ go run server.go
2) run client.go:
    $ go run client.go
3) Enter the size of the matrices you would like to multiply.
    Both matrices will have the same dimensions.
4) Manually enter integer values row by row.
    Both matrices will be displayed.
5) The product matrix is displayed

----- Notes -----
Developed in GoLand 2019.3.2
             Build #GO-193.6015.58, built on February 3, 2020
             Runtime version: 11.0.4+10-b520.11 amd64
             VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o
             Windows 10 10.0

             using:
             go version go1.13.7 windows/amd64