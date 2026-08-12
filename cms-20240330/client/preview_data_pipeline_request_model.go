// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDataPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PreviewDataPipelineRequest
	GetDescription() *string
	SetFromTime(v int64) *PreviewDataPipelineRequest
	GetFromTime() *int64
	SetKind(v string) *PreviewDataPipelineRequest
	GetKind() *string
	SetOutputs(v []*PreviewDataPipelineRequestOutputs) *PreviewDataPipelineRequest
	GetOutputs() []*PreviewDataPipelineRequestOutputs
	SetPipelineName(v string) *PreviewDataPipelineRequest
	GetPipelineName() *string
	SetProcessors(v []*PreviewDataPipelineRequestProcessors) *PreviewDataPipelineRequest
	GetProcessors() []*PreviewDataPipelineRequestProcessors
	SetSinks(v []*PreviewDataPipelineRequestSinks) *PreviewDataPipelineRequest
	GetSinks() []*PreviewDataPipelineRequestSinks
	SetSource(v *PreviewDataPipelineRequestSource) *PreviewDataPipelineRequest
	GetSource() *PreviewDataPipelineRequestSource
	SetToTime(v int64) *PreviewDataPipelineRequest
	GetToTime() *int64
}

type PreviewDataPipelineRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// Preview error span routing.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The start time of the preview.
	//
	// example:
	//
	// 1784563200
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The pipeline type.
	//
	// example:
	//
	// custom
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The named outputs.
	Outputs []*PreviewDataPipelineRequestOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The pipeline name.
	//
	// example:
	//
	// trace-archive-routing
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The common processors.
	Processors []*PreviewDataPipelineRequestProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The output destinations.
	Sinks []*PreviewDataPipelineRequestSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	Source *PreviewDataPipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The end time of the preview.
	//
	// example:
	//
	// 1784566800
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s PreviewDataPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequest) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *PreviewDataPipelineRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *PreviewDataPipelineRequest) GetKind() *string {
	return s.Kind
}

func (s *PreviewDataPipelineRequest) GetOutputs() []*PreviewDataPipelineRequestOutputs {
	return s.Outputs
}

func (s *PreviewDataPipelineRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *PreviewDataPipelineRequest) GetProcessors() []*PreviewDataPipelineRequestProcessors {
	return s.Processors
}

func (s *PreviewDataPipelineRequest) GetSinks() []*PreviewDataPipelineRequestSinks {
	return s.Sinks
}

func (s *PreviewDataPipelineRequest) GetSource() *PreviewDataPipelineRequestSource {
	return s.Source
}

func (s *PreviewDataPipelineRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *PreviewDataPipelineRequest) SetDescription(v string) *PreviewDataPipelineRequest {
	s.Description = &v
	return s
}

func (s *PreviewDataPipelineRequest) SetFromTime(v int64) *PreviewDataPipelineRequest {
	s.FromTime = &v
	return s
}

func (s *PreviewDataPipelineRequest) SetKind(v string) *PreviewDataPipelineRequest {
	s.Kind = &v
	return s
}

func (s *PreviewDataPipelineRequest) SetOutputs(v []*PreviewDataPipelineRequestOutputs) *PreviewDataPipelineRequest {
	s.Outputs = v
	return s
}

func (s *PreviewDataPipelineRequest) SetPipelineName(v string) *PreviewDataPipelineRequest {
	s.PipelineName = &v
	return s
}

func (s *PreviewDataPipelineRequest) SetProcessors(v []*PreviewDataPipelineRequestProcessors) *PreviewDataPipelineRequest {
	s.Processors = v
	return s
}

func (s *PreviewDataPipelineRequest) SetSinks(v []*PreviewDataPipelineRequestSinks) *PreviewDataPipelineRequest {
	s.Sinks = v
	return s
}

func (s *PreviewDataPipelineRequest) SetSource(v *PreviewDataPipelineRequestSource) *PreviewDataPipelineRequest {
	s.Source = v
	return s
}

func (s *PreviewDataPipelineRequest) SetToTime(v int64) *PreviewDataPipelineRequest {
	s.ToTime = &v
	return s
}

func (s *PreviewDataPipelineRequest) Validate() error {
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Processors != nil {
		for _, item := range s.Processors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Sinks != nil {
		for _, item := range s.Sinks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*PreviewDataPipelineRequestOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s PreviewDataPipelineRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputs) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputs) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineRequestOutputs) GetProcessors() []*PreviewDataPipelineRequestOutputsProcessors {
	return s.Processors
}

func (s *PreviewDataPipelineRequestOutputs) SetName(v string) *PreviewDataPipelineRequestOutputs {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputs) SetProcessors(v []*PreviewDataPipelineRequestOutputsProcessors) *PreviewDataPipelineRequestOutputs {
	s.Processors = v
	return s
}

func (s *PreviewDataPipelineRequestOutputs) Validate() error {
	if s.Processors != nil {
		for _, item := range s.Processors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewDataPipelineRequestOutputsProcessors struct {
	// The processor configuration.
	Config *PreviewDataPipelineRequestOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The processor name.
	//
	// example:
	//
	// drop-health-check
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processor type.
	//
	// example:
	//
	// filter
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewDataPipelineRequestOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputsProcessors) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputsProcessors) GetConfig() *PreviewDataPipelineRequestOutputsProcessorsConfig {
	return s.Config
}

func (s *PreviewDataPipelineRequestOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineRequestOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *PreviewDataPipelineRequestOutputsProcessors) SetConfig(v *PreviewDataPipelineRequestOutputsProcessorsConfig) *PreviewDataPipelineRequestOutputsProcessors {
	s.Config = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessors) SetName(v string) *PreviewDataPipelineRequestOutputsProcessors {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessors) SetType(v string) *PreviewDataPipelineRequestOutputsProcessors {
	s.Type = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestOutputsProcessorsConfig struct {
	// The application list.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*PreviewDataPipelineRequestOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *PreviewDataPipelineRequestOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *PreviewDataPipelineRequestOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetRules() []*PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetSelector() *PreviewDataPipelineRequestOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) GetTarget() *PreviewDataPipelineRequestOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetApplications(v []*string) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetExpression(v string) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetFields(v []*string) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetRules(v []*PreviewDataPipelineRequestOutputsProcessorsConfigRules) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetScript(v string) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetSelector(v *PreviewDataPipelineRequestOutputsProcessorsConfigSelector) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) SetTarget(v *PreviewDataPipelineRequestOutputsProcessorsConfigTarget) *PreviewDataPipelineRequestOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfig) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Selector != nil {
		if err := s.Selector.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestOutputsProcessorsConfigRules struct {
	// The prefix length to retain.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The suffix length to retain.
	//
	// example:
	//
	// 2
	KeepSuffix *int32 `json:"keepSuffix,omitempty" xml:"keepSuffix,omitempty"`
	// The sensitive keywords.
	Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	// The mask character.
	//
	// example:
	//
	// *
	MaskChar *string `json:"maskChar,omitempty" xml:"maskChar,omitempty"`
	// The masking mode.
	//
	// example:
	//
	// keyword
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The built-in sensitive types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetKeys(v []*string) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetMaskChar(v string) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetMode(v string) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) SetTypes(v []*string) *PreviewDataPipelineRequestOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestOutputsProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *PreviewDataPipelineRequestOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigTarget) SetWorkspace(v string) *PreviewDataPipelineRequestOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *PreviewDataPipelineRequestOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestProcessors struct {
	// The processor configuration.
	Config *PreviewDataPipelineRequestProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The processor name.
	//
	// example:
	//
	// drop-health-check
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processor type.
	//
	// example:
	//
	// filter
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewDataPipelineRequestProcessors) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestProcessors) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestProcessors) GetConfig() *PreviewDataPipelineRequestProcessorsConfig {
	return s.Config
}

func (s *PreviewDataPipelineRequestProcessors) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineRequestProcessors) GetType() *string {
	return s.Type
}

func (s *PreviewDataPipelineRequestProcessors) SetConfig(v *PreviewDataPipelineRequestProcessorsConfig) *PreviewDataPipelineRequestProcessors {
	s.Config = v
	return s
}

func (s *PreviewDataPipelineRequestProcessors) SetName(v string) *PreviewDataPipelineRequestProcessors {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessors) SetType(v string) *PreviewDataPipelineRequestProcessors {
	s.Type = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestProcessorsConfig struct {
	// The application list.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*PreviewDataPipelineRequestProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *PreviewDataPipelineRequestProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *PreviewDataPipelineRequestProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s PreviewDataPipelineRequestProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestProcessorsConfig) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetRules() []*PreviewDataPipelineRequestProcessorsConfigRules {
	return s.Rules
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetSelector() *PreviewDataPipelineRequestProcessorsConfigSelector {
	return s.Selector
}

func (s *PreviewDataPipelineRequestProcessorsConfig) GetTarget() *PreviewDataPipelineRequestProcessorsConfigTarget {
	return s.Target
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetApplications(v []*string) *PreviewDataPipelineRequestProcessorsConfig {
	s.Applications = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetExpression(v string) *PreviewDataPipelineRequestProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetFields(v []*string) *PreviewDataPipelineRequestProcessorsConfig {
	s.Fields = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetRules(v []*PreviewDataPipelineRequestProcessorsConfigRules) *PreviewDataPipelineRequestProcessorsConfig {
	s.Rules = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetScript(v string) *PreviewDataPipelineRequestProcessorsConfig {
	s.Script = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetSelector(v *PreviewDataPipelineRequestProcessorsConfigSelector) *PreviewDataPipelineRequestProcessorsConfig {
	s.Selector = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) SetTarget(v *PreviewDataPipelineRequestProcessorsConfigTarget) *PreviewDataPipelineRequestProcessorsConfig {
	s.Target = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfig) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Selector != nil {
		if err := s.Selector.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestProcessorsConfigRules struct {
	// The prefix length to retain.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The suffix length to retain.
	//
	// example:
	//
	// 2
	KeepSuffix *int32 `json:"keepSuffix,omitempty" xml:"keepSuffix,omitempty"`
	// The sensitive keywords.
	Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	// The mask character.
	//
	// example:
	//
	// *
	MaskChar *string `json:"maskChar,omitempty" xml:"maskChar,omitempty"`
	// The masking mode.
	//
	// example:
	//
	// keyword
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The built-in sensitive types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s PreviewDataPipelineRequestProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetKeepPrefix(v int32) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetKeepSuffix(v int32) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetKeys(v []*string) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetMaskChar(v string) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetMode(v string) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) SetTypes(v []*string) *PreviewDataPipelineRequestProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s PreviewDataPipelineRequestProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *PreviewDataPipelineRequestProcessorsConfigSelector) SetServiceNames(v []*string) *PreviewDataPipelineRequestProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s PreviewDataPipelineRequestProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *PreviewDataPipelineRequestProcessorsConfigTarget) SetWorkspace(v string) *PreviewDataPipelineRequestProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *PreviewDataPipelineRequestProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestSinks struct {
	// The list of datasets.
	Datasets []*string `json:"datasets,omitempty" xml:"datasets,omitempty" type:"Repeated"`
	// SLS Logstore
	//
	// example:
	//
	// error-spans
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the output destination.
	//
	// example:
	//
	// error-archive
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// SLS Project
	//
	// example:
	//
	// customer-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The type of the output destination.
	//
	// example:
	//
	// logstore
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewDataPipelineRequestSinks) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestSinks) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *PreviewDataPipelineRequestSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *PreviewDataPipelineRequestSinks) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineRequestSinks) GetProject() *string {
	return s.Project
}

func (s *PreviewDataPipelineRequestSinks) GetType() *string {
	return s.Type
}

func (s *PreviewDataPipelineRequestSinks) SetDatasets(v []*string) *PreviewDataPipelineRequestSinks {
	s.Datasets = v
	return s
}

func (s *PreviewDataPipelineRequestSinks) SetLogstore(v string) *PreviewDataPipelineRequestSinks {
	s.Logstore = &v
	return s
}

func (s *PreviewDataPipelineRequestSinks) SetName(v string) *PreviewDataPipelineRequestSinks {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineRequestSinks) SetProject(v string) *PreviewDataPipelineRequestSinks {
	s.Project = &v
	return s
}

func (s *PreviewDataPipelineRequestSinks) SetType(v string) *PreviewDataPipelineRequestSinks {
	s.Type = &v
	return s
}

func (s *PreviewDataPipelineRequestSinks) Validate() error {
	return dara.Validate(s)
}

type PreviewDataPipelineRequestSource struct {
	// The datasource config.
	Config *PreviewDataPipelineRequestSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The data source type.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewDataPipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestSource) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestSource) GetConfig() *PreviewDataPipelineRequestSourceConfig {
	return s.Config
}

func (s *PreviewDataPipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *PreviewDataPipelineRequestSource) SetConfig(v *PreviewDataPipelineRequestSourceConfig) *PreviewDataPipelineRequestSource {
	s.Config = v
	return s
}

func (s *PreviewDataPipelineRequestSource) SetType(v string) *PreviewDataPipelineRequestSource {
	s.Type = &v
	return s
}

func (s *PreviewDataPipelineRequestSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestSourceConfig struct {
	// The run mode.
	//
	// example:
	//
	// continuous
	RunMode *string `json:"runMode,omitempty" xml:"runMode,omitempty"`
	// The read start point.
	//
	// example:
	//
	// latest
	StartFrom *string `json:"startFrom,omitempty" xml:"startFrom,omitempty"`
	// The backfill time range.
	TimeRange *PreviewDataPipelineRequestSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s PreviewDataPipelineRequestSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *PreviewDataPipelineRequestSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *PreviewDataPipelineRequestSourceConfig) GetTimeRange() *PreviewDataPipelineRequestSourceConfigTimeRange {
	return s.TimeRange
}

func (s *PreviewDataPipelineRequestSourceConfig) SetRunMode(v string) *PreviewDataPipelineRequestSourceConfig {
	s.RunMode = &v
	return s
}

func (s *PreviewDataPipelineRequestSourceConfig) SetStartFrom(v string) *PreviewDataPipelineRequestSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *PreviewDataPipelineRequestSourceConfig) SetTimeRange(v *PreviewDataPipelineRequestSourceConfigTimeRange) *PreviewDataPipelineRequestSourceConfig {
	s.TimeRange = v
	return s
}

func (s *PreviewDataPipelineRequestSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDataPipelineRequestSourceConfigTimeRange struct {
	// The start time.
	//
	// example:
	//
	// 1722844800
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1722848400
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s PreviewDataPipelineRequestSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineRequestSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineRequestSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *PreviewDataPipelineRequestSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *PreviewDataPipelineRequestSourceConfigTimeRange) SetFrom(v int64) *PreviewDataPipelineRequestSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *PreviewDataPipelineRequestSourceConfigTimeRange) SetTo(v int64) *PreviewDataPipelineRequestSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *PreviewDataPipelineRequestSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
