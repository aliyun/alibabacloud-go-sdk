// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *SubmitAuditNoteRequest
	GetFileKey() *string
	SetNoteId(v string) *SubmitAuditNoteRequest
	GetNoteId() *string
	SetWorkspaceId(v string) *SubmitAuditNoteRequest
	GetWorkspaceId() *string
}

type SubmitAuditNoteRequest struct {
	// The FileKey of your rule library file stored in Alibaba Cloud OSS. For how to generate a FileKey, see [Common APIs: GenerateUploadConfig for File Upload and Download](https://next.api.aliyun.com/document/AiMiaoBi/2023-08-01/GenerateUploadConfig?spm=openapi-amp.newDocPublishment.0.0.18fc281fOiiBil). Your rule library file must be in DOCX, XLSX, or PDF format. If you use XLSX, it must have exactly two columns. The table header must be "Proofreading Basis" and "Source". XLSX files give the best parsing results. DOCX and PDF files are also parsed automatically.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://default/your/file/key
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// ID of the rule library. If you omit this parameter, the system uses Default.
	//
	// example:
	//
	// note_id_unique
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// Unique identifier of your Model Studio workspace. To get this ID, see [Get the Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitAuditNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditNoteRequest) GoString() string {
	return s.String()
}

func (s *SubmitAuditNoteRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitAuditNoteRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *SubmitAuditNoteRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitAuditNoteRequest) SetFileKey(v string) *SubmitAuditNoteRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitAuditNoteRequest) SetNoteId(v string) *SubmitAuditNoteRequest {
	s.NoteId = &v
	return s
}

func (s *SubmitAuditNoteRequest) SetWorkspaceId(v string) *SubmitAuditNoteRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitAuditNoteRequest) Validate() error {
	return dara.Validate(s)
}
