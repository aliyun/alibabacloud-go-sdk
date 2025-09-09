// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHostsFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *RemoveHostsFromGroupRequest
	GetHostGroupId() *string
	SetHostIds(v string) *RemoveHostsFromGroupRequest
	GetHostIds() *string
	SetInstanceId(v string) *RemoveHostsFromGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *RemoveHostsFromGroupRequest
	GetRegionId() *string
}

type RemoveHostsFromGroupRequest struct {
	// The ID of the asset group from which you want to remove hosts.
	//
	// >  You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the asset group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The IDs of the hosts that you want to remove from the host group. Specify a JSON string. You can specify up to 100 host IDs.
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the host IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host whose asset group you want to manage.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to manage.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveHostsFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveHostsFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *RemoveHostsFromGroupRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *RemoveHostsFromGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveHostsFromGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveHostsFromGroupRequest) SetHostGroupId(v string) *RemoveHostsFromGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetHostIds(v string) *RemoveHostsFromGroupRequest {
	s.HostIds = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetInstanceId(v string) *RemoveHostsFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetRegionId(v string) *RemoveHostsFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
