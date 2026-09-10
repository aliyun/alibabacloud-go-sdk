// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckSqlPreviewRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCheckColumn(v string) *ExecDataCheckSqlPreviewRequest
  GetCheckColumn() *string 
  SetDataSourceId(v string) *ExecDataCheckSqlPreviewRequest
  GetDataSourceId() *string 
  SetEngineId(v string) *ExecDataCheckSqlPreviewRequest
  GetEngineId() *string 
  SetFullTableName(v string) *ExecDataCheckSqlPreviewRequest
  GetFullTableName() *string 
  SetPartitionCondition(v string) *ExecDataCheckSqlPreviewRequest
  GetPartitionCondition() *string 
  SetTaskId(v int64) *ExecDataCheckSqlPreviewRequest
  GetTaskId() *int64 
  SetWhereClause(v string) *ExecDataCheckSqlPreviewRequest
  GetWhereClause() *string 
}

type ExecDataCheckSqlPreviewRequest struct {
  // The columns to check.
  // 
  // example:
  // 
  // id,name
  CheckColumn *string `json:"checkColumn,omitempty" xml:"checkColumn,omitempty"`
  // The ID of the data source.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 230
  DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
  // The ID of the check engine. Used in Spark scenarios.
  // 
  // example:
  // 
  // 230
  EngineId *string `json:"engineId,omitempty" xml:"engineId,omitempty"`
  // The name of the table to check, in the format `schema.table`.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test_db.test_table
  FullTableName *string `json:"fullTableName,omitempty" xml:"fullTableName,omitempty"`
  // The partition condition.
  // 
  // example:
  // 
  // date_part=20240719
  PartitionCondition *string `json:"partitionCondition,omitempty" xml:"partitionCondition,omitempty"`
  // The ID of the check task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10001
  TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // The WHERE condition.
  // 
  // example:
  // 
  // id > 100
  WhereClause *string `json:"whereClause,omitempty" xml:"whereClause,omitempty"`
}

func (s ExecDataCheckSqlPreviewRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckSqlPreviewRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckSqlPreviewRequest) GetCheckColumn() *string  {
  return s.CheckColumn
}

func (s *ExecDataCheckSqlPreviewRequest) GetDataSourceId() *string  {
  return s.DataSourceId
}

func (s *ExecDataCheckSqlPreviewRequest) GetEngineId() *string  {
  return s.EngineId
}

func (s *ExecDataCheckSqlPreviewRequest) GetFullTableName() *string  {
  return s.FullTableName
}

func (s *ExecDataCheckSqlPreviewRequest) GetPartitionCondition() *string  {
  return s.PartitionCondition
}

func (s *ExecDataCheckSqlPreviewRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *ExecDataCheckSqlPreviewRequest) GetWhereClause() *string  {
  return s.WhereClause
}

func (s *ExecDataCheckSqlPreviewRequest) SetCheckColumn(v string) *ExecDataCheckSqlPreviewRequest {
  s.CheckColumn = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetDataSourceId(v string) *ExecDataCheckSqlPreviewRequest {
  s.DataSourceId = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetEngineId(v string) *ExecDataCheckSqlPreviewRequest {
  s.EngineId = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetFullTableName(v string) *ExecDataCheckSqlPreviewRequest {
  s.FullTableName = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetPartitionCondition(v string) *ExecDataCheckSqlPreviewRequest {
  s.PartitionCondition = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetTaskId(v int64) *ExecDataCheckSqlPreviewRequest {
  s.TaskId = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) SetWhereClause(v string) *ExecDataCheckSqlPreviewRequest {
  s.WhereClause = &v
  return s
}

func (s *ExecDataCheckSqlPreviewRequest) Validate() error {
  return dara.Validate(s)
}

