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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**. Default value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize        *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private or public IP address that is used to connect to the self-managed PostgreSQL instance.
	//
	// 	- If the self-managed PostgreSQL instance resides on an Elastic Compute Service (ECS) instance, enter the private IP address of the ECS instance. For more information about how to obtain the private IP address of an ECS instance, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// 	- If the self-managed PostgreSQL instance resides in an on-premises data center, enter the private IP address of the on-premises data center.
	//
	// example:
	//
	// 172.2.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The port number that is used to connect to the self-managed PostgreSQL instance. You can run the netstat -a | grep PGSQL command to obtain the port number.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The task ID. You can obtain the task ID from the response that is returned after you call the CreateCloudMigrationPrecheckTask operation to create the task.
	//
	// example:
	//
	// 439946016
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name. You can obtain the task name from the response that is returned after you call the CreateCloudMigrationPrecheckTask operation to create the task.
	//
	// example:
	//
	// slf7w7wj3g
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
