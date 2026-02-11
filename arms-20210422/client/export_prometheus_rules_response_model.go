// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPrometheusRulesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportPrometheusRulesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportPrometheusRulesResponse
  GetStatusCode() *int32 
  SetBody(v *ExportPrometheusRulesResponseBody) *ExportPrometheusRulesResponse
  GetBody() *ExportPrometheusRulesResponseBody 
}

type ExportPrometheusRulesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportPrometheusRulesResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportPrometheusRulesResponse) GoString() string {
  return s.String()
}

func (s *ExportPrometheusRulesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportPrometheusRulesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportPrometheusRulesResponse) GetBody() *ExportPrometheusRulesResponseBody  {
  return s.Body
}

func (s *ExportPrometheusRulesResponse) SetHeaders(v map[string]*string) *ExportPrometheusRulesResponse {
  s.Headers = v
  return s
}

func (s *ExportPrometheusRulesResponse) SetStatusCode(v int32) *ExportPrometheusRulesResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportPrometheusRulesResponse) SetBody(v *ExportPrometheusRulesResponseBody) *ExportPrometheusRulesResponse {
  s.Body = v
  return s
}

func (s *ExportPrometheusRulesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

