// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportStatisticalDataResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportStatisticalDataResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportStatisticalDataResponse
  GetStatusCode() *int32 
  SetBody(v *ExportStatisticalDataResponseBody) *ExportStatisticalDataResponse
  GetBody() *ExportStatisticalDataResponseBody 
}

type ExportStatisticalDataResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportStatisticalDataResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportStatisticalDataResponse) GoString() string {
  return s.String()
}

func (s *ExportStatisticalDataResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportStatisticalDataResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportStatisticalDataResponse) GetBody() *ExportStatisticalDataResponseBody  {
  return s.Body
}

func (s *ExportStatisticalDataResponse) SetHeaders(v map[string]*string) *ExportStatisticalDataResponse {
  s.Headers = v
  return s
}

func (s *ExportStatisticalDataResponse) SetStatusCode(v int32) *ExportStatisticalDataResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportStatisticalDataResponse) SetBody(v *ExportStatisticalDataResponseBody) *ExportStatisticalDataResponse {
  s.Body = v
  return s
}

func (s *ExportStatisticalDataResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

