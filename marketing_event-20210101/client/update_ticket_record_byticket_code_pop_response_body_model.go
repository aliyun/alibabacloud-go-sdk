// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketRecordByticketCodePopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateTicketRecordByticketCodePopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateTicketRecordByticketCodePopResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateTicketRecordByticketCodePopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateTicketRecordByticketCodePopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateTicketRecordByticketCodePopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTicketRecordByticketCodePopResponseBody
	GetSuccess() *bool
}

type UpdateTicketRecordByticketCodePopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetAccessDeniedDetail(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetData(v bool) *UpdateTicketRecordByticketCodePopResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetErrCode(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetErrMessage(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetHttpStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetRequestId(v string) *UpdateTicketRecordByticketCodePopResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) SetSuccess(v bool) *UpdateTicketRecordByticketCodePopResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponseBody) Validate() error {
	return dara.Validate(s)
}
