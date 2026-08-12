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
	Rules []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetRules() []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	return s.Rules
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

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetRules(v []*UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) *UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Rules = v
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

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules struct {
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

type UpdateDataPipelineResponseBodyPipelineOutputsProcessorsConfigSelector struct {
	// The service name list.
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
	Rules []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) GetRules() []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules {
	return s.Rules
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

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetExpression(v string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetFields(v []*string) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Fields = v
	return s
}

func (s *UpdateDataPipelineResponseBodyPipelineProcessorsConfig) SetRules(v []*UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules) *UpdateDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Rules = v
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

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigRules struct {
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

type UpdateDataPipelineResponseBodyPipelineProcessorsConfigSelector struct {
	// The service name list.
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
	// The data source type.
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
