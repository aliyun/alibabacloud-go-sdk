// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableAuditNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAvailableAuditNotesResponseBody
	GetCode() *string
	SetData(v *GetAvailableAuditNotesResponseBodyData) *GetAvailableAuditNotesResponseBody
	GetData() *GetAvailableAuditNotesResponseBodyData
	SetHttpStatusCode(v int32) *GetAvailableAuditNotesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAvailableAuditNotesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAvailableAuditNotesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAvailableAuditNotesResponseBody
	GetSuccess() *bool
}

type GetAvailableAuditNotesResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAvailableAuditNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvailableAuditNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableAuditNotesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvailableAuditNotesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAvailableAuditNotesResponseBody) GetData() *GetAvailableAuditNotesResponseBodyData {
	return s.Data
}

func (s *GetAvailableAuditNotesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAvailableAuditNotesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAvailableAuditNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvailableAuditNotesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAvailableAuditNotesResponseBody) SetCode(v string) *GetAvailableAuditNotesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) SetData(v *GetAvailableAuditNotesResponseBodyData) *GetAvailableAuditNotesResponseBody {
	s.Data = v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) SetHttpStatusCode(v int32) *GetAvailableAuditNotesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) SetMessage(v string) *GetAvailableAuditNotesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) SetRequestId(v string) *GetAvailableAuditNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) SetSuccess(v bool) *GetAvailableAuditNotesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAvailableAuditNotesResponseBodyData struct {
	// example:
	//
	// 23333
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Default
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// example:
	//
	// 错题本2025-07-07_解析结果
	NoteName *string `json:"NoteName,omitempty" xml:"NoteName,omitempty"`
	// example:
	//
	// 2025-07-07 11:56:30
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAvailableAuditNotesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableAuditNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvailableAuditNotesResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetAvailableAuditNotesResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetAvailableAuditNotesResponseBodyData) GetNoteId() *string {
	return s.NoteId
}

func (s *GetAvailableAuditNotesResponseBodyData) GetNoteName() *string {
	return s.NoteName
}

func (s *GetAvailableAuditNotesResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAvailableAuditNotesResponseBodyData) SetFileSize(v int64) *GetAvailableAuditNotesResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBodyData) SetId(v string) *GetAvailableAuditNotesResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBodyData) SetNoteId(v string) *GetAvailableAuditNotesResponseBodyData {
	s.NoteId = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBodyData) SetNoteName(v string) *GetAvailableAuditNotesResponseBodyData {
	s.NoteName = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBodyData) SetUpdateTime(v string) *GetAvailableAuditNotesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAvailableAuditNotesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
