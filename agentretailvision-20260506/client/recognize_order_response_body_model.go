// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RecognizeOrderResponseBody
	GetCode() *string
	SetData(v *RecognizeOrderResponseBodyData) *RecognizeOrderResponseBody
	GetData() *RecognizeOrderResponseBodyData
	SetMessage(v string) *RecognizeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecognizeOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecognizeOrderResponseBody
	GetSuccess() *bool
}

type RecognizeOrderResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RecognizeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecognizeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *RecognizeOrderResponseBody) GetData() *RecognizeOrderResponseBodyData {
	return s.Data
}

func (s *RecognizeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecognizeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecognizeOrderResponseBody) SetCode(v string) *RecognizeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeOrderResponseBody) SetData(v *RecognizeOrderResponseBodyData) *RecognizeOrderResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeOrderResponseBody) SetMessage(v string) *RecognizeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeOrderResponseBody) SetRequestId(v string) *RecognizeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeOrderResponseBody) SetSuccess(v bool) *RecognizeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *RecognizeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeOrderResponseBodyData struct {
	// example:
	//
	// ORDER_001
	OrderUniqueId *string `json:"OrderUniqueId,omitempty" xml:"OrderUniqueId,omitempty"`
	// example:
	//
	// TASK_001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// PROCESSING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s RecognizeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeOrderResponseBodyData) GetOrderUniqueId() *string {
	return s.OrderUniqueId
}

func (s *RecognizeOrderResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RecognizeOrderResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *RecognizeOrderResponseBodyData) SetOrderUniqueId(v string) *RecognizeOrderResponseBodyData {
	s.OrderUniqueId = &v
	return s
}

func (s *RecognizeOrderResponseBodyData) SetTaskId(v string) *RecognizeOrderResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RecognizeOrderResponseBodyData) SetTaskStatus(v string) *RecognizeOrderResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *RecognizeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
