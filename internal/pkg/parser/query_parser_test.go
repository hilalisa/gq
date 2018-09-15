// Copyright 2018 HouseCanary, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseSimple(t *testing.T) {
	body := `{foo}`
	expected := `query {
  foo
}`
	doc, err := ParseQuery(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	buf := &bytes.Buffer{}
	doc.MarshallGraphQL(buf)
	if buf.String() != expected {
		spew.Dump(doc)
		t.Errorf("Got: \n%q\n expected:\n%q", buf.String(), expected)
	}
}

func TestParseError(t *testing.T) {
	body := `{`
	_, err := ParseQuery(body)
	if err == nil {
		t.Fatalf("expected parse error, got success")
	}
}

func BenchmarkParseSimple(b *testing.B) {
	body := `{foo}`
	for i := 0; i < b.N; i++ {
		ParseQuery(body)
	}
}

func TestParseArgs(t *testing.T) {
	body := `query($a: Int, $b: String = "a"){foo}`
	expected := `query($a: Int, $b: String = "a") {
  foo
}`
	doc, err := ParseQuery(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	buf := &bytes.Buffer{}
	doc.MarshallGraphQL(buf)
	if buf.String() != expected {
		spew.Dump(doc)
		t.Errorf("Got: \n%q\n expected:\n%q", buf.String(), expected)
	}
}

func TestParseInlineFragment(t *testing.T) {
	body := `query{... on Foo {a b}}`
	expected := `query {
  ... on Foo {
    a
    b
  }
}`
	doc, err := ParseQuery(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	buf := &bytes.Buffer{}
	doc.MarshallGraphQL(buf)
	if buf.String() != expected {
		spew.Dump(doc)
		t.Errorf("Got: \n%q\n expected:\n%q", buf.String(), expected)
	}
}
