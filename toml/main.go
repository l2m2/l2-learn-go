package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type Config1 struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

type Config2 struct {
	Age  int
	Cats []string
	Pi   float64
	Perf []int `toml:Perfection`
	DOB  time.Time
}

func decode1() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	confPath := filepath.Join(exPath, "conf1.toml")
	fmt.Println(confPath)
	var conf Config2
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", conf)

	conf.Age = 34
	fmt.Printf("%#v", conf)
	f, _ := os.Create(confPath)
	defer f.Close()
	if err := toml.NewEncoder(f).Encode(conf); err != nil {
		panic(err)
	}
}

func main() {
	decode1()
}
