package bufresolv

import (
	"slices"
	"testing"
)

func TestBufLock(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		r, err := New(BufLock("testdata/bufroot/helloapis/buf.lock"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto found in %v", paths)
		}
		if slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto found in %v", paths)
		}
	})

	t.Run("v2", func(t *testing.T) {
		r, err := New(BufLock("testdata/bufroot_v2/buf.lock"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto found in %v", paths)
		}
		if slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto found in %v", paths)
		}
	})
}

func TestBufConfig(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		r, err := New(BufConfig("testdata/bufroot/helloapis/buf.yaml"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto found in %v", paths)
		}
		if slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto found in %v", paths)
		}
	})

	t.Run("v2", func(t *testing.T) {
		r, err := New(BufConfig("testdata/bufroot_v2/buf.yaml"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto not found in %v", paths)
		}
	})

	t.Run("v2 no modules", func(t *testing.T) {
		r, err := New(BufConfig("testdata/bufroot_v2_no_modules/buf.yaml"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "helloapis/acme/hello/v2/hello.proto") {
			t.Errorf("helloapis/acme/hello/v2/hello.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "worldapis/acme/world/v1/world.proto") {
			t.Errorf("worldapis/acme/world/v1/world.proto not found in %v", paths)
		}
	})
}

func TestModule(t *testing.T) {
	r, err := New(BufModule("buf.build/bufbuild/protovalidate/tree/b983156c5e994cc9892e0ce3e64e17e0"))
	if err != nil {
		t.Fatal(err)
	}
	paths := r.Paths()
	if !slices.Contains(paths, "buf/validate/validate.proto") {
		t.Errorf("buf/validate/validate.proto not found in %v", paths)
	}
}

func TestBufDir(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		r, err := New(BufDir("testdata/bufroot"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto not found in %v", paths)
		}
	})

	t.Run("v2", func(t *testing.T) {
		r, err := New(BufDir("testdata/bufroot_v2"))
		if err != nil {
			t.Fatal(err)
		}
		paths := r.Paths()
		if !slices.Contains(paths, "buf/validate/validate.proto") {
			t.Errorf("buf/validate/validate.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/hello/v2/hello.proto") {
			t.Errorf("acme/hello/v2/hello.proto not found in %v", paths)
		}
		if !slices.Contains(paths, "acme/world/v1/world.proto") {
			t.Errorf("acme/world/v1/world.proto not found in %v", paths)
		}
	})
}

func TestPaths(t *testing.T) {
	r, err := New(BufConfig("testdata/bufroot/helloapis/buf.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	got := len(r.Paths())
	r2, err := New(BufConfig("testdata/bufroot/helloapis/buf.yaml"), BufConfig("testdata/bufroot/helloapis/buf.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	got2 := len(r2.Paths())
	if got != got2 {
		t.Errorf("got %d, want %d", got2, got)
	}
}
