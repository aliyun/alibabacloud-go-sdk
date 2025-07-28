// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupExpansionServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *ModifyHybridCloudGroupExpansionServerRequest
	GetClusterId() *int64
	SetGroupId(v int64) *ModifyHybridCloudGroupExpansionServerRequest
	GetGroupId() *int64
	SetInstanceId(v string) *ModifyHybridCloudGroupExpansionServerRequest
	GetInstanceId() *string
	SetMids(v string) *ModifyHybridCloudGroupExpansionServerRequest
	GetMids() *string
	SetRegionId(v string) *ModifyHybridCloudGroupExpansionServerRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupExpansionServerRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyHybridCloudGroupExpansionServerRequest struct {
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
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 612929b133e7ff**0d0
	Mids *string `json:"Mids,omitempty" xml:"Mids,omitempty"`
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
}

func (s ModifyHybridCloudGroupExpansionServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupExpansionServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetMids() *string {
	return s.Mids
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetClusterId(v int64) *ModifyHybridCloudGroupExpansionServerRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetGroupId(v int64) *ModifyHybridCloudGroupExpansionServerRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetInstanceId(v string) *ModifyHybridCloudGroupExpansionServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetMids(v string) *ModifyHybridCloudGroupExpansionServerRequest {
	s.Mids = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetRegionId(v string) *ModifyHybridCloudGroupExpansionServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudGroupExpansionServerRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerRequest) Validate() error {
	return dara.Validate(s)
}
