// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeTableRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeTableRequest
	GetDatabase() *string
	SetOwnerId(v int64) *DescribeTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTableRequest
	GetRegionId() *string
	SetSchema(v string) *DescribeTableRequest
	GetSchema() *string
	SetSecretArn(v string) *DescribeTableRequest
	GetSecretArn() *string
	SetTable(v string) *DescribeTableRequest
	GetTable() *string
	SetWorkspaceId(v string) *DescribeTableRequest
	GetWorkspaceId() *string
}

type DescribeTableRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the schema to which the table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the access credential for the created Data API account. You can call the CreateSecret operation to create an access credential.
	//
	// >  To call the DescribeTable operation as a Resource Access Management (RAM) user, the RAM user must have the permissions to call the UseSecret or GetSecretValue operation on the ARN of the access credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// mytable
	Table       *string `json:"Table,omitempty" xml:"Table,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeTableRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeTableRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DescribeTableRequest) GetTable() *string {
	return s.Table
}

func (s *DescribeTableRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeTableRequest) SetDBInstanceId(v string) *DescribeTableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTableRequest) SetDatabase(v string) *DescribeTableRequest {
	s.Database = &v
	return s
}

func (s *DescribeTableRequest) SetOwnerId(v int64) *DescribeTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableRequest) SetRegionId(v string) *DescribeTableRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableRequest) SetSchema(v string) *DescribeTableRequest {
	s.Schema = &v
	return s
}

func (s *DescribeTableRequest) SetSecretArn(v string) *DescribeTableRequest {
	s.SecretArn = &v
	return s
}

func (s *DescribeTableRequest) SetTable(v string) *DescribeTableRequest {
	s.Table = &v
	return s
}

func (s *DescribeTableRequest) SetWorkspaceId(v string) *DescribeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeTableRequest) Validate() error {
	return dara.Validate(s)
}
