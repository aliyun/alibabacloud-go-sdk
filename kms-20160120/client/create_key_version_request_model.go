// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *CreateKeyVersionRequest
	GetKeyId() *string
}

type CreateKeyVersionRequest struct {
	// The key ID, which is the globally unique identifier of the key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz62f1cb66fa42qo***
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CreateKeyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateKeyVersionRequest) SetKeyId(v string) *CreateKeyVersionRequest {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionRequest) Validate() error {
	return dara.Validate(s)
}
