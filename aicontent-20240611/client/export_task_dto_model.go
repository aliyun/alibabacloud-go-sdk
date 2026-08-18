// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportTaskDTO interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedAt(v int64) *ExportTaskDTO
  GetCreatedAt() *int64 
  SetError(v string) *ExportTaskDTO
  GetError() *string 
  SetFilename(v string) *ExportTaskDTO
  GetFilename() *string 
  SetFinishedAt(v int64) *ExportTaskDTO
  GetFinishedAt() *int64 
  SetId(v string) *ExportTaskDTO
  GetId() *string 
  SetMaxRows(v int32) *ExportTaskDTO
  GetMaxRows() *int32 
  SetProgress(v int32) *ExportTaskDTO
  GetProgress() *int32 
  SetStatus(v string) *ExportTaskDTO
  GetStatus() *string 
  SetTotal(v int32) *ExportTaskDTO
  GetTotal() *int32 
  SetType(v string) *ExportTaskDTO
  GetType() *string 
}

type ExportTaskDTO struct {
  // example:
  // 
  // 1753858800
  CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
  // example:
  // 
  // MemberBalanceChangeRecords_20260730.csv
  Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
  // example:
  // 
  // 1753858860
  FinishedAt *int64 `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
  // example:
  // 
  // task_xxxxxxxx
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 50000
  MaxRows *int32 `json:"maxRows,omitempty" xml:"maxRows,omitempty"`
  // example:
  // 
  // 0
  Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
  // example:
  // 
  // pending
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 100
  Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
  // example:
  // 
  // balance_orders
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExportTaskDTO) String() string {
  return dara.Prettify(s)
}

func (s ExportTaskDTO) GoString() string {
  return s.String()
}

func (s *ExportTaskDTO) GetCreatedAt() *int64  {
  return s.CreatedAt
}

func (s *ExportTaskDTO) GetError() *string  {
  return s.Error
}

func (s *ExportTaskDTO) GetFilename() *string  {
  return s.Filename
}

func (s *ExportTaskDTO) GetFinishedAt() *int64  {
  return s.FinishedAt
}

func (s *ExportTaskDTO) GetId() *string  {
  return s.Id
}

func (s *ExportTaskDTO) GetMaxRows() *int32  {
  return s.MaxRows
}

func (s *ExportTaskDTO) GetProgress() *int32  {
  return s.Progress
}

func (s *ExportTaskDTO) GetStatus() *string  {
  return s.Status
}

func (s *ExportTaskDTO) GetTotal() *int32  {
  return s.Total
}

func (s *ExportTaskDTO) GetType() *string  {
  return s.Type
}

func (s *ExportTaskDTO) SetCreatedAt(v int64) *ExportTaskDTO {
  s.CreatedAt = &v
  return s
}

func (s *ExportTaskDTO) SetError(v string) *ExportTaskDTO {
  s.Error = &v
  return s
}

func (s *ExportTaskDTO) SetFilename(v string) *ExportTaskDTO {
  s.Filename = &v
  return s
}

func (s *ExportTaskDTO) SetFinishedAt(v int64) *ExportTaskDTO {
  s.FinishedAt = &v
  return s
}

func (s *ExportTaskDTO) SetId(v string) *ExportTaskDTO {
  s.Id = &v
  return s
}

func (s *ExportTaskDTO) SetMaxRows(v int32) *ExportTaskDTO {
  s.MaxRows = &v
  return s
}

func (s *ExportTaskDTO) SetProgress(v int32) *ExportTaskDTO {
  s.Progress = &v
  return s
}

func (s *ExportTaskDTO) SetStatus(v string) *ExportTaskDTO {
  s.Status = &v
  return s
}

func (s *ExportTaskDTO) SetTotal(v int32) *ExportTaskDTO {
  s.Total = &v
  return s
}

func (s *ExportTaskDTO) SetType(v string) *ExportTaskDTO {
  s.Type = &v
  return s
}

func (s *ExportTaskDTO) Validate() error {
  return dara.Validate(s)
}

