package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected. Expected: %q, Got %q\n", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected. Expected: %q, Got %q\n", expect, got)
	}
}

func TestGreeTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello Gopher!\n"},
		{input: " ", expect: "Hello  !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected. Expected: %q, Got %q\n", s.expect, got)
		}
	}
}

// func TestFailureType(t *testing.T) {
// 	t.Error("Error 1 signal")
// 	t.Fatal("Fatal wil fail")
// 	t.Error("Error 2 signal") //Never printed
// }
