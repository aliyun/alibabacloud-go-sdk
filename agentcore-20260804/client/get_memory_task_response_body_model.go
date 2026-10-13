// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryTaskResponseBody
	GetCode() *string
	SetData(v *GetMemoryTaskResponseBodyData) *GetMemoryTaskResponseBody
	GetData() *GetMemoryTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetMemoryTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMemoryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMemoryTaskResponseBody
	GetSuccess() *bool
}

type GetMemoryTaskResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The task details.
	Data *GetMemoryTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMemoryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryTaskResponseBody) GetData() *GetMemoryTaskResponseBodyData {
	return s.Data
}

func (s *GetMemoryTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMemoryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMemoryTaskResponseBody) SetCode(v string) *GetMemoryTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetData(v *GetMemoryTaskResponseBodyData) *GetMemoryTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryTaskResponseBody) SetHttpStatusCode(v int32) *GetMemoryTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetMessage(v string) *GetMemoryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetRequestId(v string) *GetMemoryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetSuccess(v bool) *GetMemoryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetMemoryTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryTaskResponseBodyData struct {
	// The task error returned only when the status is FAILED or UNKNOWN. This does not indicate that the current query failed.
	Error *GetMemoryTaskResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The memories returned only when the status is SUCCEEDED. An empty array is allowed.
	Memories []*GetMemoryTaskResponseBodyDataMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// The extraction task status.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-2f5f9c5a8d0e4c3b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetMemoryTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBodyData) GetError() *GetMemoryTaskResponseBodyDataError {
	return s.Error
}

func (s *GetMemoryTaskResponseBodyData) GetMemories() []*GetMemoryTaskResponseBodyDataMemories {
	return s.Memories
}

func (s *GetMemoryTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMemoryTaskResponseBodyData) SetError(v *GetMemoryTaskResponseBodyDataError) *GetMemoryTaskResponseBodyData {
	s.Error = v
	return s
}

func (s *GetMemoryTaskResponseBodyData) SetMemories(v []*GetMemoryTaskResponseBodyDataMemories) *GetMemoryTaskResponseBodyData {
	s.Memories = v
	return s
}

func (s *GetMemoryTaskResponseBodyData) SetStatus(v string) *GetMemoryTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMemoryTaskResponseBodyData) SetTaskId(v string) *GetMemoryTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetMemoryTaskResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
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

type GetMemoryTaskResponseBodyDataError struct {
	// The task error code.
	//
	// example:
	//
	// UnknownError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The task error message.
	//
	// example:
	//
	// Memory extraction failed.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetMemoryTaskResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *GetMemoryTaskResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryTaskResponseBodyDataError) SetCode(v string) *GetMemoryTaskResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *GetMemoryTaskResponseBodyDataError) SetMessage(v string) *GetMemoryTaskResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *GetMemoryTaskResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type GetMemoryTaskResponseBodyDataMemories struct {
	// The long-term memory ID.
	//
	// example:
	//
	// mem-2f5f9c5a8d0e4c3b
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
}

func (s GetMemoryTaskResponseBodyDataMemories) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBodyDataMemories) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBodyDataMemories) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryTaskResponseBodyDataMemories) SetMemoryId(v string) *GetMemoryTaskResponseBodyDataMemories {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryTaskResponseBodyDataMemories) Validate() error {
	return dara.Validate(s)
}
