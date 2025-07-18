// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListTablesRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *ListTablesRequest
	GetDatabase() *string
	SetMaxResults(v int32) *ListTablesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTablesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListTablesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTablesRequest
	GetRegionId() *string
	SetSchema(v string) *ListTablesRequest
	GetSchema() *string
	SetSecretArn(v string) *ListTablesRequest
	GetSecretArn() *string
	SetTablePattern(v string) *ListTablesRequest
	GetTablePattern() *string
}

type ListTablesRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
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
	// The maximum number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// >  To call the ListTables operation as a Resource Access Management (RAM) user, the RAM user must have the permissions to call the UseSecret or GetSecretValue operation on the ARN of the access credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The table name pattern for matching. For example, `ab%` specifies to match table names that start with ab.
	//
	// example:
	//
	// ab%
	TablePattern *string `json:"TablePattern,omitempty" xml:"TablePattern,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListTablesRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTablesRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListTablesRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *ListTablesRequest) GetTablePattern() *string {
	return s.TablePattern
}

func (s *ListTablesRequest) SetDBInstanceId(v string) *ListTablesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListTablesRequest) SetDatabase(v string) *ListTablesRequest {
	s.Database = &v
	return s
}

func (s *ListTablesRequest) SetMaxResults(v int32) *ListTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTablesRequest) SetNextToken(v string) *ListTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTablesRequest) SetOwnerId(v int64) *ListTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTablesRequest) SetRegionId(v string) *ListTablesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTablesRequest) SetSchema(v string) *ListTablesRequest {
	s.Schema = &v
	return s
}

func (s *ListTablesRequest) SetSecretArn(v string) *ListTablesRequest {
	s.SecretArn = &v
	return s
}

func (s *ListTablesRequest) SetTablePattern(v string) *ListTablesRequest {
	s.TablePattern = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
