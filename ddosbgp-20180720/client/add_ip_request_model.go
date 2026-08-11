// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddIpRequest
	GetInstanceId() *string
	SetIpList(v string) *AddIpRequest
	GetIpList() *string
	SetRegionId(v string) *AddIpRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AddIpRequest
	GetResourceGroupId() *string
}

type AddIpRequest struct {
	// The ID of the Anti-DDoS Origin instance to manage.
	//
	// > Call [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-npk1z7t9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses to add to the Anti-DDoS Origin instance for protection. The value is a string that is converted from a JSON array. Each element in the JSON array is a struct that contains the following fields:
	//
	// - **ip**: The IP address to add. This parameter is of the String type and is required.
	//
	// - **member_uid**: The ID of the member account that owns the asset. This parameter is of the String type and is optional. Specify this parameter only when you add an asset that belongs to a member account. For example: [{"ip":"121.41.XX.XX","member_uid":"120100811162\\*\\*\\*\\*"}]
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ip":"121.41.XX.XX"},{"ip":"121.42.XX.XX"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// > Call [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) to query information about all regions that Anti-DDoS Origin supports.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the resource group in Resource Management to which the Anti-DDoS Origin instance belongs. If you leave this parameter empty, the instance is added to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *AddIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddIpRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddIpRequest) SetInstanceId(v string) *AddIpRequest {
	s.InstanceId = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

func (s *AddIpRequest) SetRegionId(v string) *AddIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpRequest) SetResourceGroupId(v string) *AddIpRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddIpRequest) Validate() error {
	return dara.Validate(s)
}
