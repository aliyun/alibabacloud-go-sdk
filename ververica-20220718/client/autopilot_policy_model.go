// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutopilotPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedRules(v *AutopilotPolicyAdvancedRules) *AutopilotPolicy
	GetAdvancedRules() *AutopilotPolicyAdvancedRules
	SetLimits(v *AutopilotPolicyLimits) *AutopilotPolicy
	GetLimits() *AutopilotPolicyLimits
	SetScaleDownRules(v *AutopilotPolicyScaleDownRules) *AutopilotPolicy
	GetScaleDownRules() *AutopilotPolicyScaleDownRules
	SetScaleUpRules(v *AutopilotPolicyScaleUpRules) *AutopilotPolicy
	GetScaleUpRules() *AutopilotPolicyScaleUpRules
	SetSilentPeriodConfig(v *AutopilotPolicySilentPeriodConfig) *AutopilotPolicy
	GetSilentPeriodConfig() *AutopilotPolicySilentPeriodConfig
}

type AutopilotPolicy struct {
	AdvancedRules      *AutopilotPolicyAdvancedRules      `json:"advancedRules,omitempty" xml:"advancedRules,omitempty" type:"Struct"`
	Limits             *AutopilotPolicyLimits             `json:"limits,omitempty" xml:"limits,omitempty" type:"Struct"`
	ScaleDownRules     *AutopilotPolicyScaleDownRules     `json:"scaleDownRules,omitempty" xml:"scaleDownRules,omitempty" type:"Struct"`
	ScaleUpRules       *AutopilotPolicyScaleUpRules       `json:"scaleUpRules,omitempty" xml:"scaleUpRules,omitempty" type:"Struct"`
	SilentPeriodConfig *AutopilotPolicySilentPeriodConfig `json:"silentPeriodConfig,omitempty" xml:"silentPeriodConfig,omitempty" type:"Struct"`
}

func (s AutopilotPolicy) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicy) GoString() string {
	return s.String()
}

func (s *AutopilotPolicy) GetAdvancedRules() *AutopilotPolicyAdvancedRules {
	return s.AdvancedRules
}

func (s *AutopilotPolicy) GetLimits() *AutopilotPolicyLimits {
	return s.Limits
}

func (s *AutopilotPolicy) GetScaleDownRules() *AutopilotPolicyScaleDownRules {
	return s.ScaleDownRules
}

func (s *AutopilotPolicy) GetScaleUpRules() *AutopilotPolicyScaleUpRules {
	return s.ScaleUpRules
}

func (s *AutopilotPolicy) GetSilentPeriodConfig() *AutopilotPolicySilentPeriodConfig {
	return s.SilentPeriodConfig
}

func (s *AutopilotPolicy) SetAdvancedRules(v *AutopilotPolicyAdvancedRules) *AutopilotPolicy {
	s.AdvancedRules = v
	return s
}

func (s *AutopilotPolicy) SetLimits(v *AutopilotPolicyLimits) *AutopilotPolicy {
	s.Limits = v
	return s
}

func (s *AutopilotPolicy) SetScaleDownRules(v *AutopilotPolicyScaleDownRules) *AutopilotPolicy {
	s.ScaleDownRules = v
	return s
}

func (s *AutopilotPolicy) SetScaleUpRules(v *AutopilotPolicyScaleUpRules) *AutopilotPolicy {
	s.ScaleUpRules = v
	return s
}

func (s *AutopilotPolicy) SetSilentPeriodConfig(v *AutopilotPolicySilentPeriodConfig) *AutopilotPolicy {
	s.SilentPeriodConfig = v
	return s
}

func (s *AutopilotPolicy) Validate() error {
	if s.AdvancedRules != nil {
		if err := s.AdvancedRules.Validate(); err != nil {
			return err
		}
	}
	if s.Limits != nil {
		if err := s.Limits.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleDownRules != nil {
		if err := s.ScaleDownRules.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleUpRules != nil {
		if err := s.ScaleUpRules.Validate(); err != nil {
			return err
		}
	}
	if s.SilentPeriodConfig != nil {
		if err := s.SilentPeriodConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AutopilotPolicyAdvancedRules struct {
	Enabled    *bool              `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s AutopilotPolicyAdvancedRules) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyAdvancedRules) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyAdvancedRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyAdvancedRules) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *AutopilotPolicyAdvancedRules) SetEnabled(v bool) *AutopilotPolicyAdvancedRules {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyAdvancedRules) SetParameters(v map[string]*string) *AutopilotPolicyAdvancedRules {
	s.Parameters = v
	return s
}

func (s *AutopilotPolicyAdvancedRules) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyLimits struct {
	CoolDownMinutes   *int64   `json:"coolDownMinutes,omitempty" xml:"coolDownMinutes,omitempty"`
	JobMaxCpu         *float64 `json:"jobMaxCpu,omitempty" xml:"jobMaxCpu,omitempty"`
	JobMaxMemory      *string  `json:"jobMaxMemory,omitempty" xml:"jobMaxMemory,omitempty"`
	JobMaxParallelism *int32   `json:"jobMaxParallelism,omitempty" xml:"jobMaxParallelism,omitempty"`
	JobMinParallelism *int32   `json:"jobMinParallelism,omitempty" xml:"jobMinParallelism,omitempty"`
}

func (s AutopilotPolicyLimits) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyLimits) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyLimits) GetCoolDownMinutes() *int64 {
	return s.CoolDownMinutes
}

func (s *AutopilotPolicyLimits) GetJobMaxCpu() *float64 {
	return s.JobMaxCpu
}

func (s *AutopilotPolicyLimits) GetJobMaxMemory() *string {
	return s.JobMaxMemory
}

func (s *AutopilotPolicyLimits) GetJobMaxParallelism() *int32 {
	return s.JobMaxParallelism
}

func (s *AutopilotPolicyLimits) GetJobMinParallelism() *int32 {
	return s.JobMinParallelism
}

func (s *AutopilotPolicyLimits) SetCoolDownMinutes(v int64) *AutopilotPolicyLimits {
	s.CoolDownMinutes = &v
	return s
}

func (s *AutopilotPolicyLimits) SetJobMaxCpu(v float64) *AutopilotPolicyLimits {
	s.JobMaxCpu = &v
	return s
}

func (s *AutopilotPolicyLimits) SetJobMaxMemory(v string) *AutopilotPolicyLimits {
	s.JobMaxMemory = &v
	return s
}

func (s *AutopilotPolicyLimits) SetJobMaxParallelism(v int32) *AutopilotPolicyLimits {
	s.JobMaxParallelism = &v
	return s
}

func (s *AutopilotPolicyLimits) SetJobMinParallelism(v int32) *AutopilotPolicyLimits {
	s.JobMinParallelism = &v
	return s
}

func (s *AutopilotPolicyLimits) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleDownRules struct {
	MemoryScaleDownRule   *AutopilotPolicyScaleDownRulesMemoryScaleDownRule   `json:"memoryScaleDownRule,omitempty" xml:"memoryScaleDownRule,omitempty" type:"Struct"`
	SlotBusyScaleDownRule *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule `json:"slotBusyScaleDownRule,omitempty" xml:"slotBusyScaleDownRule,omitempty" type:"Struct"`
}

func (s AutopilotPolicyScaleDownRules) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleDownRules) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleDownRules) GetMemoryScaleDownRule() *AutopilotPolicyScaleDownRulesMemoryScaleDownRule {
	return s.MemoryScaleDownRule
}

func (s *AutopilotPolicyScaleDownRules) GetSlotBusyScaleDownRule() *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule {
	return s.SlotBusyScaleDownRule
}

func (s *AutopilotPolicyScaleDownRules) SetMemoryScaleDownRule(v *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) *AutopilotPolicyScaleDownRules {
	s.MemoryScaleDownRule = v
	return s
}

func (s *AutopilotPolicyScaleDownRules) SetSlotBusyScaleDownRule(v *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) *AutopilotPolicyScaleDownRules {
	s.SlotBusyScaleDownRule = v
	return s
}

func (s *AutopilotPolicyScaleDownRules) Validate() error {
	if s.MemoryScaleDownRule != nil {
		if err := s.MemoryScaleDownRule.Validate(); err != nil {
			return err
		}
	}
	if s.SlotBusyScaleDownRule != nil {
		if err := s.SlotBusyScaleDownRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AutopilotPolicyScaleDownRulesMemoryScaleDownRule struct {
	Enabled                         *bool    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	MemUsageScaleDownSampleInterval *string  `json:"memUsageScaleDownSampleInterval,omitempty" xml:"memUsageScaleDownSampleInterval,omitempty"`
	MemUsageScaleDownThreshold      *float64 `json:"memUsageScaleDownThreshold,omitempty" xml:"memUsageScaleDownThreshold,omitempty"`
}

func (s AutopilotPolicyScaleDownRulesMemoryScaleDownRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleDownRulesMemoryScaleDownRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) GetMemUsageScaleDownSampleInterval() *string {
	return s.MemUsageScaleDownSampleInterval
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) GetMemUsageScaleDownThreshold() *float64 {
	return s.MemUsageScaleDownThreshold
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) SetEnabled(v bool) *AutopilotPolicyScaleDownRulesMemoryScaleDownRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) SetMemUsageScaleDownSampleInterval(v string) *AutopilotPolicyScaleDownRulesMemoryScaleDownRule {
	s.MemUsageScaleDownSampleInterval = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) SetMemUsageScaleDownThreshold(v float64) *AutopilotPolicyScaleDownRulesMemoryScaleDownRule {
	s.MemUsageScaleDownThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesMemoryScaleDownRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule struct {
	Enabled                         *bool    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	SlotBusyScaleDownSampleInterval *string  `json:"slotBusyScaleDownSampleInterval,omitempty" xml:"slotBusyScaleDownSampleInterval,omitempty"`
	SlotBusyScaleDownThreshold      *float64 `json:"slotBusyScaleDownThreshold,omitempty" xml:"slotBusyScaleDownThreshold,omitempty"`
}

func (s AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) GetSlotBusyScaleDownSampleInterval() *string {
	return s.SlotBusyScaleDownSampleInterval
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) GetSlotBusyScaleDownThreshold() *float64 {
	return s.SlotBusyScaleDownThreshold
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) SetEnabled(v bool) *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) SetSlotBusyScaleDownSampleInterval(v string) *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule {
	s.SlotBusyScaleDownSampleInterval = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) SetSlotBusyScaleDownThreshold(v float64) *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule {
	s.SlotBusyScaleDownThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleDownRulesSlotBusyScaleDownRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleUpRules struct {
	DelayRule           *AutopilotPolicyScaleUpRulesDelayRule           `json:"delayRule,omitempty" xml:"delayRule,omitempty" type:"Struct"`
	GcRule              *AutopilotPolicyScaleUpRulesGcRule              `json:"gcRule,omitempty" xml:"gcRule,omitempty" type:"Struct"`
	MemoryScaleUpRule   *AutopilotPolicyScaleUpRulesMemoryScaleUpRule   `json:"memoryScaleUpRule,omitempty" xml:"memoryScaleUpRule,omitempty" type:"Struct"`
	OomScaleUpRule      *AutopilotPolicyScaleUpRulesOomScaleUpRule      `json:"oomScaleUpRule,omitempty" xml:"oomScaleUpRule,omitempty" type:"Struct"`
	SlotBusyScaleUpRule *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule `json:"slotBusyScaleUpRule,omitempty" xml:"slotBusyScaleUpRule,omitempty" type:"Struct"`
}

func (s AutopilotPolicyScaleUpRules) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRules) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRules) GetDelayRule() *AutopilotPolicyScaleUpRulesDelayRule {
	return s.DelayRule
}

func (s *AutopilotPolicyScaleUpRules) GetGcRule() *AutopilotPolicyScaleUpRulesGcRule {
	return s.GcRule
}

func (s *AutopilotPolicyScaleUpRules) GetMemoryScaleUpRule() *AutopilotPolicyScaleUpRulesMemoryScaleUpRule {
	return s.MemoryScaleUpRule
}

func (s *AutopilotPolicyScaleUpRules) GetOomScaleUpRule() *AutopilotPolicyScaleUpRulesOomScaleUpRule {
	return s.OomScaleUpRule
}

func (s *AutopilotPolicyScaleUpRules) GetSlotBusyScaleUpRule() *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule {
	return s.SlotBusyScaleUpRule
}

func (s *AutopilotPolicyScaleUpRules) SetDelayRule(v *AutopilotPolicyScaleUpRulesDelayRule) *AutopilotPolicyScaleUpRules {
	s.DelayRule = v
	return s
}

func (s *AutopilotPolicyScaleUpRules) SetGcRule(v *AutopilotPolicyScaleUpRulesGcRule) *AutopilotPolicyScaleUpRules {
	s.GcRule = v
	return s
}

func (s *AutopilotPolicyScaleUpRules) SetMemoryScaleUpRule(v *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) *AutopilotPolicyScaleUpRules {
	s.MemoryScaleUpRule = v
	return s
}

func (s *AutopilotPolicyScaleUpRules) SetOomScaleUpRule(v *AutopilotPolicyScaleUpRulesOomScaleUpRule) *AutopilotPolicyScaleUpRules {
	s.OomScaleUpRule = v
	return s
}

func (s *AutopilotPolicyScaleUpRules) SetSlotBusyScaleUpRule(v *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) *AutopilotPolicyScaleUpRules {
	s.SlotBusyScaleUpRule = v
	return s
}

func (s *AutopilotPolicyScaleUpRules) Validate() error {
	if s.DelayRule != nil {
		if err := s.DelayRule.Validate(); err != nil {
			return err
		}
	}
	if s.GcRule != nil {
		if err := s.GcRule.Validate(); err != nil {
			return err
		}
	}
	if s.MemoryScaleUpRule != nil {
		if err := s.MemoryScaleUpRule.Validate(); err != nil {
			return err
		}
	}
	if s.OomScaleUpRule != nil {
		if err := s.OomScaleUpRule.Validate(); err != nil {
			return err
		}
	}
	if s.SlotBusyScaleUpRule != nil {
		if err := s.SlotBusyScaleUpRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AutopilotPolicyScaleUpRulesDelayRule struct {
	DelaySampleInterval *string `json:"delaySampleInterval,omitempty" xml:"delaySampleInterval,omitempty"`
	DelayThreshold      *string `json:"delayThreshold,omitempty" xml:"delayThreshold,omitempty"`
	Enabled             *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s AutopilotPolicyScaleUpRulesDelayRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRulesDelayRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) GetDelaySampleInterval() *string {
	return s.DelaySampleInterval
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) GetDelayThreshold() *string {
	return s.DelayThreshold
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) SetDelaySampleInterval(v string) *AutopilotPolicyScaleUpRulesDelayRule {
	s.DelaySampleInterval = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) SetDelayThreshold(v string) *AutopilotPolicyScaleUpRulesDelayRule {
	s.DelayThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) SetEnabled(v bool) *AutopilotPolicyScaleUpRulesDelayRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesDelayRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleUpRulesGcRule struct {
	Enabled              *bool    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	GcSampleInterval     *string  `json:"gcSampleInterval,omitempty" xml:"gcSampleInterval,omitempty"`
	GcTimeRatioThreshold *float64 `json:"gcTimeRatioThreshold,omitempty" xml:"gcTimeRatioThreshold,omitempty"`
}

func (s AutopilotPolicyScaleUpRulesGcRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRulesGcRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRulesGcRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleUpRulesGcRule) GetGcSampleInterval() *string {
	return s.GcSampleInterval
}

func (s *AutopilotPolicyScaleUpRulesGcRule) GetGcTimeRatioThreshold() *float64 {
	return s.GcTimeRatioThreshold
}

func (s *AutopilotPolicyScaleUpRulesGcRule) SetEnabled(v bool) *AutopilotPolicyScaleUpRulesGcRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesGcRule) SetGcSampleInterval(v string) *AutopilotPolicyScaleUpRulesGcRule {
	s.GcSampleInterval = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesGcRule) SetGcTimeRatioThreshold(v float64) *AutopilotPolicyScaleUpRulesGcRule {
	s.GcTimeRatioThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesGcRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleUpRulesMemoryScaleUpRule struct {
	Enabled                  *bool    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	MemUsageScaleUpThreshold *float64 `json:"memUsageScaleUpThreshold,omitempty" xml:"memUsageScaleUpThreshold,omitempty"`
}

func (s AutopilotPolicyScaleUpRulesMemoryScaleUpRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRulesMemoryScaleUpRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) GetMemUsageScaleUpThreshold() *float64 {
	return s.MemUsageScaleUpThreshold
}

func (s *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) SetEnabled(v bool) *AutopilotPolicyScaleUpRulesMemoryScaleUpRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) SetMemUsageScaleUpThreshold(v float64) *AutopilotPolicyScaleUpRulesMemoryScaleUpRule {
	s.MemUsageScaleUpThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesMemoryScaleUpRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleUpRulesOomScaleUpRule struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s AutopilotPolicyScaleUpRulesOomScaleUpRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRulesOomScaleUpRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRulesOomScaleUpRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleUpRulesOomScaleUpRule) SetEnabled(v bool) *AutopilotPolicyScaleUpRulesOomScaleUpRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesOomScaleUpRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule struct {
	Enabled                       *bool    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	SlotBusyScaleUpSampleInterval *string  `json:"slotBusyScaleUpSampleInterval,omitempty" xml:"slotBusyScaleUpSampleInterval,omitempty"`
	SlotBusyScaleUpThreshold      *float64 `json:"slotBusyScaleUpThreshold,omitempty" xml:"slotBusyScaleUpThreshold,omitempty"`
}

func (s AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) GoString() string {
	return s.String()
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) GetSlotBusyScaleUpSampleInterval() *string {
	return s.SlotBusyScaleUpSampleInterval
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) GetSlotBusyScaleUpThreshold() *float64 {
	return s.SlotBusyScaleUpThreshold
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) SetEnabled(v bool) *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) SetSlotBusyScaleUpSampleInterval(v string) *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule {
	s.SlotBusyScaleUpSampleInterval = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) SetSlotBusyScaleUpThreshold(v float64) *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule {
	s.SlotBusyScaleUpThreshold = &v
	return s
}

func (s *AutopilotPolicyScaleUpRulesSlotBusyScaleUpRule) Validate() error {
	return dara.Validate(s)
}

type AutopilotPolicySilentPeriodConfig struct {
	Enabled       *bool                                             `json:"enabled,omitempty" xml:"enabled,omitempty"`
	SilentPeriods []*AutopilotPolicySilentPeriodConfigSilentPeriods `json:"silentPeriods,omitempty" xml:"silentPeriods,omitempty" type:"Repeated"`
}

func (s AutopilotPolicySilentPeriodConfig) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicySilentPeriodConfig) GoString() string {
	return s.String()
}

func (s *AutopilotPolicySilentPeriodConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *AutopilotPolicySilentPeriodConfig) GetSilentPeriods() []*AutopilotPolicySilentPeriodConfigSilentPeriods {
	return s.SilentPeriods
}

func (s *AutopilotPolicySilentPeriodConfig) SetEnabled(v bool) *AutopilotPolicySilentPeriodConfig {
	s.Enabled = &v
	return s
}

func (s *AutopilotPolicySilentPeriodConfig) SetSilentPeriods(v []*AutopilotPolicySilentPeriodConfigSilentPeriods) *AutopilotPolicySilentPeriodConfig {
	s.SilentPeriods = v
	return s
}

func (s *AutopilotPolicySilentPeriodConfig) Validate() error {
	if s.SilentPeriods != nil {
		for _, item := range s.SilentPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AutopilotPolicySilentPeriodConfigSilentPeriods struct {
	BeginTime *int64  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s AutopilotPolicySilentPeriodConfigSilentPeriods) String() string {
	return dara.Prettify(s)
}

func (s AutopilotPolicySilentPeriodConfigSilentPeriods) GoString() string {
	return s.String()
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) GetLevel() *string {
	return s.Level
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) SetBeginTime(v int64) *AutopilotPolicySilentPeriodConfigSilentPeriods {
	s.BeginTime = &v
	return s
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) SetEndTime(v int64) *AutopilotPolicySilentPeriodConfigSilentPeriods {
	s.EndTime = &v
	return s
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) SetLevel(v string) *AutopilotPolicySilentPeriodConfigSilentPeriods {
	s.Level = &v
	return s
}

func (s *AutopilotPolicySilentPeriodConfigSilentPeriods) Validate() error {
	return dara.Validate(s)
}
