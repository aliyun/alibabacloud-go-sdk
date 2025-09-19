// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRegistryExtraRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryId(v int64) *ListImageRegistryExtraRequest
	GetRegistryId() *int64
}

type ListImageRegistryExtraRequest struct {
	// Image registry ID.
	//
	// > You can obtain this parameter by calling the [PageImageRegistry](~~PageImageRegistry~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 25363
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
}

func (s ListImageRegistryExtraRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryExtraRequest) GoString() string {
	return s.String()
}

func (s *ListImageRegistryExtraRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *ListImageRegistryExtraRequest) SetRegistryId(v int64) *ListImageRegistryExtraRequest {
	s.RegistryId = &v
	return s
}

func (s *ListImageRegistryExtraRequest) Validate() error {
	return dara.Validate(s)
}
