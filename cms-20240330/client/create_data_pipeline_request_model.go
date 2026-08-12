// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDataPipelineRequest
	GetDescription() *string
	SetKind(v string) *CreateDataPipelineRequest
	GetKind() *string
	SetOutputs(v []*CreateDataPipelineRequestOutputs) *CreateDataPipelineRequest
	GetOutputs() []*CreateDataPipelineRequestOutputs
	SetPipelineName(v string) *CreateDataPipelineRequest
	GetPipelineName() *string
	SetProcessors(v []*CreateDataPipelineRequestProcessors) *CreateDataPipelineRequest
	GetProcessors() []*CreateDataPipelineRequestProcessors
	SetSinks(v []*CreateDataPipelineRequestSinks) *CreateDataPipelineRequest
	GetSinks() []*CreateDataPipelineRequestSinks
	SetSource(v *CreateDataPipelineRequestSource) *CreateDataPipelineRequest
	GetSource() *CreateDataPipelineRequestSource
}

type CreateDataPipelineRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// Export selected trace services to the target workspace.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The pipeline type.
	//
	// example:
	//
	// export
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The named outputs.
	Outputs []*CreateDataPipelineRequestOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The pipeline name.
	//
	// example:
	//
	// export-traces-to-prod
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The common processors.
	Processors []*CreateDataPipelineRequestProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The output destinations.
	Sinks []*CreateDataPipelineRequestSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	Source *CreateDataPipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequest) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataPipelineRequest) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineRequest) GetOutputs() []*CreateDataPipelineRequestOutputs {
	return s.Outputs
}

func (s *CreateDataPipelineRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *CreateDataPipelineRequest) GetProcessors() []*CreateDataPipelineRequestProcessors {
	return s.Processors
}

func (s *CreateDataPipelineRequest) GetSinks() []*CreateDataPipelineRequestSinks {
	return s.Sinks
}

func (s *CreateDataPipelineRequest) GetSource() *CreateDataPipelineRequestSource {
	return s.Source
}

func (s *CreateDataPipelineRequest) SetDescription(v string) *CreateDataPipelineRequest {
	s.Description = &v
	return s
}

func (s *CreateDataPipelineRequest) SetKind(v string) *CreateDataPipelineRequest {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineRequest) SetOutputs(v []*CreateDataPipelineRequestOutputs) *CreateDataPipelineRequest {
	s.Outputs = v
	return s
}

func (s *CreateDataPipelineRequest) SetPipelineName(v string) *CreateDataPipelineRequest {
	s.PipelineName = &v
	return s
}

func (s *CreateDataPipelineRequest) SetProcessors(v []*CreateDataPipelineRequestProcessors) *CreateDataPipelineRequest {
	s.Processors = v
	return s
}

func (s *CreateDataPipelineRequest) SetSinks(v []*CreateDataPipelineRequestSinks) *CreateDataPipelineRequest {
	s.Sinks = v
	return s
}

func (s *CreateDataPipelineRequest) SetSource(v *CreateDataPipelineRequestSource) *CreateDataPipelineRequest {
	s.Source = v
	return s
}

func (s *CreateDataPipelineRequest) Validate() error {
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

type CreateDataPipelineRequestOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*CreateDataPipelineRequestOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputs) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputs) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestOutputs) GetProcessors() []*CreateDataPipelineRequestOutputsProcessors {
	return s.Processors
}

func (s *CreateDataPipelineRequestOutputs) SetName(v string) *CreateDataPipelineRequestOutputs {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestOutputs) SetProcessors(v []*CreateDataPipelineRequestOutputsProcessors) *CreateDataPipelineRequestOutputs {
	s.Processors = v
	return s
}

func (s *CreateDataPipelineRequestOutputs) Validate() error {
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

type CreateDataPipelineRequestOutputsProcessors struct {
	// The processor configuration.
	Config *CreateDataPipelineRequestOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s CreateDataPipelineRequestOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessors) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessors) GetConfig() *CreateDataPipelineRequestOutputsProcessorsConfig {
	return s.Config
}

func (s *CreateDataPipelineRequestOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineRequestOutputsProcessors) SetConfig(v *CreateDataPipelineRequestOutputsProcessorsConfig) *CreateDataPipelineRequestOutputsProcessors {
	s.Config = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessors) SetName(v string) *CreateDataPipelineRequestOutputsProcessors {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessors) SetType(v string) *CreateDataPipelineRequestOutputsProcessors {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestOutputsProcessorsConfig struct {
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
	Rules []*CreateDataPipelineRequestOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *CreateDataPipelineRequestOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *CreateDataPipelineRequestOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetRules() []*CreateDataPipelineRequestOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetSelector() *CreateDataPipelineRequestOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetTarget() *CreateDataPipelineRequestOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetApplications(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetExpression(v string) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetFields(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetRules(v []*CreateDataPipelineRequestOutputsProcessorsConfigRules) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetScript(v string) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetSelector(v *CreateDataPipelineRequestOutputsProcessorsConfigSelector) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetTarget(v *CreateDataPipelineRequestOutputsProcessorsConfigTarget) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) Validate() error {
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

type CreateDataPipelineRequestOutputsProcessorsConfigRules struct {
	// The length of the prefix to retain.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The length of the suffix to retain.
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

func (s CreateDataPipelineRequestOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetKeys(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetMaskChar(v string) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetMode(v string) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) SetTypes(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigTarget) SetWorkspace(v string) *CreateDataPipelineRequestOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessors struct {
	// The processor configuration.
	Config *CreateDataPipelineRequestProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s CreateDataPipelineRequestProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessors) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessors) GetConfig() *CreateDataPipelineRequestProcessorsConfig {
	return s.Config
}

func (s *CreateDataPipelineRequestProcessors) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestProcessors) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineRequestProcessors) SetConfig(v *CreateDataPipelineRequestProcessorsConfig) *CreateDataPipelineRequestProcessors {
	s.Config = v
	return s
}

func (s *CreateDataPipelineRequestProcessors) SetName(v string) *CreateDataPipelineRequestProcessors {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestProcessors) SetType(v string) *CreateDataPipelineRequestProcessors {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineRequestProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestProcessorsConfig struct {
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
	Rules []*CreateDataPipelineRequestProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *CreateDataPipelineRequestProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *CreateDataPipelineRequestProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequestProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetRules() []*CreateDataPipelineRequestProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetSelector() *CreateDataPipelineRequestProcessorsConfigSelector {
	return s.Selector
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetTarget() *CreateDataPipelineRequestProcessorsConfigTarget {
	return s.Target
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetApplications(v []*string) *CreateDataPipelineRequestProcessorsConfig {
	s.Applications = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetExpression(v string) *CreateDataPipelineRequestProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetFields(v []*string) *CreateDataPipelineRequestProcessorsConfig {
	s.Fields = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetRules(v []*CreateDataPipelineRequestProcessorsConfigRules) *CreateDataPipelineRequestProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetScript(v string) *CreateDataPipelineRequestProcessorsConfig {
	s.Script = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetSelector(v *CreateDataPipelineRequestProcessorsConfigSelector) *CreateDataPipelineRequestProcessorsConfig {
	s.Selector = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetTarget(v *CreateDataPipelineRequestProcessorsConfigTarget) *CreateDataPipelineRequestProcessorsConfig {
	s.Target = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) Validate() error {
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

type CreateDataPipelineRequestProcessorsConfigRules struct {
	// The length of the prefix to retain.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The length of the suffix to retain.
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

func (s CreateDataPipelineRequestProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetKeepPrefix(v int32) *CreateDataPipelineRequestProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetKeepSuffix(v int32) *CreateDataPipelineRequestProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetKeys(v []*string) *CreateDataPipelineRequestProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetMaskChar(v string) *CreateDataPipelineRequestProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetMode(v string) *CreateDataPipelineRequestProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) SetTypes(v []*string) *CreateDataPipelineRequestProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *CreateDataPipelineRequestProcessorsConfigSelector) SetServiceNames(v []*string) *CreateDataPipelineRequestProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDataPipelineRequestProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDataPipelineRequestProcessorsConfigTarget) SetWorkspace(v string) *CreateDataPipelineRequestProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestSinks struct {
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

func (s CreateDataPipelineRequestSinks) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestSinks) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *CreateDataPipelineRequestSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateDataPipelineRequestSinks) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestSinks) GetProject() *string {
	return s.Project
}

func (s *CreateDataPipelineRequestSinks) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineRequestSinks) SetDatasets(v []*string) *CreateDataPipelineRequestSinks {
	s.Datasets = v
	return s
}

func (s *CreateDataPipelineRequestSinks) SetLogstore(v string) *CreateDataPipelineRequestSinks {
	s.Logstore = &v
	return s
}

func (s *CreateDataPipelineRequestSinks) SetName(v string) *CreateDataPipelineRequestSinks {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestSinks) SetProject(v string) *CreateDataPipelineRequestSinks {
	s.Project = &v
	return s
}

func (s *CreateDataPipelineRequestSinks) SetType(v string) *CreateDataPipelineRequestSinks {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineRequestSinks) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestSource struct {
	// The datasource config.
	Config *CreateDataPipelineRequestSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The data source type.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataPipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestSource) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestSource) GetConfig() *CreateDataPipelineRequestSourceConfig {
	return s.Config
}

func (s *CreateDataPipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineRequestSource) SetConfig(v *CreateDataPipelineRequestSourceConfig) *CreateDataPipelineRequestSource {
	s.Config = v
	return s
}

func (s *CreateDataPipelineRequestSource) SetType(v string) *CreateDataPipelineRequestSource {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineRequestSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestSourceConfig struct {
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
	TimeRange *CreateDataPipelineRequestSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequestSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *CreateDataPipelineRequestSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *CreateDataPipelineRequestSourceConfig) GetTimeRange() *CreateDataPipelineRequestSourceConfigTimeRange {
	return s.TimeRange
}

func (s *CreateDataPipelineRequestSourceConfig) SetRunMode(v string) *CreateDataPipelineRequestSourceConfig {
	s.RunMode = &v
	return s
}

func (s *CreateDataPipelineRequestSourceConfig) SetStartFrom(v string) *CreateDataPipelineRequestSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *CreateDataPipelineRequestSourceConfig) SetTimeRange(v *CreateDataPipelineRequestSourceConfigTimeRange) *CreateDataPipelineRequestSourceConfig {
	s.TimeRange = v
	return s
}

func (s *CreateDataPipelineRequestSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestSourceConfigTimeRange struct {
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

func (s CreateDataPipelineRequestSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *CreateDataPipelineRequestSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *CreateDataPipelineRequestSourceConfigTimeRange) SetFrom(v int64) *CreateDataPipelineRequestSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *CreateDataPipelineRequestSourceConfigTimeRange) SetTo(v int64) *CreateDataPipelineRequestSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *CreateDataPipelineRequestSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
