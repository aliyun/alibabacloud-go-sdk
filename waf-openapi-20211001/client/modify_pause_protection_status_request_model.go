// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPauseProtectionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyPauseProtectionStatusRequest
	GetInstanceId() *string
	SetPauseStatus(v int32) *ModifyPauseProtectionStatusRequest
	GetPauseStatus() *int32
	SetRegionId(v string) *ModifyPauseProtectionStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyPauseProtectionStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyPauseProtectionStatusRequest struct {
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
	// Specifies whether to pause WAF protection.
	//
	// 	- **0**: does not pause WAF protection. This is the default value.
	//
	// 	- **1**: pauses WAF protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PauseStatus *int32 `json:"PauseStatus,omitempty" xml:"PauseStatus,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
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

func (s ModifyPauseProtectionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPauseProtectionStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPauseProtectionStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPauseProtectionStatusRequest) GetPauseStatus() *int32 {
	return s.PauseStatus
}

func (s *ModifyPauseProtectionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPauseProtectionStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyPauseProtectionStatusRequest) SetInstanceId(v string) *ModifyPauseProtectionStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPauseProtectionStatusRequest) SetPauseStatus(v int32) *ModifyPauseProtectionStatusRequest {
	s.PauseStatus = &v
	return s
}

func (s *ModifyPauseProtectionStatusRequest) SetRegionId(v string) *ModifyPauseProtectionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPauseProtectionStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyPauseProtectionStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyPauseProtectionStatusRequest) Validate() error {
	return dara.Validate(s)
}
