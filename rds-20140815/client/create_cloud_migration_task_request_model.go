// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationTaskRequest
	GetDBInstanceName() *string
	SetResourceOwnerId(v int64) *CreateCloudMigrationTaskRequest
	GetResourceOwnerId() *int64
	SetSourceAccount(v string) *CreateCloudMigrationTaskRequest
	GetSourceAccount() *string
	SetSourceCategory(v string) *CreateCloudMigrationTaskRequest
	GetSourceCategory() *string
	SetSourceIpAddress(v string) *CreateCloudMigrationTaskRequest
	GetSourceIpAddress() *string
	SetSourcePassword(v string) *CreateCloudMigrationTaskRequest
	GetSourcePassword() *string
	SetSourcePort(v int64) *CreateCloudMigrationTaskRequest
	GetSourcePort() *int64
	SetTaskName(v string) *CreateCloudMigrationTaskRequest
	GetTaskName() *string
}

type CreateCloudMigrationTaskRequest struct {
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
	// The category of the source instance.
	//
	// - **aliyunRDS**: ApsaraDB RDS instance.
	//
	// - **other**: other.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyunRDS
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The internal or public IP address of the self-managed PostgreSQL database.
	//
	// - To migrate a self-managed PostgreSQL database on an ECS instance to the cloud, set this parameter to the private IP address of the ECS instance. For more information about how to obtain the IP address, see [View IP addresses](https://help.aliyun.com/document_detail/98677.html).
	//
	// - To migrate a self-managed PostgreSQL database in an Internet Data Center (IDC) to the cloud, set this parameter to the internal IP address of the IDC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.XX.XX
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
	// 362c6c7a-4d20-4eac-898c-1495ceab374c
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateCloudMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCloudMigrationTaskRequest) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *CreateCloudMigrationTaskRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CreateCloudMigrationTaskRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateCloudMigrationTaskRequest) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *CreateCloudMigrationTaskRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *CreateCloudMigrationTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationTaskRequest) SetDBInstanceName(v string) *CreateCloudMigrationTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetResourceOwnerId(v int64) *CreateCloudMigrationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceAccount(v string) *CreateCloudMigrationTaskRequest {
	s.SourceAccount = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceCategory(v string) *CreateCloudMigrationTaskRequest {
	s.SourceCategory = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceIpAddress(v string) *CreateCloudMigrationTaskRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourcePassword(v string) *CreateCloudMigrationTaskRequest {
	s.SourcePassword = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourcePort(v int64) *CreateCloudMigrationTaskRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetTaskName(v string) *CreateCloudMigrationTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
