// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteIpRequest
	GetInstanceId() *string
	SetIpList(v string) *DeleteIpRequest
	GetIpList() *string
	SetRegionId(v string) *DeleteIpRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteIpRequest
	GetResourceGroupId() *string
}

type DeleteIpRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-npk1z7t9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to remove from the Anti-DDoS Origin instance. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **ip**: required. The IP address that you want to remove. Data type: string.
	//
	//     **
	//
	//     **Note*	- The IP addresses that you want to remove must be protected by the Anti-DDoS Origin instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ip":"1.XX.XX.1"},{"ip":"2.XX.XX.2"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *DeleteIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteIpRequest) SetInstanceId(v string) *DeleteIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetRegionId(v string) *DeleteIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpRequest) SetResourceGroupId(v string) *DeleteIpRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteIpRequest) Validate() error {
	return dara.Validate(s)
}
