// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMeasurementDataResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportMeasurementDataResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportMeasurementDataResponse
  GetStatusCode() *int32 
  SetBody(v *ExportMeasurementDataResponseBody) *ExportMeasurementDataResponse
  GetBody() *ExportMeasurementDataResponseBody 
}

type ExportMeasurementDataResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportMeasurementDataResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportMeasurementDataResponse) GoString() string {
  return s.String()
}

func (s *ExportMeasurementDataResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportMeasurementDataResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportMeasurementDataResponse) GetBody() *ExportMeasurementDataResponseBody  {
  return s.Body
}

func (s *ExportMeasurementDataResponse) SetHeaders(v map[string]*string) *ExportMeasurementDataResponse {
  s.Headers = v
  return s
}

func (s *ExportMeasurementDataResponse) SetStatusCode(v int32) *ExportMeasurementDataResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportMeasurementDataResponse) SetBody(v *ExportMeasurementDataResponseBody) *ExportMeasurementDataResponse {
  s.Body = v
  return s
}

func (s *ExportMeasurementDataResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

