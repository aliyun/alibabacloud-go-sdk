// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteBatchTaskRequestDeleteCommand) *DeleteBatchTaskRequest
	GetDeleteCommand() *DeleteBatchTaskRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteBatchTaskRequest
	GetOpTenantId() *int64
}

type DeleteBatchTaskRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteBatchTaskRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteBatchTaskRequest) GetDeleteCommand() *DeleteBatchTaskRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBatchTaskRequest) SetDeleteCommand(v *DeleteBatchTaskRequestDeleteCommand) *DeleteBatchTaskRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteBatchTaskRequest) SetOpTenantId(v int64) *DeleteBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteBatchTaskRequestDeleteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// test task
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteBatchTaskRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchTaskRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteBatchTaskRequestDeleteCommand) GetComment() *string {
	return s.Comment
}

func (s *DeleteBatchTaskRequestDeleteCommand) GetFileId() *int64 {
	return s.FileId
}

func (s *DeleteBatchTaskRequestDeleteCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteBatchTaskRequestDeleteCommand) SetComment(v string) *DeleteBatchTaskRequestDeleteCommand {
	s.Comment = &v
	return s
}

func (s *DeleteBatchTaskRequestDeleteCommand) SetFileId(v int64) *DeleteBatchTaskRequestDeleteCommand {
	s.FileId = &v
	return s
}

func (s *DeleteBatchTaskRequestDeleteCommand) SetProjectId(v int64) *DeleteBatchTaskRequestDeleteCommand {
	s.ProjectId = &v
	return s
}

func (s *DeleteBatchTaskRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
