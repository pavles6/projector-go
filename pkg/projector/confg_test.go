package projector_test

import (
	"reflect"
	"testing"

	"github.com/pavles6/projector-go/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("Expected to get no error, but got %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("Expected args to be %v, but got %v", args, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("Expected operation was %v, but got %v", operation, config.Operation)
	}

}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}
