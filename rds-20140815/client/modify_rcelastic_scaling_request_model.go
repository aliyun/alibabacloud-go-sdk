// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCElasticScalingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyRCElasticScalingRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyRCElasticScalingRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyRCElasticScalingRequest
	GetInstanceType() *string
	SetRegionId(v string) *ModifyRCElasticScalingRequest
	GetRegionId() *string
	SetScalingEnabled(v bool) *ModifyRCElasticScalingRequest
	GetScalingEnabled() *bool
	SetScheduledRule(v string) *ModifyRCElasticScalingRequest
	GetScheduledRule() *string
}

type ModifyRCElasticScalingRequest struct {
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// mysql.x2.medium.9cm
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// mysql.x2.medium.9cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cn-hanghzou
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingEnabled *bool   `json:"ScalingEnabled,omitempty" xml:"ScalingEnabled,omitempty"`
	// example:
	//
	// {"rule":[{"beginTime":"09:00","endTime":"17:00","acu":4}]}
	ScheduledRule *string `json:"ScheduledRule,omitempty" xml:"ScheduledRule,omitempty"`
}

func (s ModifyRCElasticScalingRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCElasticScalingRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCElasticScalingRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRCElasticScalingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCElasticScalingRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyRCElasticScalingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCElasticScalingRequest) GetScalingEnabled() *bool {
	return s.ScalingEnabled
}

func (s *ModifyRCElasticScalingRequest) GetScheduledRule() *string {
	return s.ScheduledRule
}

func (s *ModifyRCElasticScalingRequest) SetDryRun(v bool) *ModifyRCElasticScalingRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) SetInstanceId(v string) *ModifyRCElasticScalingRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) SetInstanceType(v string) *ModifyRCElasticScalingRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) SetRegionId(v string) *ModifyRCElasticScalingRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) SetScalingEnabled(v bool) *ModifyRCElasticScalingRequest {
	s.ScalingEnabled = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) SetScheduledRule(v string) *ModifyRCElasticScalingRequest {
	s.ScheduledRule = &v
	return s
}

func (s *ModifyRCElasticScalingRequest) Validate() error {
	return dara.Validate(s)
}
