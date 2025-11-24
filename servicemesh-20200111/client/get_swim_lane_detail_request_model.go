// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetSwimLaneDetailRequest
	GetGroupName() *string
	SetServiceMeshId(v string) *GetSwimLaneDetailRequest
	GetServiceMeshId() *string
	SetSwimLaneName(v string) *GetSwimLaneDetailRequest
	GetSwimLaneName() *string
}

type GetSwimLaneDetailRequest struct {
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

func (s GetSwimLaneDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetSwimLaneDetailRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetSwimLaneDetailRequest) GetSwimLaneName() *string {
	return s.SwimLaneName
}

func (s *GetSwimLaneDetailRequest) SetGroupName(v string) *GetSwimLaneDetailRequest {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneDetailRequest) SetServiceMeshId(v string) *GetSwimLaneDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetSwimLaneDetailRequest) SetSwimLaneName(v string) *GetSwimLaneDetailRequest {
	s.SwimLaneName = &v
	return s
}

func (s *GetSwimLaneDetailRequest) Validate() error {
	return dara.Validate(s)
}
