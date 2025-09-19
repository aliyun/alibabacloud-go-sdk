// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryId(v int64) *DeletePrivateRegistryRequest
	GetRegistryId() *int64
}

type DeletePrivateRegistryRequest struct {
	// The ID of the image repository.
	//
	// > You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to query the IDs of image repositories.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22233
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
}

func (s DeletePrivateRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRegistryRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateRegistryRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *DeletePrivateRegistryRequest) SetRegistryId(v int64) *DeletePrivateRegistryRequest {
	s.RegistryId = &v
	return s
}

func (s *DeletePrivateRegistryRequest) Validate() error {
	return dara.Validate(s)
}
