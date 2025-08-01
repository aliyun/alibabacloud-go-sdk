// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAuditNoteResponseBody
	GetCode() *string
	SetData(v string) *DeleteAuditNoteResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteAuditNoteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAuditNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAuditNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAuditNoteResponseBody
	GetSuccess() *bool
}

type DeleteAuditNoteResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// SUCCESSED
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

func (s DeleteAuditNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditNoteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuditNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAuditNoteResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAuditNoteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAuditNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAuditNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuditNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAuditNoteResponseBody) SetCode(v string) *DeleteAuditNoteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) SetData(v string) *DeleteAuditNoteResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) SetHttpStatusCode(v int32) *DeleteAuditNoteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) SetMessage(v string) *DeleteAuditNoteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) SetRequestId(v string) *DeleteAuditNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) SetSuccess(v bool) *DeleteAuditNoteResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAuditNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
