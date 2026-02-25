// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminPassword(v string) *RestoreInstanceRequest
	GetAdminPassword() *string
	SetAutoRenew(v bool) *RestoreInstanceRequest
	GetAutoRenew() *bool
	SetBackupTaskId(v string) *RestoreInstanceRequest
	GetBackupTaskId() *string
	SetDuration(v int32) *RestoreInstanceRequest
	GetDuration() *int32
	SetInstanceName(v string) *RestoreInstanceRequest
	GetInstanceName() *string
	SetPayType(v string) *RestoreInstanceRequest
	GetPayType() *string
	SetPricingCycle(v string) *RestoreInstanceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *RestoreInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RestoreInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*RestoreInstanceRequestTags) *RestoreInstanceRequest
	GetTags() []*RestoreInstanceRequestTags
	SetVSwitches(v []*RestoreInstanceRequestVSwitches) *RestoreInstanceRequest
	GetVSwitches() []*RestoreInstanceRequestVSwitches
	SetVpcId(v string) *RestoreInstanceRequest
	GetVpcId() *string
}

type RestoreInstanceRequest struct {
	// example:
	//
	// password_example
	AdminPassword *string `json:"AdminPassword,omitempty" xml:"AdminPassword,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// bk-adskj23hd9s
	BackupTaskId *string `json:"BackupTaskId,omitempty" xml:"BackupTaskId,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// c1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// prePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzllkih7jqxxx
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*RestoreInstanceRequestTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitches       []*RestoreInstanceRequestVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// vpc ID
	//
	// example:
	//
	// vpc-bp1fll2mci6d7pw8m****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RestoreInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequest) GetAdminPassword() *string {
	return s.AdminPassword
}

func (s *RestoreInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RestoreInstanceRequest) GetBackupTaskId() *string {
	return s.BackupTaskId
}

func (s *RestoreInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RestoreInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RestoreInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *RestoreInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RestoreInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestoreInstanceRequest) GetTags() []*RestoreInstanceRequestTags {
	return s.Tags
}

func (s *RestoreInstanceRequest) GetVSwitches() []*RestoreInstanceRequestVSwitches {
	return s.VSwitches
}

func (s *RestoreInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RestoreInstanceRequest) SetAdminPassword(v string) *RestoreInstanceRequest {
	s.AdminPassword = &v
	return s
}

func (s *RestoreInstanceRequest) SetAutoRenew(v bool) *RestoreInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RestoreInstanceRequest) SetBackupTaskId(v string) *RestoreInstanceRequest {
	s.BackupTaskId = &v
	return s
}

func (s *RestoreInstanceRequest) SetDuration(v int32) *RestoreInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RestoreInstanceRequest) SetInstanceName(v string) *RestoreInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *RestoreInstanceRequest) SetPayType(v string) *RestoreInstanceRequest {
	s.PayType = &v
	return s
}

func (s *RestoreInstanceRequest) SetPricingCycle(v string) *RestoreInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RestoreInstanceRequest) SetRegionId(v string) *RestoreInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreInstanceRequest) SetResourceGroupId(v string) *RestoreInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestoreInstanceRequest) SetTags(v []*RestoreInstanceRequestTags) *RestoreInstanceRequest {
	s.Tags = v
	return s
}

func (s *RestoreInstanceRequest) SetVSwitches(v []*RestoreInstanceRequestVSwitches) *RestoreInstanceRequest {
	s.VSwitches = v
	return s
}

func (s *RestoreInstanceRequest) SetVpcId(v string) *RestoreInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *RestoreInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RestoreInstanceRequestTags struct {
	// example:
	//
	// dukang-chengdu-sgueg
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// {\\"reasons\\": [], \\"patterns\\": [], \\"aggregates\\": [], \\"event_statistic\\": {\\"statistic\\": {}}}
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RestoreInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *RestoreInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *RestoreInstanceRequestTags) SetKey(v string) *RestoreInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *RestoreInstanceRequestTags) SetValue(v string) *RestoreInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *RestoreInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}

type RestoreInstanceRequestVSwitches struct {
	// example:
	//
	// vsw-bp19mlh98tm9teyyd****
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RestoreInstanceRequestVSwitches) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceRequestVSwitches) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequestVSwitches) GetVswId() *string {
	return s.VswId
}

func (s *RestoreInstanceRequestVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *RestoreInstanceRequestVSwitches) SetVswId(v string) *RestoreInstanceRequestVSwitches {
	s.VswId = &v
	return s
}

func (s *RestoreInstanceRequestVSwitches) SetZoneId(v string) *RestoreInstanceRequestVSwitches {
	s.ZoneId = &v
	return s
}

func (s *RestoreInstanceRequestVSwitches) Validate() error {
	return dara.Validate(s)
}
