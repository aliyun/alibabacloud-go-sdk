// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloseTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CloseTicketResponseBody
	GetCode() *int32
	SetMessage(v string) *CloseTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseTicketResponseBody
	GetSuccess() *bool
}

type CloseTicketResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the request result.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is normal.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloseTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CloseTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseTicketResponseBody) SetAccessDeniedDetail(v string) *CloseTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloseTicketResponseBody) SetCode(v int32) *CloseTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTicketResponseBody) SetMessage(v string) *CloseTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTicketResponseBody) SetRequestId(v string) *CloseTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTicketResponseBody) SetSuccess(v bool) *CloseTicketResponseBody {
	s.Success = &v
	return s
}

func (s *CloseTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
