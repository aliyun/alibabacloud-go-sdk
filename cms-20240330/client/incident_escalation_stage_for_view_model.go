// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEscalationStageForView interface {
	dara.Model
	String() string
	GoString() string
	SetCycleNotifyCount(v int32) *IncidentEscalationStageForView
	GetCycleNotifyCount() *int32
	SetCycleNotifyInterval(v int32) *IncidentEscalationStageForView
	GetCycleNotifyInterval() *int32
	SetEffectTimeRange(v *EffectTimeRange) *IncidentEscalationStageForView
	GetEffectTimeRange() *EffectTimeRange
	SetIndex(v int32) *IncidentEscalationStageForView
	GetIndex() *int32
	SetNotifyChannels(v []*NotifyChannel) *IncidentEscalationStageForView
	GetNotifyChannels() []*NotifyChannel
	SetTargetIncidentState(v string) *IncidentEscalationStageForView
	GetTargetIncidentState() *string
	SetTriggerDelay(v int32) *IncidentEscalationStageForView
	GetTriggerDelay() *int32
}

type IncidentEscalationStageForView struct {
	CycleNotifyCount    *int32           `json:"cycleNotifyCount,omitempty" xml:"cycleNotifyCount,omitempty"`
	CycleNotifyInterval *int32           `json:"cycleNotifyInterval,omitempty" xml:"cycleNotifyInterval,omitempty"`
	EffectTimeRange     *EffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty"`
	// This parameter is required.
	Index               *int32           `json:"index,omitempty" xml:"index,omitempty"`
	NotifyChannels      []*NotifyChannel `json:"notifyChannels,omitempty" xml:"notifyChannels,omitempty" type:"Repeated"`
	TargetIncidentState *string          `json:"targetIncidentState,omitempty" xml:"targetIncidentState,omitempty"`
	TriggerDelay        *int32           `json:"triggerDelay,omitempty" xml:"triggerDelay,omitempty"`
}

func (s IncidentEscalationStageForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentEscalationStageForView) GoString() string {
	return s.String()
}

func (s *IncidentEscalationStageForView) GetCycleNotifyCount() *int32 {
	return s.CycleNotifyCount
}

func (s *IncidentEscalationStageForView) GetCycleNotifyInterval() *int32 {
	return s.CycleNotifyInterval
}

func (s *IncidentEscalationStageForView) GetEffectTimeRange() *EffectTimeRange {
	return s.EffectTimeRange
}

func (s *IncidentEscalationStageForView) GetIndex() *int32 {
	return s.Index
}

func (s *IncidentEscalationStageForView) GetNotifyChannels() []*NotifyChannel {
	return s.NotifyChannels
}

func (s *IncidentEscalationStageForView) GetTargetIncidentState() *string {
	return s.TargetIncidentState
}

func (s *IncidentEscalationStageForView) GetTriggerDelay() *int32 {
	return s.TriggerDelay
}

func (s *IncidentEscalationStageForView) SetCycleNotifyCount(v int32) *IncidentEscalationStageForView {
	s.CycleNotifyCount = &v
	return s
}

func (s *IncidentEscalationStageForView) SetCycleNotifyInterval(v int32) *IncidentEscalationStageForView {
	s.CycleNotifyInterval = &v
	return s
}

func (s *IncidentEscalationStageForView) SetEffectTimeRange(v *EffectTimeRange) *IncidentEscalationStageForView {
	s.EffectTimeRange = v
	return s
}

func (s *IncidentEscalationStageForView) SetIndex(v int32) *IncidentEscalationStageForView {
	s.Index = &v
	return s
}

func (s *IncidentEscalationStageForView) SetNotifyChannels(v []*NotifyChannel) *IncidentEscalationStageForView {
	s.NotifyChannels = v
	return s
}

func (s *IncidentEscalationStageForView) SetTargetIncidentState(v string) *IncidentEscalationStageForView {
	s.TargetIncidentState = &v
	return s
}

func (s *IncidentEscalationStageForView) SetTriggerDelay(v int32) *IncidentEscalationStageForView {
	s.TriggerDelay = &v
	return s
}

func (s *IncidentEscalationStageForView) Validate() error {
	if s.EffectTimeRange != nil {
		if err := s.EffectTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyChannels != nil {
		for _, item := range s.NotifyChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
