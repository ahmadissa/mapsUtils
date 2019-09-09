package Utils

import "testing"

func TestGetBool(t *testing.T) {
	data := make(map[string]interface{})
	data["found1"] = true
	data["found2"] = false
	data["not_bool"] = "false"
	found1, ok1 := GetBool(data, "found1")
	if !found1 {
		t.Errorf("incorrect found1, got: %v, want: %v.", found1, true)
	}
	if !ok1 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok1, true)
	}
	found2, ok2 := GetBool(data, "found2")
	if found2 {
		t.Errorf("incorrect found2, got: %v, want: %v.", found2, false)
	}
	if !ok2 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok2, true)
	}
	notFound, ok3 := GetBool(data, "not_found")
	if notFound {
		t.Errorf("incorrect found2, got: %v, want: %v.", notFound, false)
	}
	if ok3 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok3, false)
	}
	notBool, ok4 := GetBool(data, "not_bool")
	if notBool {
		t.Errorf("incorrect found2, got: %v, want: %v.", notBool, false)
	}
	if ok4 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok4, false)
	}
}

func TestGetInt(t *testing.T) {
	data := make(map[string]interface{})
	data["found1"] = 10
	data["not_int"] = "sd"
	found1, ok1 := GetInt(data, "found1")
	if found1 != 10 {
		t.Errorf("incorrect found1, got: %v, want: %v.", found1, 10)
	}
	if !ok1 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok1, true)
	}
	notInt, ok2 := GetInt(data, "not_int")
	if notInt != 0 {
		t.Errorf("incorrect notString, got: %v, want: %v.", notInt, 0)
	}
	if ok2 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok2, false)
	}
	notFound, ok3 := GetInt(data, "not_found")
	if notFound != 0 {
		t.Errorf("incorrect found2, got: %v, want: %v.", notFound, "")
	}
	if ok3 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok3, false)
	}
}

func TestGetString(t *testing.T) {
	data := make(map[string]interface{})
	data["found1"] = "found1"
	data["not_String"] = true
	found1, ok1 := GetString(data, "found1")
	if found1 != "found1" {
		t.Errorf("incorrect found1, got: %v, want: %v.", found1, "found1")
	}
	if !ok1 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok1, true)
	}
	notString, ok2 := GetString(data, "not_String")
	if notString != "" {
		t.Errorf("incorrect notString, got: %v, want: %v.", notString, "")
	}
	if ok2 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok2, false)
	}
	notFound, ok3 := GetString(data, "not_found")
	if notFound != "" {
		t.Errorf("incorrect found2, got: %v, want: %v.", notFound, "")
	}
	if ok3 {
		t.Errorf("incorrect found1, got: %v, want: %v.", ok3, false)
	}
}
