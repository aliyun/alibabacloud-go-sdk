// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *CreateAliasRequest
	GetAliasName() *string
	SetKeyId(v string) *CreateAliasRequest
	GetKeyId() *string
}

type CreateAliasRequest struct {
	// The alias of the CMK.
	//
	// The alias must be 1 to 255 characters in length and must contain the prefix `alias/`. The alias cannot be prefixed with the reserved word `alias/acs`.
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
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateAliasRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateAliasRequest) SetAliasName(v string) *CreateAliasRequest {
	s.AliasName = &v
	return s
}

func (s *CreateAliasRequest) SetKeyId(v string) *CreateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *CreateAliasRequest) Validate() error {
	return dara.Validate(s)
}
