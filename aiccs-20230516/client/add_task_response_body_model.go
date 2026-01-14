// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *AddTaskResponseBody
	GetCode() *int64
	SetMessage(v string) *AddTaskResponseBody
	GetMessage() *string
	SetModel(v *AddTaskResponseBodyModel) *AddTaskResponseBody
	GetModel() *AddTaskResponseBodyModel
	SetRequestId(v string) *AddTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTaskResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AddTaskResponseBody
	GetTimestamp() *int64
}

type AddTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AddTaskResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 5453cc9b-02bc-4dbb-9527-f28a9100b912
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1686225227338
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTaskResponseBody) GetModel() *AddTaskResponseBodyModel {
	return s.Model
}

func (s *AddTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTaskResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AddTaskResponseBody) SetAccessDeniedDetail(v string) *AddTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddTaskResponseBody) SetCode(v int64) *AddTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AddTaskResponseBody) SetMessage(v string) *AddTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AddTaskResponseBody) SetModel(v *AddTaskResponseBodyModel) *AddTaskResponseBody {
	s.Model = v
	return s
}

func (s *AddTaskResponseBody) SetRequestId(v string) *AddTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTaskResponseBody) SetSuccess(v bool) *AddTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AddTaskResponseBody) SetTimestamp(v int64) *AddTaskResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AddTaskResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTaskResponseBodyModel struct {
	// 任务ID
	//
	// example:
	//
	// 47
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddTaskResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AddTaskResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddTaskResponseBodyModel) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AddTaskResponseBodyModel) SetTaskId(v int64) *AddTaskResponseBodyModel {
	s.TaskId = &v
	return s
}

func (s *AddTaskResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
