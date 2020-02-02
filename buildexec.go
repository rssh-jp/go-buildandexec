package buildexec

import(
    "errors"
    "log"
    "os"
    "exec"
)

var(
    ErrNotMatchFileSize = errors.New("Not match file size")
)

func saveFile(output, data string)error{
    fd, err := os.OpenFile(tmpfile, os.O_CREATE, 0755)
    if err != nil{
        return err
    }

    n, err := fd.WriteString(src)
    if err != nil{
        return err
    }

    if n != len(src){
        return ErrNotMatchFileSize
    }

    return nil
}

func buildGoFile(path string)error{
    cmd := exec.Command("go", "run", path)

    err := cmd.Run()
    if err != nil{
        return err
    }

    return nil
}

func Run(src string)error{
    const tmpfile = "./tmp.go"
    err := saveFile(tmpfile, src)
    if err != nil{
        return err
    }

    err = buildGoFile(tmpfile)
    if err != nil{
        return err
    }

    log.Println("+++++++++")

    
}
