// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateTaskResponseBody
	GetCode() *int32
	SetData(v int32) *CreateTaskResponseBody
	GetData() *int32
	SetMessage(v string) *CreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
}

type CreateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1994
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F439A659-3B3D-50FB-A4BC-69FBF16413C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateTaskResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) SetCode(v int32) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetData(v int32) *CreateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
