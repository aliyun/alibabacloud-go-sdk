// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyRemarkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyRemarkRequest
	GetRegionId() *string
	SetRemark(v string) *ModifyRemarkRequest
	GetRemark() *string
	SetResourceGroupId(v string) *ModifyRemarkRequest
	GetResourceGroupId() *string
}

type ModifyRemarkRequest struct {
	// The ID of the Anti-DDoS Origin instance for which you want to add remarks.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks for the Anti-DDoS Origin instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRemarkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyRemarkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyRemarkRequest) SetInstanceId(v string) *ModifyRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRegionId(v string) *ModifyRemarkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRemark(v string) *ModifyRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyRemarkRequest) SetResourceGroupId(v string) *ModifyRemarkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyRemarkRequest) Validate() error {
	return dara.Validate(s)
}
