// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationPrecheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetDBInstanceName() *string
	SetResourceOwnerId(v int64) *CreateCloudMigrationPrecheckTaskRequest
	GetResourceOwnerId() *int64
	SetSourceAccount(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceAccount() *string
	SetSourceCategory(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceCategory() *string
	SetSourceIpAddress(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceIpAddress() *string
	SetSourcePassword(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourcePassword() *string
	SetSourcePort(v int64) *CreateCloudMigrationPrecheckTaskRequest
	GetSourcePort() *int64
	SetTaskName(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetTaskName() *string
}

type CreateCloudMigrationPrecheckTaskRequest struct {
	// The ID of the target instance. You can invoke the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The username. The database account created in the [Create a migration account](https://help.aliyun.com/document_detail/369500.html) step.
	//
	// This parameter is required.
	//
	// example:
	//
	// migratetest
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// The type of the self-managed PostgreSQL database. Valid values:
	//
	// - **idcOnVpc**: IDC-based self-managed PostgreSQL database (the IDC is connected to the VPC).
	//
	// - **ecsOnVpc**: ECS-based self-managed PostgreSQL database on Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecsOnVpc
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The internal IP address of the self-managed PostgreSQL database.
	//
	// - For one-click migration of an ECS-based self-managed PostgreSQL database, set this parameter to the private IP address of the ECS instance. For more information about how to obtain the IP address, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// - For one-click migration of an IDC-based self-managed PostgreSQL database, set this parameter to the internal IP address of the IDC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.2.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The password. The password of the database account created in the [Create a migration account](https://help.aliyun.com/document_detail/369500.html) step.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// The port of the self-managed PostgreSQL database. You can run the `netstat -a | grep PGSQL` command to view the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The task name. You can specify a custom name. If you do not specify this parameter, the system automatically generates a name.
	//
	// example:
	//
	// slf7w7wj3g
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateCloudMigrationPrecheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationPrecheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetResourceOwnerId(v int64) *CreateCloudMigrationPrecheckTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceAccount(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceAccount = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceCategory(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceCategory = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceIpAddress(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourcePassword(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourcePassword = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourcePort(v int64) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetTaskName(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
