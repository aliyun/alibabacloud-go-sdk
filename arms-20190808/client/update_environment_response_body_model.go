// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEnvironmentResponseBody
	GetCode() *int32
	SetData(v string) *UpdateEnvironmentResponseBody
	GetData() *string
	SetMessage(v string) *UpdateEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvironmentResponseBody
	GetRequestId() *string
}

type UpdateEnvironmentResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70675725-8F11-4817-8106-CFE0AD71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEnvironmentResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvironmentResponseBody) SetCode(v int32) *UpdateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetData(v string) *UpdateEnvironmentResponseBody {
	s.Data = &v
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
