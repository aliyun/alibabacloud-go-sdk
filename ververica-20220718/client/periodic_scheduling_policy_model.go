// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPeriodicSchedulingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetIsFinished(v bool) *PeriodicSchedulingPolicy
	GetIsFinished() *bool
	SetOnlyOnceTriggerTime(v int64) *PeriodicSchedulingPolicy
	GetOnlyOnceTriggerTime() *int64
	SetOnlyOnceTriggerTimeIsExpired(v bool) *PeriodicSchedulingPolicy
	GetOnlyOnceTriggerTimeIsExpired() *bool
	SetPeriodicSchedulingLevel(v string) *PeriodicSchedulingPolicy
	GetPeriodicSchedulingLevel() *string
	SetPeriodicSchedulingValues(v []*int32) *PeriodicSchedulingPolicy
	GetPeriodicSchedulingValues() []*int32
	SetPeriodicTriggerTime(v int64) *PeriodicSchedulingPolicy
	GetPeriodicTriggerTime() *int64
	SetResourceSetting(v *BriefResourceSetting) *PeriodicSchedulingPolicy
	GetResourceSetting() *BriefResourceSetting
}

type PeriodicSchedulingPolicy struct {
	IsFinished *bool `json:"isFinished,omitempty" xml:"isFinished,omitempty"`
	// example:
	//
	// 1723195800000
	OnlyOnceTriggerTime *int64 `json:"onlyOnceTriggerTime,omitempty" xml:"onlyOnceTriggerTime,omitempty"`
	// example:
	//
	// true
	OnlyOnceTriggerTimeIsExpired *bool `json:"onlyOnceTriggerTimeIsExpired,omitempty" xml:"onlyOnceTriggerTimeIsExpired,omitempty"`
	// example:
	//
	// DAY
	PeriodicSchedulingLevel  *string  `json:"periodicSchedulingLevel,omitempty" xml:"periodicSchedulingLevel,omitempty"`
	PeriodicSchedulingValues []*int32 `json:"periodicSchedulingValues,omitempty" xml:"periodicSchedulingValues,omitempty" type:"Repeated"`
	// example:
	//
	// 1723199340000
	PeriodicTriggerTime *int64                `json:"periodicTriggerTime,omitempty" xml:"periodicTriggerTime,omitempty"`
	ResourceSetting     *BriefResourceSetting `json:"resourceSetting,omitempty" xml:"resourceSetting,omitempty"`
}

func (s PeriodicSchedulingPolicy) String() string {
	return dara.Prettify(s)
}

func (s PeriodicSchedulingPolicy) GoString() string {
	return s.String()
}

func (s *PeriodicSchedulingPolicy) GetIsFinished() *bool {
	return s.IsFinished
}

func (s *PeriodicSchedulingPolicy) GetOnlyOnceTriggerTime() *int64 {
	return s.OnlyOnceTriggerTime
}

func (s *PeriodicSchedulingPolicy) GetOnlyOnceTriggerTimeIsExpired() *bool {
	return s.OnlyOnceTriggerTimeIsExpired
}

func (s *PeriodicSchedulingPolicy) GetPeriodicSchedulingLevel() *string {
	return s.PeriodicSchedulingLevel
}

func (s *PeriodicSchedulingPolicy) GetPeriodicSchedulingValues() []*int32 {
	return s.PeriodicSchedulingValues
}

func (s *PeriodicSchedulingPolicy) GetPeriodicTriggerTime() *int64 {
	return s.PeriodicTriggerTime
}

func (s *PeriodicSchedulingPolicy) GetResourceSetting() *BriefResourceSetting {
	return s.ResourceSetting
}

func (s *PeriodicSchedulingPolicy) SetIsFinished(v bool) *PeriodicSchedulingPolicy {
	s.IsFinished = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetOnlyOnceTriggerTime(v int64) *PeriodicSchedulingPolicy {
	s.OnlyOnceTriggerTime = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetOnlyOnceTriggerTimeIsExpired(v bool) *PeriodicSchedulingPolicy {
	s.OnlyOnceTriggerTimeIsExpired = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicSchedulingLevel(v string) *PeriodicSchedulingPolicy {
	s.PeriodicSchedulingLevel = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicSchedulingValues(v []*int32) *PeriodicSchedulingPolicy {
	s.PeriodicSchedulingValues = v
	return s
}

func (s *PeriodicSchedulingPolicy) SetPeriodicTriggerTime(v int64) *PeriodicSchedulingPolicy {
	s.PeriodicTriggerTime = &v
	return s
}

func (s *PeriodicSchedulingPolicy) SetResourceSetting(v *BriefResourceSetting) *PeriodicSchedulingPolicy {
	s.ResourceSetting = v
	return s
}

func (s *PeriodicSchedulingPolicy) Validate() error {
	if s.ResourceSetting != nil {
		if err := s.ResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
