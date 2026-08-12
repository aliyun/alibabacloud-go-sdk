// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDataPipelineRequest
	GetDescription() *string
	SetExpectedVersion(v int64) *UpdateDataPipelineRequest
	GetExpectedVersion() *int64
	SetKind(v string) *UpdateDataPipelineRequest
	GetKind() *string
	SetOutputs(v []*UpdateDataPipelineRequestOutputs) *UpdateDataPipelineRequest
	GetOutputs() []*UpdateDataPipelineRequestOutputs
	SetProcessors(v []*UpdateDataPipelineRequestProcessors) *UpdateDataPipelineRequest
	GetProcessors() []*UpdateDataPipelineRequestProcessors
	SetSinks(v []*UpdateDataPipelineRequestSinks) *UpdateDataPipelineRequest
	GetSinks() []*UpdateDataPipelineRequestSinks
	SetSource(v *UpdateDataPipelineRequestSource) *UpdateDataPipelineRequest
	GetSource() *UpdateDataPipelineRequestSource
}

type UpdateDataPipelineRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// Export selected trace services to the target workspace.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expected version.
	//
	// example:
	//
	// 3
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The pipeline type.
	//
	// example:
	//
	// export
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The named outputs.
	Outputs []*UpdateDataPipelineRequestOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The common processors.
	Processors []*UpdateDataPipelineRequestProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The output destinations.
	Sinks []*UpdateDataPipelineRequestSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	Source *UpdateDataPipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataPipelineRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateDataPipelineRequest) GetKind() *string {
	return s.Kind
}

func (s *UpdateDataPipelineRequest) GetOutputs() []*UpdateDataPipelineRequestOutputs {
	return s.Outputs
}

func (s *UpdateDataPipelineRequest) GetProcessors() []*UpdateDataPipelineRequestProcessors {
	return s.Processors
}

func (s *UpdateDataPipelineRequest) GetSinks() []*UpdateDataPipelineRequestSinks {
	return s.Sinks
}

func (s *UpdateDataPipelineRequest) GetSource() *UpdateDataPipelineRequestSource {
	return s.Source
}

func (s *UpdateDataPipelineRequest) SetDescription(v string) *UpdateDataPipelineRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataPipelineRequest) SetExpectedVersion(v int64) *UpdateDataPipelineRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateDataPipelineRequest) SetKind(v string) *UpdateDataPipelineRequest {
	s.Kind = &v
	return s
}

func (s *UpdateDataPipelineRequest) SetOutputs(v []*UpdateDataPipelineRequestOutputs) *UpdateDataPipelineRequest {
	s.Outputs = v
	return s
}

func (s *UpdateDataPipelineRequest) SetProcessors(v []*UpdateDataPipelineRequestProcessors) *UpdateDataPipelineRequest {
	s.Processors = v
	return s
}

func (s *UpdateDataPipelineRequest) SetSinks(v []*UpdateDataPipelineRequestSinks) *UpdateDataPipelineRequest {
	s.Sinks = v
	return s
}

func (s *UpdateDataPipelineRequest) SetSource(v *UpdateDataPipelineRequestSource) *UpdateDataPipelineRequest {
	s.Source = v
	return s
}

func (s *UpdateDataPipelineRequest) Validate() error {
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

type UpdateDataPipelineRequestOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*UpdateDataPipelineRequestOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineRequestOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputs) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputs) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineRequestOutputs) GetProcessors() []*UpdateDataPipelineRequestOutputsProcessors {
	return s.Processors
}

func (s *UpdateDataPipelineRequestOutputs) SetName(v string) *UpdateDataPipelineRequestOutputs {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputs) SetProcessors(v []*UpdateDataPipelineRequestOutputsProcessors) *UpdateDataPipelineRequestOutputs {
	s.Processors = v
	return s
}

func (s *UpdateDataPipelineRequestOutputs) Validate() error {
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

type UpdateDataPipelineRequestOutputsProcessors struct {
	// The processor configuration.
	Config *UpdateDataPipelineRequestOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s UpdateDataPipelineRequestOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputsProcessors) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputsProcessors) GetConfig() *UpdateDataPipelineRequestOutputsProcessorsConfig {
	return s.Config
}

func (s *UpdateDataPipelineRequestOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineRequestOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineRequestOutputsProcessors) SetConfig(v *UpdateDataPipelineRequestOutputsProcessorsConfig) *UpdateDataPipelineRequestOutputsProcessors {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessors) SetName(v string) *UpdateDataPipelineRequestOutputsProcessors {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessors) SetType(v string) *UpdateDataPipelineRequestOutputsProcessors {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineRequestOutputsProcessorsConfig struct {
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
	Rules []*UpdateDataPipelineRequestOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *UpdateDataPipelineRequestOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *UpdateDataPipelineRequestOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetRules() []*UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetSelector() *UpdateDataPipelineRequestOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) GetTarget() *UpdateDataPipelineRequestOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetApplications(v []*string) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetExpression(v string) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetRules(v []*UpdateDataPipelineRequestOutputsProcessorsConfigRules) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetScript(v string) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetSelector(v *UpdateDataPipelineRequestOutputsProcessorsConfigSelector) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) SetTarget(v *UpdateDataPipelineRequestOutputsProcessorsConfigTarget) *UpdateDataPipelineRequestOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfig) Validate() error {
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

type UpdateDataPipelineRequestOutputsProcessorsConfigRules struct {
	// The retained prefix length.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The retained suffix length.
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

func (s UpdateDataPipelineRequestOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetKeys(v []*string) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetMaskChar(v string) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetMode(v string) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) SetTypes(v []*string) *UpdateDataPipelineRequestOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestOutputsProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *UpdateDataPipelineRequestOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigTarget) SetWorkspace(v string) *UpdateDataPipelineRequestOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *UpdateDataPipelineRequestOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestProcessors struct {
	// The processor configuration.
	Config *UpdateDataPipelineRequestProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s UpdateDataPipelineRequestProcessors) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestProcessors) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestProcessors) GetConfig() *UpdateDataPipelineRequestProcessorsConfig {
	return s.Config
}

func (s *UpdateDataPipelineRequestProcessors) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineRequestProcessors) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineRequestProcessors) SetConfig(v *UpdateDataPipelineRequestProcessorsConfig) *UpdateDataPipelineRequestProcessors {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineRequestProcessors) SetName(v string) *UpdateDataPipelineRequestProcessors {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessors) SetType(v string) *UpdateDataPipelineRequestProcessors {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineRequestProcessorsConfig struct {
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
	Rules []*UpdateDataPipelineRequestProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *UpdateDataPipelineRequestProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *UpdateDataPipelineRequestProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineRequestProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestProcessorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetRules() []*UpdateDataPipelineRequestProcessorsConfigRules {
	return s.Rules
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetSelector() *UpdateDataPipelineRequestProcessorsConfigSelector {
	return s.Selector
}

func (s *UpdateDataPipelineRequestProcessorsConfig) GetTarget() *UpdateDataPipelineRequestProcessorsConfigTarget {
	return s.Target
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetApplications(v []*string) *UpdateDataPipelineRequestProcessorsConfig {
	s.Applications = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetExpression(v string) *UpdateDataPipelineRequestProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineRequestProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetRules(v []*UpdateDataPipelineRequestProcessorsConfigRules) *UpdateDataPipelineRequestProcessorsConfig {
	s.Rules = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetScript(v string) *UpdateDataPipelineRequestProcessorsConfig {
	s.Script = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetSelector(v *UpdateDataPipelineRequestProcessorsConfigSelector) *UpdateDataPipelineRequestProcessorsConfig {
	s.Selector = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) SetTarget(v *UpdateDataPipelineRequestProcessorsConfigTarget) *UpdateDataPipelineRequestProcessorsConfig {
	s.Target = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfig) Validate() error {
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

type UpdateDataPipelineRequestProcessorsConfigRules struct {
	// The retained prefix length.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The retained suffix length.
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

func (s UpdateDataPipelineRequestProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetKeepPrefix(v int32) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetKeepSuffix(v int32) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetKeys(v []*string) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetMaskChar(v string) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetMode(v string) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) SetTypes(v []*string) *UpdateDataPipelineRequestProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestProcessorsConfigSelector struct {
	// The service name list.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineRequestProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *UpdateDataPipelineRequestProcessorsConfigSelector) SetServiceNames(v []*string) *UpdateDataPipelineRequestProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDataPipelineRequestProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDataPipelineRequestProcessorsConfigTarget) SetWorkspace(v string) *UpdateDataPipelineRequestProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *UpdateDataPipelineRequestProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestSinks struct {
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

func (s UpdateDataPipelineRequestSinks) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestSinks) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *UpdateDataPipelineRequestSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateDataPipelineRequestSinks) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineRequestSinks) GetProject() *string {
	return s.Project
}

func (s *UpdateDataPipelineRequestSinks) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineRequestSinks) SetDatasets(v []*string) *UpdateDataPipelineRequestSinks {
	s.Datasets = v
	return s
}

func (s *UpdateDataPipelineRequestSinks) SetLogstore(v string) *UpdateDataPipelineRequestSinks {
	s.Logstore = &v
	return s
}

func (s *UpdateDataPipelineRequestSinks) SetName(v string) *UpdateDataPipelineRequestSinks {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineRequestSinks) SetProject(v string) *UpdateDataPipelineRequestSinks {
	s.Project = &v
	return s
}

func (s *UpdateDataPipelineRequestSinks) SetType(v string) *UpdateDataPipelineRequestSinks {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineRequestSinks) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineRequestSource struct {
	// The datasource config.
	Config *UpdateDataPipelineRequestSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The data source type.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDataPipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestSource) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestSource) GetConfig() *UpdateDataPipelineRequestSourceConfig {
	return s.Config
}

func (s *UpdateDataPipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineRequestSource) SetConfig(v *UpdateDataPipelineRequestSourceConfig) *UpdateDataPipelineRequestSource {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineRequestSource) SetType(v string) *UpdateDataPipelineRequestSource {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineRequestSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineRequestSourceConfig struct {
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
	TimeRange *UpdateDataPipelineRequestSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineRequestSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *UpdateDataPipelineRequestSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *UpdateDataPipelineRequestSourceConfig) GetTimeRange() *UpdateDataPipelineRequestSourceConfigTimeRange {
	return s.TimeRange
}

func (s *UpdateDataPipelineRequestSourceConfig) SetRunMode(v string) *UpdateDataPipelineRequestSourceConfig {
	s.RunMode = &v
	return s
}

func (s *UpdateDataPipelineRequestSourceConfig) SetStartFrom(v string) *UpdateDataPipelineRequestSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *UpdateDataPipelineRequestSourceConfig) SetTimeRange(v *UpdateDataPipelineRequestSourceConfigTimeRange) *UpdateDataPipelineRequestSourceConfig {
	s.TimeRange = v
	return s
}

func (s *UpdateDataPipelineRequestSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineRequestSourceConfigTimeRange struct {
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

func (s UpdateDataPipelineRequestSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineRequestSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineRequestSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *UpdateDataPipelineRequestSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *UpdateDataPipelineRequestSourceConfigTimeRange) SetFrom(v int64) *UpdateDataPipelineRequestSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *UpdateDataPipelineRequestSourceConfigTimeRange) SetTo(v int64) *UpdateDataPipelineRequestSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *UpdateDataPipelineRequestSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
