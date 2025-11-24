// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *GetCaCertRequest
	GetServiceMeshId() *string
}

type GetCaCertRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c7894c929677643a5bfe1a732d778a****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetCaCertRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertRequest) GoString() string {
	return s.String()
}

func (s *GetCaCertRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetCaCertRequest) SetServiceMeshId(v string) *GetCaCertRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetCaCertRequest) Validate() error {
	return dara.Validate(s)
}
