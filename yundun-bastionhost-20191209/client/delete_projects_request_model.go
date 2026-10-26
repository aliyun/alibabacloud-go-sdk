// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteProjectsRequest
	GetInstanceId() *string
	SetProjectIds(v string) *DeleteProjectsRequest
	GetProjectIds() *string
	SetRegionId(v string) *DeleteProjectsRequest
	GetRegionId() *string
}

type DeleteProjectsRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the DescribeInstances operation to obtain the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the projects that you want to delete.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2]
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteProjectsRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *DeleteProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteProjectsRequest) SetInstanceId(v string) *DeleteProjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteProjectsRequest) SetProjectIds(v string) *DeleteProjectsRequest {
	s.ProjectIds = &v
	return s
}

func (s *DeleteProjectsRequest) SetRegionId(v string) *DeleteProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProjectsRequest) Validate() error {
	return dara.Validate(s)
}
