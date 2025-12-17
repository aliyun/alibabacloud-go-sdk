// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateServiceResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceResponseBody
	GetRequestId() *string
}

type UpdateServiceResponseBody struct {
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
	// B725275B-50C6-5A49-A9FD-F0332FCB3351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceResponseBody) SetCode(v string) *UpdateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceResponseBody) SetMessage(v string) *UpdateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
