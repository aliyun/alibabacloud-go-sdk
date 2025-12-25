// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReplyTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ReplyTicketResponseBody
	GetCode() *int32
	SetData(v string) *ReplyTicketResponseBody
	GetData() *string
	SetMessage(v string) *ReplyTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReplyTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReplyTicketResponseBody
	GetSuccess() *bool
}

type ReplyTicketResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned value is the dialogID of the response to the ticket.
	//
	// example:
	//
	// 46434030
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the API request. The requestid is unique for each call.
	//
	// example:
	//
	// C1DA4C6F-963E-5741-AB57-67A554D102FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is normal.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplyTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplyTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReplyTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReplyTicketResponseBody) GetData() *string {
	return s.Data
}

func (s *ReplyTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplyTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplyTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReplyTicketResponseBody) SetAccessDeniedDetail(v string) *ReplyTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReplyTicketResponseBody) SetCode(v int32) *ReplyTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ReplyTicketResponseBody) SetData(v string) *ReplyTicketResponseBody {
	s.Data = &v
	return s
}

func (s *ReplyTicketResponseBody) SetMessage(v string) *ReplyTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ReplyTicketResponseBody) SetRequestId(v string) *ReplyTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyTicketResponseBody) SetSuccess(v bool) *ReplyTicketResponseBody {
	s.Success = &v
	return s
}

func (s *ReplyTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
