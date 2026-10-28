// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomerQuotaRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportCustomerQuotaRecordRequest
  GetEndDate() *string 
  SetEndUserPk(v int64) *ExportCustomerQuotaRecordRequest
  GetEndUserPk() *int64 
  SetLanguage(v string) *ExportCustomerQuotaRecordRequest
  GetLanguage() *string 
  SetOperationType(v string) *ExportCustomerQuotaRecordRequest
  GetOperationType() *string 
  SetStartDate(v string) *ExportCustomerQuotaRecordRequest
  GetStartDate() *string 
}

type ExportCustomerQuotaRecordRequest struct {
  // The end date. Format: yyyy-MM-dd.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2023-12-24
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // The UID of the customer.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 5113766248601929
  EndUserPk *int64 `json:"EndUserPk,omitempty" xml:"EndUserPk,omitempty"`
  // The language parameter. Default value: en. Valid values:
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
  // The operation type. Valid values:
  // 
  // - all: All operations.
  // 
  // - quota_create: Creation.
  // 
  // - quota_amount_adjust: Quota amount adjustment.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // all
  OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
  // The start date. Format: yyyy-MM-dd.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2023-11-10
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportCustomerQuotaRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomerQuotaRecordRequest) GoString() string {
  return s.String()
}

func (s *ExportCustomerQuotaRecordRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportCustomerQuotaRecordRequest) GetEndUserPk() *int64  {
  return s.EndUserPk
}

func (s *ExportCustomerQuotaRecordRequest) GetLanguage() *string  {
  return s.Language
}

func (s *ExportCustomerQuotaRecordRequest) GetOperationType() *string  {
  return s.OperationType
}

func (s *ExportCustomerQuotaRecordRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportCustomerQuotaRecordRequest) SetEndDate(v string) *ExportCustomerQuotaRecordRequest {
  s.EndDate = &v
  return s
}

func (s *ExportCustomerQuotaRecordRequest) SetEndUserPk(v int64) *ExportCustomerQuotaRecordRequest {
  s.EndUserPk = &v
  return s
}

func (s *ExportCustomerQuotaRecordRequest) SetLanguage(v string) *ExportCustomerQuotaRecordRequest {
  s.Language = &v
  return s
}

func (s *ExportCustomerQuotaRecordRequest) SetOperationType(v string) *ExportCustomerQuotaRecordRequest {
  s.OperationType = &v
  return s
}

func (s *ExportCustomerQuotaRecordRequest) SetStartDate(v string) *ExportCustomerQuotaRecordRequest {
  s.StartDate = &v
  return s
}

func (s *ExportCustomerQuotaRecordRequest) Validate() error {
  return dara.Validate(s)
}

