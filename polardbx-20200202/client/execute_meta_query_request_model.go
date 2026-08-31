// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMetaQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceName(v string) *ExecuteMetaQueryRequest
  GetDBInstanceName() *string 
  SetMaxResultRows(v int64) *ExecuteMetaQueryRequest
  GetMaxResultRows() *int64 
  SetRegionId(v string) *ExecuteMetaQueryRequest
  GetRegionId() *string 
  SetSql(v string) *ExecuteMetaQueryRequest
  GetSql() *string 
  SetStorageInstId(v string) *ExecuteMetaQueryRequest
  GetStorageInstId() *string 
}

type ExecuteMetaQueryRequest struct {
  // The primary instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pxsp-*********
  DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
  // The maximum number of rows to return. Default value: 100. Valid values: 1 to 1000. The actual number of returned rows is the minimum value among the code hard limit, the Biz DB limit, and the outermost LIMIT clause in the SQL statement. To retrieve data continuously, implement pagination in the SQL statement.
  MaxResultRows *int64 `json:"MaxResultRows,omitempty" xml:"MaxResultRows,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The SQL statement to execute.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // show databases;
  Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
  // The data node (DN) instance ID.
  // 
  // example:
  // 
  // pxc-xdb-s-*
  StorageInstId *string `json:"StorageInstId,omitempty" xml:"StorageInstId,omitempty"`
}

func (s ExecuteMetaQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMetaQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteMetaQueryRequest) GetDBInstanceName() *string  {
  return s.DBInstanceName
}

func (s *ExecuteMetaQueryRequest) GetMaxResultRows() *int64  {
  return s.MaxResultRows
}

func (s *ExecuteMetaQueryRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteMetaQueryRequest) GetSql() *string  {
  return s.Sql
}

func (s *ExecuteMetaQueryRequest) GetStorageInstId() *string  {
  return s.StorageInstId
}

func (s *ExecuteMetaQueryRequest) SetDBInstanceName(v string) *ExecuteMetaQueryRequest {
  s.DBInstanceName = &v
  return s
}

func (s *ExecuteMetaQueryRequest) SetMaxResultRows(v int64) *ExecuteMetaQueryRequest {
  s.MaxResultRows = &v
  return s
}

func (s *ExecuteMetaQueryRequest) SetRegionId(v string) *ExecuteMetaQueryRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteMetaQueryRequest) SetSql(v string) *ExecuteMetaQueryRequest {
  s.Sql = &v
  return s
}

func (s *ExecuteMetaQueryRequest) SetStorageInstId(v string) *ExecuteMetaQueryRequest {
  s.StorageInstId = &v
  return s
}

func (s *ExecuteMetaQueryRequest) Validate() error {
  return dara.Validate(s)
}

