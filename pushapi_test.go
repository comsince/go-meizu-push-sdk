package goPushSdk

import (
	"testing"
	"fmt"
)

func TestRegister(t *testing.T){
	message := Register("100999","80355073480594a99470dcacccd8cf2c","862891030007404")
	fmt.Print("response message ",message.message)
	if message.code != 200 {
		t.Error("Status Code not 200")
	}
}
