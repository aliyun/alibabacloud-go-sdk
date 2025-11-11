// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAsyncTaskResponseBody
	GetCode() *string
	SetData(v *SubmitAsyncTaskResponseBodyData) *SubmitAsyncTaskResponseBody
	GetData() *SubmitAsyncTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitAsyncTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitAsyncTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAsyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAsyncTaskResponseBody
	GetSuccess() *bool
}

type SubmitAsyncTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAsyncTaskResponseBody) GetData() *SubmitAsyncTaskResponseBodyData {
	return s.Data
}

func (s *SubmitAsyncTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAsyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAsyncTaskResponseBody) SetCode(v string) *SubmitAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetData(v *SubmitAsyncTaskResponseBodyData) *SubmitAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetHttpStatusCode(v int32) *SubmitAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetMessage(v string) *SubmitAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetRequestId(v string) *SubmitAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetSuccess(v bool) *SubmitAsyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAsyncTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {}
	TaskIntermediateResult interface{} `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitAsyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitAsyncTaskResponseBodyData) GetTaskIntermediateResult() interface{} {
	return s.TaskIntermediateResult
}

func (s *SubmitAsyncTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskId(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskIntermediateResult(v interface{}) *SubmitAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskName(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
