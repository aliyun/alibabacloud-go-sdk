// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DiduiAreaDeductionProResponseBody
	GetCode() *string
	SetData(v *DiduiAreaDeductionProResponseBodyData) *DiduiAreaDeductionProResponseBody
	GetData() *DiduiAreaDeductionProResponseBodyData
	SetMessage(v string) *DiduiAreaDeductionProResponseBody
	GetMessage() *string
	SetRequestId(v string) *DiduiAreaDeductionProResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DiduiAreaDeductionProResponseBody
	GetSuccess() *bool
}

type DiduiAreaDeductionProResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submit status of the asynchronous task.
	Data *DiduiAreaDeductionProResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message or failure description.
	//
	// example:
	//
	// Task submitted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiduiAreaDeductionProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionProResponseBody) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionProResponseBody) GetCode() *string {
	return s.Code
}

func (s *DiduiAreaDeductionProResponseBody) GetData() *DiduiAreaDeductionProResponseBodyData {
	return s.Data
}

func (s *DiduiAreaDeductionProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DiduiAreaDeductionProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiduiAreaDeductionProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DiduiAreaDeductionProResponseBody) SetCode(v string) *DiduiAreaDeductionProResponseBody {
	s.Code = &v
	return s
}

func (s *DiduiAreaDeductionProResponseBody) SetData(v *DiduiAreaDeductionProResponseBodyData) *DiduiAreaDeductionProResponseBody {
	s.Data = v
	return s
}

func (s *DiduiAreaDeductionProResponseBody) SetMessage(v string) *DiduiAreaDeductionProResponseBody {
	s.Message = &v
	return s
}

func (s *DiduiAreaDeductionProResponseBody) SetRequestId(v string) *DiduiAreaDeductionProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiduiAreaDeductionProResponseBody) SetSuccess(v bool) *DiduiAreaDeductionProResponseBody {
	s.Success = &v
	return s
}

func (s *DiduiAreaDeductionProResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiduiAreaDeductionProResponseBodyData struct {
	// The ID of the asynchronous task, which is used to call QueryAsyncTaskResult to query the task result.
	//
	// example:
	//
	// task_778xxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DiduiAreaDeductionProResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionProResponseBodyData) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionProResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DiduiAreaDeductionProResponseBodyData) SetTaskId(v string) *DiduiAreaDeductionProResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DiduiAreaDeductionProResponseBodyData) Validate() error {
	return dara.Validate(s)
}
