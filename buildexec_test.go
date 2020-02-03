package buildexec

import (
	"os"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	src := `
        package main
        import(
            "log"
        )

        func main(){
            log.Println("START")
            defer log.Println("END")

            log.Println("++++++++++++++")
        }
    `

	res, err := Run(src)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", res)
}

func TestBuildSuccess(t *testing.T) {
	src := `
        package main
        import(
            "log"
        )

        func main(){
            log.Println("START")
            defer log.Println("END")

            log.Println("++++++++++++++")
        }
    `

	const outfile = "./out"
	res, err := Build(src, outfile)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := os.Remove(outfile)
		if err != nil {
			t.Fatal(err)
		}
	}()

	t.Logf("%+v\n", res)
}
