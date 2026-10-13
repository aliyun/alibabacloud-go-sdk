// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogoutDeveloperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LogoutDeveloperResponseBody
	GetCode() *string
	SetMessage(v string) *LogoutDeveloperResponseBody
	GetMessage() *string
	SetRequestId(v string) *LogoutDeveloperResponseBody
	GetRequestId() *string
}

type LogoutDeveloperResponseBody struct {
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

func (s LogoutDeveloperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LogoutDeveloperResponseBody) GoString() string {
	return s.String()
}

func (s *LogoutDeveloperResponseBody) GetCode() *string {
	return s.Code
}

func (s *LogoutDeveloperResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LogoutDeveloperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LogoutDeveloperResponseBody) SetCode(v string) *LogoutDeveloperResponseBody {
	s.Code = &v
	return s
}

func (s *LogoutDeveloperResponseBody) SetMessage(v string) *LogoutDeveloperResponseBody {
	s.Message = &v
	return s
}

func (s *LogoutDeveloperResponseBody) SetRequestId(v string) *LogoutDeveloperResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogoutDeveloperResponseBody) Validate() error {
	return dara.Validate(s)
}
