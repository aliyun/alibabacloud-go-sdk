// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNoteId(v string) *DownloadAuditNoteRequest
	GetNoteId() *string
	SetTaskId(v string) *DownloadAuditNoteRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *DownloadAuditNoteRequest
	GetWorkspaceId() *string
}

type DownloadAuditNoteRequest struct {
	// Rule library ID. If left blank, the default is Default.
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// The task ID obtained by calling the SubmitAuditNote API. This is the unique identifier for the custom rule library task index. Store it securely when using it. When using this API, if the input parameters include \\`taskId\\`, you can retrieve the structured rule library that you successfully uploaded via the SubmitAuditNote API but has not undergone post-processing by ConfirmAndProcessAuditNote. If the input parameters do not include \\`taskId\\`, you can retrieve the structured rule library that has undergone post-processing and is available for auditing.
	//
	// example:
	//
	// xxx_Default_1241541251241
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique identifier of the Alibaba Cloud Model Studio workspace. Get the [Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DownloadAuditNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadAuditNoteRequest) GoString() string {
	return s.String()
}

func (s *DownloadAuditNoteRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *DownloadAuditNoteRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DownloadAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DownloadAuditNoteRequest) SetNoteId(v string) *DownloadAuditNoteRequest {
	s.NoteId = &v
	return s
}

func (s *DownloadAuditNoteRequest) SetTaskId(v string) *DownloadAuditNoteRequest {
	s.TaskId = &v
	return s
}

func (s *DownloadAuditNoteRequest) SetWorkspaceId(v string) *DownloadAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DownloadAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
