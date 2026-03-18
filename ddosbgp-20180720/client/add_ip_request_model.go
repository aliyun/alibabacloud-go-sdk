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
	// The IP addresses that you want to add to the Anti-DDoS Origin instance. This parameter is a string consisting of JSON arrays. Each element in a JSON array is a JSON struct that includes the following fields:
	//
	// 	- **ip**: required. The IP address that you want to add. Data type: string.
	//
	// 	- **member_uid**: optional. The member to which the asset belongs. Data type: string. This field is required only if the asset of a member is queried. Example: [{"ip":"121.41.XX.XX","member_uid":"120100811162\\*\\*\\*\\*"}].
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
