// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickAddTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuickAddTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuickAddTaskResponseBody
	GetCode() *string
	SetMessage(v string) *QuickAddTaskResponseBody
	GetMessage() *string
	SetModel(v *QuickAddTaskResponseBodyModel) *QuickAddTaskResponseBody
	GetModel() *QuickAddTaskResponseBodyModel
	SetRequestId(v string) *QuickAddTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuickAddTaskResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *QuickAddTaskResponseBody
	GetTimestamp() *int64
}

type QuickAddTaskResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QuickAddTaskResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 77
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QuickAddTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QuickAddTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuickAddTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuickAddTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuickAddTaskResponseBody) GetModel() *QuickAddTaskResponseBodyModel {
	return s.Model
}

func (s *QuickAddTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuickAddTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuickAddTaskResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QuickAddTaskResponseBody) SetAccessDeniedDetail(v string) *QuickAddTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuickAddTaskResponseBody) SetCode(v string) *QuickAddTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QuickAddTaskResponseBody) SetMessage(v string) *QuickAddTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QuickAddTaskResponseBody) SetModel(v *QuickAddTaskResponseBodyModel) *QuickAddTaskResponseBody {
	s.Model = v
	return s
}

func (s *QuickAddTaskResponseBody) SetRequestId(v string) *QuickAddTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuickAddTaskResponseBody) SetSuccess(v bool) *QuickAddTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QuickAddTaskResponseBody) SetTimestamp(v int64) *QuickAddTaskResponseBody {
	s.Timestamp = &v
	return s
}

func (s *QuickAddTaskResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuickAddTaskResponseBodyModel struct {
	// 任务id
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QuickAddTaskResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QuickAddTaskResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QuickAddTaskResponseBodyModel) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QuickAddTaskResponseBodyModel) SetTaskId(v int64) *QuickAddTaskResponseBodyModel {
	s.TaskId = &v
	return s
}

func (s *QuickAddTaskResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
