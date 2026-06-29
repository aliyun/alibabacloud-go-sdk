// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *UpdatePipelineRequestContext) *UpdatePipelineRequest
	GetContext() *UpdatePipelineRequestContext
	SetOpTenantId(v int64) *UpdatePipelineRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdatePipelineRequestUpdateCommand) *UpdatePipelineRequest
	GetUpdateCommand() *UpdatePipelineRequestUpdateCommand
}

type UpdatePipelineRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *UpdatePipelineRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The configuration for updating the pipeline or workflow node.
	//
	// This parameter is required.
	UpdateCommand *UpdatePipelineRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) GetContext() *UpdatePipelineRequestContext {
	return s.Context
}

func (s *UpdatePipelineRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdatePipelineRequest) GetUpdateCommand() *UpdatePipelineRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdatePipelineRequest) SetContext(v *UpdatePipelineRequestContext) *UpdatePipelineRequest {
	s.Context = v
	return s
}

func (s *UpdatePipelineRequest) SetOpTenantId(v int64) *UpdatePipelineRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdatePipelineRequest) SetUpdateCommand(v *UpdatePipelineRequestUpdateCommand) *UpdatePipelineRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdatePipelineRequest) Validate() error {
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

type UpdatePipelineRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: the development environment.
	//
	// - PROD: the production environment. For workflow nodes, only PROD is supported.
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

func (s UpdatePipelineRequestContext) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestContext) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestContext) GetEnv() *string {
	return s.Env
}

func (s *UpdatePipelineRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdatePipelineRequestContext) SetEnv(v string) *UpdatePipelineRequestContext {
	s.Env = &v
	return s
}

func (s *UpdatePipelineRequestContext) SetProjectId(v int64) *UpdatePipelineRequestContext {
	s.ProjectId = &v
	return s
}

func (s *UpdatePipelineRequestContext) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestUpdateCommand struct {
	// The remarks.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The integration pipeline configuration mode. Valid values:
	//
	// - PIPELINE (default): pipeline mode.
	//
	// - JSON: script mode.
	//
	// This parameter is not applicable to workflow nodes.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The basic information about the integration pipeline or workflow node.
	//
	// This parameter is required.
	NodeInfo *UpdatePipelineRequestUpdateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The integration pipeline component or workflow operator configuration.
	//
	// This parameter is required.
	PipelineConfig *UpdatePipelineRequestUpdateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// The integration pipeline configuration in JSON string format for script mode. Workflow nodes do not support script mode.
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// The node type. Valid values:
	//
	// - 0 (default): batch integration.
	//
	// - 1: real-time integration.
	//
	// - 14: workflow node.
	//
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// The scheduling configuration in JSON string format. Refer to the toJsonString method of the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The channel configuration in JSON string format. Refer to the toJsonString method of the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting.
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// Specifies whether to submit the node. Default value: true.
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
}

func (s UpdatePipelineRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestUpdateCommand) GetComment() *string {
	return s.Comment
}

func (s *UpdatePipelineRequestUpdateCommand) GetMode() *string {
	return s.Mode
}

func (s *UpdatePipelineRequestUpdateCommand) GetNodeInfo() *UpdatePipelineRequestUpdateCommandNodeInfo {
	return s.NodeInfo
}

func (s *UpdatePipelineRequestUpdateCommand) GetPipelineConfig() *UpdatePipelineRequestUpdateCommandPipelineConfig {
	return s.PipelineConfig
}

func (s *UpdatePipelineRequestUpdateCommand) GetPipelineJson() *string {
	return s.PipelineJson
}

func (s *UpdatePipelineRequestUpdateCommand) GetPipelineType() *int32 {
	return s.PipelineType
}

func (s *UpdatePipelineRequestUpdateCommand) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *UpdatePipelineRequestUpdateCommand) GetSettings() *string {
	return s.Settings
}

func (s *UpdatePipelineRequestUpdateCommand) GetSubmit() *bool {
	return s.Submit
}

func (s *UpdatePipelineRequestUpdateCommand) SetComment(v string) *UpdatePipelineRequestUpdateCommand {
	s.Comment = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetMode(v string) *UpdatePipelineRequestUpdateCommand {
	s.Mode = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetNodeInfo(v *UpdatePipelineRequestUpdateCommandNodeInfo) *UpdatePipelineRequestUpdateCommand {
	s.NodeInfo = v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetPipelineConfig(v *UpdatePipelineRequestUpdateCommandPipelineConfig) *UpdatePipelineRequestUpdateCommand {
	s.PipelineConfig = v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetPipelineJson(v string) *UpdatePipelineRequestUpdateCommand {
	s.PipelineJson = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetPipelineType(v int32) *UpdatePipelineRequestUpdateCommand {
	s.PipelineType = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetScheduleConfig(v string) *UpdatePipelineRequestUpdateCommand {
	s.ScheduleConfig = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetSettings(v string) *UpdatePipelineRequestUpdateCommand {
	s.Settings = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) SetSubmit(v bool) *UpdatePipelineRequestUpdateCommand {
	s.Submit = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommand) Validate() error {
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

type UpdatePipelineRequestUpdateCommandNodeInfo struct {
	// The folder of the integration pipeline or workflow node (defaults to the root folder). The folder must exist. If it does not exist, call the relevant API operation to create a folder of type offlinePipeline (or unstructuredPipeline for workflows).
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The file ID of the pipeline or workflow. Leave this parameter empty for initial creation. When updating a pipeline or workflow node, specify at least one of pipelineId, fileId, or nodeId.
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
	// The ID of the pipeline or workflow node. Leave this parameter empty for initial creation. When updating a pipeline or workflow node, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s UpdatePipelineRequestUpdateCommandNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestUpdateCommandNodeInfo) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) GetDirectory() *string {
	return s.Directory
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) SetDirectory(v string) *UpdatePipelineRequestUpdateCommandNodeInfo {
	s.Directory = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) SetFileId(v int64) *UpdatePipelineRequestUpdateCommandNodeInfo {
	s.FileId = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) SetNodeId(v string) *UpdatePipelineRequestUpdateCommandNodeInfo {
	s.NodeId = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) SetNodeName(v string) *UpdatePipelineRequestUpdateCommandNodeInfo {
	s.NodeName = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) SetPipelineId(v int64) *UpdatePipelineRequestUpdateCommandNodeInfo {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandNodeInfo) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestUpdateCommandPipelineConfig struct {
	// The DAG (directed acyclic graph) link configuration that describes the connection relationships among all components or operators.
	//
	// This parameter is required.
	Hops []*UpdatePipelineRequestUpdateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The component or operator configurations, including the detailed configurations of all components or operators used.
	//
	// This parameter is required.
	Steps []*UpdatePipelineRequestUpdateCommandPipelineConfigSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
}

func (s UpdatePipelineRequestUpdateCommandPipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestUpdateCommandPipelineConfig) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfig) GetHops() []*UpdatePipelineRequestUpdateCommandPipelineConfigHops {
	return s.Hops
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfig) GetSteps() []*UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	return s.Steps
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfig) SetHops(v []*UpdatePipelineRequestUpdateCommandPipelineConfigHops) *UpdatePipelineRequestUpdateCommandPipelineConfig {
	s.Hops = v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfig) SetSteps(v []*UpdatePipelineRequestUpdateCommandPipelineConfigSteps) *UpdatePipelineRequestUpdateCommandPipelineConfig {
	s.Steps = v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfig) Validate() error {
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

type UpdatePipelineRequestUpdateCommandPipelineConfigHops struct {
	// Specifies whether the downstream condition is true for a conditional distribution component. Set this parameter to true if the downstream condition is true, or false otherwise. This parameter is not applicable to workflow nodes.
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

func (s UpdatePipelineRequestUpdateCommandPipelineConfigHops) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestUpdateCommandPipelineConfigHops) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) GetSendTo() *bool {
	return s.SendTo
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) GetSource() *string {
	return s.Source
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) GetTarget() *string {
	return s.Target
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) SetSendTo(v bool) *UpdatePipelineRequestUpdateCommandPipelineConfigHops {
	s.SendTo = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) SetSource(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigHops {
	s.Source = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) SetTarget(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigHops {
	s.Target = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigHops) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestUpdateCommandPipelineConfigSteps struct {
	// Specifies the data distribution method when the current component has multiple downstream components. Valid values:
	//
	// - true (default): The data from the current component is distributed to all downstream components in a round-robin manner. For example, if the current component has 100 records and two downstream components, each downstream component receives 50 records.
	//
	// - false: The full data from the current component is sent to all downstream components. For example, if the current component has 100 records and two downstream components, each downstream component receives 100 records.
	//
	// This parameter is not applicable to workflow nodes.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// The plugin ID. Each plugin or operator has a unique identifier. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepKey. Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the configuration created on the Dataphin console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The specific component configuration in JSON string format. Refer to the toJsonString method of the relevant subclasses of the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig (or com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.unstructured.BaseOAUnstructuredNeuronConfig for workflow operators). Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the node configuration created on the Dataphin console.
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
	// For workflow nodes, this parameter specifies the operator type, such as image for images and text for text. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepType. Developers should inherit the component or operator configuration class and implement the corresponding component or operator configuration. Each component or operator configuration has the same structure as the configuration created on the Dataphin console.
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s UpdatePipelineRequestUpdateCommandPipelineConfigSteps) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GetIsDistribute() *bool {
	return s.IsDistribute
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GetKey() *string {
	return s.Key
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GetStepName() *string {
	return s.StepName
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) GetStepType() *string {
	return s.StepType
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) SetIsDistribute(v bool) *UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	s.IsDistribute = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) SetKey(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	s.Key = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) SetPluginConfig(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	s.PluginConfig = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) SetStepName(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	s.StepName = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) SetStepType(v string) *UpdatePipelineRequestUpdateCommandPipelineConfigSteps {
	s.StepType = &v
	return s
}

func (s *UpdatePipelineRequestUpdateCommandPipelineConfigSteps) Validate() error {
	return dara.Validate(s)
}
