// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAICPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyName(v string) *DeleteAICPublicKeyRequest
	GetKeyName() *string
}

type DeleteAICPublicKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s DeleteAICPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAICPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAICPublicKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *DeleteAICPublicKeyRequest) SetKeyName(v string) *DeleteAICPublicKeyRequest {
	s.KeyName = &v
	return s
}

func (s *DeleteAICPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
