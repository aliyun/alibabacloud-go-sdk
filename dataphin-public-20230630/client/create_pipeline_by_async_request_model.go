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
	// This parameter is required.
	Context *CreatePipelineByAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// This parameter is required.
	CreateCommand *CreatePipelineByAsyncRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
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
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	NodeInfo *CreatePipelineByAsyncRequestCreateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// This parameter is required.
	PipelineConfig *CreatePipelineByAsyncRequestCreateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// example:
	//
	// 0
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"cronExpression":"0 0 0 	- 	- ?"}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	Submit   *bool   `json:"Submit,omitempty" xml:"Submit,omitempty"`
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
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	// This parameter is required.
	Hops []*CreatePipelineByAsyncRequestCreateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
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
	SendTo *bool `json:"SendTo,omitempty" xml:"SendTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
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
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
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
