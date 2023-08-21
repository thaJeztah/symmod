package symmod

import "testing"

func TestHelloWorld(t *testing.T) {
	if HelloWorld != "hello world" {
		t.Errorf("expected 'hello world'")
	}
}
