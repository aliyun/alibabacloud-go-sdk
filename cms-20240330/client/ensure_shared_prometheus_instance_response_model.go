// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnsureSharedPrometheusInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnsureSharedPrometheusInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnsureSharedPrometheusInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *EnsureSharedPrometheusInstanceResponseBody) *EnsureSharedPrometheusInstanceResponse
  GetBody() *EnsureSharedPrometheusInstanceResponseBody 
}

type EnsureSharedPrometheusInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnsureSharedPrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnsureSharedPrometheusInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnsureSharedPrometheusInstanceResponse) GoString() string {
  return s.String()
}

func (s *EnsureSharedPrometheusInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnsureSharedPrometheusInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnsureSharedPrometheusInstanceResponse) GetBody() *EnsureSharedPrometheusInstanceResponseBody  {
  return s.Body
}

func (s *EnsureSharedPrometheusInstanceResponse) SetHeaders(v map[string]*string) *EnsureSharedPrometheusInstanceResponse {
  s.Headers = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponse) SetStatusCode(v int32) *EnsureSharedPrometheusInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponse) SetBody(v *EnsureSharedPrometheusInstanceResponseBody) *EnsureSharedPrometheusInstanceResponse {
  s.Body = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

