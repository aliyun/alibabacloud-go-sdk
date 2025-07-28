// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncProductInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SyncProductInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *SyncProductInstanceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *SyncProductInstanceRequest
	GetResourceManagerResourceGroupId() *string
}

type SyncProductInstanceRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zsk****fb09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
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
	// rg-acfm4co****f5qa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s SyncProductInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncProductInstanceRequest) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncProductInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SyncProductInstanceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *SyncProductInstanceRequest) SetInstanceId(v string) *SyncProductInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncProductInstanceRequest) SetRegionId(v string) *SyncProductInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *SyncProductInstanceRequest) SetResourceManagerResourceGroupId(v string) *SyncProductInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *SyncProductInstanceRequest) Validate() error {
	return dara.Validate(s)
}
