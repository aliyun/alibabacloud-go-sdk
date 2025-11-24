// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmAppMeshInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *GetVmAppMeshInfoRequest
	GetServiceMeshId() *string
}

type GetVmAppMeshInfoRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce51a7de4a5144db88a864****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetVmAppMeshInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVmAppMeshInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetVmAppMeshInfoRequest) SetServiceMeshId(v string) *GetVmAppMeshInfoRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetVmAppMeshInfoRequest) Validate() error {
	return dara.Validate(s)
}
