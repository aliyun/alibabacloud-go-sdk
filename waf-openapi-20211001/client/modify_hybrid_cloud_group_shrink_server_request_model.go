// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupShrinkServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *ModifyHybridCloudGroupShrinkServerRequest
	GetClusterId() *int64
	SetGroupId(v int64) *ModifyHybridCloudGroupShrinkServerRequest
	GetGroupId() *int64
	SetInstanceId(v string) *ModifyHybridCloudGroupShrinkServerRequest
	GetInstanceId() *string
	SetMids(v string) *ModifyHybridCloudGroupShrinkServerRequest
	GetMids() *string
	SetRegionId(v string) *ModifyHybridCloudGroupShrinkServerRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupShrinkServerRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyHybridCloudGroupShrinkServerRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldb****05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The machine ID of the node that you want to remove from the node group. You can specify multiple machine IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// b1bf3f544f30c1de0b72d91290***bbbb
	Mids *string `json:"Mids,omitempty" xml:"Mids,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm2th****v6ay
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyHybridCloudGroupShrinkServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupShrinkServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetMids() *string {
	return s.Mids
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetClusterId(v int64) *ModifyHybridCloudGroupShrinkServerRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetGroupId(v int64) *ModifyHybridCloudGroupShrinkServerRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetInstanceId(v string) *ModifyHybridCloudGroupShrinkServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetMids(v string) *ModifyHybridCloudGroupShrinkServerRequest {
	s.Mids = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetRegionId(v string) *ModifyHybridCloudGroupShrinkServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupShrinkServerRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerRequest) Validate() error {
	return dara.Validate(s)
}
