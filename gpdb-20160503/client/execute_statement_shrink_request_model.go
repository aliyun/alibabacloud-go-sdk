// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceId(v string) *ExecuteStatementShrinkRequest
  GetDBInstanceId() *string 
  SetDatabase(v string) *ExecuteStatementShrinkRequest
  GetDatabase() *string 
  SetOwnerId(v int64) *ExecuteStatementShrinkRequest
  GetOwnerId() *int64 
  SetParametersShrink(v string) *ExecuteStatementShrinkRequest
  GetParametersShrink() *string 
  SetRagWorkspaceCollectionShrink(v string) *ExecuteStatementShrinkRequest
  GetRagWorkspaceCollectionShrink() *string 
  SetRegionId(v string) *ExecuteStatementShrinkRequest
  GetRegionId() *string 
  SetRunType(v string) *ExecuteStatementShrinkRequest
  GetRunType() *string 
  SetSecretArn(v string) *ExecuteStatementShrinkRequest
  GetSecretArn() *string 
  SetSql(v string) *ExecuteStatementShrinkRequest
  GetSql() *string 
  SetSqlsShrink(v string) *ExecuteStatementShrinkRequest
  GetSqlsShrink() *string 
  SetStatementName(v string) *ExecuteStatementShrinkRequest
  GetStatementName() *string 
  SetWorkspaceId(v string) *ExecuteStatementShrinkRequest
  GetWorkspaceId() *string 
}

type ExecuteStatementShrinkRequest struct {
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
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The configuration parameters.
  ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
  RagWorkspaceCollectionShrink *string `json:"RagWorkspaceCollection,omitempty" xml:"RagWorkspaceCollection,omitempty"`
  // The region ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The execution type. Valid values:
  // 
  // 	- synchronous
  // 
  // 	- asynchronous (not supported)
  // 
  // example:
  // 
  // synchronous
  RunType *string `json:"RunType,omitempty" xml:"RunType,omitempty"`
  // The Alibaba Cloud Resource Name (ARN) of the access credential for the created Data API account. You can call the CreateSecret operation to create an access credential.
  // 
  // >  To call the ExecuteStatement operation as a Resource Access Management (RAM) user, the RAM user must have the permissions to call the UseSecret or GetSecretValue operation on the ARN of the access credential.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
  SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
  // The SQL statements that you want to execute.
  // 
  // example:
  // 
  // select 	- from table1
  Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
  // The SQL statements.
  SqlsShrink *string `json:"Sqls,omitempty" xml:"Sqls,omitempty"`
  // The name of the set of SQL statements that you want to execute. This parameter takes effect when the RunType parameter is set to asynchronous.
  // 
  // example:
  // 
  // test
  StatementName *string `json:"StatementName,omitempty" xml:"StatementName,omitempty"`
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExecuteStatementShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteStatementShrinkRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *ExecuteStatementShrinkRequest) GetDatabase() *string  {
  return s.Database
}

func (s *ExecuteStatementShrinkRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExecuteStatementShrinkRequest) GetParametersShrink() *string  {
  return s.ParametersShrink
}

func (s *ExecuteStatementShrinkRequest) GetRagWorkspaceCollectionShrink() *string  {
  return s.RagWorkspaceCollectionShrink
}

func (s *ExecuteStatementShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteStatementShrinkRequest) GetRunType() *string  {
  return s.RunType
}

func (s *ExecuteStatementShrinkRequest) GetSecretArn() *string  {
  return s.SecretArn
}

func (s *ExecuteStatementShrinkRequest) GetSql() *string  {
  return s.Sql
}

func (s *ExecuteStatementShrinkRequest) GetSqlsShrink() *string  {
  return s.SqlsShrink
}

func (s *ExecuteStatementShrinkRequest) GetStatementName() *string  {
  return s.StatementName
}

func (s *ExecuteStatementShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecuteStatementShrinkRequest) SetDBInstanceId(v string) *ExecuteStatementShrinkRequest {
  s.DBInstanceId = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetDatabase(v string) *ExecuteStatementShrinkRequest {
  s.Database = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetOwnerId(v int64) *ExecuteStatementShrinkRequest {
  s.OwnerId = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetParametersShrink(v string) *ExecuteStatementShrinkRequest {
  s.ParametersShrink = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetRagWorkspaceCollectionShrink(v string) *ExecuteStatementShrinkRequest {
  s.RagWorkspaceCollectionShrink = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetRegionId(v string) *ExecuteStatementShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetRunType(v string) *ExecuteStatementShrinkRequest {
  s.RunType = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetSecretArn(v string) *ExecuteStatementShrinkRequest {
  s.SecretArn = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetSql(v string) *ExecuteStatementShrinkRequest {
  s.Sql = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetSqlsShrink(v string) *ExecuteStatementShrinkRequest {
  s.SqlsShrink = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetStatementName(v string) *ExecuteStatementShrinkRequest {
  s.StatementName = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) SetWorkspaceId(v string) *ExecuteStatementShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecuteStatementShrinkRequest) Validate() error {
  return dara.Validate(s)
}

