// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetSwimLaneListRequest
	GetGroupName() *string
	SetServiceMeshId(v string) *GetSwimLaneListRequest
	GetServiceMeshId() *string
}

type GetSwimLaneListRequest struct {
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

func (s GetSwimLaneListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneListRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetSwimLaneListRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetSwimLaneListRequest) SetGroupName(v string) *GetSwimLaneListRequest {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneListRequest) SetServiceMeshId(v string) *GetSwimLaneListRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetSwimLaneListRequest) Validate() error {
	return dara.Validate(s)
}
