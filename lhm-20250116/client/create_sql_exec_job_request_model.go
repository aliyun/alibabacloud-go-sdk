// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlExecJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *CreateSqlExecJobRequest
	GetConcurrency() *int32
	SetTaskId(v int64) *CreateSqlExecJobRequest
	GetTaskId() *int64
}

type CreateSqlExecJobRequest struct {
	// The concurrency level. This is an optional parameter.
	//
	// example:
	//
	// 5
	Concurrency *int32 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateSqlExecJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlExecJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlExecJobRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateSqlExecJobRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateSqlExecJobRequest) SetConcurrency(v int32) *CreateSqlExecJobRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateSqlExecJobRequest) SetTaskId(v int64) *CreateSqlExecJobRequest {
	s.TaskId = &v
	return s
}

func (s *CreateSqlExecJobRequest) Validate() error {
	return dara.Validate(s)
}
