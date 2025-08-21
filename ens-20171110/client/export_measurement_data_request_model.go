// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMeasurementDataRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportMeasurementDataRequest
  GetEndDate() *string 
  SetStartDate(v string) *ExportMeasurementDataRequest
  GetStartDate() *string 
}

type ExportMeasurementDataRequest struct {
  // The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2021-10-30T16:00:00Z
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // The beginning of the time range to query. Specify the time in the yyyy-mm-ddthh:mm:ssz format. The time must be in UTC.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2019-06-01T00:00:00Z
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportMeasurementDataRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportMeasurementDataRequest) GoString() string {
  return s.String()
}

func (s *ExportMeasurementDataRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportMeasurementDataRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportMeasurementDataRequest) SetEndDate(v string) *ExportMeasurementDataRequest {
  s.EndDate = &v
  return s
}

func (s *ExportMeasurementDataRequest) SetStartDate(v string) *ExportMeasurementDataRequest {
  s.StartDate = &v
  return s
}

func (s *ExportMeasurementDataRequest) Validate() error {
  return dara.Validate(s)
}

