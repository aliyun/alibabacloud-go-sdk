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
	SetProjectId(v int64) *SetPolicyAccessTimeRangeConfigRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyAccessTimeRangeConfigRequest
	GetRegionId() *string
}

type SetPolicyAccessTimeRangeConfigRequest struct {
	// The logon time limit configuration.
	//
	// This parameter is required.
	AccessTimeRangeConfig *SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig `json:"AccessTimeRangeConfig,omitempty" xml:"AccessTimeRangeConfig,omitempty" type:"Struct"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *SetPolicyAccessTimeRangeConfigRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *SetPolicyAccessTimeRangeConfigRequest) SetProjectId(v int64) *SetPolicyAccessTimeRangeConfigRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) SetRegionId(v string) *SetPolicyAccessTimeRangeConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigRequest) Validate() error {
	if s.AccessTimeRangeConfig != nil {
		if err := s.AccessTimeRangeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfig struct {
	// The details of the allowed logon periods.
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
	if s.EffectiveTime != nil {
		for _, item := range s.EffectiveTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetPolicyAccessTimeRangeConfigRequestAccessTimeRangeConfigEffectiveTime struct {
	// The days of the week when logon is allowed.
	Days []*int32 `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The hours of the day when logon is allowed.
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
