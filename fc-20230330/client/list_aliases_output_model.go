// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetAliases(v []*Alias) *ListAliasesOutput
	GetAliases() []*Alias
	SetNextToken(v string) *ListAliasesOutput
	GetNextToken() *string
}

type ListAliasesOutput struct {
	Aliases []*Alias `json:"aliases" xml:"aliases" type:"Repeated"`
	// example:
	//
	// test
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAliasesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesOutput) GoString() string {
	return s.String()
}

func (s *ListAliasesOutput) GetAliases() []*Alias {
	return s.Aliases
}

func (s *ListAliasesOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAliasesOutput) SetAliases(v []*Alias) *ListAliasesOutput {
	s.Aliases = v
	return s
}

func (s *ListAliasesOutput) SetNextToken(v string) *ListAliasesOutput {
	s.NextToken = &v
	return s
}

func (s *ListAliasesOutput) Validate() error {
	return dara.Validate(s)
}
