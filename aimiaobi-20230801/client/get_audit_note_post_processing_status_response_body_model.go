// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNotePostProcessingStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAuditNotePostProcessingStatusResponseBody
	GetCode() *string
	SetData(v *GetAuditNotePostProcessingStatusResponseBodyData) *GetAuditNotePostProcessingStatusResponseBody
	GetData() *GetAuditNotePostProcessingStatusResponseBodyData
	SetHttpStatusCode(v int32) *GetAuditNotePostProcessingStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAuditNotePostProcessingStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuditNotePostProcessingStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuditNotePostProcessingStatusResponseBody
	GetSuccess() *bool
}

type GetAuditNotePostProcessingStatusResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// if can be null:
	// true
	Data *GetAuditNotePostProcessingStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuditNotePostProcessingStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNotePostProcessingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetData() *GetAuditNotePostProcessingStatusResponseBodyData {
	return s.Data
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditNotePostProcessingStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetCode(v string) *GetAuditNotePostProcessingStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetData(v *GetAuditNotePostProcessingStatusResponseBodyData) *GetAuditNotePostProcessingStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetHttpStatusCode(v int32) *GetAuditNotePostProcessingStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetMessage(v string) *GetAuditNotePostProcessingStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetRequestId(v string) *GetAuditNotePostProcessingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) SetSuccess(v bool) *GetAuditNotePostProcessingStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuditNotePostProcessingStatusResponseBodyData struct {
	// example:
	//
	// 1970-01-01 23:34:45
	CompletionTime *string `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	// example:
	//
	// 1970-01-01 12:34:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// Default
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// example:
	//
	// 233
	ProcessedLines *int32 `json:"ProcessedLines,omitempty" xml:"ProcessedLines,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 666
	TotalLines *int32 `json:"TotalLines,omitempty" xml:"TotalLines,omitempty"`
}

func (s GetAuditNotePostProcessingStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNotePostProcessingStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetCompletionTime() *string {
	return s.CompletionTime
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetNoteId() *string {
	return s.NoteId
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetProcessedLines() *int32 {
	return s.ProcessedLines
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) GetTotalLines() *int32 {
	return s.TotalLines
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetCompletionTime(v string) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetCreateTime(v string) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetErrorMessage(v string) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetNoteId(v string) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.NoteId = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetProcessedLines(v int32) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.ProcessedLines = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetStatus(v string) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) SetTotalLines(v int32) *GetAuditNotePostProcessingStatusResponseBodyData {
	s.TotalLines = &v
	return s
}

func (s *GetAuditNotePostProcessingStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
