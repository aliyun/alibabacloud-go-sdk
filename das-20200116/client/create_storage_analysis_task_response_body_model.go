// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateStorageAnalysisTaskResponseBody
	GetCode() *int64
	SetData(v *CreateStorageAnalysisTaskResponseBodyData) *CreateStorageAnalysisTaskResponseBody
	GetData() *CreateStorageAnalysisTaskResponseBodyData
	SetMessage(v string) *CreateStorageAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStorageAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStorageAnalysisTaskResponseBody
	GetSuccess() *bool
}

type CreateStorageAnalysisTaskResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateStorageAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request is successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStorageAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageAnalysisTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateStorageAnalysisTaskResponseBody) GetData() *CreateStorageAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *CreateStorageAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStorageAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStorageAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStorageAnalysisTaskResponseBody) SetCode(v int64) *CreateStorageAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBody) SetData(v *CreateStorageAnalysisTaskResponseBodyData) *CreateStorageAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBody) SetMessage(v string) *CreateStorageAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBody) SetRequestId(v string) *CreateStorageAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBody) SetSuccess(v bool) *CreateStorageAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStorageAnalysisTaskResponseBodyData struct {
	// Indicates whether the task is created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CreateTaskSuccess *bool `json:"CreateTaskSuccess,omitempty" xml:"CreateTaskSuccess,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// unknown error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 910f83f4b96df0524ddc5749f61539f8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateStorageAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStorageAnalysisTaskResponseBodyData) GetCreateTaskSuccess() *bool {
	return s.CreateTaskSuccess
}

func (s *CreateStorageAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateStorageAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateStorageAnalysisTaskResponseBodyData) SetCreateTaskSuccess(v bool) *CreateStorageAnalysisTaskResponseBodyData {
	s.CreateTaskSuccess = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBodyData) SetErrorMessage(v string) *CreateStorageAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBodyData) SetTaskId(v string) *CreateStorageAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateStorageAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
