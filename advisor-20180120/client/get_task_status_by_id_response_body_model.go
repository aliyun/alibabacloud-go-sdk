// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTaskStatusByIdResponseBodyData) *GetTaskStatusByIdResponseBody
	GetData() *GetTaskStatusByIdResponseBodyData
	SetRequestId(v string) *GetTaskStatusByIdResponseBody
	GetRequestId() *string
}

type GetTaskStatusByIdResponseBody struct {
	Data *GetTaskStatusByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskStatusByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponseBody) GetData() *GetTaskStatusByIdResponseBodyData {
	return s.Data
}

func (s *GetTaskStatusByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatusByIdResponseBody) SetData(v *GetTaskStatusByIdResponseBodyData) *GetTaskStatusByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskStatusByIdResponseBody) SetRequestId(v string) *GetTaskStatusByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskStatusByIdResponseBodyData struct {
	// example:
	//
	// 95906135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskStatusByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetTaskStatusByIdResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetTaskStatusByIdResponseBodyData) SetTaskId(v int64) *GetTaskStatusByIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusByIdResponseBodyData) SetTaskStatus(v int32) *GetTaskStatusByIdResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetTaskStatusByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
