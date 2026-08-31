// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *CreatePipelineRequestContext) *CreatePipelineRequest
	GetContext() *CreatePipelineRequestContext
	SetCreateCommand(v *CreatePipelineRequestCreateCommand) *CreatePipelineRequest
	GetCreateCommand() *CreatePipelineRequestCreateCommand
	SetOpTenantId(v int64) *CreatePipelineRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreatePipelineRequest
	GetOpUserId() *string
}

type CreatePipelineRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *CreatePipelineRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The configuration for creating a pipeline or workflow node.
	//
	// This parameter is required.
	CreateCommand *CreatePipelineRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) GetContext() *CreatePipelineRequestContext {
	return s.Context
}

func (s *CreatePipelineRequest) GetCreateCommand() *CreatePipelineRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreatePipelineRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreatePipelineRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreatePipelineRequest) SetContext(v *CreatePipelineRequestContext) *CreatePipelineRequest {
	s.Context = v
	return s
}

func (s *CreatePipelineRequest) SetCreateCommand(v *CreatePipelineRequestCreateCommand) *CreatePipelineRequest {
	s.CreateCommand = v
	return s
}

func (s *CreatePipelineRequest) SetOpTenantId(v int64) *CreatePipelineRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreatePipelineRequest) SetOpUserId(v string) *CreatePipelineRequest {
	s.OpUserId = &v
	return s
}

func (s *CreatePipelineRequest) Validate() error {
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

type CreatePipelineRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: the development environment.
	//
	// - PROD: the production environment.
	//
	// For workflow nodes, only PROD is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the integration pipeline or workflow node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreatePipelineRequestContext) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestContext) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestContext) GetEnv() *string {
	return s.Env
}

func (s *CreatePipelineRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineRequestContext) SetEnv(v string) *CreatePipelineRequestContext {
	s.Env = &v
	return s
}

func (s *CreatePipelineRequestContext) SetProjectId(v int64) *CreatePipelineRequestContext {
	s.ProjectId = &v
	return s
}

func (s *CreatePipelineRequestContext) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestCreateCommand struct {
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The integration pipeline configuration mode. Valid values:
	//
	// - PIPELINE: pipeline mode. This is the default value.
	//
	// - JSON: script mode.
	//
	// If the node is a workflow node, you can ignore this parameter.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The basic information of the integration pipeline or workflow node.
	//
	// This parameter is required.
	NodeInfo *CreatePipelineRequestCreateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The integration pipeline component or workflow operator configuration.
	//
	// This parameter is required.
	PipelineConfig *CreatePipelineRequestCreateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// The integration pipeline configuration in JSON string format for script mode. Workflow nodes do not support script mode.
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// The node type. Valid values:
	//
	// - 0: offline integration. Default value: 0.
	//
	// - 1: real-time integration.
	//
	// - 14: offline workflow node.
	//
	// - 15: real-time workflow.
	//
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// The scheduling configuration in JSON string format. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig#toJsonString method.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The channel configuration in JSON string format. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting#toJsonString method.
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// Specifies whether to submit the node. The node is submitted by default.
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
}

func (s CreatePipelineRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestCreateCommand) GetComment() *string {
	return s.Comment
}

func (s *CreatePipelineRequestCreateCommand) GetMode() *string {
	return s.Mode
}

func (s *CreatePipelineRequestCreateCommand) GetNodeInfo() *CreatePipelineRequestCreateCommandNodeInfo {
	return s.NodeInfo
}

func (s *CreatePipelineRequestCreateCommand) GetPipelineConfig() *CreatePipelineRequestCreateCommandPipelineConfig {
	return s.PipelineConfig
}

func (s *CreatePipelineRequestCreateCommand) GetPipelineJson() *string {
	return s.PipelineJson
}

func (s *CreatePipelineRequestCreateCommand) GetPipelineType() *int32 {
	return s.PipelineType
}

func (s *CreatePipelineRequestCreateCommand) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *CreatePipelineRequestCreateCommand) GetSettings() *string {
	return s.Settings
}

func (s *CreatePipelineRequestCreateCommand) GetSubmit() *bool {
	return s.Submit
}

func (s *CreatePipelineRequestCreateCommand) SetComment(v string) *CreatePipelineRequestCreateCommand {
	s.Comment = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetMode(v string) *CreatePipelineRequestCreateCommand {
	s.Mode = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetNodeInfo(v *CreatePipelineRequestCreateCommandNodeInfo) *CreatePipelineRequestCreateCommand {
	s.NodeInfo = v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetPipelineConfig(v *CreatePipelineRequestCreateCommandPipelineConfig) *CreatePipelineRequestCreateCommand {
	s.PipelineConfig = v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetPipelineJson(v string) *CreatePipelineRequestCreateCommand {
	s.PipelineJson = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetPipelineType(v int32) *CreatePipelineRequestCreateCommand {
	s.PipelineType = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetScheduleConfig(v string) *CreatePipelineRequestCreateCommand {
	s.ScheduleConfig = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetSettings(v string) *CreatePipelineRequestCreateCommand {
	s.Settings = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) SetSubmit(v bool) *CreatePipelineRequestCreateCommand {
	s.Submit = &v
	return s
}

func (s *CreatePipelineRequestCreateCommand) Validate() error {
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

type CreatePipelineRequestCreateCommandNodeInfo struct {
	// The folder of the integration pipeline or workflow node. Default value: root folder. The folder must already exist. If it does not exist, call the relevant API operation to create a folder of type offlinePipeline (or unstructuredPipeline for workflows).
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The pipeline or workflow file ID. Leave this parameter empty for initial creation. When updating a pipeline or workflow node, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The scheduling node ID of the pipeline or workflow node. Leave this parameter empty for initial creation. When updating a pipeline or workflow node, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the integration pipeline or workflow node.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The pipeline or workflow node ID. Leave this parameter empty for initial creation. When updating a pipeline or workflow node, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s CreatePipelineRequestCreateCommandNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestCreateCommandNodeInfo) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) GetDirectory() *string {
	return s.Directory
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) SetDirectory(v string) *CreatePipelineRequestCreateCommandNodeInfo {
	s.Directory = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) SetFileId(v int64) *CreatePipelineRequestCreateCommandNodeInfo {
	s.FileId = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) SetNodeId(v string) *CreatePipelineRequestCreateCommandNodeInfo {
	s.NodeId = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) SetNodeName(v string) *CreatePipelineRequestCreateCommandNodeInfo {
	s.NodeName = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) SetPipelineId(v int64) *CreatePipelineRequestCreateCommandNodeInfo {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandNodeInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestCreateCommandPipelineConfig struct {
	// The directed acyclic graph (DAG) link configuration that describes the connections between all components or operators.
	//
	// This parameter is required.
	Hops []*CreatePipelineRequestCreateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The component or operator configurations, which contain the detailed configurations of all components or operators used.
	//
	// This parameter is required.
	Steps []*CreatePipelineRequestCreateCommandPipelineConfigSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
}

func (s CreatePipelineRequestCreateCommandPipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestCreateCommandPipelineConfig) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestCreateCommandPipelineConfig) GetHops() []*CreatePipelineRequestCreateCommandPipelineConfigHops {
	return s.Hops
}

func (s *CreatePipelineRequestCreateCommandPipelineConfig) GetSteps() []*CreatePipelineRequestCreateCommandPipelineConfigSteps {
	return s.Steps
}

func (s *CreatePipelineRequestCreateCommandPipelineConfig) SetHops(v []*CreatePipelineRequestCreateCommandPipelineConfigHops) *CreatePipelineRequestCreateCommandPipelineConfig {
	s.Hops = v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfig) SetSteps(v []*CreatePipelineRequestCreateCommandPipelineConfigSteps) *CreatePipelineRequestCreateCommandPipelineConfig {
	s.Steps = v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfig) Validate() error {
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

type CreatePipelineRequestCreateCommandPipelineConfigHops struct {
	// Specifies the downstream condition for a conditional distribution component. Set this parameter to true if the downstream condition is true, or false otherwise. If the node is a workflow node, you can ignore this parameter.
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

func (s CreatePipelineRequestCreateCommandPipelineConfigHops) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestCreateCommandPipelineConfigHops) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) GetSendTo() *bool {
	return s.SendTo
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) GetSource() *string {
	return s.Source
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) GetTarget() *string {
	return s.Target
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) SetSendTo(v bool) *CreatePipelineRequestCreateCommandPipelineConfigHops {
	s.SendTo = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) SetSource(v string) *CreatePipelineRequestCreateCommandPipelineConfigHops {
	s.Source = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) SetTarget(v string) *CreatePipelineRequestCreateCommandPipelineConfigHops {
	s.Target = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigHops) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestCreateCommandPipelineConfigSteps struct {
	// Specifies the data distribution mode when the current component has multiple downstream components. Valid values:
	//
	// - true: The data of the current component is distributed to all downstream components in a round-robin manner. For example, if the current component has 100 records and two downstream components, each downstream component receives 50 records. Default value: true.
	//
	// - false: The full data of the current component is sent to all downstream components. For example, if the current component has 100 records and two downstream components, both downstream components receive 100 records.
	//
	// If the node is a workflow node, you can ignore this parameter.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// The plugin ID. Each plugin or operator has a unique identifier. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepKey. Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the configuration created on the Dataphin console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The specific component configuration in JSON string format. Refer to the toJsonString method of the relevant subclasses of the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig (for workflow operators, use com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.unstructured.BaseOAUnstructuredNeuronConfig). Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the node configuration created on the Dataphin console.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// The step name. Step names must be unique within the same pipeline node.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The component type. Valid values:
	//
	// - input: an input component.
	//
	// - output: an output component.
	//
	// - transfrom: a transform component.
	//
	// - process: a flow control component.
	//
	// For workflow nodes, this parameter specifies the operator type, such as image or text. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepType. Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the configuration created on the Dataphin console.
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s CreatePipelineRequestCreateCommandPipelineConfigSteps) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestCreateCommandPipelineConfigSteps) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) GetIsDistribute() *bool {
	return s.IsDistribute
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) GetKey() *string {
	return s.Key
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) GetStepName() *string {
	return s.StepName
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) GetStepType() *string {
	return s.StepType
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) SetIsDistribute(v bool) *CreatePipelineRequestCreateCommandPipelineConfigSteps {
	s.IsDistribute = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) SetKey(v string) *CreatePipelineRequestCreateCommandPipelineConfigSteps {
	s.Key = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) SetPluginConfig(v string) *CreatePipelineRequestCreateCommandPipelineConfigSteps {
	s.PluginConfig = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) SetStepName(v string) *CreatePipelineRequestCreateCommandPipelineConfigSteps {
	s.StepName = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) SetStepType(v string) *CreatePipelineRequestCreateCommandPipelineConfigSteps {
	s.StepType = &v
	return s
}

func (s *CreatePipelineRequestCreateCommandPipelineConfigSteps) Validate() error {
	return dara.Validate(s)
}
