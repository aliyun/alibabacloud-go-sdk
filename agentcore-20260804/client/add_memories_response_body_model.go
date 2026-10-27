// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddMemoriesResponseBody
	GetCode() *string
	SetData(v *AddMemoriesResponseBodyData) *AddMemoriesResponseBody
	GetData() *AddMemoriesResponseBodyData
	SetHttpStatusCode(v int32) *AddMemoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddMemoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMemoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMemoriesResponseBody
	GetSuccess() *bool
}

type AddMemoriesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result of the write operation.
	Data *AddMemoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddMemoriesResponseBody) GetData() *AddMemoriesResponseBodyData {
	return s.Data
}

func (s *AddMemoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddMemoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMemoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMemoriesResponseBody) SetCode(v string) *AddMemoriesResponseBody {
	s.Code = &v
	return s
}

func (s *AddMemoriesResponseBody) SetData(v *AddMemoriesResponseBodyData) *AddMemoriesResponseBody {
	s.Data = v
	return s
}

func (s *AddMemoriesResponseBody) SetHttpStatusCode(v int32) *AddMemoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMemoriesResponseBody) SetMessage(v string) *AddMemoriesResponseBody {
	s.Message = &v
	return s
}

func (s *AddMemoriesResponseBody) SetRequestId(v string) *AddMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemoriesResponseBody) SetSuccess(v bool) *AddMemoriesResponseBody {
	s.Success = &v
	return s
}

func (s *AddMemoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMemoriesResponseBodyData struct {
	// The memory references returned only when sync is set to true. An empty array is allowed. This parameter is not returned together with taskId.
	Memories []*AddMemoriesResponseBodyDataMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// The ID of the task returned only when sync is set to false. This parameter is not returned together with memories.
	//
	// example:
	//
	// task-2f5f9c5a8d0e4c3b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AddMemoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponseBodyData) GetMemories() []*AddMemoriesResponseBodyDataMemories {
	return s.Memories
}

func (s *AddMemoriesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AddMemoriesResponseBodyData) SetMemories(v []*AddMemoriesResponseBodyDataMemories) *AddMemoriesResponseBodyData {
	s.Memories = v
	return s
}

func (s *AddMemoriesResponseBodyData) SetTaskId(v string) *AddMemoriesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AddMemoriesResponseBodyData) Validate() error {
	if s.Memories != nil {
		for _, item := range s.Memories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddMemoriesResponseBodyDataMemories struct {
	// The ID of the long-term memory.
	//
	// example:
	//
	// mem-2f5f9c5a8d0e4c3b
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
}

func (s AddMemoriesResponseBodyDataMemories) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesResponseBodyDataMemories) GoString() string {
	return s.String()
}

func (s *AddMemoriesResponseBodyDataMemories) GetMemoryId() *string {
	return s.MemoryId
}

func (s *AddMemoriesResponseBodyDataMemories) SetMemoryId(v string) *AddMemoriesResponseBodyDataMemories {
	s.MemoryId = &v
	return s
}

func (s *AddMemoriesResponseBodyDataMemories) Validate() error {
	return dara.Validate(s)
}
