// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*SendMessageRequestAttachments) *SendMessageRequest
	GetAttachments() []*SendMessageRequestAttachments
	SetContent(v string) *SendMessageRequest
	GetContent() *string
	SetResume(v []*SendMessageRequestResume) *SendMessageRequest
	GetResume() []*SendMessageRequestResume
	SetRunId(v string) *SendMessageRequest
	GetRunId() *string
	SetType(v string) *SendMessageRequest
	GetType() *string
}

type SendMessageRequest struct {
	// The list of attachments.
	//
	// example:
	//
	// []
	Attachments []*SendMessageRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The message text in MESSAGE mode. This parameter cannot be empty. The type parameter can be omitted in this mode.
	//
	// example:
	//
	// Query the status of the sample resource without making changes
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The list of interrupt replies. This parameter is required when type is set to RESUME. You must reply to all pending interrupts at once. Each interruptId can appear only once.
	Resume []*SendMessageRequestResume `json:"resume,omitempty" xml:"resume,omitempty" type:"Repeated"`
	// The ID of the active run to resubscribe to. This parameter is required in RECONNECT mode.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The request type. Valid values: MESSAGE (sends a message), RESUME (resumes a run), and RECONNECT (reconnects to a run). Default value: MESSAGE.
	//
	// example:
	//
	// MESSAGE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetAttachments() []*SendMessageRequestAttachments {
	return s.Attachments
}

func (s *SendMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendMessageRequest) GetResume() []*SendMessageRequestResume {
	return s.Resume
}

func (s *SendMessageRequest) GetRunId() *string {
	return s.RunId
}

func (s *SendMessageRequest) GetType() *string {
	return s.Type
}

func (s *SendMessageRequest) SetAttachments(v []*SendMessageRequestAttachments) *SendMessageRequest {
	s.Attachments = v
	return s
}

func (s *SendMessageRequest) SetContent(v string) *SendMessageRequest {
	s.Content = &v
	return s
}

func (s *SendMessageRequest) SetResume(v []*SendMessageRequestResume) *SendMessageRequest {
	s.Resume = v
	return s
}

func (s *SendMessageRequest) SetRunId(v string) *SendMessageRequest {
	s.RunId = &v
	return s
}

func (s *SendMessageRequest) SetType(v string) *SendMessageRequest {
	s.Type = &v
	return s
}

func (s *SendMessageRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resume != nil {
		for _, item := range s.Resume {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendMessageRequestAttachments struct {
	// The file name extension.
	//
	// example:
	//
	// xlsx
	FileExtension *string `json:"fileExtension,omitempty" xml:"fileExtension,omitempty"`
	// The ID of the file after it is uploaded.
	//
	// example:
	//
	// att_12345678
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// The original name of the file.
	//
	// example:
	//
	// file.xlsx
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The workspace path.
	//
	// example:
	//
	// source/att_12345678.xlsx
	WorkspacePath *string `json:"workspacePath,omitempty" xml:"workspacePath,omitempty"`
}

func (s SendMessageRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequestAttachments) GoString() string {
	return s.String()
}

func (s *SendMessageRequestAttachments) GetFileExtension() *string {
	return s.FileExtension
}

func (s *SendMessageRequestAttachments) GetFileId() *string {
	return s.FileId
}

func (s *SendMessageRequestAttachments) GetSourceName() *string {
	return s.SourceName
}

func (s *SendMessageRequestAttachments) GetWorkspacePath() *string {
	return s.WorkspacePath
}

func (s *SendMessageRequestAttachments) SetFileExtension(v string) *SendMessageRequestAttachments {
	s.FileExtension = &v
	return s
}

func (s *SendMessageRequestAttachments) SetFileId(v string) *SendMessageRequestAttachments {
	s.FileId = &v
	return s
}

func (s *SendMessageRequestAttachments) SetSourceName(v string) *SendMessageRequestAttachments {
	s.SourceName = &v
	return s
}

func (s *SendMessageRequestAttachments) SetWorkspacePath(v string) *SendMessageRequestAttachments {
	s.WorkspacePath = &v
	return s
}

func (s *SendMessageRequestAttachments) Validate() error {
	return dara.Validate(s)
}

type SendMessageRequestResume struct {
	// The ID of the interrupt to be replied to. Obtain this value from data.outcome.interrupts[].id in the SSE RUN_FINISHED event or the GetRun response.
	//
	// example:
	//
	// interrupt_example_001
	InterruptId *string `json:"interruptId,omitempty" xml:"interruptId,omitempty"`
	// The interrupt handling result. This parameter is required when status is set to resolved and must conform to the responseSchema of the corresponding interrupt. Do not pass this parameter when status is set to cancelled.
	//
	// example:
	//
	// {"region":"cn-hangzhou"}
	Payload interface{} `json:"payload,omitempty" xml:"payload,omitempty"`
	// The reply method for the interrupt. Valid values: resolved (submits the answer, confirmation, or tool result required by the interrupt, and must include payload) and cancelled (cancels the interrupt without passing payload).
	//
	// example:
	//
	// resolved
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SendMessageRequestResume) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequestResume) GoString() string {
	return s.String()
}

func (s *SendMessageRequestResume) GetInterruptId() *string {
	return s.InterruptId
}

func (s *SendMessageRequestResume) GetPayload() interface{} {
	return s.Payload
}

func (s *SendMessageRequestResume) GetStatus() *string {
	return s.Status
}

func (s *SendMessageRequestResume) SetInterruptId(v string) *SendMessageRequestResume {
	s.InterruptId = &v
	return s
}

func (s *SendMessageRequestResume) SetPayload(v interface{}) *SendMessageRequestResume {
	s.Payload = v
	return s
}

func (s *SendMessageRequestResume) SetStatus(v string) *SendMessageRequestResume {
	s.Status = &v
	return s
}

func (s *SendMessageRequestResume) Validate() error {
	return dara.Validate(s)
}
