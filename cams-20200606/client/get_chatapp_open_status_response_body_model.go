// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappOpenStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChatappOpenStatusResponseBody
	GetCode() *string
	SetMessage(v string) *GetChatappOpenStatusResponseBody
	GetMessage() *string
	SetOpenStatus(v bool) *GetChatappOpenStatusResponseBody
	GetOpenStatus() *bool
	SetRequestId(v string) *GetChatappOpenStatusResponseBody
	GetRequestId() *string
}

type GetChatappOpenStatusResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	OpenStatus *bool `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappOpenStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappOpenStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappOpenStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappOpenStatusResponseBody) GetOpenStatus() *bool {
	return s.OpenStatus
}

func (s *GetChatappOpenStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappOpenStatusResponseBody) SetCode(v string) *GetChatappOpenStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappOpenStatusResponseBody) SetMessage(v string) *GetChatappOpenStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappOpenStatusResponseBody) SetOpenStatus(v bool) *GetChatappOpenStatusResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *GetChatappOpenStatusResponseBody) SetRequestId(v string) *GetChatappOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappOpenStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
