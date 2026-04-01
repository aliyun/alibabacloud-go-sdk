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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1trqb4p1xd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID. You must set this parameter to the ID of the task that you create by calling the **CreateReplicationLink*	- operation for the disaster recovery instance.
	//
	// example:
	//
	// 8413252
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. You must set this parameter to the name of the task that you create by calling the **CreateReplicationLink*	- operation for the disaster recovery instance.
	//
	// example:
	//
	// test01
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **create**: creates a synchronization link.
	//
	// 	- **create-dryrun**: performs a precheck before a synchronization link is created.
	//
	// Valid values:
	//
	// 	- create: creates a replication link.
	//
	// 	- create-dryrun: performs a precheck before a replication link is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// create
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
