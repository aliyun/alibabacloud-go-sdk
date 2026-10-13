// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateOidcAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitiateOidcAuthResponseBody
	GetCode() *string
	SetMessage(v string) *InitiateOidcAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitiateOidcAuthResponseBody
	GetRequestId() *string
}

type InitiateOidcAuthResponseBody struct {
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

func (s InitiateOidcAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiateOidcAuthResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateOidcAuthResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitiateOidcAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitiateOidcAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitiateOidcAuthResponseBody) SetCode(v string) *InitiateOidcAuthResponseBody {
	s.Code = &v
	return s
}

func (s *InitiateOidcAuthResponseBody) SetMessage(v string) *InitiateOidcAuthResponseBody {
	s.Message = &v
	return s
}

func (s *InitiateOidcAuthResponseBody) SetRequestId(v string) *InitiateOidcAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiateOidcAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
