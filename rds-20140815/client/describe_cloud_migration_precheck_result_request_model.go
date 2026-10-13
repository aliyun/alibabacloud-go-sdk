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
	// The instance ID of the target instance. You can invoke the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The page number. Valid values: integers greater than 0 and up to the maximum value of the Integer type. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**.
	//
	// Default value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize        *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The internal IP address or public IP address of the self-managed PostgreSQL database.
	//
	// - In the one-click cloud migration scenario for a self-managed PostgreSQL database on an ECS instance, set this parameter to the private IP address of the ECS instance. For more information about how to obtain the IP address, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// - In the one-click cloud migration scenario for a self-managed PostgreSQL database in an on-premises data center, set this parameter to the internal IP address of the on-premises data center.
	//
	// example:
	//
	// 172.2.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The port of the self-managed PostgreSQL database. You can run the netstat -a | grep PGSQL command to view the port.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The ID of the task. You can obtain the task ID from the response message after you call the CreateCloudMigrationPrecheckTask operation to create a pre-check task for one-click cloud migration.
	//
	// example:
	//
	// 439946016
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task. You can obtain the task name from the response message after you call the CreateCloudMigrationPrecheckTask operation to create a pre-check task for one-click cloud migration.
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
