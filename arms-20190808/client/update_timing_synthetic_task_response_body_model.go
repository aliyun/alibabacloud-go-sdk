// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v *UpdateTimingSyntheticTaskResponseBodyData) *UpdateTimingSyntheticTaskResponseBody
	GetData() *UpdateTimingSyntheticTaskResponseBodyData
	SetMessage(v string) *UpdateTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTimingSyntheticTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTimingSyntheticTaskResponseBody
	GetSuccess() *bool
}

type UpdateTimingSyntheticTaskResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *UpdateTimingSyntheticTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateTimingSyntheticTaskResponseBody) GetData() *UpdateTimingSyntheticTaskResponseBodyData {
	return s.Data
}

func (s *UpdateTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTimingSyntheticTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTimingSyntheticTaskResponseBody) SetCode(v int64) *UpdateTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBody) SetData(v *UpdateTimingSyntheticTaskResponseBodyData) *UpdateTimingSyntheticTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBody) SetMessage(v string) *UpdateTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBody) SetRequestId(v string) *UpdateTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBody) SetSuccess(v bool) *UpdateTimingSyntheticTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 1eeb351722c84e05b52c82fd0dc9953e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTimingSyntheticTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTimingSyntheticTaskResponseBodyData) SetTaskId(v string) *UpdateTimingSyntheticTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
