// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConnectionResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConnectionResponseBody
	GetRequestId() *string
}

type UpdateConnectionResponseBody struct {
	// The response code. Valid value: 200, which indicates that the request was successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8346BE8F-40F3-533D-A0B8-1359C31BD5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectionResponseBody) SetCode(v string) *UpdateConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetMessage(v string) *UpdateConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetRequestId(v string) *UpdateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
