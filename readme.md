[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-f4981d0f882b2a3f0472912d15f9806d57e124e0fc890972558857b51b24a6f9.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=10739941)
In assignment 3, we will simulate a spring motion. We will make a program that creates a gif file that simulates the oscillation of a spring with a mass, based on the inputs.

```
> go run main.go

Drawing the graph of the spring 
oscillation using the following IVP.
mx'' + bx' + kx = 0, x(0) = x0, x'(0) = x1, 0<t<T

Enter the spring constant(k): 1.2
Enter the damping contact(b): .8
Enter the mass of the object(m): 2.1
Enter the initial displacement(x0): .5
Enter the initial velocity(x1): -.6
Enter the duration(T): 12
```


In the `main.go` file, I already put code that simulates the spring motion.
It uses a custom package called `spring`. 
You can find the package in the `/spring` folder.
The package contains two go files `spring.go` and `func.go`.
The `spring.go` file contains the `Spring` struct.
The `func.go` file contains the `Func` struct.
There are two functions whose contents are missing.
* `FindSolution` in `/spring/spring.go`
* `Graph` in `/spring/func.go`

The main task is to fill in the missing parts of these two functions
so that the program can animate the spring motion properly.
The following gif shows the simplest way to animate the spring motion.
You may add additional features to the animation such as 
* the x and y axes
* labels and ticks on the axes
* animated figure of the mass and the spring

Any additional features are welcome.
Try to be creative and feel free to ask any questions.
Most importantly, work together and have fun!

![](./spring.gif)
