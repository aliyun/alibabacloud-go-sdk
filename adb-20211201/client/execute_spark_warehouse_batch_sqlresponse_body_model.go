// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkWarehouseBatchSQLResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *SparkBatchSQL) *ExecuteSparkWarehouseBatchSQLResponseBody
  GetData() *SparkBatchSQL 
  SetRequestId(v string) *ExecuteSparkWarehouseBatchSQLResponseBody
  GetRequestId() *string 
}

type ExecuteSparkWarehouseBatchSQLResponseBody struct {
  // The returned data.
  // 
  // example:
  // 
  // true
  Data *SparkBatchSQL `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // CBE843D8-964D-5EA3-9D31-822125611B6E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSparkWarehouseBatchSQLResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkWarehouseBatchSQLResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSparkWarehouseBatchSQLResponseBody) GetData() *SparkBatchSQL  {
  return s.Data
}

func (s *ExecuteSparkWarehouseBatchSQLResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSparkWarehouseBatchSQLResponseBody) SetData(v *SparkBatchSQL) *ExecuteSparkWarehouseBatchSQLResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLResponseBody) SetRequestId(v string) *ExecuteSparkWarehouseBatchSQLResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSparkWarehouseBatchSQLResponseBody) Validate() error {
  return dara.Validate(s)
}

