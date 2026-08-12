// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDataPipelinesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataPipelinesResponseBody
	GetNextToken() *string
	SetPipelines(v []*ListDataPipelinesResponseBodyPipelines) *ListDataPipelinesResponseBody
	GetPipelines() []*ListDataPipelinesResponseBodyPipelines
	SetRequestId(v string) *ListDataPipelinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDataPipelinesResponseBody
	GetTotalCount() *int64
}

type ListDataPipelinesResponseBody struct {
	// The maximum number of results per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of data pipelines.
	Pipelines []*ListDataPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of data pipelines.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataPipelinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataPipelinesResponseBody) GetPipelines() []*ListDataPipelinesResponseBodyPipelines {
	return s.Pipelines
}

func (s *ListDataPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataPipelinesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataPipelinesResponseBody) SetMaxResults(v int32) *ListDataPipelinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataPipelinesResponseBody) SetNextToken(v string) *ListDataPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataPipelinesResponseBody) SetPipelines(v []*ListDataPipelinesResponseBodyPipelines) *ListDataPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListDataPipelinesResponseBody) SetRequestId(v string) *ListDataPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataPipelinesResponseBody) SetTotalCount(v int64) *ListDataPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataPipelinesResponseBody) Validate() error {
	if s.Pipelines != nil {
		for _, item := range s.Pipelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataPipelinesResponseBodyPipelines struct {
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
	// The description of the pipeline.
	//
	// example:
	//
	// Export selected trace services to the target workspace.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the pipeline.
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
	Outputs []*ListDataPipelinesResponseBodyPipelinesOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The name of the pipeline.
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
	Processors []*ListDataPipelinesResponseBodyPipelinesProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
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
	Sinks []*ListDataPipelinesResponseBodyPipelinesSinks `json:"sinks,omitempty" xml:"sinks,omitempty" type:"Repeated"`
	// The data source.
	//
	// This parameter is required.
	Source *ListDataPipelinesResponseBodyPipelinesSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
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
	// The time when the pipeline was last updated.
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

func (s ListDataPipelinesResponseBodyPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataPipelinesResponseBodyPipelines) GetDescription() *string {
	return s.Description
}

func (s *ListDataPipelinesResponseBodyPipelines) GetKind() *string {
	return s.Kind
}

func (s *ListDataPipelinesResponseBodyPipelines) GetOutputs() []*ListDataPipelinesResponseBodyPipelinesOutputs {
	return s.Outputs
}

func (s *ListDataPipelinesResponseBodyPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListDataPipelinesResponseBodyPipelines) GetProcessors() []*ListDataPipelinesResponseBodyPipelinesProcessors {
	return s.Processors
}

func (s *ListDataPipelinesResponseBodyPipelines) GetSignalType() *string {
	return s.SignalType
}

func (s *ListDataPipelinesResponseBodyPipelines) GetSinks() []*ListDataPipelinesResponseBodyPipelinesSinks {
	return s.Sinks
}

func (s *ListDataPipelinesResponseBodyPipelines) GetSource() *ListDataPipelinesResponseBodyPipelinesSource {
	return s.Source
}

func (s *ListDataPipelinesResponseBodyPipelines) GetStatus() *string {
	return s.Status
}

func (s *ListDataPipelinesResponseBodyPipelines) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListDataPipelinesResponseBodyPipelines) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDataPipelinesResponseBodyPipelines) GetVersion() *int64 {
	return s.Version
}

func (s *ListDataPipelinesResponseBodyPipelines) SetCreateTime(v string) *ListDataPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetDescription(v string) *ListDataPipelinesResponseBodyPipelines {
	s.Description = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetKind(v string) *ListDataPipelinesResponseBodyPipelines {
	s.Kind = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetOutputs(v []*ListDataPipelinesResponseBodyPipelinesOutputs) *ListDataPipelinesResponseBodyPipelines {
	s.Outputs = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListDataPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetProcessors(v []*ListDataPipelinesResponseBodyPipelinesProcessors) *ListDataPipelinesResponseBodyPipelines {
	s.Processors = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetSignalType(v string) *ListDataPipelinesResponseBodyPipelines {
	s.SignalType = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetSinks(v []*ListDataPipelinesResponseBodyPipelinesSinks) *ListDataPipelinesResponseBodyPipelines {
	s.Sinks = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetSource(v *ListDataPipelinesResponseBodyPipelinesSource) *ListDataPipelinesResponseBodyPipelines {
	s.Source = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetStatus(v string) *ListDataPipelinesResponseBodyPipelines {
	s.Status = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetStatusMessage(v string) *ListDataPipelinesResponseBodyPipelines {
	s.StatusMessage = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetUpdateTime(v string) *ListDataPipelinesResponseBodyPipelines {
	s.UpdateTime = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) SetVersion(v int64) *ListDataPipelinesResponseBodyPipelines {
	s.Version = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelines) Validate() error {
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

type ListDataPipelinesResponseBodyPipelinesOutputs struct {
	// The output name.
	//
	// example:
	//
	// checkout_route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The branch processors.
	Processors []*ListDataPipelinesResponseBodyPipelinesOutputsProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s ListDataPipelinesResponseBodyPipelinesOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputs) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputs) GetName() *string {
	return s.Name
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputs) GetProcessors() []*ListDataPipelinesResponseBodyPipelinesOutputsProcessors {
	return s.Processors
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputs) SetName(v string) *ListDataPipelinesResponseBodyPipelinesOutputs {
	s.Name = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputs) SetProcessors(v []*ListDataPipelinesResponseBodyPipelinesOutputsProcessors) *ListDataPipelinesResponseBodyPipelinesOutputs {
	s.Processors = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputs) Validate() error {
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

type ListDataPipelinesResponseBodyPipelinesOutputsProcessors struct {
	// The processor configuration.
	Config *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the processor.
	//
	// example:
	//
	// drop-health-check
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the processor.
	//
	// example:
	//
	// filter
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessors) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessors) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) GetConfig() *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	return s.Config
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) GetName() *string {
	return s.Name
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) GetType() *string {
	return s.Type
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) SetConfig(v *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) *ListDataPipelinesResponseBodyPipelinesOutputsProcessors {
	s.Config = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) SetName(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessors {
	s.Name = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) SetType(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessors {
	s.Type = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The list of fields.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetRules() []*ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	return s.Rules
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetSelector() *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector {
	return s.Selector
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) GetTarget() *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget {
	return s.Target
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetApplications(v []*string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Applications = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetExpression(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetFields(v []*string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Fields = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetRules(v []*ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Rules = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetScript(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Script = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetSelector(v *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Selector = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) SetTarget(v *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig {
	s.Target = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfig) Validate() error {
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

type ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules struct {
	// The length of the prefix to keep.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The length of the suffix to keep.
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

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetKeepPrefix(v int32) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetKeepSuffix(v int32) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetKeys(v []*string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetMaskChar(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetMode(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) SetTypes(v []*string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) SetServiceNames(v []*string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) SetWorkspace(v string) *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesOutputsProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesProcessors struct {
	// The processor configuration.
	Config *ListDataPipelinesResponseBodyPipelinesProcessorsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the processor.
	//
	// example:
	//
	// drop-health-check
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the processor.
	//
	// example:
	//
	// filter
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataPipelinesResponseBodyPipelinesProcessors) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesProcessors) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) GetConfig() *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	return s.Config
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) GetName() *string {
	return s.Name
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) GetType() *string {
	return s.Type
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) SetConfig(v *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) *ListDataPipelinesResponseBodyPipelinesProcessors {
	s.Config = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) SetName(v string) *ListDataPipelinesResponseBodyPipelinesProcessors {
	s.Name = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) SetType(v string) *ListDataPipelinesResponseBodyPipelinesProcessors {
	s.Type = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessors) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataPipelinesResponseBodyPipelinesProcessorsConfig struct {
	// The list of applications.
	Applications []*string `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// The filter expression.
	//
	// example:
	//
	// attributes["http.route"] != "/health"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The list of fields.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The list of masking rules.
	Rules []*ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	// The SPL script.
	//
	// example:
	//
	// 	- | extend latency_ms = duration / 1000000
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
	// The service selector.
	Selector *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector `json:"selector,omitempty" xml:"selector,omitempty" type:"Struct"`
	// The processing target.
	Target *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetApplications() []*string {
	return s.Applications
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetExpression() *string {
	return s.Expression
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetFields() []*string {
	return s.Fields
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetRules() []*ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	return s.Rules
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetScript() *string {
	return s.Script
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetSelector() *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector {
	return s.Selector
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) GetTarget() *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget {
	return s.Target
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetApplications(v []*string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Applications = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetExpression(v string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Expression = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetFields(v []*string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Fields = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetRules(v []*ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Rules = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetScript(v string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Script = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetSelector(v *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Selector = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) SetTarget(v *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) *ListDataPipelinesResponseBodyPipelinesProcessorsConfig {
	s.Target = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfig) Validate() error {
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

type ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules struct {
	// The length of the prefix to keep.
	//
	// example:
	//
	// 2
	KeepPrefix *int32 `json:"keepPrefix,omitempty" xml:"keepPrefix,omitempty"`
	// The length of the suffix to keep.
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

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetKeepPrefix() *int32 {
	return s.KeepPrefix
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetKeepSuffix() *int32 {
	return s.KeepSuffix
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetKeys() []*string {
	return s.Keys
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetMaskChar() *string {
	return s.MaskChar
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetMode() *string {
	return s.Mode
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) GetTypes() []*string {
	return s.Types
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetKeepPrefix(v int32) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.KeepPrefix = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetKeepSuffix(v int32) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.KeepSuffix = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetKeys(v []*string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.Keys = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetMaskChar(v string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.MaskChar = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetMode(v string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.Mode = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) SetTypes(v []*string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules {
	s.Types = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigRules) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector struct {
	// The list of service names.
	//
	// example:
	//
	// ["checkout-*","order-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) SetServiceNames(v []*string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector {
	s.ServiceNames = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigSelector) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget struct {
	// The target workspace.
	//
	// example:
	//
	// target-checkout-ws
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) SetWorkspace(v string) *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget {
	s.Workspace = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesProcessorsConfigTarget) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesSinks struct {
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

func (s ListDataPipelinesResponseBodyPipelinesSinks) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesSinks) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) GetDatasets() []*string {
	return s.Datasets
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) GetLogstore() *string {
	return s.Logstore
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) GetName() *string {
	return s.Name
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) GetProject() *string {
	return s.Project
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) GetType() *string {
	return s.Type
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) SetDatasets(v []*string) *ListDataPipelinesResponseBodyPipelinesSinks {
	s.Datasets = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) SetLogstore(v string) *ListDataPipelinesResponseBodyPipelinesSinks {
	s.Logstore = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) SetName(v string) *ListDataPipelinesResponseBodyPipelinesSinks {
	s.Name = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) SetProject(v string) *ListDataPipelinesResponseBodyPipelinesSinks {
	s.Project = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) SetType(v string) *ListDataPipelinesResponseBodyPipelinesSinks {
	s.Type = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSinks) Validate() error {
	return dara.Validate(s)
}

type ListDataPipelinesResponseBodyPipelinesSource struct {
	// The datasource config.
	Config *ListDataPipelinesResponseBodyPipelinesSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The type of the data source.
	//
	// example:
	//
	// traces-default
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataPipelinesResponseBodyPipelinesSource) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesSource) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesSource) GetConfig() *ListDataPipelinesResponseBodyPipelinesSourceConfig {
	return s.Config
}

func (s *ListDataPipelinesResponseBodyPipelinesSource) GetType() *string {
	return s.Type
}

func (s *ListDataPipelinesResponseBodyPipelinesSource) SetConfig(v *ListDataPipelinesResponseBodyPipelinesSourceConfig) *ListDataPipelinesResponseBodyPipelinesSource {
	s.Config = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSource) SetType(v string) *ListDataPipelinesResponseBodyPipelinesSource {
	s.Type = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSource) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataPipelinesResponseBodyPipelinesSourceConfig struct {
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
	TimeRange *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s ListDataPipelinesResponseBodyPipelinesSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesSourceConfig) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) GetRunMode() *string {
	return s.RunMode
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) GetStartFrom() *string {
	return s.StartFrom
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) GetTimeRange() *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange {
	return s.TimeRange
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) SetRunMode(v string) *ListDataPipelinesResponseBodyPipelinesSourceConfig {
	s.RunMode = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) SetStartFrom(v string) *ListDataPipelinesResponseBodyPipelinesSourceConfig {
	s.StartFrom = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) SetTimeRange(v *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) *ListDataPipelinesResponseBodyPipelinesSourceConfig {
	s.TimeRange = v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfig) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange struct {
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

func (s ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) GetFrom() *int64 {
	return s.From
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) GetTo() *int64 {
	return s.To
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) SetFrom(v int64) *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange {
	s.From = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) SetTo(v int64) *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange {
	s.To = &v
	return s
}

func (s *ListDataPipelinesResponseBodyPipelinesSourceConfigTimeRange) Validate() error {
	return dara.Validate(s)
}
