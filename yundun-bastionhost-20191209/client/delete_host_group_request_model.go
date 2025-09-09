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
	SetRegionId(v string) *DeleteHostGroupRequest
	GetRegionId() *string
}

type DeleteHostGroupRequest struct {
	// The ID of the asset group that you want to delete.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the asset group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host whose asset group you want to delete.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to delete.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *DeleteHostGroupRequest) SetRegionId(v string) *DeleteHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
