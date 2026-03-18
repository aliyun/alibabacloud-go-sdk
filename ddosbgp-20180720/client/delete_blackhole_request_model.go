// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackholeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteBlackholeRequest
	GetInstanceId() *string
	SetIp(v string) *DeleteBlackholeRequest
	GetIp() *string
	SetRegionId(v string) *DeleteBlackholeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteBlackholeRequest
	GetResourceGroupId() *string
}

type DeleteBlackholeRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address for which you want to deactivate blackhole filtering.
	//
	// >  You can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query all the IP addresses that are protected by the Anti-DDoS Origin instance and the protection status of the IP addresses. For example, you can query whether blackhole filtering is triggered for an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBlackholeRequest) GetIp() *string {
	return s.Ip
}

func (s *DeleteBlackholeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBlackholeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteBlackholeRequest) SetInstanceId(v string) *DeleteBlackholeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetRegionId(v string) *DeleteBlackholeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetResourceGroupId(v string) *DeleteBlackholeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteBlackholeRequest) Validate() error {
	return dara.Validate(s)
}
