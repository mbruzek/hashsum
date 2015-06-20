/* Copyright 2015 Matthew Bruzek */
package main

import (
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
  "fmt"
  "io/ioutil"
  "os"
)

func readFile(filename string) []byte {
  var rawBytes []byte
  var err error
  if filename != "" {
    rawBytes, err = ioutil.ReadFile(filename)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }
  return rawBytes
}

func Hashsum(algorithm string, filename string) string {
  var hashString string
  rawBytes := readFile(filename)
  if algorithm == "md5" {
    md5sum := md5.Sum(rawBytes)
    hashString = fmt.Sprintf("%x", md5sum)
  } else if algorithm == "sha1" {
    sha1sum := sha1.Sum(rawBytes)
    hashString = fmt.Sprintf("%x", sha1sum)
  } else if algorithm == "sha256" {
    sha256sum := sha256.Sum256(rawBytes)
    hashString = fmt.Sprintf("%x", sha256sum)
  } else {
    fmt.Printf("%s algorithm not implemented.\n", algorithm)
    os.Exit(1)
  }
  return hashString
}

func main() {
  if len(os.Args) > 2 {
    algorithm := os.Args[1]
    filename := os.Args[2]
    hashsum := Hashsum(algorithm, filename)
    fmt.Printf("%s  %s\n", hashsum, filename)
  }
}
