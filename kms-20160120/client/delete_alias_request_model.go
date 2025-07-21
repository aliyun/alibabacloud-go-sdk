// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DeleteAliasRequest
	GetAliasName() *string
}

type DeleteAliasRequest struct {
	// The alias that you want to delete.
	//
	// The value must be 1 to 255 characters in length and must include the alias/ prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// alias/example
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DeleteAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAliasRequest) GoString() string {
	return s.String()
}

func (s *DeleteAliasRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DeleteAliasRequest) SetAliasName(v string) *DeleteAliasRequest {
	s.AliasName = &v
	return s
}

func (s *DeleteAliasRequest) Validate() error {
	return dara.Validate(s)
}
