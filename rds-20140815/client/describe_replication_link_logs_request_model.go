// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationLinkLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReplicationLinkLogsRequest
	GetDBInstanceId() *string
	SetPageNumber(v int64) *DescribeReplicationLinkLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeReplicationLinkLogsRequest
	GetPageSize() *int64
	SetTaskId(v int64) *DescribeReplicationLinkLogsRequest
	GetTaskId() *int64
	SetTaskName(v string) *DescribeReplicationLinkLogsRequest
	GetTaskName() *string
	SetTaskType(v string) *DescribeReplicationLinkLogsRequest
	GetTaskType() *string
}

type DescribeReplicationLinkLogsRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeReplicationLinkLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationLinkLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicationLinkLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReplicationLinkLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeReplicationLinkLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeReplicationLinkLogsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeReplicationLinkLogsRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeReplicationLinkLogsRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeReplicationLinkLogsRequest) SetDBInstanceId(v string) *DescribeReplicationLinkLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) SetPageNumber(v int64) *DescribeReplicationLinkLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) SetPageSize(v int64) *DescribeReplicationLinkLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) SetTaskId(v int64) *DescribeReplicationLinkLogsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) SetTaskName(v string) *DescribeReplicationLinkLogsRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) SetTaskType(v string) *DescribeReplicationLinkLogsRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeReplicationLinkLogsRequest) Validate() error {
	return dara.Validate(s)
}
