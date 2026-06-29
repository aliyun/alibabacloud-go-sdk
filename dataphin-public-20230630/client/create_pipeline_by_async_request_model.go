// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineByAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *CreatePipelineByAsyncRequestContext) *CreatePipelineByAsyncRequest
	GetContext() *CreatePipelineByAsyncRequestContext
	SetCreateCommand(v *CreatePipelineByAsyncRequestCreateCommand) *CreatePipelineByAsyncRequest
	GetCreateCommand() *CreatePipelineByAsyncRequestCreateCommand
	SetOpTenantId(v int64) *CreatePipelineByAsyncRequest
	GetOpTenantId() *int64
}

type CreatePipelineByAsyncRequest struct {
	// Request context information
	//
	// This parameter is required.
	Context *CreatePipelineByAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// Create pipeline/workflow task configuration
	//
	// This parameter is required.
	CreateCommand *CreatePipelineByAsyncRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreatePipelineByAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequest) GetContext() *CreatePipelineByAsyncRequestContext {
	return s.Context
}

func (s *CreatePipelineByAsyncRequest) GetCreateCommand() *CreatePipelineByAsyncRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreatePipelineByAsyncRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineByAsyncRequest) SetContext(v *CreatePipelineByAsyncRequestContext) *CreatePipelineByAsyncRequest {
	s.Context = v
	return s
}

func (s *CreatePipelineByAsyncRequest) SetCreateCommand(v *CreatePipelineByAsyncRequestCreateCommand) *CreatePipelineByAsyncRequest {
	s.CreateCommand = v
	return s
}

func (s *CreatePipelineByAsyncRequest) SetOpTenantId(v int64) *CreatePipelineByAsyncRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineByAsyncRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineByAsyncRequestContext struct {
	// Current operating environment env: DEV - indicates the development environment, PROD - indicates the production environment (for workflows, only PROD is currently supported)
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Project ID to which the integration pipeline/workflow task belongs
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreatePipelineByAsyncRequestContext) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestContext) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestContext) GetEnv() *string {
	return s.Env
}

func (s *CreatePipelineByAsyncRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineByAsyncRequestContext) SetEnv(v string) *CreatePipelineByAsyncRequestContext {
	s.Env = &v
	return s
}

func (s *CreatePipelineByAsyncRequestContext) SetProjectId(v int64) *CreatePipelineByAsyncRequestContext {
	s.ProjectId = &v
	return s
}

func (s *CreatePipelineByAsyncRequestContext) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineByAsyncRequestCreateCommand struct {
	// Comment
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Integration pipeline configuration mode: PIPELINE - indicates pipeline mode (default), JSON - indicates script mode. This is not applicable to workflows.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Integration pipeline task basic information
	//
	// This parameter is required.
	NodeInfo *CreatePipelineByAsyncRequestCreateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// Integration pipeline component/workflow operator configuration
	//
	// This parameter is required.
	PipelineConfig *CreatePipelineByAsyncRequestCreateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// In script mode: integration pipeline configuration (in JSON string format). Workflow tasks do not support script mode
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// Task type: 0 - indicates offline integration (default), 1 - indicates real-time integration, 14 - indicates a workflow task
	//
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// Scheduling configuration in JSON string format. Refer to the utility class: com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig#toJsonString method
	//
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// Channel configuration in JSON string format. Refer to the utility class: com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting#toJsonString method
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// Whether to submit. The default is to submit
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
}

func (s CreatePipelineByAsyncRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetComment() *string {
	return s.Comment
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetMode() *string {
	return s.Mode
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetNodeInfo() *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	return s.NodeInfo
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetPipelineConfig() *CreatePipelineByAsyncRequestCreateCommandPipelineConfig {
	return s.PipelineConfig
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetPipelineJson() *string {
	return s.PipelineJson
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetPipelineType() *int32 {
	return s.PipelineType
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetSettings() *string {
	return s.Settings
}

func (s *CreatePipelineByAsyncRequestCreateCommand) GetSubmit() *bool {
	return s.Submit
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetComment(v string) *CreatePipelineByAsyncRequestCreateCommand {
	s.Comment = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetMode(v string) *CreatePipelineByAsyncRequestCreateCommand {
	s.Mode = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetNodeInfo(v *CreatePipelineByAsyncRequestCreateCommandNodeInfo) *CreatePipelineByAsyncRequestCreateCommand {
	s.NodeInfo = v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetPipelineConfig(v *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) *CreatePipelineByAsyncRequestCreateCommand {
	s.PipelineConfig = v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetPipelineJson(v string) *CreatePipelineByAsyncRequestCreateCommand {
	s.PipelineJson = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetPipelineType(v int32) *CreatePipelineByAsyncRequestCreateCommand {
	s.PipelineType = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetScheduleConfig(v string) *CreatePipelineByAsyncRequestCreateCommand {
	s.ScheduleConfig = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetSettings(v string) *CreatePipelineByAsyncRequestCreateCommand {
	s.Settings = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) SetSubmit(v bool) *CreatePipelineByAsyncRequestCreateCommand {
	s.Submit = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommand) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PipelineConfig != nil {
		if err := s.PipelineConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineByAsyncRequestCreateCommandNodeInfo struct {
	// Integration pipeline task node directory (defaults to the root directory). The directory must exist. If it does not exist, call the relevant API to create a directory of type offlinePipeline
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Pipeline file ID. Leave empty for the first creation. When updating a pipeline task, at least one of pipelineId, fileId, or nodeId must be configured
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Pipeline task scheduling node ID. Leave empty for the first creation. When updating a pipeline task, at least one of pipelineId, fileId, or nodeId must be configured
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Integration pipeline task name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Pipeline task ID. Leave empty for the first creation. When updating a pipeline task, at least one of pipelineId, fileId, or nodeId must be configured
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s CreatePipelineByAsyncRequestCreateCommandNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestCreateCommandNodeInfo) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) GetDirectory() *string {
	return s.Directory
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) SetDirectory(v string) *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	s.Directory = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) SetFileId(v int64) *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	s.FileId = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) SetNodeId(v string) *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	s.NodeId = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) SetNodeName(v string) *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	s.NodeName = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) SetPipelineId(v int64) *CreatePipelineByAsyncRequestCreateCommandNodeInfo {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandNodeInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineByAsyncRequestCreateCommandPipelineConfig struct {
	// DAG (Directed Acyclic Graph) link configuration: describes the connection relationships of all components
	//
	// This parameter is required.
	Hops []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// Component/operator configuration: contains detailed configurations of all components/operators used
	//
	// This parameter is required.
	Steps []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfig) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) GetHops() []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops {
	return s.Hops
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) GetSteps() []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	return s.Steps
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) SetHops(v []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) *CreatePipelineByAsyncRequestCreateCommandPipelineConfig {
	s.Hops = v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) SetSteps(v []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) *CreatePipelineByAsyncRequestCreateCommandPipelineConfig {
	s.Steps = v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfig) Validate() error {
	if s.Hops != nil {
		for _, item := range s.Hops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops struct {
	// For conditional distribution components, set to true when the downstream connection condition is true, otherwise set to false. This is not applicable to workflow tasks.
	SendTo *bool `json:"SendTo,omitempty" xml:"SendTo,omitempty"`
	// Input step name, i.e., Steps[*].StepName
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Output step name, i.e., Steps[*].StepName
	//
	// This parameter is required.
	//
	// example:
	//
	// odps_writer
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) GetSendTo() *bool {
	return s.SendTo
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) GetSource() *string {
	return s.Source
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) GetTarget() *string {
	return s.Target
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) SetSendTo(v bool) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops {
	s.SendTo = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) SetSource(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops {
	s.Source = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) SetTarget(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops {
	s.Target = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps struct {
	// Indicates the data distribution method when the current component has multiple downstream components: true - indicates that the data of the current component is sent to all downstream components in a round-robin manner. For example, if the current component has 100 records and two downstream components, each downstream component receives 50 records. The default value is true. false - indicates that the data of the current component is sent in full to all downstream components. For example, if the current component has 100 records and two downstream components, both downstream components receive 100 records. This value is not applicable to workflow tasks.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// Plugin ID. Each plugin/operator has a unique identifier. Refer to the utility class: com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepKey. Developers should extend the component/operator configuration class to implement the corresponding component/operator configuration. Each component/operator configuration has the same structure as a configuration created on the Dataphin page
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specific component configuration in JSON string format. Refer to the utility class: subclasses of com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig (for workflow operators, use com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.unstructured.BaseOAUnstructuredNeuronConfig) and their toJsonString methods. Developers should extend the component/operator configuration class to implement the corresponding component/operator configuration. Each component/operator configuration has the same structure as a task configuration created on the Dataphin page
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// Step name. Step names must be unique within the same pipeline task
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// Component type: input - indicates an input component, output - indicates an output component, transform - indicates a transform component, process - indicates a flow control component. For workflow tasks, it indicates the operator type, for example: image - image, text - text. Refer to the utility class: com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepType. Developers should extend the component/operator configuration class to implement the corresponding component/operator configuration. Each component/operator configuration has the same structure as a configuration created on the Dataphin page
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GetIsDistribute() *bool {
	return s.IsDistribute
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GetKey() *string {
	return s.Key
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GetStepName() *string {
	return s.StepName
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) GetStepType() *string {
	return s.StepType
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) SetIsDistribute(v bool) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	s.IsDistribute = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) SetKey(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	s.Key = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) SetPluginConfig(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	s.PluginConfig = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) SetStepName(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	s.StepName = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) SetStepType(v string) *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps {
	s.StepType = &v
	return s
}

func (s *CreatePipelineByAsyncRequestCreateCommandPipelineConfigSteps) Validate() error {
	return dara.Validate(s)
}
