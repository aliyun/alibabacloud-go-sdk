// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVMFromServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsId(v string) *RemoveVMFromServiceMeshRequest
	GetEcsId() *string
	SetServiceMeshId(v string) *RemoveVMFromServiceMeshRequest
	GetServiceMeshId() *string
}

type RemoveVMFromServiceMeshRequest struct {
	// The ID of the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze90ts4edj3650****
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

func (s RemoveVMFromServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVMFromServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshRequest) GetEcsId() *string {
	return s.EcsId
}

func (s *RemoveVMFromServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *RemoveVMFromServiceMeshRequest) SetEcsId(v string) *RemoveVMFromServiceMeshRequest {
	s.EcsId = &v
	return s
}

func (s *RemoveVMFromServiceMeshRequest) SetServiceMeshId(v string) *RemoveVMFromServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *RemoveVMFromServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
