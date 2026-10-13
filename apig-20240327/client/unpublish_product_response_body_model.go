// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnpublishProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnpublishProductResponseBody
	GetCode() *string
	SetMessage(v string) *UnpublishProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnpublishProductResponseBody
	GetRequestId() *string
}

type UnpublishProductResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UnpublishProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnpublishProductResponseBody) GoString() string {
	return s.String()
}

func (s *UnpublishProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnpublishProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnpublishProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnpublishProductResponseBody) SetCode(v string) *UnpublishProductResponseBody {
	s.Code = &v
	return s
}

func (s *UnpublishProductResponseBody) SetMessage(v string) *UnpublishProductResponseBody {
	s.Message = &v
	return s
}

func (s *UnpublishProductResponseBody) SetRequestId(v string) *UnpublishProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnpublishProductResponseBody) Validate() error {
	return dara.Validate(s)
}
