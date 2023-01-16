package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string   //cp:classpath
	class       string   //.class文件
	args        []string //参数数组
}

func parseCmd() *Cmd {
	cmd := &Cmd{} //cmd存放Cmd结构体的地址 等同于var cmd *Cmd = &cmd{}
	flag.Usage = printUsage
	//BoolVar(p *bool, name string, value bool, usage string)
	//BoolVar defines a bool flag with specified name, default value, and usage string.
	//The argument p points to a bool variable in which to store the value of the flag.
	//&cmd.helpFlag是存放bool变量的地址
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	//os.Args[0]是要运行的文件名,如ch01.exe
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
