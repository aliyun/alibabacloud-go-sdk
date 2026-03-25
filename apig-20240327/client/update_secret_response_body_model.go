// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecretResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecretResponseBody
	GetRequestId() *string
}

type UpdateSecretResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// B01CABA5-9AA1-51A2-A5A7-8BD83BA062F7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecretResponseBody) SetCode(v string) *UpdateSecretResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecretResponseBody) SetMessage(v string) *UpdateSecretResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecretResponseBody) SetRequestId(v string) *UpdateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
