// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateHostGroupRequest
	GetComment() *string
	SetHostGroupName(v string) *CreateHostGroupRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *CreateHostGroupRequest
	GetInstanceId() *string
	SetProjectId(v int64) *CreateHostGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *CreateHostGroupRequest
	GetRegionId() *string
}

type CreateHostGroupRequest struct {
	// The description of the asset group. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Local host group
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the asset group. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// HostGroup01
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The instance ID of the bastion host where you want to create the asset group.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host instance where you want to create the asset group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateHostGroupRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *CreateHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHostGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHostGroupRequest) SetComment(v string) *CreateHostGroupRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostGroupRequest) SetHostGroupName(v string) *CreateHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *CreateHostGroupRequest) SetInstanceId(v string) *CreateHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostGroupRequest) SetProjectId(v int64) *CreateHostGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateHostGroupRequest) SetRegionId(v string) *CreateHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
