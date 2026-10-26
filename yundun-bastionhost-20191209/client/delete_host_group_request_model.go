// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *DeleteHostGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *DeleteHostGroupRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteHostGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteHostGroupRequest
	GetRegionId() *string
}

type DeleteHostGroupRequest struct {
	// The ID of the asset group to be deleted.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The instance ID of the bastion host to which the asset group to be deleted belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
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
	// The region ID of the bastion host to which the asset group to be deleted belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *DeleteHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHostGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostGroupRequest) SetHostGroupId(v string) *DeleteHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *DeleteHostGroupRequest) SetInstanceId(v string) *DeleteHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostGroupRequest) SetProjectId(v int64) *DeleteHostGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteHostGroupRequest) SetRegionId(v string) *DeleteHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
