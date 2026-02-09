// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEnvironmentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvironmentResponseBody
	GetRequestId() *string
}

type UpdateEnvironmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 52FB803B-3CD8-5FF8-AAE9-C2B841F6A483
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvironmentResponseBody) SetCode(v string) *UpdateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetMessage(v string) *UpdateEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetRequestId(v string) *UpdateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
