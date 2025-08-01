// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAndPostProcessAuditNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfirmAndPostProcessAuditNoteResponseBody
	GetCode() *string
	SetData(v string) *ConfirmAndPostProcessAuditNoteResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ConfirmAndPostProcessAuditNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ConfirmAndPostProcessAuditNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmAndPostProcessAuditNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmAndPostProcessAuditNoteResponseBody
	GetSuccess() *bool
}

type ConfirmAndPostProcessAuditNoteResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 33
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ConfirmAndPostProcessAuditNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAndPostProcessAuditNoteResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetData() *string {
	return s.Data
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetCode(v string) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetData(v string) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.Data = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetHttpStatusCode(v int32) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetMessage(v string) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetRequestId(v string) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) SetSuccess(v bool) *ConfirmAndPostProcessAuditNoteResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmAndPostProcessAuditNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
