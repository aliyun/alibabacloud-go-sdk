// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceId(v string) *ExecuteStatementRequest
  GetDBInstanceId() *string 
  SetDatabase(v string) *ExecuteStatementRequest
  GetDatabase() *string 
  SetOwnerId(v int64) *ExecuteStatementRequest
  GetOwnerId() *int64 
  SetParameters(v []interface{}) *ExecuteStatementRequest
  GetParameters() []interface{} 
  SetRagWorkspaceCollection(v *ExecuteStatementRequestRagWorkspaceCollection) *ExecuteStatementRequest
  GetRagWorkspaceCollection() *ExecuteStatementRequestRagWorkspaceCollection 
  SetRegionId(v string) *ExecuteStatementRequest
  GetRegionId() *string 
  SetRunType(v string) *ExecuteStatementRequest
  GetRunType() *string 
  SetSecretArn(v string) *ExecuteStatementRequest
  GetSecretArn() *string 
  SetSql(v string) *ExecuteStatementRequest
  GetSql() *string 
  SetSqls(v []*string) *ExecuteStatementRequest
  GetSqls() []*string 
  SetStatementName(v string) *ExecuteStatementRequest
  GetStatementName() *string 
  SetWorkspaceId(v string) *ExecuteStatementRequest
  GetWorkspaceId() *string 
}

type ExecuteStatementRequest struct {
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
  Parameters []interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
  RagWorkspaceCollection *ExecuteStatementRequestRagWorkspaceCollection `json:"RagWorkspaceCollection,omitempty" xml:"RagWorkspaceCollection,omitempty" type:"Struct"`
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
  Sqls []*string `json:"Sqls,omitempty" xml:"Sqls,omitempty" type:"Repeated"`
  // The name of the set of SQL statements that you want to execute. This parameter takes effect when the RunType parameter is set to asynchronous.
  // 
  // example:
  // 
  // test
  StatementName *string `json:"StatementName,omitempty" xml:"StatementName,omitempty"`
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExecuteStatementRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementRequest) GoString() string {
  return s.String()
}

func (s *ExecuteStatementRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *ExecuteStatementRequest) GetDatabase() *string  {
  return s.Database
}

func (s *ExecuteStatementRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExecuteStatementRequest) GetParameters() []interface{}  {
  return s.Parameters
}

func (s *ExecuteStatementRequest) GetRagWorkspaceCollection() *ExecuteStatementRequestRagWorkspaceCollection  {
  return s.RagWorkspaceCollection
}

func (s *ExecuteStatementRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteStatementRequest) GetRunType() *string  {
  return s.RunType
}

func (s *ExecuteStatementRequest) GetSecretArn() *string  {
  return s.SecretArn
}

func (s *ExecuteStatementRequest) GetSql() *string  {
  return s.Sql
}

func (s *ExecuteStatementRequest) GetSqls() []*string  {
  return s.Sqls
}

func (s *ExecuteStatementRequest) GetStatementName() *string  {
  return s.StatementName
}

func (s *ExecuteStatementRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExecuteStatementRequest) SetDBInstanceId(v string) *ExecuteStatementRequest {
  s.DBInstanceId = &v
  return s
}

func (s *ExecuteStatementRequest) SetDatabase(v string) *ExecuteStatementRequest {
  s.Database = &v
  return s
}

func (s *ExecuteStatementRequest) SetOwnerId(v int64) *ExecuteStatementRequest {
  s.OwnerId = &v
  return s
}

func (s *ExecuteStatementRequest) SetParameters(v []interface{}) *ExecuteStatementRequest {
  s.Parameters = v
  return s
}

func (s *ExecuteStatementRequest) SetRagWorkspaceCollection(v *ExecuteStatementRequestRagWorkspaceCollection) *ExecuteStatementRequest {
  s.RagWorkspaceCollection = v
  return s
}

func (s *ExecuteStatementRequest) SetRegionId(v string) *ExecuteStatementRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteStatementRequest) SetRunType(v string) *ExecuteStatementRequest {
  s.RunType = &v
  return s
}

func (s *ExecuteStatementRequest) SetSecretArn(v string) *ExecuteStatementRequest {
  s.SecretArn = &v
  return s
}

func (s *ExecuteStatementRequest) SetSql(v string) *ExecuteStatementRequest {
  s.Sql = &v
  return s
}

func (s *ExecuteStatementRequest) SetSqls(v []*string) *ExecuteStatementRequest {
  s.Sqls = v
  return s
}

func (s *ExecuteStatementRequest) SetStatementName(v string) *ExecuteStatementRequest {
  s.StatementName = &v
  return s
}

func (s *ExecuteStatementRequest) SetWorkspaceId(v string) *ExecuteStatementRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExecuteStatementRequest) Validate() error {
  return dara.Validate(s)
}

type ExecuteStatementRequestRagWorkspaceCollection struct {
  Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ExecuteStatementRequestRagWorkspaceCollection) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementRequestRagWorkspaceCollection) GoString() string {
  return s.String()
}

func (s *ExecuteStatementRequestRagWorkspaceCollection) GetCollection() *string  {
  return s.Collection
}

func (s *ExecuteStatementRequestRagWorkspaceCollection) GetNamespace() *string  {
  return s.Namespace
}

func (s *ExecuteStatementRequestRagWorkspaceCollection) SetCollection(v string) *ExecuteStatementRequestRagWorkspaceCollection {
  s.Collection = &v
  return s
}

func (s *ExecuteStatementRequestRagWorkspaceCollection) SetNamespace(v string) *ExecuteStatementRequestRagWorkspaceCollection {
  s.Namespace = &v
  return s
}

func (s *ExecuteStatementRequestRagWorkspaceCollection) Validate() error {
  return dara.Validate(s)
}

