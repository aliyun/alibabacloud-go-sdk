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
	// The time when the pipeline was created.
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
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) GetRules() []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules {
	return s.Rules
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

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetExpression(v string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetFields(v []*string) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig) SetRules(v []*GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules) *GetDataPipelineResponseBodyPipelineOutputsProcessorsConfig {
	s.Rules = v
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

type GetDataPipelineResponseBodyPipelineOutputsProcessorsConfigRules struct {
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
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The field list.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The masking rule list.
	Rules []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) GetRules() []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules {
	return s.Rules
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

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetExpression(v string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetFields(v []*string) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Fields = v
	return s
}

func (s *GetDataPipelineResponseBodyPipelineProcessorsConfig) SetRules(v []*GetDataPipelineResponseBodyPipelineProcessorsConfigRules) *GetDataPipelineResponseBodyPipelineProcessorsConfig {
	s.Rules = v
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

type GetDataPipelineResponseBodyPipelineProcessorsConfigRules struct {
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
	// The data source type.
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
