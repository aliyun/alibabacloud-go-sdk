// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlAuditRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDbName(v string) *EnableSqlAuditRequest
  GetDbName() *string 
  SetDrdsInstanceId(v string) *EnableSqlAuditRequest
  GetDrdsInstanceId() *string 
  SetIsRecall(v bool) *EnableSqlAuditRequest
  GetIsRecall() *bool 
  SetRecallEndTimestamp(v string) *EnableSqlAuditRequest
  GetRecallEndTimestamp() *string 
  SetRecallStartTimestamp(v string) *EnableSqlAuditRequest
  GetRecallStartTimestamp() *string 
}

type EnableSqlAuditRequest struct {
  // The name of the database for which you want to enable the SQL audit feature.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test
  DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
  // The ID of the PolarDB-X 1.0 instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // drds***********
  DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
  // Specifies whether to backtrack historical SQL statements for auditing.
  // 
  // example:
  // 
  // true
  IsRecall *bool `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
  // The timestamp that indicates when the backtracking ends. Unit: milliseconds.
  // 
  // > The end time of the backtracking must be later than the start time of the backtracking.
  // 
  // example:
  // 
  // 1568875132000
  RecallEndTimestamp *string `json:"RecallEndTimestamp,omitempty" xml:"RecallEndTimestamp,omitempty"`
  // The timestamp that indicates when the backtracking starts. Unit: milliseconds.
  // 
  // example:
  // 
  // 1568875132000
  RecallStartTimestamp *string `json:"RecallStartTimestamp,omitempty" xml:"RecallStartTimestamp,omitempty"`
}

func (s EnableSqlAuditRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlAuditRequest) GoString() string {
  return s.String()
}

func (s *EnableSqlAuditRequest) GetDbName() *string  {
  return s.DbName
}

func (s *EnableSqlAuditRequest) GetDrdsInstanceId() *string  {
  return s.DrdsInstanceId
}

func (s *EnableSqlAuditRequest) GetIsRecall() *bool  {
  return s.IsRecall
}

func (s *EnableSqlAuditRequest) GetRecallEndTimestamp() *string  {
  return s.RecallEndTimestamp
}

func (s *EnableSqlAuditRequest) GetRecallStartTimestamp() *string  {
  return s.RecallStartTimestamp
}

func (s *EnableSqlAuditRequest) SetDbName(v string) *EnableSqlAuditRequest {
  s.DbName = &v
  return s
}

func (s *EnableSqlAuditRequest) SetDrdsInstanceId(v string) *EnableSqlAuditRequest {
  s.DrdsInstanceId = &v
  return s
}

func (s *EnableSqlAuditRequest) SetIsRecall(v bool) *EnableSqlAuditRequest {
  s.IsRecall = &v
  return s
}

func (s *EnableSqlAuditRequest) SetRecallEndTimestamp(v string) *EnableSqlAuditRequest {
  s.RecallEndTimestamp = &v
  return s
}

func (s *EnableSqlAuditRequest) SetRecallStartTimestamp(v string) *EnableSqlAuditRequest {
  s.RecallStartTimestamp = &v
  return s
}

func (s *EnableSqlAuditRequest) Validate() error {
  return dara.Validate(s)
}

