// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *UpdateDataPipelineResponseBodyPipeline) *UpdateDataPipelineResponseBody
	GetPipeline() *UpdateDataPipelineResponseBodyPipeline
	SetRequestId(v string) *UpdateDataPipelineResponseBody
	GetRequestId() *string
}

type UpdateDataPipelineResponseBody struct {
	// The data pipeline.
	Pipeline *UpdateDataPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDataPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBody) GetPipeline() *UpdateDataPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *UpdateDataPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataPipelineResponseBody) SetPipeline(v *UpdateDataPipelineResponseBodyPipeline) *UpdateDataPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *UpdateDataPipelineResponseBody) SetRequestId(v string) *UpdateDataPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataPipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipeline struct {
	// The time when the data pipeline was created.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-10T05:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// Export selected trace services to the target workspace.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The pipeline type.
	//
	// This parameter is required.
	//
	// example:
	//
	// export
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The named outputs.
	//
	// This parameter is required.
	Outputs []*UpdateDataPipelineResponseBodyPipelineOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The pipeline name.
	//
	// This parameter is required.
	//
	// example:
	//
	// export-traces-to-prod
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The common processors.
	//
	// This parameter is required.
	Processors []*UpdateDataPipelineResponseBodyPipelineProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The signal type.
	//
	// This parameter is required.
	//
	// example:
	//
	// traces
	SignalType *string `json:"signalType,omitempty" xml:"signalType,omitempty"`
	// The output destinations.
	//
	// This parameter is required.
	Sinks []*UpdateDataPipelineResponseBodyPipelineSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	//
	// This parameter is required.
	Source *UpdateDataPipelineResponseBodyPipelineSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The running status.
	//
	// This parameter is required.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status message.
	//
	// example:
	//
	// Pipeline is running.
	StatusMessage *string `json:"statusMessage,omitempty" xml:"statusMessage,omitempty"`
	// The update time.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-10T05:10:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The configuration version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateDataPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetKind() *string {
	return s.Kind
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetOutputs() []*UpdateDataPipelineResponseBodyPipelineOutputs {
	return s.Outputs
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetPipelineName() *string {
	return s.PipelineName
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetProcessors() []*UpdateDataPipelineResponseBodyPipelineProcessors {
	return s.Processors
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetSignalType() *string {
	return s.SignalType
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetSinks() []*UpdateDataPipelineResponseBodyPipelineSinks {
	return s.Sinks
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetSource() *UpdateDataPipelineResponseBodyPipelineSource {
	return s.Source
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateDataPipelineResponseBodyPipeline) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetCreateTime(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetDescription(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.Description = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetKind(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.Kind = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetOutputs(v []*UpdateDataPipelineResponseBodyPipelineOutputs) *UpdateDataPipelineResponseBodyPipeline {
	s.Outputs = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetPipelineName(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.PipelineName = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetProcessors(v []*UpdateDataPipelineResponseBodyPipelineProcessors) *UpdateDataPipelineResponseBodyPipeline {
	s.Processors = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetSignalType(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.SignalType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetSinks(v []*UpdateDataPipelineResponseBodyPipelineSinks) *UpdateDataPipelineResponseBodyPipeline {
	s.Sinks = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetSource(v *UpdateDataPipelineResponseBodyPipelineSource) *UpdateDataPipelineResponseBodyPipeline {
	s.Source = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetStatus(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetStatusMessage(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.StatusMessage = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetUpdateTime(v string) *UpdateDataPipelineResponseBodyPipeline {
	s.UpdateTime = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) SetVersion(v int64) *UpdateDataPipelineResponseBodyPipeline {
	s.Version = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipeline) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*UpdateDataPipelineResponseBodyPipelineOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputs) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputs) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputs) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputs) GetProcessors() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessors {
	return s.Processors
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputs) SetName(v string) *UpdateDataPipelineResponseBodyPipelineOutputs {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputs) SetProcessors(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessors) *UpdateDataPipelineResponseBodyPipelineOutputs {
	s.Processors = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputs) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineOutputsProcessors struct {
	// The processor configuration.
	Config *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessors) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) GetConfig() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	return s.Config
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) SetConfig(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) *UpdateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) SetName(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) SetType(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
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
	Projections []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetAssignments() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	return s.Assignments
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetProjections() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	return s.Projections
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetRules() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScope() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	return s.Scope
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetSelector() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetTarget() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetApplications(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetAssignments(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetParameters(v map[string]interface{}) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetProjections(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Projections = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetRules(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScope(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Scope = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScript(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetSelector(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetTarget(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments struct {
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

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetField(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections struct {
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

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetSource(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetTarget(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules struct {
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

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeys(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMaskChar(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMode(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetTypes(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetConditions() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetMetricName() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetServiceName() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetConditions(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetMetricName(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetServiceName(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetField() *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetField(v *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField struct {
	// The JSON object container.
	//
	// example:
	//
	// {}
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
	// instanceId
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetContainer(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetKind(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetName(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetPath(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName struct {
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

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) SetWorkspace(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessors struct {
	// The processor configuration.
	Config *UpdateDataPipelineResponseBodyPipelineProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s UpdateDataPipelineResponseBodyPipelineProcessors) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessors) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) GetConfig() *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	return s.Config
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) SetConfig(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) *UpdateDataPipelineResponseBodyPipelineProcessors {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) SetName(v string) *UpdateDataPipelineResponseBodyPipelineProcessors {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) SetType(v string) *UpdateDataPipelineResponseBodyPipelineProcessors {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
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
	Projections []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetAssignments() []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	return s.Assignments
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetProjections() []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	return s.Projections
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetRules() []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	return s.Rules
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetScope() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	return s.Scope
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetSelector() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	return s.Selector
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetTarget() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	return s.Target
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetApplications(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Applications = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetAssignments(v []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetParameters(v map[string]interface{}) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetProjections(v []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Projections = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetRules(v []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Rules = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetScope(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Scope = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetScript(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Script = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetSelector(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Selector = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetTarget(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Target = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments struct {
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

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetField(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections struct {
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

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetSource(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetTarget(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules struct {
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

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepPrefix(v int32) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepSuffix(v int32) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeys(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMaskChar(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMode(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetTypes(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetConditions() []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetMetricName() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetServiceName() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetConditions(v []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetMetricName(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetServiceName(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScope) Validate() error {
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

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetField() *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetField(v *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField struct {
	// The JSON object container.
	//
	// example:
	//
	// {}
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
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetContainer(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetKind(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetName(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetPath(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName struct {
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

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetMatchType(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetValues(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) SetServiceNames(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) SetWorkspace(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineSinks struct {
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

func (s UpdateDataPipelineResponseBodyPipelineSinks) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineSinks) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) GetName() *string {
	return s.Name
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) GetProject() *string {
	return s.Project
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) SetDatasets(v []*string) *UpdateDataPipelineResponseBodyPipelineSinks {
	s.Datasets = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) SetLogstore(v string) *UpdateDataPipelineResponseBodyPipelineSinks {
	s.Logstore = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) SetName(v string) *UpdateDataPipelineResponseBodyPipelineSinks {
	s.Name = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) SetProject(v string) *UpdateDataPipelineResponseBodyPipelineSinks {
	s.Project = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) SetType(v string) *UpdateDataPipelineResponseBodyPipelineSinks {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSinks) Validate() error {
	return dara.Validate(s)
}

type UpdateDataPipelineResponseBodyPipelineSource struct {
	// The datasource config.
	Config *UpdateDataPipelineResponseBodyPipelineSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The type of the data source.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDataPipelineResponseBodyPipelineSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineSource) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineSource) GetConfig() *UpdateDataPipelineResponseBodyPipelineSourceConfig {
	return s.Config
}

func (s *UpdateDataPipelineResponseBodyPipelineSource) GetType() *string {
	return s.Type
}

func (s *UpdateDataPipelineResponseBodyPipelineSource) SetConfig(v *UpdateDataPipelineResponseBodyPipelineSourceConfig) *UpdateDataPipelineResponseBodyPipelineSource {
	s.Config = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSource) SetType(v string) *UpdateDataPipelineResponseBodyPipelineSource {
	s.Type = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineSourceConfig struct {
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
	TimeRange *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s UpdateDataPipelineResponseBodyPipelineSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) GetTimeRange() *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	return s.TimeRange
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) SetRunMode(v string) *UpdateDataPipelineResponseBodyPipelineSourceConfig {
	s.RunMode = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) SetStartFrom(v string) *UpdateDataPipelineResponseBodyPipelineSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) SetTimeRange(v *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) *UpdateDataPipelineResponseBodyPipelineSourceConfig {
	s.TimeRange = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange struct {
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

func (s UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetFrom(v int64) *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetTo(v int64) *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
