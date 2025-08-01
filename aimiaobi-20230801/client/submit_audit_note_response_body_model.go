// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAuditNoteResponseBody
	GetCode() *string
	SetData(v string) *SubmitAuditNoteResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *SubmitAuditNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitAuditNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAuditNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAuditNoteResponseBody
	GetSuccess() *bool
}

type SubmitAuditNoteResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx_Default_89123748917
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

func (s SubmitAuditNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditNoteResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAuditNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAuditNoteResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitAuditNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitAuditNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAuditNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAuditNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAuditNoteResponseBody) SetCode(v string) *SubmitAuditNoteResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) SetData(v string) *SubmitAuditNoteResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) SetHttpStatusCode(v int32) *SubmitAuditNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) SetMessage(v string) *SubmitAuditNoteResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) SetRequestId(v string) *SubmitAuditNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) SetSuccess(v bool) *SubmitAuditNoteResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAuditNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
