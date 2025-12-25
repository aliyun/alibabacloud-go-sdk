// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReopenTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ReopenTicketResponseBody
	GetCode() *int32
	SetData(v string) *ReopenTicketResponseBody
	GetData() *string
	SetMessage(v string) *ReopenTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReopenTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReopenTicketResponseBody
	GetSuccess() *bool
}

type ReopenTicketResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the request result.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned after the call succeeds.
	//
	// example:
	//
	// 0005PYGCW
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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

func (s ReopenTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReopenTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReopenTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReopenTicketResponseBody) GetData() *string {
	return s.Data
}

func (s *ReopenTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReopenTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReopenTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReopenTicketResponseBody) SetAccessDeniedDetail(v string) *ReopenTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReopenTicketResponseBody) SetCode(v int32) *ReopenTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ReopenTicketResponseBody) SetData(v string) *ReopenTicketResponseBody {
	s.Data = &v
	return s
}

func (s *ReopenTicketResponseBody) SetMessage(v string) *ReopenTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ReopenTicketResponseBody) SetRequestId(v string) *ReopenTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenTicketResponseBody) SetSuccess(v bool) *ReopenTicketResponseBody {
	s.Success = &v
	return s
}

func (s *ReopenTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
