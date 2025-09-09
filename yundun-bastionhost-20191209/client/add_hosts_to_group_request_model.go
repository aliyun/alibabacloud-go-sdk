// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHostsToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *AddHostsToGroupRequest
	GetHostGroupId() *string
	SetHostIds(v string) *AddHostsToGroupRequest
	GetHostIds() *string
	SetInstanceId(v string) *AddHostsToGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AddHostsToGroupRequest
	GetRegionId() *string
}

type AddHostsToGroupRequest struct {
	// The ID of the asset group to which you want to add hosts.
	//
	// >You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the asset group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The IDs of the hosts that you want to add to the asset group. Specify a JSON string. You can specify up to 100 host IDs.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the host IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host whose asset group you want to add hosts to.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to add hosts to.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddHostsToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHostsToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AddHostsToGroupRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *AddHostsToGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHostsToGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddHostsToGroupRequest) SetHostGroupId(v string) *AddHostsToGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetHostIds(v string) *AddHostsToGroupRequest {
	s.HostIds = &v
	return s
}

func (s *AddHostsToGroupRequest) SetInstanceId(v string) *AddHostsToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetRegionId(v string) *AddHostsToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddHostsToGroupRequest) Validate() error {
	return dara.Validate(s)
}
