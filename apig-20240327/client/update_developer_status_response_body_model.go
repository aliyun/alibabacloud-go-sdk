// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeveloperStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDeveloperStatusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateDeveloperStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDeveloperStatusResponseBody
	GetRequestId() *string
}

type UpdateDeveloperStatusResponseBody struct {
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

func (s UpdateDeveloperStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeveloperStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeveloperStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDeveloperStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDeveloperStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeveloperStatusResponseBody) SetCode(v string) *UpdateDeveloperStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeveloperStatusResponseBody) SetMessage(v string) *UpdateDeveloperStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeveloperStatusResponseBody) SetRequestId(v string) *UpdateDeveloperStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeveloperStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
