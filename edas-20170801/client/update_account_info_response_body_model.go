// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAccountInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateAccountInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAccountInfoResponseBody
	GetRequestId() *string
}

type UpdateAccountInfoResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// edit successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ECD1D6FC-4307-4583-BA6F-215F3857EAF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAccountInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountInfoResponseBody) SetCode(v int32) *UpdateAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountInfoResponseBody) SetMessage(v string) *UpdateAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountInfoResponseBody) SetRequestId(v string) *UpdateAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
