// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *GetSqlConversionProgressRequest
	GetTaskId() *int64
}

type GetSqlConversionProgressRequest struct {
	// The task ID that uniquely identifies a task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSqlConversionProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionProgressRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConversionProgressRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetSqlConversionProgressRequest) SetTaskId(v int64) *GetSqlConversionProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetSqlConversionProgressRequest) Validate() error {
	return dara.Validate(s)
}
