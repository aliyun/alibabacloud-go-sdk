// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DisableKeyRequest
	GetKeyId() *string
}

type DisableKeyRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DisableKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableKeyRequest) GoString() string {
	return s.String()
}

func (s *DisableKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *DisableKeyRequest) SetKeyId(v string) *DisableKeyRequest {
	s.KeyId = &v
	return s
}

func (s *DisableKeyRequest) Validate() error {
	return dara.Validate(s)
}
