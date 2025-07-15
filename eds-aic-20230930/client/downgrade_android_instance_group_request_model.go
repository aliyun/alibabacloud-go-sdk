// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeAndroidInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DowngradeAndroidInstanceGroupRequest
	GetAndroidInstanceIds() []*string
	SetAutoPay(v bool) *DowngradeAndroidInstanceGroupRequest
	GetAutoPay() *bool
	SetInstanceGroupId(v string) *DowngradeAndroidInstanceGroupRequest
	GetInstanceGroupId() *string
}

type DowngradeAndroidInstanceGroupRequest struct {
	// The IDs of the cloud phone instances that you want to delete.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enable the auto-payment feature. Default value: false.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Ensure your account has sufficient balance to use this feature.
	//
	// 	- false: disables the auto-payment feature. This requires manual payment each time you place an order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the instance group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ag-cuv4scs4obxhs****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
}

func (s DowngradeAndroidInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DowngradeAndroidInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *DowngradeAndroidInstanceGroupRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DowngradeAndroidInstanceGroupRequest) SetAndroidInstanceIds(v []*string) *DowngradeAndroidInstanceGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DowngradeAndroidInstanceGroupRequest) SetAutoPay(v bool) *DowngradeAndroidInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *DowngradeAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
