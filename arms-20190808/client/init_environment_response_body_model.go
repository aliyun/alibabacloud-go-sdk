// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InitEnvironmentResponseBody
	GetCode() *int32
	SetData(v string) *InitEnvironmentResponseBody
	GetData() *string
	SetMessage(v string) *InitEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitEnvironmentResponseBody
	GetRequestId() *string
}

type InitEnvironmentResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
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
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InitEnvironmentResponseBody) GetData() *string {
	return s.Data
}

func (s *InitEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitEnvironmentResponseBody) SetCode(v int32) *InitEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *InitEnvironmentResponseBody) SetData(v string) *InitEnvironmentResponseBody {
	s.Data = &v
	return s
}

func (s *InitEnvironmentResponseBody) SetMessage(v string) *InitEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *InitEnvironmentResponseBody) SetRequestId(v string) *InitEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
