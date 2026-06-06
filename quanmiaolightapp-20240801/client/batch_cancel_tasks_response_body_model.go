// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCancelTasksResponseBody
	GetCode() *string
	SetData(v *BatchCancelTasksResponseBodyData) *BatchCancelTasksResponseBody
	GetData() *BatchCancelTasksResponseBodyData
	SetHttpStatusCode(v int32) *BatchCancelTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchCancelTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCancelTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCancelTasksResponseBody
	GetSuccess() *bool
}

type BatchCancelTasksResponseBody struct {
	// example:
	//
	// Success
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchCancelTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-GHIJ-KLMNOPQRST
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchCancelTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCancelTasksResponseBody) GetData() *BatchCancelTasksResponseBodyData {
	return s.Data
}

func (s *BatchCancelTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchCancelTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCancelTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCancelTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCancelTasksResponseBody) SetCode(v string) *BatchCancelTasksResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCancelTasksResponseBody) SetData(v *BatchCancelTasksResponseBodyData) *BatchCancelTasksResponseBody {
	s.Data = v
	return s
}

func (s *BatchCancelTasksResponseBody) SetHttpStatusCode(v int32) *BatchCancelTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchCancelTasksResponseBody) SetMessage(v string) *BatchCancelTasksResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCancelTasksResponseBody) SetRequestId(v string) *BatchCancelTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCancelTasksResponseBody) SetSuccess(v bool) *BatchCancelTasksResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCancelTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCancelTasksResponseBodyData struct {
	Results []*BatchCancelTasksResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s BatchCancelTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksResponseBodyData) GetResults() []*BatchCancelTasksResponseBodyDataResults {
	return s.Results
}

func (s *BatchCancelTasksResponseBodyData) SetResults(v []*BatchCancelTasksResponseBodyDataResults) *BatchCancelTasksResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchCancelTasksResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCancelTasksResponseBodyDataResults struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchCancelTasksResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksResponseBodyDataResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchCancelTasksResponseBodyDataResults) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCancelTasksResponseBodyDataResults) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchCancelTasksResponseBodyDataResults) SetErrorMessage(v string) *BatchCancelTasksResponseBodyDataResults {
	s.ErrorMessage = &v
	return s
}

func (s *BatchCancelTasksResponseBodyDataResults) SetSuccess(v bool) *BatchCancelTasksResponseBodyDataResults {
	s.Success = &v
	return s
}

func (s *BatchCancelTasksResponseBodyDataResults) SetTaskId(v string) *BatchCancelTasksResponseBodyDataResults {
	s.TaskId = &v
	return s
}

func (s *BatchCancelTasksResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
