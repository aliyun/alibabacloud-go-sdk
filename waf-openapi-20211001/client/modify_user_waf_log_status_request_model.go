// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserWafLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserWafLogStatusRequest
	GetInstanceId() *string
	SetLogRegionId(v string) *ModifyUserWafLogStatusRequest
	GetLogRegionId() *string
	SetLogStatus(v int32) *ModifyUserWafLogStatusRequest
	GetLogStatus() *int32
	SetRegionId(v string) *ModifyUserWafLogStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyUserWafLogStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyUserWafLogStatusRequest struct {
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the log storage region. If this parameter is not specified, Simple Log Service is enabled in the region where the WAF instance resides.
	//
	// - **cn-hangzhou**: the default region where Simple Log Service is enabled for a WAF instance in the Chinese mainland.
	//
	// - **ap-southeast-1**: the default region where Simple Log Service is enabled for a WAF instance outside the Chinese mainland.
	//
	// > Call [DescribeUserSlsLogRegions](https://help.aliyun.com/document_detail/2712598.html) to query the available log storage regions.
	//
	// example:
	//
	// cn-beijing
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// Indicates whether Simple Log Service is enabled. Valid values:
	//
	// - **0**: Simple Log Service is disabled.
	//
	// - **1**: Simple Log Service is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LogStatus *int32 `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyUserWafLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserWafLogStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserWafLogStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserWafLogStatusRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ModifyUserWafLogStatusRequest) GetLogStatus() *int32 {
	return s.LogStatus
}

func (s *ModifyUserWafLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserWafLogStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyUserWafLogStatusRequest) SetInstanceId(v string) *ModifyUserWafLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserWafLogStatusRequest) SetLogRegionId(v string) *ModifyUserWafLogStatusRequest {
	s.LogRegionId = &v
	return s
}

func (s *ModifyUserWafLogStatusRequest) SetLogStatus(v int32) *ModifyUserWafLogStatusRequest {
	s.LogStatus = &v
	return s
}

func (s *ModifyUserWafLogStatusRequest) SetRegionId(v string) *ModifyUserWafLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserWafLogStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyUserWafLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyUserWafLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
