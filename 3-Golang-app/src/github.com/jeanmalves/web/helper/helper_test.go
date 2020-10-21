package helper

import (
    "strings"
    "testing"
)

func TestGreeting(t *testing.T) {
    bold_string := Greeting("Code.education Rocks!")
    same_string := strings.Compare("<b>Code.education Rocks!</b>", bold_string)

    if same_string != 0 {
        t.Errorf("A string %q, e diferente de '<b>Code.education Rocks!</b>'", bold_string)
    }
}