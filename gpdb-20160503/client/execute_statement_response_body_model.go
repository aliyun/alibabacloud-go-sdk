// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedAt(v string) *ExecuteStatementResponseBody
  GetCreatedAt() *string 
  SetDBInstanceId(v string) *ExecuteStatementResponseBody
  GetDBInstanceId() *string 
  SetData(v *ExecuteStatementResponseBodyData) *ExecuteStatementResponseBody
  GetData() *ExecuteStatementResponseBodyData 
  SetDatabase(v string) *ExecuteStatementResponseBody
  GetDatabase() *string 
  SetId(v string) *ExecuteStatementResponseBody
  GetId() *string 
  SetMessage(v string) *ExecuteStatementResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteStatementResponseBody
  GetRequestId() *string 
  SetSecretArn(v string) *ExecuteStatementResponseBody
  GetSecretArn() *string 
  SetStatus(v string) *ExecuteStatementResponseBody
  GetStatus() *string 
}

type ExecuteStatementResponseBody struct {
  // The time when the SQL statements were created.
  // 
  // example:
  // 
  // 2023-12-04T10:08:47+0800
  CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
  // The instance ID.
  // 
  // example:
  // 
  // gp-xxxxxxxxx
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The returned results of the synchronous call.
  Data *ExecuteStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The name of the database.
  // 
  // example:
  // 
  // adbtest
  Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
  // The ID of the job for asynchronously executing the SQL statements.
  // 
  // example:
  // 
  // ABB39CC3
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // ABB39CC3-4488-4857-905D-2E4A051D0521
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The ARN of the access credential for the created Data API account.
  // 
  // example:
  // 
  // acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
  SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
  // The status of the operation. Valid values:
  // 
  // 	- **success**
  // 
  // 	- **fail**
  // 
  // example:
  // 
  // success
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ExecuteStatementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBody) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *ExecuteStatementResponseBody) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *ExecuteStatementResponseBody) GetData() *ExecuteStatementResponseBodyData  {
  return s.Data
}

func (s *ExecuteStatementResponseBody) GetDatabase() *string  {
  return s.Database
}

func (s *ExecuteStatementResponseBody) GetId() *string  {
  return s.Id
}

func (s *ExecuteStatementResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteStatementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteStatementResponseBody) GetSecretArn() *string  {
  return s.SecretArn
}

func (s *ExecuteStatementResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *ExecuteStatementResponseBody) SetCreatedAt(v string) *ExecuteStatementResponseBody {
  s.CreatedAt = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetDBInstanceId(v string) *ExecuteStatementResponseBody {
  s.DBInstanceId = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetData(v *ExecuteStatementResponseBodyData) *ExecuteStatementResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteStatementResponseBody) SetDatabase(v string) *ExecuteStatementResponseBody {
  s.Database = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetId(v string) *ExecuteStatementResponseBody {
  s.Id = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetMessage(v string) *ExecuteStatementResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetRequestId(v string) *ExecuteStatementResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetSecretArn(v string) *ExecuteStatementResponseBody {
  s.SecretArn = &v
  return s
}

func (s *ExecuteStatementResponseBody) SetStatus(v string) *ExecuteStatementResponseBody {
  s.Status = &v
  return s
}

func (s *ExecuteStatementResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteStatementResponseBodyData struct {
  // The metadata of the columns.
  ColumnMetadata *ExecuteStatementResponseBodyDataColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Struct"`
  // The rows of data.
  Records *ExecuteStatementResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
  // The total number of entries returned.
  // 
  // example:
  // 
  // 10
  TotalNumRows *int64 `json:"TotalNumRows,omitempty" xml:"TotalNumRows,omitempty"`
}

func (s ExecuteStatementResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyData) GetColumnMetadata() *ExecuteStatementResponseBodyDataColumnMetadata  {
  return s.ColumnMetadata
}

func (s *ExecuteStatementResponseBodyData) GetRecords() *ExecuteStatementResponseBodyDataRecords  {
  return s.Records
}

func (s *ExecuteStatementResponseBodyData) GetTotalNumRows() *int64  {
  return s.TotalNumRows
}

func (s *ExecuteStatementResponseBodyData) SetColumnMetadata(v *ExecuteStatementResponseBodyDataColumnMetadata) *ExecuteStatementResponseBodyData {
  s.ColumnMetadata = v
  return s
}

func (s *ExecuteStatementResponseBodyData) SetRecords(v *ExecuteStatementResponseBodyDataRecords) *ExecuteStatementResponseBodyData {
  s.Records = v
  return s
}

func (s *ExecuteStatementResponseBodyData) SetTotalNumRows(v int64) *ExecuteStatementResponseBodyData {
  s.TotalNumRows = &v
  return s
}

func (s *ExecuteStatementResponseBodyData) Validate() error {
  if s.ColumnMetadata != nil {
    if err := s.ColumnMetadata.Validate(); err != nil {
      return err
    }
  }
  if s.Records != nil {
    if err := s.Records.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteStatementResponseBodyDataColumnMetadata struct {
  ColumnMetadata []*ColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyDataColumnMetadata) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataColumnMetadata) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) GetColumnMetadata() []*ColumnMetadata  {
  return s.ColumnMetadata
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) SetColumnMetadata(v []*ColumnMetadata) *ExecuteStatementResponseBodyDataColumnMetadata {
  s.ColumnMetadata = v
  return s
}

func (s *ExecuteStatementResponseBodyDataColumnMetadata) Validate() error {
  if s.ColumnMetadata != nil {
    for _, item := range s.ColumnMetadata {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteStatementResponseBodyDataRecords struct {
  Records []*ExecuteStatementResponseBodyDataRecordsRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyDataRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataRecords) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyDataRecords) GetRecords() []*ExecuteStatementResponseBodyDataRecordsRecords  {
  return s.Records
}

func (s *ExecuteStatementResponseBodyDataRecords) SetRecords(v []*ExecuteStatementResponseBodyDataRecordsRecords) *ExecuteStatementResponseBodyDataRecords {
  s.Records = v
  return s
}

func (s *ExecuteStatementResponseBodyDataRecords) Validate() error {
  if s.Records != nil {
    for _, item := range s.Records {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteStatementResponseBodyDataRecordsRecords struct {
  Record []*Field `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResponseBodyDataRecordsRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponseBodyDataRecordsRecords) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponseBodyDataRecordsRecords) GetRecord() []*Field  {
  return s.Record
}

func (s *ExecuteStatementResponseBodyDataRecordsRecords) SetRecord(v []*Field) *ExecuteStatementResponseBodyDataRecordsRecords {
  s.Record = v
  return s
}

func (s *ExecuteStatementResponseBodyDataRecordsRecords) Validate() error {
  if s.Record != nil {
    for _, item := range s.Record {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

