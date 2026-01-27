// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeImportTaskRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeImportTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImportTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeImportTaskRequest
	GetTaskId() *string
}

type DescribeImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 159****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeImportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImportTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeImportTaskRequest) SetDBInstanceId(v string) *DescribeImportTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeImportTaskRequest) SetOwnerId(v int64) *DescribeImportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImportTaskRequest) SetRegionId(v string) *DescribeImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImportTaskRequest) SetTaskId(v string) *DescribeImportTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
