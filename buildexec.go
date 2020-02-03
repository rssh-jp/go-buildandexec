package buildexec

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	ErrNotMatchFileSize = errors.New("Not match file size")
)

func saveFile(output, data string) error {
	fd, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	n, err := fd.WriteString(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return ErrNotMatchFileSize
	}

	return nil
}

type Response struct {
	OutputString string
	ErrorString  string
}

func runGoFile(path string) (*Response, error) {
	cmd := exec.Command("go", "run", path, "1>&2")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	outbuf, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	errbuf, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, err
	}

	res := new(Response)
	res.OutputString = string(outbuf)
	res.ErrorString = string(errbuf)

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildGoFile(path, destpath string) error {
	cmd := exec.Command("go", "build", "-o", destpath, path)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func removeFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func Run(src string) (*Response, error) {
	const tmpfile = "./tmp.go"
	err := saveFile(tmpfile, src)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = removeFile(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
	}()

	res, err := runGoFile(tmpfile)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Build(src, destpath string) (*Response, error) {
	const tmpfile = "./tmp.go"
	err := saveFile(tmpfile, src)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = removeFile(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = buildGoFile(tmpfile, destpath)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
