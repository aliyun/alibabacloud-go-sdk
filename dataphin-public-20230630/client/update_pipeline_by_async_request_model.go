// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineByAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *UpdatePipelineByAsyncRequestContext) *UpdatePipelineByAsyncRequest
	GetContext() *UpdatePipelineByAsyncRequestContext
	SetOpTenantId(v int64) *UpdatePipelineByAsyncRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdatePipelineByAsyncRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdatePipelineByAsyncRequestUpdateCommand) *UpdatePipelineByAsyncRequest
	GetUpdateCommand() *UpdatePipelineByAsyncRequestUpdateCommand
}

type UpdatePipelineByAsyncRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *UpdatePipelineByAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update configuration for the pipeline task.
	//
	// This parameter is required.
	UpdateCommand *UpdatePipelineByAsyncRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdatePipelineByAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequest) GetContext() *UpdatePipelineByAsyncRequestContext {
	return s.Context
}

func (s *UpdatePipelineByAsyncRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdatePipelineByAsyncRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdatePipelineByAsyncRequest) GetUpdateCommand() *UpdatePipelineByAsyncRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdatePipelineByAsyncRequest) SetContext(v *UpdatePipelineByAsyncRequestContext) *UpdatePipelineByAsyncRequest {
	s.Context = v
	return s
}

func (s *UpdatePipelineByAsyncRequest) SetOpTenantId(v int64) *UpdatePipelineByAsyncRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequest) SetOpUserId(v string) *UpdatePipelineByAsyncRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequest) SetUpdateCommand(v *UpdatePipelineByAsyncRequestUpdateCommand) *UpdatePipelineByAsyncRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdatePipelineByAsyncRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineByAsyncRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD: production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the integration pipeline task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdatePipelineByAsyncRequestContext) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestContext) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestContext) GetEnv() *string {
	return s.Env
}

func (s *UpdatePipelineByAsyncRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdatePipelineByAsyncRequestContext) SetEnv(v string) *UpdatePipelineByAsyncRequestContext {
	s.Env = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestContext) SetProjectId(v int64) *UpdatePipelineByAsyncRequestContext {
	s.ProjectId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestContext) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineByAsyncRequestUpdateCommand struct {
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The configuration mode of the integration pipeline. Valid values:
	//
	// - PIPELINE: pipeline mode (default).
	//
	// - JSON: script mode.
	//
	// Ignore this parameter for workflow tasks.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The basic information of the integration pipeline task.
	//
	// This parameter is required.
	NodeInfo *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The component configuration for the integration pipeline.
	//
	// This parameter is required.
	PipelineConfig *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// The integration pipeline configuration in JSON string format for script mode. Workflow tasks do not support script mode.
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// The task type. Valid values:
	//
	// - 0: offline integration (default).
	//
	// - 1: real-time integration.
	//
	// - 14: workflow task.
	//
	// - 15: real-time workflow.
	//
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// The schedule configuration in JSON string format. Refer to the com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig#toJsonString method in the utility class.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The channel settings in JSON string format. Refer to the com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting#toJsonString method in the utility class.
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// Specifies whether to submit the task. Default value: true.
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
}

func (s UpdatePipelineByAsyncRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetComment() *string {
	return s.Comment
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetMode() *string {
	return s.Mode
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetNodeInfo() *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	return s.NodeInfo
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetPipelineConfig() *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig {
	return s.PipelineConfig
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetPipelineJson() *string {
	return s.PipelineJson
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetPipelineType() *int32 {
	return s.PipelineType
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetSettings() *string {
	return s.Settings
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) GetSubmit() *bool {
	return s.Submit
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetComment(v string) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.Comment = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetMode(v string) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.Mode = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetNodeInfo(v *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.NodeInfo = v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetPipelineConfig(v *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.PipelineConfig = v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetPipelineJson(v string) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.PipelineJson = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetPipelineType(v int32) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.PipelineType = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetScheduleConfig(v string) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.ScheduleConfig = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetSettings(v string) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.Settings = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) SetSubmit(v bool) *UpdatePipelineByAsyncRequestUpdateCommand {
	s.Submit = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommand) Validate() error {
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

type UpdatePipelineByAsyncRequestUpdateCommandNodeInfo struct {
	// The directory of the integration pipeline task node. The default value is the root directory. The directory must exist. If it does not exist, call the relevant API to create a directory of the offlinePipeline type.
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The ID of the pipeline file. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the scheduling node for the pipeline task. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the integration pipeline task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the pipeline task. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GetDirectory() *string {
	return s.Directory
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) SetDirectory(v string) *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	s.Directory = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) SetFileId(v int64) *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	s.FileId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) SetNodeId(v string) *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	s.NodeId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) SetNodeName(v string) *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	s.NodeName = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) SetPipelineId(v int64) *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandNodeInfo) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig struct {
	// The directed acyclic graph (DAG) hop configurations that describe the connection relationships among all components.
	//
	// This parameter is required.
	Hops []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The step configurations, including the detailed configurations of all used components.
	//
	// This parameter is required.
	Steps []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) GetHops() []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops {
	return s.Hops
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) GetSteps() []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	return s.Steps
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) SetHops(v []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig {
	s.Hops = v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) SetSteps(v []*UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig {
	s.Steps = v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfig) Validate() error {
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

type UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops struct {
	// Specifies whether to connect to the downstream component when the condition is met for a conditional distribution component. Set to true if the downstream condition is true, and false otherwise. Ignore this parameter for workflow tasks.
	SendTo *bool `json:"SendTo,omitempty" xml:"SendTo,omitempty"`
	// The name of the input step, which corresponds to Steps[*].StepName.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the output step, which corresponds to Steps[*].StepName.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps_writer
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) GetSendTo() *bool {
	return s.SendTo
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) GetSource() *string {
	return s.Source
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) GetTarget() *string {
	return s.Target
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) SetSendTo(v bool) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops {
	s.SendTo = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) SetSource(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops {
	s.Source = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) SetTarget(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops {
	s.Target = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigHops) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps struct {
	// Specifies the data distribution method when the current component has multiple downstream components. Valid values:
	//
	// - true: distributes data to all downstream components in turn. For example, if the current component has 100 data records and two downstream components, each downstream component receives 50 records. This is the default value.
	//
	// - false: sends all data to all downstream components. For example, if the current component has 100 data records and two downstream components, both downstream components receive all 100 records.
	//
	// Ignore this parameter for workflow tasks.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// The plugin ID. Each plugin or operator has a unique identifier. Refer to com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepKey. Developers must inherit the component or operator configuration class to implement the corresponding configuration. The structure of each configuration must be identical to the configuration structure created on the Dataphin page.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The specific component or operator configuration in JSON string format. Refer to the toJsonString method of the relevant subclasses of com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig (or com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.unstructured.BaseOAUnstructuredNeuronConfig for workflow operators). Developers must inherit this component or operator configuration class to implement the corresponding configuration. The structure of each component or operator configuration must be identical to the task configuration structure created on the Dataphin page.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// The step name. Step names must be unique within the same pipeline task.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The step type. For pipeline tasks, valid values:
	//
	// - input: input component.
	//
	// - output: output component.
	//
	// - transfrom: transform component.
	//
	// - process: flow control component.
	//
	// For workflow tasks, this indicates the operator type, such as image: image or text: text. Refer to com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepType. Developers must inherit the component or operator configuration class to implement the corresponding configuration. The structure of each configuration must be identical to the configuration structure created on the Dataphin page.
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GetIsDistribute() *bool {
	return s.IsDistribute
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GetKey() *string {
	return s.Key
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GetStepName() *string {
	return s.StepName
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) GetStepType() *string {
	return s.StepType
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) SetIsDistribute(v bool) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	s.IsDistribute = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) SetKey(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	s.Key = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) SetPluginConfig(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	s.PluginConfig = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) SetStepName(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	s.StepName = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) SetStepType(v string) *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps {
	s.StepType = &v
	return s
}

func (s *UpdatePipelineByAsyncRequestUpdateCommandPipelineConfigSteps) Validate() error {
	return dara.Validate(s)
}
