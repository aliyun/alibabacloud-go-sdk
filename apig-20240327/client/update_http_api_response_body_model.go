// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHttpApiResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHttpApiResponseBody
	GetRequestId() *string
}

type UpdateHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpApiResponseBody) SetCode(v string) *UpdateHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiResponseBody) SetMessage(v string) *UpdateHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiResponseBody) SetRequestId(v string) *UpdateHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpApiResponseBody) Validate() error {
	return dara.Validate(s)
}
