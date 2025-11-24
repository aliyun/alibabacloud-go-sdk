// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVMIntoServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsId(v string) *AddVMIntoServiceMeshRequest
	GetEcsId() *string
	SetServiceMeshId(v string) *AddVMIntoServiceMeshRequest
	GetServiceMeshId() *string
}

type AddVMIntoServiceMeshRequest struct {
	// The ID of the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze90ts4e7dj3650****
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccb37ff104caf419fbf48fb38e6f3****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s AddVMIntoServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVMIntoServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshRequest) GetEcsId() *string {
	return s.EcsId
}

func (s *AddVMIntoServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *AddVMIntoServiceMeshRequest) SetEcsId(v string) *AddVMIntoServiceMeshRequest {
	s.EcsId = &v
	return s
}

func (s *AddVMIntoServiceMeshRequest) SetServiceMeshId(v string) *AddVMIntoServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *AddVMIntoServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
