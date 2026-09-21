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
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*CreateDataPipelineRequestOutputsProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The list of fields.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The extended parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The list of field projections.
	Projections []*CreateDataPipelineRequestOutputsProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*CreateDataPipelineRequestOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The pipeline processing scope.
	Scope *CreateDataPipelineRequestOutputsProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
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

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetAssignments() []*CreateDataPipelineRequestOutputsProcessorsConfigAssignments {
	return s.Assignments
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetProjections() []*CreateDataPipelineRequestOutputsProcessorsConfigProjections {
	return s.Projections
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetRules() []*CreateDataPipelineRequestOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) GetScope() *CreateDataPipelineRequestOutputsProcessorsConfigScope {
	return s.Scope
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

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetAssignments(v []*CreateDataPipelineRequestOutputsProcessorsConfigAssignments) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Assignments = v
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

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetParameters(v map[string]interface{}) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetProjections(v []*CreateDataPipelineRequestOutputsProcessorsConfigProjections) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Projections = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetRules(v []*CreateDataPipelineRequestOutputsProcessorsConfigRules) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfig) SetScope(v *CreateDataPipelineRequestOutputsProcessorsConfigScope) *CreateDataPipelineRequestOutputsProcessorsConfig {
	s.Scope = v
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
	if s.Assignments != nil {
		for _, item := range s.Assignments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Projections != nil {
		for _, item := range s.Projections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
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

type CreateDataPipelineRequestOutputsProcessorsConfigAssignments struct {
	// The assignment expression.
	//
	// example:
	//
	// duration / 1000000.0
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The output field.
	//
	// example:
	//
	// latency_ms
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigAssignments) SetExpression(v string) *CreateDataPipelineRequestOutputsProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigAssignments) SetField(v string) *CreateDataPipelineRequestOutputsProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigProjections struct {
	// The source field.
	//
	// example:
	//
	// serviceName
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target field.
	//
	// example:
	//
	// service
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigProjections) SetSource(v string) *CreateDataPipelineRequestOutputsProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigProjections) SetTarget(v string) *CreateDataPipelineRequestOutputsProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
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

type CreateDataPipelineRequestOutputsProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) GetConditions() []*CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) GetMetricName() *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) GetServiceName() *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) SetConditions(v []*CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) *CreateDataPipelineRequestOutputsProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) SetMetricName(v *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) *CreateDataPipelineRequestOutputsProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) SetServiceName(v *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) *CreateDataPipelineRequestOutputsProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScope) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetricName != nil {
		if err := s.MetricName.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceName != nil {
		if err := s.ServiceName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) GetField() *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) SetField(v *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) SetMatchType(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) SetValues(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField struct {
	// The JSON object container.
	//
	// example:
	//
	// attributes
	Container *string `json:"container,omitempty" xml:"container,omitempty"`
	// The reference data type.
	//
	// example:
	//
	// label
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The field or dimension name.
	//
	// example:
	//
	// serviceName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) SetContainer(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) SetKind(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) SetName(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) SetPath(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The metric names.
	//
	// example:
	//
	// ["http_requests_total"]
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) SetMatchType(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) SetValues(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) SetMatchType(v string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) SetValues(v []*string) *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestOutputsProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestOutputsProcessorsConfigSelector struct {
	// The list of service names.
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
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*CreateDataPipelineRequestProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The list of fields.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The extended parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The list of field projections.
	Projections []*CreateDataPipelineRequestProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*CreateDataPipelineRequestProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The pipeline processing scope in which the processor takes effect.
	Scope *CreateDataPipelineRequestProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
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

func (s *CreateDataPipelineRequestProcessorsConfig) GetAssignments() []*CreateDataPipelineRequestProcessorsConfigAssignments {
	return s.Assignments
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetProjections() []*CreateDataPipelineRequestProcessorsConfigProjections {
	return s.Projections
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetRules() []*CreateDataPipelineRequestProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineRequestProcessorsConfig) GetScope() *CreateDataPipelineRequestProcessorsConfigScope {
	return s.Scope
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

func (s *CreateDataPipelineRequestProcessorsConfig) SetAssignments(v []*CreateDataPipelineRequestProcessorsConfigAssignments) *CreateDataPipelineRequestProcessorsConfig {
	s.Assignments = v
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

func (s *CreateDataPipelineRequestProcessorsConfig) SetParameters(v map[string]interface{}) *CreateDataPipelineRequestProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetProjections(v []*CreateDataPipelineRequestProcessorsConfigProjections) *CreateDataPipelineRequestProcessorsConfig {
	s.Projections = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetRules(v []*CreateDataPipelineRequestProcessorsConfigRules) *CreateDataPipelineRequestProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfig) SetScope(v *CreateDataPipelineRequestProcessorsConfigScope) *CreateDataPipelineRequestProcessorsConfig {
	s.Scope = v
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
	if s.Assignments != nil {
		for _, item := range s.Assignments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Projections != nil {
		for _, item := range s.Projections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
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

type CreateDataPipelineRequestProcessorsConfigAssignments struct {
	// The assignment expression.
	//
	// example:
	//
	// duration / 1000000.0
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The output field.
	//
	// example:
	//
	// latency_ms
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateDataPipelineRequestProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineRequestProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *CreateDataPipelineRequestProcessorsConfigAssignments) SetExpression(v string) *CreateDataPipelineRequestProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigAssignments) SetField(v string) *CreateDataPipelineRequestProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigProjections struct {
	// The source field.
	//
	// example:
	//
	// serviceName
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target field.
	//
	// example:
	//
	// service
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s CreateDataPipelineRequestProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *CreateDataPipelineRequestProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *CreateDataPipelineRequestProcessorsConfigProjections) SetSource(v string) *CreateDataPipelineRequestProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigProjections) SetTarget(v string) *CreateDataPipelineRequestProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
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

type CreateDataPipelineRequestProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*CreateDataPipelineRequestProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *CreateDataPipelineRequestProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *CreateDataPipelineRequestProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s CreateDataPipelineRequestProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) GetConditions() []*CreateDataPipelineRequestProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) GetMetricName() *CreateDataPipelineRequestProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) GetServiceName() *CreateDataPipelineRequestProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) SetConditions(v []*CreateDataPipelineRequestProcessorsConfigScopeConditions) *CreateDataPipelineRequestProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) SetMetricName(v *CreateDataPipelineRequestProcessorsConfigScopeMetricName) *CreateDataPipelineRequestProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) SetServiceName(v *CreateDataPipelineRequestProcessorsConfigScopeServiceName) *CreateDataPipelineRequestProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScope) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetricName != nil {
		if err := s.MetricName.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceName != nil {
		if err := s.ServiceName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *CreateDataPipelineRequestProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) GetField() *CreateDataPipelineRequestProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) SetField(v *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) *CreateDataPipelineRequestProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) SetMatchType(v string) *CreateDataPipelineRequestProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) SetValues(v []*string) *CreateDataPipelineRequestProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineRequestProcessorsConfigScopeConditionsField struct {
	// The JSON object container.
	//
	// example:
	//
	// attributes
	Container *string `json:"container,omitempty" xml:"container,omitempty"`
	// The reference data type.
	//
	// example:
	//
	// field
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The field or dimension name.
	//
	// example:
	//
	// serviceName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) SetContainer(v string) *CreateDataPipelineRequestProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) SetKind(v string) *CreateDataPipelineRequestProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) SetName(v string) *CreateDataPipelineRequestProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) SetPath(v []*string) *CreateDataPipelineRequestProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigScopeMetricName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The metric names.
	//
	// example:
	//
	// ["http_requests_total"]
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeMetricName) SetMatchType(v string) *CreateDataPipelineRequestProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeMetricName) SetValues(v []*string) *CreateDataPipelineRequestProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineRequestProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineRequestProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeServiceName) SetMatchType(v string) *CreateDataPipelineRequestProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeServiceName) SetValues(v []*string) *CreateDataPipelineRequestProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineRequestProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineRequestProcessorsConfigSelector struct {
	// The list of service names.
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
	// The type of the data source.
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
