// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteSwimLaneGroupRequest
	GetGroupName() *string
	SetServiceMeshId(v string) *DeleteSwimLaneGroupRequest
	GetServiceMeshId() *string
}

type DeleteSwimLaneGroupRequest struct {
	// The name of the lane group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteSwimLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteSwimLaneGroupRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteSwimLaneGroupRequest) SetGroupName(v string) *DeleteSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteSwimLaneGroupRequest) SetServiceMeshId(v string) *DeleteSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteSwimLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
