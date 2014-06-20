package gosseract_test

import "github.com/otiai10/gosseract"
import . "github.com/otiai10/mint"
import "testing"

func Test_Greet(t *testing.T) {
	Expect(t, gosseract.Greet()).ToBe("Hello,Gosseract.")
}

func Test_Must(t *testing.T) {
    Expect(t, gosseract.Must(gosseract.Params{})).ToBe("gosseract")
}

func Test_NewClient(t *testing.T) {
	client, e := gosseract.NewClient()
	Expect(t, e).ToBe(nil)
	Expect(t, client).TypeOf("*gosseract.Client")
}
