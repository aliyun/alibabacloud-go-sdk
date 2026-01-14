// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCancelCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TaskCancelCallResponseBody
	GetCode() *string
	SetMessage(v string) *TaskCancelCallResponseBody
	GetMessage() *string
	SetModel(v *TaskCancelCallResponseBodyModel) *TaskCancelCallResponseBody
	GetModel() *TaskCancelCallResponseBodyModel
	SetRequestId(v string) *TaskCancelCallResponseBody
	GetRequestId() *string
	SetSuccess(v string) *TaskCancelCallResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *TaskCancelCallResponseBody
	GetTimestamp() *int64
}

type TaskCancelCallResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCancelCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCancelCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskCancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *TaskCancelCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskCancelCallResponseBody) GetModel() *TaskCancelCallResponseBodyModel {
	return s.Model
}

func (s *TaskCancelCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskCancelCallResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *TaskCancelCallResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskCancelCallResponseBody) SetCode(v string) *TaskCancelCallResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetMessage(v string) *TaskCancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetModel(v *TaskCancelCallResponseBodyModel) *TaskCancelCallResponseBody {
	s.Model = v
	return s
}

func (s *TaskCancelCallResponseBody) SetRequestId(v string) *TaskCancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetSuccess(v string) *TaskCancelCallResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetTimestamp(v int64) *TaskCancelCallResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskCancelCallResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TaskCancelCallResponseBodyModel struct {
	// 错误手机号
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s TaskCancelCallResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TaskCancelCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponseBodyModel) GetUnHandleNumbers() []*string {
	return s.UnHandleNumbers
}

func (s *TaskCancelCallResponseBodyModel) SetUnHandleNumbers(v []*string) *TaskCancelCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

func (s *TaskCancelCallResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
