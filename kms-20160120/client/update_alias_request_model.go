// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *UpdateAliasRequest
	GetAliasName() *string
	SetKeyId(v string) *UpdateAliasRequest
	GetKeyId() *string
}

type UpdateAliasRequest struct {
	// The alias that you want to bind.
	//
	// The value must be 1 to 255 characters in length and must include the alias/ prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// alias/example
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdateAliasRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *UpdateAliasRequest) SetAliasName(v string) *UpdateAliasRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateAliasRequest) SetKeyId(v string) *UpdateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateAliasRequest) Validate() error {
	return dara.Validate(s)
}
