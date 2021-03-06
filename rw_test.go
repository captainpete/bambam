package main

import (
	"os"
	"os/exec"
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test003WriteReadThroughGeneratedTranslationCode(t *testing.T) {

	cv.Convey("Given bambam generated go bindings", t, func() {
		cv.Convey("then we should be able to write to disk, and read back the same structure", func() {

			tdir := NewTempDir()
			// comment the defer out to debug any rw test failures.
			defer tdir.Cleanup()

			MainArgs([]string{os.Args[0], "-o", tdir.DirPath, "rw.go.txt"})

			err := exec.Command("cp", "rw.go.txt", tdir.DirPath+"/rw.go").Run()
			cv.So(err, cv.ShouldEqual, nil)

			tdir.MoveTo()

			err = exec.Command("capnpc", "-ogo", "schema.capnp").Run()
			cv.So(err, cv.ShouldEqual, nil)

			err = exec.Command("go", "build").Run()
			cv.So(err, cv.ShouldEqual, nil)

			// run it
			err = exec.Command("./" + tdir.DirPath).Run()
			cv.So(err, cv.ShouldEqual, nil)

		})
	})
}
