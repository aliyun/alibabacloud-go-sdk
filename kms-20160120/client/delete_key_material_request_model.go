// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DeleteKeyMaterialRequest
	GetKeyId() *string
}

type DeleteKeyMaterialRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DeleteKeyMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *DeleteKeyMaterialRequest) SetKeyId(v string) *DeleteKeyMaterialRequest {
	s.KeyId = &v
	return s
}

func (s *DeleteKeyMaterialRequest) Validate() error {
	return dara.Validate(s)
}
