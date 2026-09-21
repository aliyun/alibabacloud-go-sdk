// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *GetDataPipelineResponseBodyPipeline) *GetDataPipelineResponseBody
	GetPipeline() *GetDataPipelineResponseBodyPipeline
	SetRequestId(v string) *GetDataPipelineResponseBody
	GetRequestId() *string
}

type GetDataPipelineResponseBody struct {
	// The data pipeline.
	Pipeline *GetDataPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBody) GetPipeline() *GetDataPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *GetDataPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataPipelineResponseBody) SetPipeline(v *GetDataPipelineResponseBodyPipeline) *GetDataPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetDataPipelineResponseBody) SetRequestId(v string) *GetDataPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataPipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipeline struct {
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
	Outputs []*GetDataPipelineResponseBodyPipelineOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
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
	Processors []*GetDataPipelineResponseBodyPipelineProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	Sinks []*GetDataPipelineResponseBodyPipelineSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	//
	// This parameter is required.
	Source *GetDataPipelineResponseBodyPipelineSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
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

func (s GetDataPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipeline) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDataPipelineResponseBodyPipeline) GetDescription() *string {
	return s.Description
}

func (s *GetDataPipelineResponseBodyPipeline) GetKind() *string {
	return s.Kind
}

func (s *GetDataPipelineResponseBodyPipeline) GetOutputs() []*GetDataPipelineResponseBodyPipelineOutputs {
	return s.Outputs
}

func (s *GetDataPipelineResponseBodyPipeline) GetPipelineName() *string {
	return s.PipelineName
}

func (s *GetDataPipelineResponseBodyPipeline) GetProcessors() []*GetDataPipelineResponseBodyPipelineProcessors {
	return s.Processors
}

func (s *GetDataPipelineResponseBodyPipeline) GetSignalType() *string {
	return s.SignalType
}

func (s *GetDataPipelineResponseBodyPipeline) GetSinks() []*GetDataPipelineResponseBodyPipelineSinks {
	return s.Sinks
}

func (s *GetDataPipelineResponseBodyPipeline) GetSource() *GetDataPipelineResponseBodyPipelineSource {
	return s.Source
}

func (s *GetDataPipelineResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *GetDataPipelineResponseBodyPipeline) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetDataPipelineResponseBodyPipeline) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDataPipelineResponseBodyPipeline) GetVersion() *int64 {
	return s.Version
}

func (s *GetDataPipelineResponseBodyPipeline) SetCreateTime(v string) *GetDataPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetDescription(v string) *GetDataPipelineResponseBodyPipeline {
	s.Description = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetKind(v string) *GetDataPipelineResponseBodyPipeline {
	s.Kind = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetOutputs(v []*GetDataPipelineResponseBodyPipelineOutputs) *GetDataPipelineResponseBodyPipeline {
	s.Outputs = v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetPipelineName(v string) *GetDataPipelineResponseBodyPipeline {
	s.PipelineName = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetProcessors(v []*GetDataPipelineResponseBodyPipelineProcessors) *GetDataPipelineResponseBodyPipeline {
	s.Processors = v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetSignalType(v string) *GetDataPipelineResponseBodyPipeline {
	s.SignalType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetSinks(v []*GetDataPipelineResponseBodyPipelineSinks) *GetDataPipelineResponseBodyPipeline {
	s.Sinks = v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetSource(v *GetDataPipelineResponseBodyPipelineSource) *GetDataPipelineResponseBodyPipeline {
	s.Source = v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetStatus(v string) *GetDataPipelineResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetStatusMessage(v string) *GetDataPipelineResponseBodyPipeline {
	s.StatusMessage = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetUpdateTime(v string) *GetDataPipelineResponseBodyPipeline {
	s.UpdateTime = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) SetVersion(v int64) *GetDataPipelineResponseBodyPipeline {
	s.Version = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipeline) Validate() error {
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

type GetDataPipelineResponseBodyPipelineOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*GetDataPipelineResponseBodyPipelineOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputs) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputs) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineOutputs) GetProcessors() []*GetDataPipelineResponseBodyPipelineOutputsProcessors {
	return s.Processors
}

func (s *GetDataPipelineResponseBodyPipelineOutputs) SetName(v string) *GetDataPipelineResponseBodyPipelineOutputs {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputs) SetProcessors(v []*GetDataPipelineResponseBodyPipelineOutputsProcessors) *GetDataPipelineResponseBodyPipelineOutputs {
	s.Processors = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputs) Validate() error {
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

type GetDataPipelineResponseBodyPipelineOutputsProcessors struct {
	// The processor configuration.
	Config *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s GetDataPipelineResponseBodyPipelineOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessors) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) GetConfig() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	return s.Config
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) SetConfig(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) *GetDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Config = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) SetName(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) SetType(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessors {
	s.Type = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig struct {
	// The application list.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The field assignment list.
	Assignments []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The extended parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The field projection list.
	Projections []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetAssignments() []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	return s.Assignments
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetProjections() []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	return s.Projections
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetRules() []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScope() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	return s.Scope
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetSelector() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetTarget() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetApplications(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetAssignments(v []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetExpression(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetFields(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetParameters(v map[string]interface{}) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetProjections(v []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Projections = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetRules(v []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScope(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Scope = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetScript(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetSelector(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetTarget(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) Validate() error {
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

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments struct {
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

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetExpression(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) SetField(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections struct {
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

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetSource(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) SetTarget(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules struct {
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

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetKeys(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMaskChar(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetMode(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) SetTypes(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetConditions() []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetMetricName() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) GetServiceName() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetConditions(v []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetMetricName(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) SetServiceName(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScope) Validate() error {
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

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetField() *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetField(v *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField struct {
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
	// azone
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetContainer(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetKind(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetName(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) SetPath(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName struct {
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The metric name.
	//
	// example:
	//
	// ["http_requests_total"]
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName struct {
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) SetWorkspace(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessors struct {
	// The processor configuration.
	Config *GetDataPipelineResponseBodyPipelineProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s GetDataPipelineResponseBodyPipelineProcessors) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessors) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) GetConfig() *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	return s.Config
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) GetType() *string {
	return s.Type
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) SetConfig(v *GetDataPipelineResponseBodyPipelineProcessorsConfig) *GetDataPipelineResponseBodyPipelineProcessors {
	s.Config = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) SetName(v string) *GetDataPipelineResponseBodyPipelineProcessors {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) SetType(v string) *GetDataPipelineResponseBodyPipelineProcessors {
	s.Type = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineProcessorsConfig struct {
	// The application list.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The field assignment list.
	Assignments []*GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments `json:"assignments,omitempty" xml:"assignments,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The extended parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The field projection list.
	Projections []*GetDataPipelineResponseBodyPipelineProcessorsConfigProjections `json:"projections,omitempty" xml:"projections,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The scope in which the pipeline processing processor takes effect.
	Scope *GetDataPipelineResponseBodyPipelineProcessorsConfigScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfig) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetAssignments() []*GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	return s.Assignments
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetProjections() []*GetDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	return s.Projections
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetRules() []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	return s.Rules
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetScope() *GetDataPipelineResponseBodyPipelineProcessorsConfigScope {
	return s.Scope
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetSelector() *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	return s.Selector
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetTarget() *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	return s.Target
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetApplications(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Applications = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetAssignments(v []*GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Assignments = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetExpression(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetFields(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Fields = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetParameters(v map[string]interface{}) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Parameters = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetProjections(v []*GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Projections = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetRules(v []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Rules = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetScope(v *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Scope = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetScript(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Script = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetSelector(v *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Selector = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetTarget(v *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Target = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) Validate() error {
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

type GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments struct {
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

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) GetField() *string {
	return s.Field
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetExpression(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) SetField(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments {
	s.Field = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigAssignments) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigProjections struct {
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

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetSource() *string {
	return s.Source
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) GetTarget() *string {
	return s.Target
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetSource(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Source = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) SetTarget(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections {
	s.Target = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigProjections) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigRules struct {
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

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepPrefix(v int32) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeepSuffix(v int32) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetKeys(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMaskChar(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetMode(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) SetTypes(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigScope struct {
	// The additional field conditions.
	Conditions []*GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The metric name scope.
	MetricName *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName `json:"metricName,omitempty" xml:"metricName,omitempty" type:"Struct"`
	// The service name scope.
	ServiceName *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName `json:"serviceName,omitempty" xml:"serviceName,omitempty" type:"Struct"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScope) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScope) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) GetConditions() []*GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	return s.Conditions
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) GetMetricName() *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	return s.MetricName
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) GetServiceName() *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	return s.ServiceName
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) SetConditions(v []*GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) *GetDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.Conditions = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) SetMetricName(v *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) *GetDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.MetricName = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) SetServiceName(v *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) *GetDataPipelineResponseBodyPipelineProcessorsConfigScope {
	s.ServiceName = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScope) Validate() error {
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

type GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions struct {
	// The field reference.
	Field *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField `json:"field,omitempty" xml:"field,omitempty" type:"Struct"`
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetField() *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	return s.Field
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetField(v *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Field = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditions) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField struct {
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
	// explorer_link
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The field or dimension name.
	//
	// example:
	//
	// site_check_monitor
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON literal key path.
	Path []*string `json:"path,omitempty" xml:"path,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetContainer() *string {
	return s.Container
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetKind() *string {
	return s.Kind
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) GetPath() []*string {
	return s.Path
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetContainer(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Container = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetKind(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Kind = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetName(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) SetPath(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField {
	s.Path = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeConditionsField) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName struct {
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The metric name.
	//
	// example:
	//
	// ["http_requests_total"]
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeMetricName) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName struct {
	// The match type.
	//
	// example:
	//
	// EXACT
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetMatchType() *string {
	return s.MatchType
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) GetValues() []*string {
	return s.Values
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetMatchType(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.MatchType = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) SetValues(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName {
	s.Values = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigScopeServiceName) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) SetServiceNames(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) SetWorkspace(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineSinks struct {
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

func (s GetDataPipelineResponseBodyPipelineSinks) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineSinks) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *GetDataPipelineResponseBodyPipelineSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *GetDataPipelineResponseBodyPipelineSinks) GetName() *string {
	return s.Name
}

func (s *GetDataPipelineResponseBodyPipelineSinks) GetProject() *string {
	return s.Project
}

func (s *GetDataPipelineResponseBodyPipelineSinks) GetType() *string {
	return s.Type
}

func (s *GetDataPipelineResponseBodyPipelineSinks) SetDatasets(v []*string) *GetDataPipelineResponseBodyPipelineSinks {
	s.Datasets = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSinks) SetLogstore(v string) *GetDataPipelineResponseBodyPipelineSinks {
	s.Logstore = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSinks) SetName(v string) *GetDataPipelineResponseBodyPipelineSinks {
	s.Name = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSinks) SetProject(v string) *GetDataPipelineResponseBodyPipelineSinks {
	s.Project = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSinks) SetType(v string) *GetDataPipelineResponseBodyPipelineSinks {
	s.Type = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSinks) Validate() error {
	return dara.Validate(s)
}

type GetDataPipelineResponseBodyPipelineSource struct {
	// The datasource config.
	Config *GetDataPipelineResponseBodyPipelineSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The type of the data source.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDataPipelineResponseBodyPipelineSource) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineSource) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineSource) GetConfig() *GetDataPipelineResponseBodyPipelineSourceConfig {
	return s.Config
}

func (s *GetDataPipelineResponseBodyPipelineSource) GetType() *string {
	return s.Type
}

func (s *GetDataPipelineResponseBodyPipelineSource) SetConfig(v *GetDataPipelineResponseBodyPipelineSourceConfig) *GetDataPipelineResponseBodyPipelineSource {
	s.Config = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSource) SetType(v string) *GetDataPipelineResponseBodyPipelineSource {
	s.Type = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineSourceConfig struct {
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
	TimeRange *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetDataPipelineResponseBodyPipelineSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineSourceConfig) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) GetTimeRange() *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	return s.TimeRange
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) SetRunMode(v string) *GetDataPipelineResponseBodyPipelineSourceConfig {
	s.RunMode = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) SetStartFrom(v string) *GetDataPipelineResponseBodyPipelineSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) SetTimeRange(v *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) *GetDataPipelineResponseBodyPipelineSourceConfig {
	s.TimeRange = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataPipelineResponseBodyPipelineSourceConfigTimeRange struct {
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

func (s GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetFrom(v int64) *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) SetTo(v int64) *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
