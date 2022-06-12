package digest

import (
	"crypto/sha512"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/abergmeier/buildkit_ex/pkg/digest/options"
)

func init() {
	os.Chdir("testdata")
}

func TestDigestOfFileAndAllInputsWithDefaultOptions(t *testing.T) {
	defer os.Remove("foo.txt")
	err := os.WriteFile("foo.txt", []byte("foo"), 0666)
	if err != nil {
		panic(err)
	}
	firstDigest, err := DigestOfFileAndAllInputs("Containerfile")
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest == [sha512.Size]byte{} {
		t.Fatal("Digest not calculated")
	}
	err = os.WriteFile("foo.txt", []byte("fooandbar"), 0666)
	if err != nil {
		panic(err)
	}
	secondDigest, err := DigestOfFileAndAllInputs("Containerfile")
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest == secondDigest {
		t.Fatal("Second digest not properly calculated")
	}
}

func TestDigestOfFileAndAllInputsDir(t *testing.T) {
	digest, err := DigestOfFileAndAllInputs("Containerfile.dir")
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if digest == [sha512.Size]byte{} {
		t.Fatal("Digest not calculated")
	}

	hash := sha512.New()
	copyFileToHash(t, "Containerfile.dir", hash)
	copyFileToHash(t, "includeme/1.txt", hash)
	copyFileToHash(t, "includeme/2.txt", hash)

	expected := hash.Sum(nil)
	if !reflect.DeepEqual(digest[:], expected) {
		t.Fatalf("Unexpected digest:\n%X\nExpected:\n%X\n", digest, expected)
	}
}

func TestDigestOfFileAndAllInputsWithRemoteChanged(t *testing.T) {
	t.Parallel()

	firstDigest, err := DigestOfFileAndAllInputs("Containerfile.remote", options.TreatHttpAlwaysChanged())
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest == [sha512.Size]byte{} {
		t.Fatal("Digest not calculated")
	}
	secondDigest, err := DigestOfFileAndAllInputs("Containerfile.remote", options.TreatHttpAlwaysChanged())
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest == secondDigest {
		t.Fatal("Second digest not properly calculated")
	}
}

func TestDigestOfFileAndAllInputsWithRemoteUnchanged(t *testing.T) {
	t.Parallel()

	firstDigest, err := DigestOfFileAndAllInputs("Containerfile.remote", options.TreatHttpAlwaysUnchanged())
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest == [sha512.Size]byte{} {
		t.Fatal("Digest not calculated")
	}
	secondDigest, err := DigestOfFileAndAllInputs("Containerfile.remote", options.TreatHttpAlwaysUnchanged())
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if firstDigest != secondDigest {
		t.Fatal("Second digest not properly calculated")
	}
}

func TestDigestOfFileAndAllInputsPrecise(t *testing.T) {
	digest, err := DigestOfFileAndAllInputs("Containerfile.advanced")
	if err != nil {
		t.Fatal("DigestOfFileAndAllInputs failed:", err)
	}
	if digest == [sha512.Size]byte{} {
		t.Fatal("Digest not calculated")
	}
	hash := sha512.New()
	copyFileToHash(t, "Containerfile.advanced", hash)
	copyFileToHash(t, "aybabtu.txt", hash)
	copyFileToHash(t, "dummy.txt", hash)

	expected := hash.Sum(nil)
	if !reflect.DeepEqual(digest[:], expected) {
		t.Fatalf("Unexpected digest:\n%X\nExpected:\n%X\n", digest, expected)
	}
}

func copyFileToHash(t *testing.T, filename string, w io.Writer) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		panic(err)
	}
}
