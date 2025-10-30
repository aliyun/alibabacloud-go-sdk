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
	// This parameter is required.
	Context *UpdatePipelineRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	NodeInfo *UpdatePipelineRequestUpdateCommandNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// This parameter is required.
	PipelineConfig *UpdatePipelineRequestUpdateCommandPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
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
	// This parameter is required.
	Hops []*UpdatePipelineRequestUpdateCommandPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
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
