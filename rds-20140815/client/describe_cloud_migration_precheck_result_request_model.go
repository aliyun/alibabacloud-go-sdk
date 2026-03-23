// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationPrecheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeCloudMigrationPrecheckResultRequest
	GetDBInstanceName() *string
	SetPageNumber(v int64) *DescribeCloudMigrationPrecheckResultRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudMigrationPrecheckResultRequest
	GetPageSize() *int64
	SetResourceOwnerId(v int64) *DescribeCloudMigrationPrecheckResultRequest
	GetResourceOwnerId() *int64
	SetSourceIpAddress(v string) *DescribeCloudMigrationPrecheckResultRequest
	GetSourceIpAddress() *string
	SetSourcePort(v int64) *DescribeCloudMigrationPrecheckResultRequest
	GetSourcePort() *int64
	SetTaskId(v int64) *DescribeCloudMigrationPrecheckResultRequest
	GetTaskId() *int64
	SetTaskName(v string) *DescribeCloudMigrationPrecheckResultRequest
	GetTaskName() *string
}

type DescribeCloudMigrationPrecheckResultRequest struct {
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

func (s DescribeCloudMigrationPrecheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationPrecheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeCloudMigrationPrecheckResultRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetDBInstanceName(v string) *DescribeCloudMigrationPrecheckResultRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetPageNumber(v int64) *DescribeCloudMigrationPrecheckResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetPageSize(v int64) *DescribeCloudMigrationPrecheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetResourceOwnerId(v int64) *DescribeCloudMigrationPrecheckResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetSourceIpAddress(v string) *DescribeCloudMigrationPrecheckResultRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetSourcePort(v int64) *DescribeCloudMigrationPrecheckResultRequest {
	s.SourcePort = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetTaskId(v int64) *DescribeCloudMigrationPrecheckResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) SetTaskName(v string) *DescribeCloudMigrationPrecheckResultRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultRequest) Validate() error {
	return dara.Validate(s)
}
