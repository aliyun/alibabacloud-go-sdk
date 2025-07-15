// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAndroidInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpgradeAndroidInstanceGroupRequest
	GetAutoPay() *bool
	SetIncreaseNumberOfInstance(v int32) *UpgradeAndroidInstanceGroupRequest
	GetIncreaseNumberOfInstance() *int32
	SetInstanceGroupId(v string) *UpgradeAndroidInstanceGroupRequest
	GetInstanceGroupId() *string
}

type UpgradeAndroidInstanceGroupRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Make sure that your Alibaba Cloud account has sufficient balance.
	//
	// 	- false: disables the auto-payment feature. You need to manually complete the payment process.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The number of instances that you want to increase.
	//
	// example:
	//
	// 10
	IncreaseNumberOfInstance *int32 `json:"IncreaseNumberOfInstance,omitempty" xml:"IncreaseNumberOfInstance,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-asguicdjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
}

func (s UpgradeAndroidInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAndroidInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpgradeAndroidInstanceGroupRequest) GetIncreaseNumberOfInstance() *int32 {
	return s.IncreaseNumberOfInstance
}

func (s *UpgradeAndroidInstanceGroupRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *UpgradeAndroidInstanceGroupRequest) SetAutoPay(v bool) *UpgradeAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupRequest) SetIncreaseNumberOfInstance(v int32) *UpgradeAndroidInstanceGroupRequest {
	s.IncreaseNumberOfInstance = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *UpgradeAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *UpgradeAndroidInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
