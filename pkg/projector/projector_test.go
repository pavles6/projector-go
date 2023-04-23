package projector_test

import (
	"reflect"
	"testing"

	"github.com/pavles6/projector-go/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is great",
			},
			"/foo": {
				"foo": "baz",
				"bar": "baz",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "asdf",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)

	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}

	if v != value {
		t.Errorf("expected value \"%s\" got \"%s\"", value, v)
	}

}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is great")

}

func TestGetValues(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	expected := map[string]string{
		"foo": "bar3",
		"bar": "baz",
		"fem": "is great",
	}

	if ok := reflect.DeepEqual(proj.GetValues(), expected); !ok {
		t.Error("get values failed")
	}

}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")

	proj.SetValue("foo", "bar4")

	test(t, proj, "foo", "bar4")

	proj = getProjector("/", data)

	test(t, proj, "foo", "bar1")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")

	proj.RemoveValue("foo")

	test(t, proj, "foo", "baz")
}
