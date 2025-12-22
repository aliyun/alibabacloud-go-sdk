// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoCreationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVideoCreationTaskResponseBody
	GetCode() *string
	SetData(v *CreateVideoCreationTaskResponseBodyData) *CreateVideoCreationTaskResponseBody
	GetData() *CreateVideoCreationTaskResponseBodyData
	SetMessage(v string) *CreateVideoCreationTaskResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *CreateVideoCreationTaskResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *CreateVideoCreationTaskResponseBody
	GetSuccess() *bool
}

type CreateVideoCreationTaskResponseBody struct {
	// example:
	//
	// 0
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateVideoCreationTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool   `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateVideoCreationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVideoCreationTaskResponseBody) GetData() *CreateVideoCreationTaskResponseBodyData {
	return s.Data
}

func (s *CreateVideoCreationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVideoCreationTaskResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *CreateVideoCreationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVideoCreationTaskResponseBody) SetCode(v string) *CreateVideoCreationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBody) SetData(v *CreateVideoCreationTaskResponseBodyData) *CreateVideoCreationTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateVideoCreationTaskResponseBody) SetMessage(v string) *CreateVideoCreationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBody) SetRetryAble(v bool) *CreateVideoCreationTaskResponseBody {
	s.RetryAble = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBody) SetSuccess(v bool) *CreateVideoCreationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoCreationTaskResponseBodyData struct {
	// example:
	//
	// 1
	EstimatedWaitTimeInSeconds *int64 `json:"estimatedWaitTimeInSeconds,omitempty" xml:"estimatedWaitTimeInSeconds,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateVideoCreationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoCreationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVideoCreationTaskResponseBodyData) GetEstimatedWaitTimeInSeconds() *int64 {
	return s.EstimatedWaitTimeInSeconds
}

func (s *CreateVideoCreationTaskResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoCreationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoCreationTaskResponseBodyData) SetEstimatedWaitTimeInSeconds(v int64) *CreateVideoCreationTaskResponseBodyData {
	s.EstimatedWaitTimeInSeconds = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBodyData) SetRequestId(v string) *CreateVideoCreationTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBodyData) SetTaskId(v string) *CreateVideoCreationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateVideoCreationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
