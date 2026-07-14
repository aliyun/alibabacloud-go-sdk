// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenChatappServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenChatappServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenChatappServiceResponseBody
	GetMessage() *string
	SetOpenStatus(v bool) *OpenChatappServiceResponseBody
	GetOpenStatus() *bool
	SetRequestId(v string) *OpenChatappServiceResponseBody
	GetRequestId() *string
}

type OpenChatappServiceResponseBody struct {
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
	// 3D2FFEE6-368D-532D-87AA-F45B02DD28B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenChatappServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenChatappServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenChatappServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenChatappServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenChatappServiceResponseBody) GetOpenStatus() *bool {
	return s.OpenStatus
}

func (s *OpenChatappServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenChatappServiceResponseBody) SetCode(v string) *OpenChatappServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenChatappServiceResponseBody) SetMessage(v string) *OpenChatappServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenChatappServiceResponseBody) SetOpenStatus(v bool) *OpenChatappServiceResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *OpenChatappServiceResponseBody) SetRequestId(v string) *OpenChatappServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenChatappServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
