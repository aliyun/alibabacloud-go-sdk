// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAccessTimeRangeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTimeRangeConfig(v *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) *SetPolicyAccessTimeRangeConfigRequest
	GetAccessTimeRangeConfig() *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig
	SetInstanceId(v string) *SetPolicyAccessTimeRangeConfigRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyAccessTimeRangeConfigRequest
	GetPolicyId() *string
	SetRegionId(v string) *SetPolicyAccessTimeRangeConfigRequest
	GetRegionId() *string
}

type SetPolicyAccessTimeRangeConfigRequest struct {
	// The logon period limits.
	//
	// This parameter is required.
	AccessTimeRangeConfig *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig `json:"AccessTimeRangeConfig,omitempty" xml:"AccessTimeRangeConfig,omitempty" type:"Struct"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The control policy ID.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyAccessTimeRangeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigRequest) GetAccessTimeRangeConfig() *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig {
	return s.AccessTimeRangeConfig
}

func (s *SetPolicyAccessTimeRangeConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyAccessTimeRangeConfigRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyAccessTimeRangeConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyAccessTimeRangeConfigRequest) SetAccessTimeRangeConfig(v *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) *SetPolicyAccessTimeRangeConfigRequest {
	s.AccessTimeRangeConfig = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) SetInstanceId(v string) *SetPolicyAccessTimeRangeConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) SetPolicyId(v string) *SetPolicyAccessTimeRangeConfigRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) SetRegionId(v string) *SetPolicyAccessTimeRangeConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) Validate() error {
	return dara.Validate(s)
}

type SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig struct {
	// The details about the periods during which users can log on to the assets.
	EffectiveTime []*SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty" type:"Repeated"`
}

func (s SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) GetEffectiveTime() []*SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime {
	return s.EffectiveTime
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) SetEffectiveTime(v []*SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig {
	s.EffectiveTime = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig) Validate() error {
	return dara.Validate(s)
}

type SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime struct {
	// The days of the week during which users can log on to the assets.
	Days []*int32 `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The time periods of the day during which users can log on to the assets.
	Hours []*int32 `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
}

func (s SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) GetDays() []*int32 {
	return s.Days
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) GetHours() []*int32 {
	return s.Hours
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) SetDays(v []*int32) *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime {
	s.Days = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) SetHours(v []*int32) *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime {
	s.Hours = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime) Validate() error {
	return dara.Validate(s)
}
