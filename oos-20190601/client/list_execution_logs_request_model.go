// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *ListExecutionLogsRequest
	GetExecutionId() *string
	SetLogType(v string) *ListExecutionLogsRequest
	GetLogType() *string
	SetMaxResults(v int32) *ListExecutionLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExecutionLogsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListExecutionLogsRequest
	GetRegionId() *string
	SetTaskExecutionId(v string) *ListExecutionLogsRequest
	GetTaskExecutionId() *string
}

type ListExecutionLogsRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The type of the log.
	//
	// example:
	//
	// System
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which you want to query the logs of the execution.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// exec-1234567zxcvb.t0010
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
}

func (s ListExecutionLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionLogsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListExecutionLogsRequest) GetLogType() *string {
	return s.LogType
}

func (s *ListExecutionLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutionLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExecutionLogsRequest) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListExecutionLogsRequest) SetExecutionId(v string) *ListExecutionLogsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetLogType(v string) *ListExecutionLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsRequest) SetMaxResults(v int32) *ListExecutionLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsRequest) SetNextToken(v string) *ListExecutionLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionLogsRequest) SetRegionId(v string) *ListExecutionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetTaskExecutionId(v string) *ListExecutionLogsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionLogsRequest) Validate() error {
	return dara.Validate(s)
}
