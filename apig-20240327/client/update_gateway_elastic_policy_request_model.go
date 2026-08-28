// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayElasticPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticPolicy(v *UpdateGatewayElasticPolicyRequestElasticPolicy) *UpdateGatewayElasticPolicyRequest
	GetElasticPolicy() *UpdateGatewayElasticPolicyRequestElasticPolicy
}

type UpdateGatewayElasticPolicyRequest struct {
	ElasticPolicy *UpdateGatewayElasticPolicyRequestElasticPolicy `json:"elasticPolicy,omitempty" xml:"elasticPolicy,omitempty" type:"Struct"`
}

func (s UpdateGatewayElasticPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyRequest) GetElasticPolicy() *UpdateGatewayElasticPolicyRequestElasticPolicy {
	return s.ElasticPolicy
}

func (s *UpdateGatewayElasticPolicyRequest) SetElasticPolicy(v *UpdateGatewayElasticPolicyRequestElasticPolicy) *UpdateGatewayElasticPolicyRequest {
	s.ElasticPolicy = v
	return s
}

func (s *UpdateGatewayElasticPolicyRequest) Validate() error {
	if s.ElasticPolicy != nil {
		if err := s.ElasticPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayElasticPolicyRequestElasticPolicy struct {
	// example:
	//
	// true
	ElasticEnabled *bool `json:"elasticEnabled,omitempty" xml:"elasticEnabled,omitempty"`
	// example:
	//
	// CronHPA
	ElasticType               *string                                                                    `json:"elasticType,omitempty" xml:"elasticType,omitempty"`
	EnableScaleTimePolicyList []*UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList `json:"enableScaleTimePolicyList,omitempty" xml:"enableScaleTimePolicyList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	LoadWarningThreshold *bool `json:"loadWarningThreshold,omitempty" xml:"loadWarningThreshold,omitempty"`
	// example:
	//
	// 10
	MaxUnits       *int32                                                          `json:"maxUnits,omitempty" xml:"maxUnits,omitempty"`
	TimePolicyList []*UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList `json:"timePolicyList,omitempty" xml:"timePolicyList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicy) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetElasticEnabled() *bool {
	return s.ElasticEnabled
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetElasticType() *string {
	return s.ElasticType
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetEnableScaleTimePolicyList() []*UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList {
	return s.EnableScaleTimePolicyList
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetLoadWarningThreshold() *bool {
	return s.LoadWarningThreshold
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetMaxUnits() *int32 {
	return s.MaxUnits
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) GetTimePolicyList() []*UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList {
	return s.TimePolicyList
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetElasticEnabled(v bool) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.ElasticEnabled = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetElasticType(v string) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.ElasticType = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetEnableScaleTimePolicyList(v []*UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.EnableScaleTimePolicyList = v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetLoadWarningThreshold(v bool) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.LoadWarningThreshold = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetMaxUnits(v int32) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.MaxUnits = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) SetTimePolicyList(v []*UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) *UpdateGatewayElasticPolicyRequestElasticPolicy {
	s.TimePolicyList = v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicy) Validate() error {
	if s.EnableScaleTimePolicyList != nil {
		for _, item := range s.EnableScaleTimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimePolicyList != nil {
		for _, item := range s.TimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList struct {
	// example:
	//
	// 18:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 09:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) SetEndTime(v string) *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) SetStartTime(v string) *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyEnableScaleTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList struct {
	// example:
	//
	// 06:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 4
	Units *int32 `json:"units,omitempty" xml:"units,omitempty"`
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) GetUnits() *int32 {
	return s.Units
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) SetEndTime(v string) *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) SetStartTime(v string) *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) SetUnits(v int32) *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList {
	s.Units = &v
	return s
}

func (s *UpdateGatewayElasticPolicyRequestElasticPolicyTimePolicyList) Validate() error {
	return dara.Validate(s)
}
