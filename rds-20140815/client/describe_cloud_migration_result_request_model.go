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
	// The target instance ID. You can invoke the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The internal IP address of the self-managed PostgreSQL database.
	//
	// - For a one-click cloud migration of a self-managed PostgreSQL database on an ECS instance, set this parameter to the private IP address of the ECS instance. For more information, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// - For a one-click cloud migration of a self-managed PostgreSQL database in an IDC, set this parameter to the internal IP address of the IDC.
	//
	// example:
	//
	// 172.16.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The port of the self-managed PostgreSQL database. You can run the netstat -a | grep PGSQL command to query the port.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The task ID. You can obtain the task ID from the response of the CreateCloudMigrationTask operation when you create an RDS PostgreSQL cloud migration task.
	//
	// example:
	//
	// 440437220
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. You can obtain the task name from the response of the CreateCloudMigrationTask operation when you create an RDS PostgreSQL cloud migration task.
	//
	// example:
	//
	// 362c6c7a-4d20-4eac-898c-1495ceab374c
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
