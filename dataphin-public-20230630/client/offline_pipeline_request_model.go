// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *OfflinePipelineRequestContext) *OfflinePipelineRequest
	GetContext() *OfflinePipelineRequestContext
	SetOfflineCommand(v *OfflinePipelineRequestOfflineCommand) *OfflinePipelineRequest
	GetOfflineCommand() *OfflinePipelineRequestOfflineCommand
	SetOpTenantId(v int64) *OfflinePipelineRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *OfflinePipelineRequest
	GetOpUserId() *string
}

type OfflinePipelineRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *OfflinePipelineRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The command to take the pipeline task offline.
	//
	// This parameter is required.
	OfflineCommand *OfflinePipelineRequestOfflineCommand `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty" type:"Struct"`
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

func (s OfflinePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineRequest) GoString() string {
	return s.String()
}

func (s *OfflinePipelineRequest) GetContext() *OfflinePipelineRequestContext {
	return s.Context
}

func (s *OfflinePipelineRequest) GetOfflineCommand() *OfflinePipelineRequestOfflineCommand {
	return s.OfflineCommand
}

func (s *OfflinePipelineRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflinePipelineRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *OfflinePipelineRequest) SetContext(v *OfflinePipelineRequestContext) *OfflinePipelineRequest {
	s.Context = v
	return s
}

func (s *OfflinePipelineRequest) SetOfflineCommand(v *OfflinePipelineRequestOfflineCommand) *OfflinePipelineRequest {
	s.OfflineCommand = v
	return s
}

func (s *OfflinePipelineRequest) SetOpTenantId(v int64) *OfflinePipelineRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflinePipelineRequest) SetOpUserId(v string) *OfflinePipelineRequest {
	s.OpUserId = &v
	return s
}

func (s *OfflinePipelineRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.OfflineCommand != nil {
		if err := s.OfflineCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OfflinePipelineRequestContext struct {
	// The current operation environment. Valid values: DEV: development environment. PROD: production environment.
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

func (s OfflinePipelineRequestContext) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineRequestContext) GoString() string {
	return s.String()
}

func (s *OfflinePipelineRequestContext) GetEnv() *string {
	return s.Env
}

func (s *OfflinePipelineRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *OfflinePipelineRequestContext) SetEnv(v string) *OfflinePipelineRequestContext {
	s.Env = &v
	return s
}

func (s *OfflinePipelineRequestContext) SetProjectId(v int64) *OfflinePipelineRequestContext {
	s.ProjectId = &v
	return s
}

func (s *OfflinePipelineRequestContext) Validate() error {
	return dara.Validate(s)
}

type OfflinePipelineRequestOfflineCommand struct {
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether to delete the task.
	//
	// This parameter is required.
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The ID of the integration task file. You can specify pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the scheduling node for the integration task. You can specify pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The primary key of the integration pipeline. You can specify pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s OfflinePipelineRequestOfflineCommand) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineRequestOfflineCommand) GoString() string {
	return s.String()
}

func (s *OfflinePipelineRequestOfflineCommand) GetComment() *string {
	return s.Comment
}

func (s *OfflinePipelineRequestOfflineCommand) GetDelete() *bool {
	return s.Delete
}

func (s *OfflinePipelineRequestOfflineCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *OfflinePipelineRequestOfflineCommand) GetNodeId() *string {
	return s.NodeId
}

func (s *OfflinePipelineRequestOfflineCommand) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *OfflinePipelineRequestOfflineCommand) SetComment(v string) *OfflinePipelineRequestOfflineCommand {
	s.Comment = &v
	return s
}

func (s *OfflinePipelineRequestOfflineCommand) SetDelete(v bool) *OfflinePipelineRequestOfflineCommand {
	s.Delete = &v
	return s
}

func (s *OfflinePipelineRequestOfflineCommand) SetFileId(v int64) *OfflinePipelineRequestOfflineCommand {
	s.FileId = &v
	return s
}

func (s *OfflinePipelineRequestOfflineCommand) SetNodeId(v string) *OfflinePipelineRequestOfflineCommand {
	s.NodeId = &v
	return s
}

func (s *OfflinePipelineRequestOfflineCommand) SetPipelineId(v int64) *OfflinePipelineRequestOfflineCommand {
	s.PipelineId = &v
	return s
}

func (s *OfflinePipelineRequestOfflineCommand) Validate() error {
	return dara.Validate(s)
}
