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
	// Instance ID of the Anti-DDoS Origin instance.
	//
	// > You can call [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) to query instance IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-npk1z7t9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses to remove from the protected objects of the Anti-DDoS Origin instance. The value is a string converted from a JSON array. Each element in the JSON array is a struct that contains the following fields:
	//
	// - **ip**: String | Required | The IP address to remove.
	//
	//
	//
	//     > The IP address must already be in the protection list of the Anti-DDoS Origin instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ip":"1.XX.XX.1"},{"ip":"2.XX.XX.2"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// > You can call [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) to query all regions supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. Default value: empty, which indicates the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
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
