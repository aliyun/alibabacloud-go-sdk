// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkFlowByJsonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *CreateWorkFlowByJsonRequestContext) *CreateWorkFlowByJsonRequest
	GetContext() *CreateWorkFlowByJsonRequestContext
	SetCreateCommand(v *CreateWorkFlowByJsonRequestCreateCommand) *CreateWorkFlowByJsonRequest
	GetCreateCommand() *CreateWorkFlowByJsonRequestCreateCommand
	SetOpTenantId(v int64) *CreateWorkFlowByJsonRequest
	GetOpTenantId() *int64
}

type CreateWorkFlowByJsonRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *CreateWorkFlowByJsonRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The JSON script command for creating a workflow.
	//
	// This parameter is required.
	CreateCommand *CreateWorkFlowByJsonRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateWorkFlowByJsonRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonRequest) GetContext() *CreateWorkFlowByJsonRequestContext {
	return s.Context
}

func (s *CreateWorkFlowByJsonRequest) GetCreateCommand() *CreateWorkFlowByJsonRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateWorkFlowByJsonRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateWorkFlowByJsonRequest) SetContext(v *CreateWorkFlowByJsonRequestContext) *CreateWorkFlowByJsonRequest {
	s.Context = v
	return s
}

func (s *CreateWorkFlowByJsonRequest) SetCreateCommand(v *CreateWorkFlowByJsonRequestCreateCommand) *CreateWorkFlowByJsonRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateWorkFlowByJsonRequest) SetOpTenantId(v int64) *CreateWorkFlowByJsonRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateWorkFlowByJsonRequest) Validate() error {
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

type CreateWorkFlowByJsonRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: the development environment.
	//
	// - PROD: the production environment.
	//
	// The current version supports only BASIC mode, so set this parameter to PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the workflow node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 789
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateWorkFlowByJsonRequestContext) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonRequestContext) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonRequestContext) GetEnv() *string {
	return s.Env
}

func (s *CreateWorkFlowByJsonRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateWorkFlowByJsonRequestContext) SetEnv(v string) *CreateWorkFlowByJsonRequestContext {
	s.Env = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestContext) SetProjectId(v int64) *CreateWorkFlowByJsonRequestContext {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestContext) Validate() error {
	return dara.Validate(s)
}

type CreateWorkFlowByJsonRequestCreateCommand struct {
	// The description of the node.
	//
	// example:
	//
	// cooment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The folder to which the node belongs. If this parameter is left empty, the root folder is used.
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The schedule configuration. This parameter is required for periodic nodes. The value is a JSON string. Refer to the utility class: com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAScheduleConfig#toJsonString method.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// Specifies whether to submit the node. Default value: true.
	Submit *bool `json:"Submit,omitempty" xml:"Submit,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// workflow_name
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// - 1: periodic scheduling.
	//
	// - 3: manual scheduling.
	//
	// - 5: real-time node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5372881
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The workflow JSON.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"pipelineDTO":{"hops":[],"steps":[{"name":"xxx","x":305,"y":144,"id":"c404a7c6-8a75-4ed8-b348-0785423ad859","webConfig":{},"type":"text","key":"special_character_removal","pluginConfig":{"neuronParameters":{},"neuronInput":{},"neuronOutput":{},"setting":{}}}]}}
	WorkFlowJson *string `json:"WorkFlowJson,omitempty" xml:"WorkFlowJson,omitempty"`
}

func (s CreateWorkFlowByJsonRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetSubmit() *bool {
	return s.Submit
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) GetWorkFlowJson() *string {
	return s.WorkFlowJson
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetDescription(v string) *CreateWorkFlowByJsonRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetDirectory(v string) *CreateWorkFlowByJsonRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetScheduleConfig(v string) *CreateWorkFlowByJsonRequestCreateCommand {
	s.ScheduleConfig = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetSubmit(v bool) *CreateWorkFlowByJsonRequestCreateCommand {
	s.Submit = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetTaskName(v string) *CreateWorkFlowByJsonRequestCreateCommand {
	s.TaskName = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetTaskType(v int32) *CreateWorkFlowByJsonRequestCreateCommand {
	s.TaskType = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) SetWorkFlowJson(v string) *CreateWorkFlowByJsonRequestCreateCommand {
	s.WorkFlowJson = &v
	return s
}

func (s *CreateWorkFlowByJsonRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
