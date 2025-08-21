// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyName(v string) *DeletePublicKeyRequest
	GetKeyName() *string
}

type DeletePublicKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// appKey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s DeletePublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *DeletePublicKeyRequest) SetKeyName(v string) *DeletePublicKeyRequest {
	s.KeyName = &v
	return s
}

func (s *DeletePublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
