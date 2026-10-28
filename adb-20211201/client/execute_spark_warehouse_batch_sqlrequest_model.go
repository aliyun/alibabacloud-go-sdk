// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkWarehouseBatchSQLRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgency(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetAgency() *string 
  SetDBClusterId(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetDBClusterId() *string 
  SetExecuteResultLimit(v int64) *ExecuteSparkWarehouseBatchSQLRequest
  GetExecuteResultLimit() *int64 
  SetExecuteTimeLimitInSeconds(v int64) *ExecuteSparkWarehouseBatchSQLRequest
  GetExecuteTimeLimitInSeconds() *int64 
  SetQuery(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetQuery() *string 
  SetResourceGroupName(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetResourceGroupName() *string 
  SetRuntimeConfig(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetRuntimeConfig() *string 
  SetSchema(v string) *ExecuteSparkWarehouseBatchSQLRequest
  GetSchema() *string 
}

type ExecuteSparkWarehouseBatchSQLRequest struct {
  // The name of the client.
  // 
  // example:
  // 
  // DataWorks
  Agency *string `json:"Agency,omitempty" xml:"Agency,omitempty"`
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // amv-bp11q28kvl688****
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // The maximum amount of execution result data that can be written to Object Storage Service (OSS). Unit: MB. Default value: 4096. The size of compressed objects is difficult to estimate. The data that is actually written to OSS is smaller than the specified value.
  // 
  // example:
  // 
  // 4096
  ExecuteResultLimit *int64 `json:"ExecuteResultLimit,omitempty" xml:"ExecuteResultLimit,omitempty"`
  // The maximum execution duration. Unit: seconds. If a set of SQL statements fail to be executed for the specified period of time after submission, they are marked as a timeout error. The default value is 360000 seconds, which is equivalent to 100 hours.
  // 
  // example:
  // 
  // 3600000
  ExecuteTimeLimitInSeconds *int64 `json:"ExecuteTimeLimitInSeconds,omitempty" xml:"ExecuteTimeLimitInSeconds,omitempty"`
  // The SQL statements that you want to execute in batches. Separate multiple SQL statements with semicolons (;). The execution engine executes the SQL statements in sequence in the same session.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // CREATE TABLE user(id INT, name STRING);
  // 
  // INSERT INTO t VALUE(1, \\"Bob\\");
  // 
  // SELECT 	- FROM t;
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  // The name of the resource group.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test
  ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
  // The additional runtime parameter. Specify the parameter in the JSON format.
  // 
  // example:
  // 
  // {
  // 
  //  "OSSURL": "oss://testBucketname/"
  // 
  // }
  RuntimeConfig *string `json:"RuntimeConfig,omitempty" xml:"RuntimeConfig,omitempty"`
  // The name of the database.
  // 
  // example:
  // 
  // adb_demo
  Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ExecuteSparkWarehouseBatchSQLRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkWarehouseBatchSQLRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetAgency() *string  {
  return s.Agency
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetExecuteResultLimit() *int64  {
  return s.ExecuteResultLimit
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetExecuteTimeLimitInSeconds() *int64  {
  return s.ExecuteTimeLimitInSeconds
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetResourceGroupName() *string  {
  return s.ResourceGroupName
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetRuntimeConfig() *string  {
  return s.RuntimeConfig
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) GetSchema() *string  {
  return s.Schema
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetAgency(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.Agency = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetDBClusterId(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.DBClusterId = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetExecuteResultLimit(v int64) *ExecuteSparkWarehouseBatchSQLRequest {
  s.ExecuteResultLimit = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetExecuteTimeLimitInSeconds(v int64) *ExecuteSparkWarehouseBatchSQLRequest {
  s.ExecuteTimeLimitInSeconds = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetQuery(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.Query = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetResourceGroupName(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.ResourceGroupName = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetRuntimeConfig(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.RuntimeConfig = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) SetSchema(v string) *ExecuteSparkWarehouseBatchSQLRequest {
  s.Schema = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLRequest) Validate() error {
  return dara.Validate(s)
}

