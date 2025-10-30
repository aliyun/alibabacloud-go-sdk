// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineByAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *OfflinePipelineByAsyncRequestContext) *OfflinePipelineByAsyncRequest
	GetContext() *OfflinePipelineByAsyncRequestContext
	SetOfflineCommand(v *OfflinePipelineByAsyncRequestOfflineCommand) *OfflinePipelineByAsyncRequest
	GetOfflineCommand() *OfflinePipelineByAsyncRequestOfflineCommand
	SetOpTenantId(v int64) *OfflinePipelineByAsyncRequest
	GetOpTenantId() *int64
}

type OfflinePipelineByAsyncRequest struct {
	// This parameter is required.
	Context *OfflinePipelineByAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// This parameter is required.
	OfflineCommand *OfflinePipelineByAsyncRequestOfflineCommand `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OfflinePipelineByAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncRequest) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncRequest) GetContext() *OfflinePipelineByAsyncRequestContext {
	return s.Context
}

func (s *OfflinePipelineByAsyncRequest) GetOfflineCommand() *OfflinePipelineByAsyncRequestOfflineCommand {
	return s.OfflineCommand
}

func (s *OfflinePipelineByAsyncRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflinePipelineByAsyncRequest) SetContext(v *OfflinePipelineByAsyncRequestContext) *OfflinePipelineByAsyncRequest {
	s.Context = v
	return s
}

func (s *OfflinePipelineByAsyncRequest) SetOfflineCommand(v *OfflinePipelineByAsyncRequestOfflineCommand) *OfflinePipelineByAsyncRequest {
	s.OfflineCommand = v
	return s
}

func (s *OfflinePipelineByAsyncRequest) SetOpTenantId(v int64) *OfflinePipelineByAsyncRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflinePipelineByAsyncRequest) Validate() error {
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

type OfflinePipelineByAsyncRequestContext struct {
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

func (s OfflinePipelineByAsyncRequestContext) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncRequestContext) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncRequestContext) GetEnv() *string {
	return s.Env
}

func (s *OfflinePipelineByAsyncRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *OfflinePipelineByAsyncRequestContext) SetEnv(v string) *OfflinePipelineByAsyncRequestContext {
	s.Env = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestContext) SetProjectId(v int64) *OfflinePipelineByAsyncRequestContext {
	s.ProjectId = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestContext) Validate() error {
	return dara.Validate(s)
}

type OfflinePipelineByAsyncRequestOfflineCommand struct {
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s OfflinePipelineByAsyncRequestOfflineCommand) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncRequestOfflineCommand) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) GetComment() *string {
	return s.Comment
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) GetDelete() *bool {
	return s.Delete
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) GetNodeId() *string {
	return s.NodeId
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) SetComment(v string) *OfflinePipelineByAsyncRequestOfflineCommand {
	s.Comment = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) SetDelete(v bool) *OfflinePipelineByAsyncRequestOfflineCommand {
	s.Delete = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) SetFileId(v int64) *OfflinePipelineByAsyncRequestOfflineCommand {
	s.FileId = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) SetNodeId(v string) *OfflinePipelineByAsyncRequestOfflineCommand {
	s.NodeId = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) SetPipelineId(v int64) *OfflinePipelineByAsyncRequestOfflineCommand {
	s.PipelineId = &v
	return s
}

func (s *OfflinePipelineByAsyncRequestOfflineCommand) Validate() error {
	return dara.Validate(s)
}
