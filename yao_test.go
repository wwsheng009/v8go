package v8go_test

import (
	"fmt"
	"testing"

	v8 "rogchap.com/v8go"
)

func TestYaoVersionPrint(t *testing.T) {
	v := v8.Version()
	fmt.Println("Version", v)
}

func TestYaoInit(t *testing.T) {
	t.Parallel()
	v8.YaoInit(1024)
	defer v8.YaoDispose()

	iso := v8.YaoNewIsolate()
	defer iso.Dispose()

	// stat := iso.GetHeapStatistics()
	// if stat.HeapSizeLimit > 1100*1024*1024 {
	// 	t.Errorf("HeapSizeLimit error %d", stat.HeapSizeLimit)
	// }
}

func TestYaoCopyIsolate(t *testing.T) {
	t.Parallel()
	v8.YaoInit(1024)
	defer v8.YaoDispose()

	iso := v8.YaoNewIsolate()
	defer iso.Dispose()

	new, err := iso.Copy()
	if err != nil {
		t.Error(err)
	}
	defer new.Dispose()
}

func TestYaoIsolateAsGlobal(t *testing.T) {

	v8.YaoInit(1024)
	defer v8.YaoDispose()

	iso := v8.YaoNewIsolate()
	defer iso.Dispose()

	iso.AsGlobal()
}

func TestYaoNewIsolateFromGlobal(t *testing.T) {
	v8.YaoInit(1024)
	defer v8.YaoDispose()

	iso := v8.YaoNewIsolate()
	iso.AsGlobal()
	iso.Dispose()
	iso = nil

	new, err := v8.YaoNewIsolateFromGlobal()
	if err != nil {
		t.Error(err)
	}

	if new == nil {
		t.Errorf("new is nil")
	}
	defer new.Dispose()
}
