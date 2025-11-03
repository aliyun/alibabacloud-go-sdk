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
	// The ID of the image repository.
	//
	// >  You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to obtain the ID.
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
