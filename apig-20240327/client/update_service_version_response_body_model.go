// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateServiceVersionResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceVersionResponseBody
	GetRequestId() *string
}

type UpdateServiceVersionResponseBody struct {
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
	// A60EE5CA-1294-532A-9775-8D2FD1C6EFBF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceVersionResponseBody) SetCode(v string) *UpdateServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) SetMessage(v string) *UpdateServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) SetRequestId(v string) *UpdateServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
