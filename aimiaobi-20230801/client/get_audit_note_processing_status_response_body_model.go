// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditNoteProcessingStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAuditNoteProcessingStatusResponseBody
	GetCode() *string
	SetData(v *GetAuditNoteProcessingStatusResponseBodyData) *GetAuditNoteProcessingStatusResponseBody
	GetData() *GetAuditNoteProcessingStatusResponseBodyData
	SetHttpStatusCode(v int32) *GetAuditNoteProcessingStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAuditNoteProcessingStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuditNoteProcessingStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuditNoteProcessingStatusResponseBody
	GetSuccess() *bool
}

type GetAuditNoteProcessingStatusResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAuditNoteProcessingStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAuditNoteProcessingStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNoteProcessingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetData() *GetAuditNoteProcessingStatusResponseBodyData {
	return s.Data
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditNoteProcessingStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetCode(v string) *GetAuditNoteProcessingStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetData(v *GetAuditNoteProcessingStatusResponseBodyData) *GetAuditNoteProcessingStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetHttpStatusCode(v int32) *GetAuditNoteProcessingStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetMessage(v string) *GetAuditNoteProcessingStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetRequestId(v string) *GetAuditNoteProcessingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) SetSuccess(v bool) *GetAuditNoteProcessingStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuditNoteProcessingStatusResponseBodyData struct {
	// example:
	//
	// oss://default/path/to/audit/note
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// 504
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 错题本2025-07-07_解析结果
	NoteName *string `json:"NoteName,omitempty" xml:"NoteName,omitempty"`
	// example:
	//
	// SUCCESSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 111_Default_20250708142918
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2024-11-25 11:40:50
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAuditNoteProcessingStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAuditNoteProcessingStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetNoteName() *string {
	return s.NoteName
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetFileKey(v string) *GetAuditNoteProcessingStatusResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetFileSize(v int64) *GetAuditNoteProcessingStatusResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetNoteName(v string) *GetAuditNoteProcessingStatusResponseBodyData {
	s.NoteName = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetStatus(v string) *GetAuditNoteProcessingStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetTaskId(v string) *GetAuditNoteProcessingStatusResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) SetUpdateTime(v int64) *GetAuditNoteProcessingStatusResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAuditNoteProcessingStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
