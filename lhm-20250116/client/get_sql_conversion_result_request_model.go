// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *GetSqlConversionResultRequest
	GetPage() *int32
	SetSize(v int32) *GetSqlConversionResultRequest
	GetSize() *int32
	SetTaskId(v int64) *GetSqlConversionResultRequest
	GetTaskId() *int64
}

type GetSqlConversionResultRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The task ID that uniquely identifies a task.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSqlConversionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionResultRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConversionResultRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetSqlConversionResultRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetSqlConversionResultRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetSqlConversionResultRequest) SetPage(v int32) *GetSqlConversionResultRequest {
	s.Page = &v
	return s
}

func (s *GetSqlConversionResultRequest) SetSize(v int32) *GetSqlConversionResultRequest {
	s.Size = &v
	return s
}

func (s *GetSqlConversionResultRequest) SetTaskId(v int64) *GetSqlConversionResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetSqlConversionResultRequest) Validate() error {
	return dara.Validate(s)
}
