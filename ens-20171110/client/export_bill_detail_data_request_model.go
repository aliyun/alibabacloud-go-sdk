// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportBillDetailDataRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportBillDetailDataRequest
  GetEndDate() *string 
  SetStartDate(v string) *ExportBillDetailDataRequest
  GetStartDate() *string 
}

type ExportBillDetailDataRequest struct {
  // The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2020-08-30T00:00:00Z
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2020-06-01T00:00:00Z
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportBillDetailDataRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportBillDetailDataRequest) GoString() string {
  return s.String()
}

func (s *ExportBillDetailDataRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportBillDetailDataRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportBillDetailDataRequest) SetEndDate(v string) *ExportBillDetailDataRequest {
  s.EndDate = &v
  return s
}

func (s *ExportBillDetailDataRequest) SetStartDate(v string) *ExportBillDetailDataRequest {
  s.StartDate = &v
  return s
}

func (s *ExportBillDetailDataRequest) Validate() error {
  return dara.Validate(s)
}

