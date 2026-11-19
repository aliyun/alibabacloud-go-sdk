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
	SetOpUserId(v string) *CreatePipelineByAsyncRequest
	GetOpUserId() *string
}

type CreatePipelineByAsyncRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *CreatePipelineByAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The configuration for creating a pipeline or workflow task.
	//
	// This parameter is required.
	CreateCommand *CreatePipelineByAsyncRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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

func (s *CreatePipelineByAsyncRequest) GetOpUserId() *string {
	return s.OpUserId
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

func (s *CreatePipelineByAsyncRequest) SetOpUserId(v string) *CreatePipelineByAsyncRequest {
	s.OpUserId = &v
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
	// The current operation environment. Valid values: DEV: development environment. PROD: production environment. For workflows, only PROD is currently supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the integration pipeline or workflow task belongs.
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
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The integration pipeline configuration mode. Valid values: PIPELINE: pipeline mode (default). JSON: script mode. Ignore this parameter for workflows.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The basic information of the integration pipeline task.
	//
	// This parameter is required.
	NodeInfo *CreatePipelineByAsyncRequestCreateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The integration pipeline component or workflow operator configuration.
	//
	// This parameter is required.
	PipelineConfig *CreatePipelineByAsyncRequestCreateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// The integration pipeline configuration in JSON string format for script mode. Workflow tasks do not support script mode.
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// The task type. Valid values: 0: offline integration (default). 1: real-time integration. 14: offline workflow task. 15: real-time workflow.
	//
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// The schedule configuration in JSON string format. Refer to the toJsonString method in the com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig utility class.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The channel configuration in JSON string format. Refer to the toJsonString method in the com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting utility class.
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// Specifies whether to submit the task. By default, the task is submitted.
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
	// The node directory of the integration pipeline task. Default value: root directory. The directory must exist. If it does not exist, call the relevant API operation to create a directory of the offlinePipeline type.
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The ID of the pipeline file. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must configure at least one of the following parameters: pipelineId, fileId, and nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the schedule node for the pipeline task. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must configure at least one of the following parameters: pipelineId, fileId, and nodeId.
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
	// The ID of the pipeline task. Leave this parameter empty when creating a task for the first time. When updating a pipeline task, you must configure at least one of the following parameters: pipelineId, fileId, and nodeId.
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
	// The directed acyclic graph (DAG) link configuration that describes the connection relationships among all components.
	//
	// This parameter is required.
	Hops []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The component or operator configurations, including the detailed configurations of all used components or operators.
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
	// For conditional distribution components, set this parameter to true if the downstream connection condition is true. Otherwise, set it to false. Ignore this parameter for workflow tasks.
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
	// The data distribution method when the current component has multiple downstream components. Valid values: true: distributes data from the current component to all downstream components in turn. For example, if the current component has 100 data records and two downstream components, each downstream component receives 50 data records. Default value: true. false: sends all data from the current component to all downstream components. For example, if the current component has 100 data records and two downstream components, both downstream components receive all 100 data records. Ignore this parameter for workflow tasks.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// The plugin ID. Each plugin or operator has a unique identifier. Refer to the stepKey field in the com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig utility class. Developers must inherit this component or operator configuration class to implement the corresponding configuration. The structure of each component or operator configuration must be identical to the configuration structure created on the Dataphin page.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The specific component configuration in JSON string format. Refer to the toJsonString method of the relevant subclasses of com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig (for workflow operators, use com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.unstructured.BaseOAUnstructuredNeuronConfig). Developers must inherit this component or operator configuration class to implement the corresponding configuration. The structure of each component or operator configuration must be identical to the task configuration structure created on the Dataphin page.
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
	// The component type. Valid values: input: input component. output: output component. transfrom: transform component. process: flow control component. For workflow tasks, this indicates the operator type, such as image for images and text for text. Refer to the stepType field in the com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig utility class. Developers must inherit this component or operator configuration class to implement the corresponding configuration. The structure of each component or operator configuration must be identical to the configuration structure created on the Dataphin page.
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
