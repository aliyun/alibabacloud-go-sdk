// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListDatabasesRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *ListDatabasesRequest
	GetDatabase() *string
	SetMaxResults(v int32) *ListDatabasesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatabasesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListDatabasesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDatabasesRequest
	GetRegionId() *string
	SetSecretArn(v string) *ListDatabasesRequest
	GetSecretArn() *string
}

type ListDatabasesRequest struct {
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
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
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
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListDatabasesRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListDatabasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatabasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatabasesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDatabasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabasesRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *ListDatabasesRequest) SetDBInstanceId(v string) *ListDatabasesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListDatabasesRequest) SetDatabase(v string) *ListDatabasesRequest {
	s.Database = &v
	return s
}

func (s *ListDatabasesRequest) SetMaxResults(v int32) *ListDatabasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatabasesRequest) SetNextToken(v string) *ListDatabasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatabasesRequest) SetOwnerId(v int64) *ListDatabasesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDatabasesRequest) SetRegionId(v string) *ListDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabasesRequest) SetSecretArn(v string) *ListDatabasesRequest {
	s.SecretArn = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
