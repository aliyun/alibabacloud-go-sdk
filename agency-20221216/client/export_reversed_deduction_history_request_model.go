// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportReversedDeductionHistoryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportReversedDeductionHistoryRequest
  GetEndDate() *string 
  SetExportUid(v int64) *ExportReversedDeductionHistoryRequest
  GetExportUid() *int64 
  SetLanguage(v string) *ExportReversedDeductionHistoryRequest
  GetLanguage() *string 
  SetStartDate(v string) *ExportReversedDeductionHistoryRequest
  GetStartDate() *string 
}

type ExportReversedDeductionHistoryRequest struct {
  // The end date. Format: yyyy-MM-dd.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2023-05-01
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // The UID for which the reversed deduction details are exported.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 123
  ExportUid *int64 `json:"ExportUid,omitempty" xml:"ExportUid,omitempty"`
  // The language of the exported content. Valid values:
  // 
  // - en: English.
  // 
  // - zh: Chinese.
  // 
  // - ja: Japanese.
  // 
  // example:
  // 
  // en
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  // The start date. Format: yyyy-MM-dd.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2023-01-01
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportReversedDeductionHistoryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportReversedDeductionHistoryRequest) GoString() string {
  return s.String()
}

func (s *ExportReversedDeductionHistoryRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportReversedDeductionHistoryRequest) GetExportUid() *int64  {
  return s.ExportUid
}

func (s *ExportReversedDeductionHistoryRequest) GetLanguage() *string  {
  return s.Language
}

func (s *ExportReversedDeductionHistoryRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportReversedDeductionHistoryRequest) SetEndDate(v string) *ExportReversedDeductionHistoryRequest {
  s.EndDate = &v
  return s
}

func (s *ExportReversedDeductionHistoryRequest) SetExportUid(v int64) *ExportReversedDeductionHistoryRequest {
  s.ExportUid = &v
  return s
}

func (s *ExportReversedDeductionHistoryRequest) SetLanguage(v string) *ExportReversedDeductionHistoryRequest {
  s.Language = &v
  return s
}

func (s *ExportReversedDeductionHistoryRequest) SetStartDate(v string) *ExportReversedDeductionHistoryRequest {
  s.StartDate = &v
  return s
}

func (s *ExportReversedDeductionHistoryRequest) Validate() error {
  return dara.Validate(s)
}

