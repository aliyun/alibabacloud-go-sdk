// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ModifyAliasRequest
	GetAlias() *string
	SetIndex(v string) *ModifyAliasRequest
	GetIndex() *string
}

type ModifyAliasRequest struct {
	// alias name
	//
	// example:
	//
	// test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// index name
	//
	// example:
	//
	// index
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
}

func (s ModifyAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAliasRequest) GoString() string {
	return s.String()
}

func (s *ModifyAliasRequest) GetAlias() *string {
	return s.Alias
}

func (s *ModifyAliasRequest) GetIndex() *string {
	return s.Index
}

func (s *ModifyAliasRequest) SetAlias(v string) *ModifyAliasRequest {
	s.Alias = &v
	return s
}

func (s *ModifyAliasRequest) SetIndex(v string) *ModifyAliasRequest {
	s.Index = &v
	return s
}

func (s *ModifyAliasRequest) Validate() error {
	return dara.Validate(s)
}
