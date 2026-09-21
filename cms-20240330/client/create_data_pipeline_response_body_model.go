// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *CreateDataPipelineResponseBodyPipeline) *CreateDataPipelineResponseBody
	GetPipeline() *CreateDataPipelineResponseBodyPipeline
	SetRequestId(v string) *CreateDataPipelineResponseBody
	GetRequestId() *string
}

type CreateDataPipelineResponseBody struct {
	// The data pipeline.
	Pipeline *CreateDataPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDataPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBody) GetPipeline() *CreateDataPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *CreateDataPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataPipelineResponseBody) SetPipeline(v *CreateDataPipelineResponseBodyPipeline) *CreateDataPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *CreateDataPipelineResponseBody) SetRequestId(v string) *CreateDataPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataPipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipeline struct {
	// The creation time.
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
	Outputs []*CreateDataPipelineResponseBodyPipelineOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
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
	Processors []*CreateDataPipelineResponseBodyPipelineProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	Sinks []*CreateDataPipelineResponseBodyPipelineSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	//
	// This parameter is required.
	Source *CreateDataPipelineResponseBodyPipelineSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
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

func (s CreateDataPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipeline) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateDataPipelineResponseBodyPipeline) GetDescription() *string {
	return s.Description
}

func (s *CreateDataPipelineResponseBodyPipeline) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineResponseBodyPipeline) GetOutputs() []*CreateDataPipelineResponseBodyPipelineOutputs {
	return s.Outputs
}

func (s *CreateDataPipelineResponseBodyPipeline) GetPipelineName() *string {
	return s.PipelineName
}

func (s *CreateDataPipelineResponseBodyPipeline) GetProcessors() []*CreateDataPipelineResponseBodyPipelineProcessors {
	return s.Processors
}

func (s *CreateDataPipelineResponseBodyPipeline) GetSignalType() *string {
	return s.SignalType
}

func (s *CreateDataPipelineResponseBodyPipeline) GetSinks() []*CreateDataPipelineResponseBodyPipelineSinks {
	return s.Sinks
}

func (s *CreateDataPipelineResponseBodyPipeline) GetSource() *CreateDataPipelineResponseBodyPipelineSource {
	return s.Source
}

func (s *CreateDataPipelineResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *CreateDataPipelineResponseBodyPipeline) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *CreateDataPipelineResponseBodyPipeline) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateDataPipelineResponseBodyPipeline) GetVersion() *int64 {
	return s.Version
}

func (s *CreateDataPipelineResponseBodyPipeline) SetCreateTime(v string) *CreateDataPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetDescription(v string) *CreateDataPipelineResponseBodyPipeline {
	s.Description = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetKind(v string) *CreateDataPipelineResponseBodyPipeline {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetOutputs(v []*CreateDataPipelineResponseBodyPipelineOutputs) *CreateDataPipelineResponseBodyPipeline {
	s.Outputs = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetPipelineName(v string) *CreateDataPipelineResponseBodyPipeline {
	s.PipelineName = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetProcessors(v []*CreateDataPipelineResponseBodyPipelineProcessors) *CreateDataPipelineResponseBodyPipeline {
	s.Processors = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetSignalType(v string) *CreateDataPipelineResponseBodyPipeline {
	s.SignalType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetSinks(v []*CreateDataPipelineResponseBodyPipelineSinks) *CreateDataPipelineResponseBodyPipeline {
	s.Sinks = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetSource(v *CreateDataPipelineResponseBodyPipelineSource) *CreateDataPipelineResponseBodyPipeline {
	s.Source = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetStatus(v string) *CreateDataPipelineResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetStatusMessage(v string) *CreateDataPipelineResponseBodyPipeline {
	s.StatusMessage = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetUpdateTime(v string) *CreateDataPipelineResponseBodyPipeline {
	s.UpdateTime = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) SetVersion(v int64) *CreateDataPipelineResponseBodyPipeline {
	s.Version = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipeline) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*CreateDataPipelineResponseBodyPipelineOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputs) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputs) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputs) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineOutputs) GetProcessors() []*CreateDataPipelineResponseBodyPipelineOutputsProcessors {
	return s.Processors
}

func (s *CreateDataPipelineResponseBodyPipelineOutputs) SetName(v string) *CreateDataPipelineResponseBodyPipelineOutputs {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputs) SetProcessors(v []*CreateDataPipelineResponseBodyPipelineOutputsProcessors) *CreateDataPipelineResponseBodyPipelineOutputs {
	s.Processors = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputs) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineOutputsProcessors struct {
	// The processor configuration.
	Config *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessors) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) GetConfig() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	return s.Config
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) SetConfig(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) *CreateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Config = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) SetName(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) SetType(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
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
	Projections []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetAssignments() []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	return s.Assignments
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetProjections() []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	return s.Projections
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetRules() []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScope() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	return s.Scope
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetSelector() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetTarget() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetApplications(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetAssignments(v []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetExpression(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetFields(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetParameters(v map[string]interface{}) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetProjections(v []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Projections = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetRules(v []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScope(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Scope = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScript(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetSelector(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetTarget(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments struct {
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetExpression(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetField(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections struct {
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetSource(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetTarget(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules struct {
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeys(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMaskChar(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMode(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetTypes(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetConditions() []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetMetricName() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetServiceName() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetConditions(v []*CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetMetricName(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetServiceName(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetField() *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetField(v *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField struct {
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetContainer(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetKind(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetName(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetPath(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName struct {
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

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) SetWorkspace(v string) *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessors struct {
	// The processor configuration.
	Config *CreateDataPipelineResponseBodyPipelineProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s CreateDataPipelineResponseBodyPipelineProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessors) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) GetConfig() *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	return s.Config
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) SetConfig(v *CreateDataPipelineResponseBodyPipelineProcessorsConfig) *CreateDataPipelineResponseBodyPipelineProcessors {
	s.Config = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) SetName(v string) *CreateDataPipelineResponseBodyPipelineProcessors {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) SetType(v string) *CreateDataPipelineResponseBodyPipelineProcessors {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The list of field assignments.
	Assignments []*CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
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
	Projections []*CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*CreateDataPipelineResponseBodyPipelineProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetAssignments() []*CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	return s.Assignments
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetProjections() []*CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	return s.Projections
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetRules() []*CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	return s.Rules
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetScope() *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	return s.Scope
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetSelector() *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	return s.Selector
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) GetTarget() *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	return s.Target
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetApplications(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Applications = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetAssignments(v []*CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetExpression(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetFields(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Fields = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetParameters(v map[string]interface{}) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetProjections(v []*CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Projections = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetRules(v []*CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Rules = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetScope(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Scope = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetScript(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Script = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetSelector(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Selector = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) SetTarget(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) *CreateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Target = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfig) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments struct {
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

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetExpression(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetField(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections struct {
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

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetSource(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetTarget(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigRules struct {
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

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepPrefix(v int32) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepSuffix(v int32) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeys(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMaskChar(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMode(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) SetTypes(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetConditions() []*CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetMetricName() *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) GetServiceName() *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetConditions(v []*CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetMetricName(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) SetServiceName(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScope) Validate() error {
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

type CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The matching method.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetField() *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetField(v *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField struct {
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
	// service.name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetContainer(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetKind(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetName(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetPath(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName struct {
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

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName struct {
	// The matching method.
	//
	// example:
	//
	// GLOB
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetMatchType(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetValues(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) SetServiceNames(v []*string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) SetWorkspace(v string) *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineSinks struct {
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

func (s CreateDataPipelineResponseBodyPipelineSinks) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineSinks) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) GetName() *string {
	return s.Name
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) GetProject() *string {
	return s.Project
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) SetDatasets(v []*string) *CreateDataPipelineResponseBodyPipelineSinks {
	s.Datasets = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) SetLogstore(v string) *CreateDataPipelineResponseBodyPipelineSinks {
	s.Logstore = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) SetName(v string) *CreateDataPipelineResponseBodyPipelineSinks {
	s.Name = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) SetProject(v string) *CreateDataPipelineResponseBodyPipelineSinks {
	s.Project = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) SetType(v string) *CreateDataPipelineResponseBodyPipelineSinks {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSinks) Validate() error {
	return dara.Validate(s)
}

type CreateDataPipelineResponseBodyPipelineSource struct {
	// The datasource config.
	Config *CreateDataPipelineResponseBodyPipelineSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The type of the data source.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataPipelineResponseBodyPipelineSource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineSource) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineSource) GetConfig() *CreateDataPipelineResponseBodyPipelineSourceConfig {
	return s.Config
}

func (s *CreateDataPipelineResponseBodyPipelineSource) GetType() *string {
	return s.Type
}

func (s *CreateDataPipelineResponseBodyPipelineSource) SetConfig(v *CreateDataPipelineResponseBodyPipelineSourceConfig) *CreateDataPipelineResponseBodyPipelineSource {
	s.Config = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSource) SetType(v string) *CreateDataPipelineResponseBodyPipelineSource {
	s.Type = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineSourceConfig struct {
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
	TimeRange *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s CreateDataPipelineResponseBodyPipelineSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) GetTimeRange() *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	return s.TimeRange
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) SetRunMode(v string) *CreateDataPipelineResponseBodyPipelineSourceConfig {
	s.RunMode = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) SetStartFrom(v string) *CreateDataPipelineResponseBodyPipelineSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) SetTimeRange(v *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) *CreateDataPipelineResponseBodyPipelineSourceConfig {
	s.TimeRange = v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange struct {
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

func (s CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetFrom(v int64) *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetTo(v int64) *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *CreateDataPipelineResponseBodyPipelineSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
