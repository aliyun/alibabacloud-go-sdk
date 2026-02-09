// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHttpApiOperationResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHttpApiOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHttpApiOperationResponseBody
	GetRequestId() *string
}

type UpdateHttpApiOperationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHttpApiOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHttpApiOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpApiOperationResponseBody) SetCode(v string) *UpdateHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiOperationResponseBody) SetMessage(v string) *UpdateHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiOperationResponseBody) SetRequestId(v string) *UpdateHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpApiOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
