// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterBypassStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterResourceId(v string) *ModifyHybridCloudClusterBypassStatusRequest
	GetClusterResourceId() *string
	SetInstanceId(v string) *ModifyHybridCloudClusterBypassStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHybridCloudClusterBypassStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterBypassStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleStatus(v string) *ModifyHybridCloudClusterBypassStatusRequest
	GetRuleStatus() *string
}

type ModifyHybridCloudClusterBypassStatusRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdbc_cluster_****
	ClusterResourceId *string `json:"ClusterResourceId,omitempty" xml:"ClusterResourceId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// **
	//
	// **You can call the **DescribeInstanceInfo[ operation to obtain the ID of the WAF instance.](https://help.aliyun.com/document_detail/140857.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The status of manual bypass. Valid values:
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) GetClusterResourceId() *string {
	return s.ClusterResourceId
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetClusterResourceId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.ClusterResourceId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetInstanceId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetRegionId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetRuleStatus(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) Validate() error {
	return dara.Validate(s)
}
