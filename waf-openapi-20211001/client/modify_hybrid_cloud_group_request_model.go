// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *ModifyHybridCloudGroupRequest
	GetClusterId() *int64
	SetGroupId(v int64) *ModifyHybridCloudGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *ModifyHybridCloudGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *ModifyHybridCloudGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHybridCloudGroupRequest
	GetRegionId() *string
	SetRemark(v string) *ModifyHybridCloudGroupRequest
	GetRemark() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyHybridCloudGroupRequest struct {
	// The ID of the cluster.
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
	// The name of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid value:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyHybridCloudGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *ModifyHybridCloudGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyHybridCloudGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyHybridCloudGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyHybridCloudGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudGroupRequest) SetClusterId(v int64) *ModifyHybridCloudGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetGroupId(v int64) *ModifyHybridCloudGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetGroupName(v string) *ModifyHybridCloudGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetInstanceId(v string) *ModifyHybridCloudGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetRegionId(v string) *ModifyHybridCloudGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetRemark(v string) *ModifyHybridCloudGroupRequest {
	s.Remark = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupRequest) Validate() error {
	return dara.Validate(s)
}
