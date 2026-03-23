// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCloudMigrationResultRequest
	GetDBInstanceName() *string
	SetPageNumber(v int64) *DescribeCloudMigrationResultRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudMigrationResultRequest
	GetPageSize() *int64
	SetResourceOwnerId(v int64) *DescribeCloudMigrationResultRequest
	GetResourceOwnerId() *int64
	SetSourceIpAddress(v string) *DescribeCloudMigrationResultRequest
	GetSourceIpAddress() *string
	SetSourcePort(v int64) *DescribeCloudMigrationResultRequest
	GetSourcePort() *int64
	SetTaskId(v int64) *DescribeCloudMigrationResultRequest
	GetTaskId() *int64
	SetTaskName(v string) *DescribeCloudMigrationResultRequest
	GetTaskName() *string
}

type DescribeCloudMigrationResultRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	SourcePort      *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeCloudMigrationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationResultRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCloudMigrationResultRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudMigrationResultRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudMigrationResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCloudMigrationResultRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeCloudMigrationResultRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeCloudMigrationResultRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeCloudMigrationResultRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCloudMigrationResultRequest) SetDBInstanceName(v string) *DescribeCloudMigrationResultRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetPageNumber(v int64) *DescribeCloudMigrationResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetPageSize(v int64) *DescribeCloudMigrationResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetResourceOwnerId(v int64) *DescribeCloudMigrationResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetSourceIpAddress(v string) *DescribeCloudMigrationResultRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetSourcePort(v int64) *DescribeCloudMigrationResultRequest {
	s.SourcePort = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetTaskId(v int64) *DescribeCloudMigrationResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) SetTaskName(v string) *DescribeCloudMigrationResultRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeCloudMigrationResultRequest) Validate() error {
	return dara.Validate(s)
}
