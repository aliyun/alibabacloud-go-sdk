// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteSwimLaneRequest
	GetGroupName() *string
	SetServiceMeshId(v string) *DeleteSwimLaneRequest
	GetServiceMeshId() *string
	SetSwimLaneName(v string) *DeleteSwimLaneRequest
	GetSwimLaneName() *string
}

type DeleteSwimLaneRequest struct {
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
	// The name of the lane.
	//
	// example:
	//
	// s1
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s DeleteSwimLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteSwimLaneRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteSwimLaneRequest) GetSwimLaneName() *string {
	return s.SwimLaneName
}

func (s *DeleteSwimLaneRequest) SetGroupName(v string) *DeleteSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteSwimLaneRequest) SetServiceMeshId(v string) *DeleteSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteSwimLaneRequest) SetSwimLaneName(v string) *DeleteSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

func (s *DeleteSwimLaneRequest) Validate() error {
	return dara.Validate(s)
}
