// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListSchemasRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *ListSchemasRequest
	GetDatabase() *string
	SetMaxResults(v int32) *ListSchemasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSchemasRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListSchemasRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListSchemasRequest
	GetRegionId() *string
	SetSchemaPattern(v string) *ListSchemasRequest
	GetSchemaPattern() *string
	SetSecretArn(v string) *ListSchemasRequest
	GetSecretArn() *string
}

type ListSchemasRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// 20
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
	// The schema name pattern for matching. For example, `ab%` specifies to match schema names that start with ab.
	//
	// example:
	//
	// aaa%
	SchemaPattern *string `json:"SchemaPattern,omitempty" xml:"SchemaPattern,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the access credential for the created Data API account. You can call the CreateSecret operation to create an access credential.
	//
	// >  To call the ListSchemas operation as a Resource Access Management (RAM) user, the RAM user must have the permissions to call the UseSecret or GetSecretValue operation on the ARN of the access credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
}

func (s ListSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListSchemasRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListSchemasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSchemasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchemasRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSchemasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSchemasRequest) GetSchemaPattern() *string {
	return s.SchemaPattern
}

func (s *ListSchemasRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *ListSchemasRequest) SetDBInstanceId(v string) *ListSchemasRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListSchemasRequest) SetDatabase(v string) *ListSchemasRequest {
	s.Database = &v
	return s
}

func (s *ListSchemasRequest) SetMaxResults(v int32) *ListSchemasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSchemasRequest) SetNextToken(v string) *ListSchemasRequest {
	s.NextToken = &v
	return s
}

func (s *ListSchemasRequest) SetOwnerId(v int64) *ListSchemasRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSchemasRequest) SetRegionId(v string) *ListSchemasRequest {
	s.RegionId = &v
	return s
}

func (s *ListSchemasRequest) SetSchemaPattern(v string) *ListSchemasRequest {
	s.SchemaPattern = &v
	return s
}

func (s *ListSchemasRequest) SetSecretArn(v string) *ListSchemasRequest {
	s.SecretArn = &v
	return s
}

func (s *ListSchemasRequest) Validate() error {
	return dara.Validate(s)
}
