// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetProjectRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetProjectRequest
	GetRegionId() *string
}

type GetProjectRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the project to query.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetProjectRequest) SetInstanceId(v string) *GetProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *GetProjectRequest) SetProjectId(v int64) *GetProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRequest) SetRegionId(v string) *GetProjectRequest {
	s.RegionId = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
