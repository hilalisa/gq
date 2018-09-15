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
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/housecanary/gq/internal/pkg/parser/gen"
)

func safeParse(input string, cb func(*gen.GraphqlParser)) (err ParseError) {
	err = nil
	is := antlr.NewInputStream(input)
	lexer := gen.NewGraphqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := gen.NewGraphqlParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(panicListener{})
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case ParseError:
				err = t
			default:
				panic(r)
			}
		}
	}()
	cb(p)
	return
}
