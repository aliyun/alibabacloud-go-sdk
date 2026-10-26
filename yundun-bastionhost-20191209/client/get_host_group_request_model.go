// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *GetHostGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *GetHostGroupRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetHostGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetHostGroupRequest
	GetRegionId() *string
}

type GetHostGroupRequest struct {
	// The ID of the asset group that you want to query.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The instance ID of the bastion host to which the asset group belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The region ID of the bastion host to which the asset group belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupRequest) GoString() string {
	return s.String()
}

func (s *GetHostGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *GetHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostGroupRequest) SetHostGroupId(v string) *GetHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *GetHostGroupRequest) SetInstanceId(v string) *GetHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostGroupRequest) SetProjectId(v int64) *GetHostGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *GetHostGroupRequest) SetRegionId(v string) *GetHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
