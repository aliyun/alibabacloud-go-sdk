// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityArchiveTableProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityArchiveTableProgressResponseBody
	GetCode() *string
	SetData(v *GetQualityArchiveTableProgressResponseBodyData) *GetQualityArchiveTableProgressResponseBody
	GetData() *GetQualityArchiveTableProgressResponseBodyData
	SetHttpStatusCode(v int32) *GetQualityArchiveTableProgressResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityArchiveTableProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityArchiveTableProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityArchiveTableProgressResponseBody
	GetSuccess() *bool
}

type GetQualityArchiveTableProgressResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task progress details.
	Data *GetQualityArchiveTableProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityArchiveTableProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityArchiveTableProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityArchiveTableProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityArchiveTableProgressResponseBody) GetData() *GetQualityArchiveTableProgressResponseBodyData {
	return s.Data
}

func (s *GetQualityArchiveTableProgressResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityArchiveTableProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityArchiveTableProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityArchiveTableProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityArchiveTableProgressResponseBody) SetCode(v string) *GetQualityArchiveTableProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) SetData(v *GetQualityArchiveTableProgressResponseBodyData) *GetQualityArchiveTableProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) SetHttpStatusCode(v int32) *GetQualityArchiveTableProgressResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) SetMessage(v string) *GetQualityArchiveTableProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) SetRequestId(v string) *GetQualityArchiveTableProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) SetSuccess(v bool) *GetQualityArchiveTableProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityArchiveTableProgressResponseBodyData struct {
	// The archived table ID. This parameter is returned after the task succeeds and can be used to switch the active table.
	//
	// example:
	//
	// 88012
	ArchiveTableId *int64 `json:"ArchiveTableId,omitempty" xml:"ArchiveTableId,omitempty"`
	// The full name of the archived table. This parameter is returned after the task succeeds. When creating a table, the name includes the automatically appended _exception_data suffix.
	//
	// example:
	//
	// Train.a01_reanme3_exception_data
	ArchiveTableName *string `json:"ArchiveTableName,omitempty" xml:"ArchiveTableName,omitempty"`
	// The reason for the task failure. This parameter is returned only when Status is FAILED.
	//
	// example:
	//
	// The archived table name already exists!
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task status. Valid values:
	//
	// - PROGRESS: In progress.
	//
	// - SUCCESS: Succeeded.
	//
	// - FAILED: Failed.
	//
	// - CANCEL: Canceled.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetQualityArchiveTableProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityArchiveTableProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityArchiveTableProgressResponseBodyData) GetArchiveTableId() *int64 {
	return s.ArchiveTableId
}

func (s *GetQualityArchiveTableProgressResponseBodyData) GetArchiveTableName() *string {
	return s.ArchiveTableName
}

func (s *GetQualityArchiveTableProgressResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetQualityArchiveTableProgressResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetQualityArchiveTableProgressResponseBodyData) SetArchiveTableId(v int64) *GetQualityArchiveTableProgressResponseBodyData {
	s.ArchiveTableId = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBodyData) SetArchiveTableName(v string) *GetQualityArchiveTableProgressResponseBodyData {
	s.ArchiveTableName = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBodyData) SetErrorMessage(v string) *GetQualityArchiveTableProgressResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBodyData) SetStatus(v string) *GetQualityArchiveTableProgressResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
