// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatementResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *GetStatementResultRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *GetStatementResultRequest
	GetDatabase() *string
	SetId(v string) *GetStatementResultRequest
	GetId() *string
	SetOwnerId(v int64) *GetStatementResultRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetStatementResultRequest
	GetRegionId() *string
	SetSecretArn(v string) *GetStatementResultRequest
	GetSecretArn() *string
}

type GetStatementResultRequest struct {
	// Instance ID. Can be obtained by calling DescribeDBInstances.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Database name.
	//
	// example:
	//
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// Task ID for asynchronous SQL execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9A920F47-416A-4044-817C-7C2A72AD16D3
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Access credential. Created through the CreateSecret interface.
	//
	// > When accessing this interface with a sub-account, the sub-account must have the UseSecret or GetSecretValue permission for this SecretArn.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
}

func (s GetStatementResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStatementResultRequest) GoString() string {
	return s.String()
}

func (s *GetStatementResultRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetStatementResultRequest) GetDatabase() *string {
	return s.Database
}

func (s *GetStatementResultRequest) GetId() *string {
	return s.Id
}

func (s *GetStatementResultRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetStatementResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStatementResultRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *GetStatementResultRequest) SetDBInstanceId(v string) *GetStatementResultRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetStatementResultRequest) SetDatabase(v string) *GetStatementResultRequest {
	s.Database = &v
	return s
}

func (s *GetStatementResultRequest) SetId(v string) *GetStatementResultRequest {
	s.Id = &v
	return s
}

func (s *GetStatementResultRequest) SetOwnerId(v int64) *GetStatementResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetStatementResultRequest) SetRegionId(v string) *GetStatementResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetStatementResultRequest) SetSecretArn(v string) *GetStatementResultRequest {
	s.SecretArn = &v
	return s
}

func (s *GetStatementResultRequest) Validate() error {
	return dara.Validate(s)
}
