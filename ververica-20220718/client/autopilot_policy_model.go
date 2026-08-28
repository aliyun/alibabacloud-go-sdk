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
	// The advanced rule configuration. This includes advanced parameters such as chain-break optimization, minimum parallelism, and TM CPU scaling. Disabled by default and must be explicitly enabled.
	AdvancedRules *AutopilotPolicyAdvancedRules `json:"advancedRules,omitempty" xml:"advancedRules,omitempty" type:"Struct"`
	// The upper and lower limits for tuning resources.
	Limits *AutopilotPolicyLimits `json:"limits,omitempty" xml:"limits,omitempty" type:"Struct"`
	// The scale-down rule configuration.
	ScaleDownRules *AutopilotPolicyScaleDownRules `json:"scaleDownRules,omitempty" xml:"scaleDownRules,omitempty" type:"Struct"`
	// The scale-up rule configuration.
	ScaleUpRules *AutopilotPolicyScaleUpRules `json:"scaleUpRules,omitempty" xml:"scaleUpRules,omitempty" type:"Struct"`
	// The silent period configuration. Automatic tuning operations are not performed during silent periods.
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
	// Specifies whether to enable advanced rules.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The advanced rule parameters. An empty map indicates that internal default parameters are used. You can override specific internal parameters by using key-value pairs. The entire map is replaced.
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
	// The minimum cool-down time between two tuning operations, in minutes.
	//
	// example:
	//
	// 10
	CoolDownMinutes *int64 `json:"coolDownMinutes,omitempty" xml:"coolDownMinutes,omitempty"`
	// The maximum CPU.
	//
	// example:
	//
	// 16
	JobMaxCpu *float64 `json:"jobMaxCpu,omitempty" xml:"jobMaxCpu,omitempty"`
	// The maximum memory. Format examples: 4Gi, 256GiB.
	//
	// example:
	//
	// 64GiB
	JobMaxMemory *string `json:"jobMaxMemory,omitempty" xml:"jobMaxMemory,omitempty"`
	// The maximum parallelism.
	//
	// example:
	//
	// 10
	JobMaxParallelism *int32 `json:"jobMaxParallelism,omitempty" xml:"jobMaxParallelism,omitempty"`
	// The minimum parallelism.
	//
	// example:
	//
	// 1
	JobMinParallelism *int32 `json:"jobMinParallelism,omitempty" xml:"jobMinParallelism,omitempty"`
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
	// The memory scale-down rule. Scale-down is triggered when memory usage falls below the threshold.
	MemoryScaleDownRule *AutopilotPolicyScaleDownRulesMemoryScaleDownRule `json:"memoryScaleDownRule,omitempty" xml:"memoryScaleDownRule,omitempty" type:"Struct"`
	// The slot idle scale-down rule. Scale-down is triggered when the slot busy ratio falls below the threshold.
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
	// Specifies whether to enable memory scale-down.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The memory scale-down sampling interval. Format examples: 4h, 5m.
	//
	// example:
	//
	// 25h
	MemUsageScaleDownSampleInterval *string `json:"memUsageScaleDownSampleInterval,omitempty" xml:"memUsageScaleDownSampleInterval,omitempty"`
	// The memory scale-down threshold. Valid values: 0.0 to 1.0. Scale-down is triggered when memory usage falls below this value. This value must be less than the scale-up threshold.
	//
	// example:
	//
	// 0.3
	MemUsageScaleDownThreshold *float64 `json:"memUsageScaleDownThreshold,omitempty" xml:"memUsageScaleDownThreshold,omitempty"`
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
	// Specifies whether to enable slot idle scale-down.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The slot idle sampling interval. Format examples: 4h, 5m.
	//
	// example:
	//
	// 24h
	SlotBusyScaleDownSampleInterval *string `json:"slotBusyScaleDownSampleInterval,omitempty" xml:"slotBusyScaleDownSampleInterval,omitempty"`
	// The slot idle scale-down threshold. Valid values: 0.0 to 1.0. Scale-down is triggered when the slot busy ratio falls below this value. This value must be less than the scale-up threshold.
	//
	// example:
	//
	// 0.2
	SlotBusyScaleDownThreshold *float64 `json:"slotBusyScaleDownThreshold,omitempty" xml:"slotBusyScaleDownThreshold,omitempty"`
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
	// The delay detection scale-up rule. Scale-up is triggered when the job delay exceeds the threshold.
	DelayRule *AutopilotPolicyScaleUpRulesDelayRule `json:"delayRule,omitempty" xml:"delayRule,omitempty" type:"Struct"`
	// The GC tuning rule. Scale-up is triggered when the GC time ratio exceeds the threshold.
	GcRule *AutopilotPolicyScaleUpRulesGcRule `json:"gcRule,omitempty" xml:"gcRule,omitempty" type:"Struct"`
	// The memory scale-up rule. Scale-up is triggered when memory usage exceeds the threshold.
	MemoryScaleUpRule *AutopilotPolicyScaleUpRulesMemoryScaleUpRule `json:"memoryScaleUpRule,omitempty" xml:"memoryScaleUpRule,omitempty" type:"Struct"`
	// The OOM scale-up rule. Scale-up is triggered when an OOM risk is detected.
	OomScaleUpRule *AutopilotPolicyScaleUpRulesOomScaleUpRule `json:"oomScaleUpRule,omitempty" xml:"oomScaleUpRule,omitempty" type:"Struct"`
	// The slot busy scale-up rule. Scale-up is triggered when the slot busy ratio exceeds the threshold.
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
	// The delay sampling interval. Format examples: 3min, 5m, 1h.
	//
	// example:
	//
	// 3min
	DelaySampleInterval *string `json:"delaySampleInterval,omitempty" xml:"delaySampleInterval,omitempty"`
	// The latency threshold. Format examples: 1min, 10m. Scale-up is triggered when the delay continuously exceeds this threshold.
	//
	// example:
	//
	// 1min
	DelayThreshold *string `json:"delayThreshold,omitempty" xml:"delayThreshold,omitempty"`
	// Specifies whether to enable delay detection scale-up.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
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
	// Specifies whether to enable GC tuning.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The GC sampling interval. Format examples: 3min, 5m.
	//
	// example:
	//
	// 3min
	GcSampleInterval *string `json:"gcSampleInterval,omitempty" xml:"gcSampleInterval,omitempty"`
	// The GC time ratio threshold. Valid values: 0.0 to 1.0. Scale-up is triggered when the GC time ratio exceeds this value.
	//
	// example:
	//
	// 0.2
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
	// Specifies whether to enable memory scale-up.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The memory scale-up threshold. Valid values: 0.0 to 1.0. Scale-up is triggered when memory usage exceeds this value.
	//
	// example:
	//
	// 0.95
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
	// Specifies whether to enable OOM scale-up.
	//
	// example:
	//
	// true
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
	// Specifies whether to enable slot busy scale-up.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The slot busy sampling interval. Format examples: 6min, 5m.
	//
	// example:
	//
	// 6min
	SlotBusyScaleUpSampleInterval *string `json:"slotBusyScaleUpSampleInterval,omitempty" xml:"slotBusyScaleUpSampleInterval,omitempty"`
	// The slot busy scale-up threshold. Valid values: 0.0 to 1.0. Scale-up is triggered when the slot busy ratio exceeds this value.
	//
	// example:
	//
	// 0.8
	SlotBusyScaleUpThreshold *float64 `json:"slotBusyScaleUpThreshold,omitempty" xml:"slotBusyScaleUpThreshold,omitempty"`
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
	// Specifies whether to enable silent periods.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of silent periods. This is a full replacement, not an append operation.
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
	// The start time. For the DAY level: 0-1439, representing the minute offset of the day (for example, 540 represents 9:00). For the WEEK level: 1-7, representing the day of the week (ISO 8601, 1=Monday, 7=Sunday).
	//
	// example:
	//
	// 540
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end time. The format is the same as beginTime. For the WEEK level, if endTime is less than beginTime, it indicates a cross-week period (for example, beginTime=6, endTime=2 means silent from Saturday to the following Tuesday).
	//
	// example:
	//
	// 1080
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The silent level. DAY indicates daily repetition. WEEK indicates weekly repetition.
	//
	// example:
	//
	// DAY
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
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
