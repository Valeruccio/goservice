//тест фаил должен называться с окончанием blabla_test
//go test
//go mod init

package hello

import "testing"

func TestHello(t *testing.T)  {
	wnt := "Hello"

	if got := Hello(); got != wnt {
		t.Errorf("Хуйня твой Hello() = %q, want %q", got, wnt)
	}

}