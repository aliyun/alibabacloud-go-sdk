// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableProducingTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetMetaTableProducingTasksResponseBodyData) *GetMetaTableProducingTasksResponseBody
	GetData() []*GetMetaTableProducingTasksResponseBodyData
	SetErrorCode(v string) *GetMetaTableProducingTasksResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableProducingTasksResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableProducingTasksResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableProducingTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableProducingTasksResponseBody
	GetSuccess() *bool
}

type GetMetaTableProducingTasksResponseBody struct {
	Data           []*GetMetaTableProducingTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode      *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableProducingTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableProducingTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableProducingTasksResponseBody) GetData() []*GetMetaTableProducingTasksResponseBodyData {
	return s.Data
}

func (s *GetMetaTableProducingTasksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableProducingTasksResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableProducingTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableProducingTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableProducingTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableProducingTasksResponseBody) SetData(v []*GetMetaTableProducingTasksResponseBodyData) *GetMetaTableProducingTasksResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) SetErrorCode(v string) *GetMetaTableProducingTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) SetErrorMessage(v string) *GetMetaTableProducingTasksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) SetHttpStatusCode(v int32) *GetMetaTableProducingTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) SetRequestId(v string) *GetMetaTableProducingTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) SetSuccess(v bool) *GetMetaTableProducingTasksResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMetaTableProducingTasksResponseBodyData struct {
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetMetaTableProducingTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableProducingTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableProducingTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMetaTableProducingTasksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *GetMetaTableProducingTasksResponseBodyData) SetTaskId(v string) *GetMetaTableProducingTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBodyData) SetTaskName(v string) *GetMetaTableProducingTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetMetaTableProducingTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
