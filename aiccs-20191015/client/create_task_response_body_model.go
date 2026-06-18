// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskResponseBody
	GetCode() *string
	SetData(v int64) *CreateTaskResponseBody
	GetData() *int64
	SetMessage(v string) *CreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskResponseBody
	GetSuccess() *bool
}

type CreateTaskResponseBody struct {
	// Request status code. A return value of "OK" indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Unique job ID of the robot outbound calling task.
	//
	// example:
	//
	// 12****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskResponseBody) SetCode(v string) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetData(v int64) *CreateTaskResponseBody {
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

func (s *CreateTaskResponseBody) SetSuccess(v bool) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
