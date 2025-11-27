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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The username of the account that is used to connect to the self-managed PostgreSQL instance. Enter the username of the account that you created in the [Create an account for cloud migration on a self-managed PostgreSQL instance](https://help.aliyun.com/document_detail/369500.html) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// migratetest
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// The environment in which the self-managed PostgreSQL instance runs.
	//
	// 	- **idcOnVpc**: The self-managed PostgreSQL instance resides in a data center. The data center can communicate with the VPC to which the ApsaraDB RDS for PostgreSQL instance belongs.
	//
	// 	- **ecsOnVpc**: The self-managed PostgreSQL instance resides on an ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecsOnVpc
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The private IP address that is used to connect to the self-managed PostgreSQL instance.
	//
	// 	- If the self-managed PostgreSQL instance resides on an ECS instance, enter the private IP address of the ECS instance. For more information about how to obtain the private IP address of an ECS instance, see [View IP addresses](https://help.aliyun.com/document_detail/273914.html).
	//
	// 	- If the self-managed PostgreSQL instance resides in an on-premises data center, enter the private IP address of the on-premises data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.2.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The password of the account that is used to connect to the self-managed PostgreSQL instance. Enter the password of the account that you created in the [Create an account for cloud migration on a self-managed PostgreSQL instance](https://help.aliyun.com/document_detail/369500.html) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// The port that is used to connect to the self-managed PostgreSQL instance. You can run the `netstat -a | grep PGSQL` command to view the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The name of the task. If you do not specify this parameter, ApsaraDB RDS automatically generates a name for the cloud migration task.
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
