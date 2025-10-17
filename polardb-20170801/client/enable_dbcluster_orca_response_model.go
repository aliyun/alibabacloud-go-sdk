// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterOrcaResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBClusterOrcaResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBClusterOrcaResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBClusterOrcaResponseBody) *EnableDBClusterOrcaResponse
  GetBody() *EnableDBClusterOrcaResponseBody 
}

type EnableDBClusterOrcaResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBClusterOrcaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBClusterOrcaResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterOrcaResponse) GoString() string {
  return s.String()
}

func (s *EnableDBClusterOrcaResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBClusterOrcaResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBClusterOrcaResponse) GetBody() *EnableDBClusterOrcaResponseBody  {
  return s.Body
}

func (s *EnableDBClusterOrcaResponse) SetHeaders(v map[string]*string) *EnableDBClusterOrcaResponse {
  s.Headers = v
  return s
}

func (s *EnableDBClusterOrcaResponse) SetStatusCode(v int32) *EnableDBClusterOrcaResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBClusterOrcaResponse) SetBody(v *EnableDBClusterOrcaResponseBody) *EnableDBClusterOrcaResponse {
  s.Body = v
  return s
}

func (s *EnableDBClusterOrcaResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

