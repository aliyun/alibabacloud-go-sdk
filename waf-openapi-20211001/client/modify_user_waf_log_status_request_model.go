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
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LogStatus *int32 `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
