// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *CreateAliasRequest
	GetAlias() *string
	SetIndex(v string) *CreateAliasRequest
	GetIndex() *string
	SetNewMode(v bool) *CreateAliasRequest
	GetNewMode() *bool
}

type CreateAliasRequest struct {
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
	// Specifies whether the OpenSearch Vector Search Edition instance is of the new version.
	//
	// example:
	//
	// true
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) GetAlias() *string {
	return s.Alias
}

func (s *CreateAliasRequest) GetIndex() *string {
	return s.Index
}

func (s *CreateAliasRequest) GetNewMode() *bool {
	return s.NewMode
}

func (s *CreateAliasRequest) SetAlias(v string) *CreateAliasRequest {
	s.Alias = &v
	return s
}

func (s *CreateAliasRequest) SetIndex(v string) *CreateAliasRequest {
	s.Index = &v
	return s
}

func (s *CreateAliasRequest) SetNewMode(v bool) *CreateAliasRequest {
	s.NewMode = &v
	return s
}

func (s *CreateAliasRequest) Validate() error {
	return dara.Validate(s)
}
