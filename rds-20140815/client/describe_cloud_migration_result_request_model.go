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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize        *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private IP address that is used to connect to the self-managed PostgreSQL instance.
	//
	// 	- If the self-managed PostgreSQL instance resides on an Elastic Compute Service (ECS) instance, enter the private IP address of the ECS instance. For more information about how to obtain the private IP address of an ECS instance, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// 	- If the self-managed PostgreSQL instance resides in a data center, enter the private IP address of the data center.
	//
	// example:
	//
	// 172.16.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The port number that is used to connect to the self-managed PostgreSQL instance. You can run the netstat -a | grep PGSQL command to obtain the port number.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The task ID. You can obtain the task ID from the response that is returned when you call the CreateCloudMigrationTask operation to create the task.
	//
	// example:
	//
	// 440437220
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. You can obtain the task name from the response that is returned when you call the CreateCloudMigrationTask operation to create the task.
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
