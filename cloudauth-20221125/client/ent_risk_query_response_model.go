// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntRiskQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntRiskQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntRiskQueryResponse
  GetStatusCode() *int32 
  SetBody(v *EntRiskQueryResponseBody) *EntRiskQueryResponse
  GetBody() *EntRiskQueryResponseBody 
}

type EntRiskQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntRiskQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntRiskQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s EntRiskQueryResponse) GoString() string {
  return s.String()
}

func (s *EntRiskQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntRiskQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntRiskQueryResponse) GetBody() *EntRiskQueryResponseBody  {
  return s.Body
}

func (s *EntRiskQueryResponse) SetHeaders(v map[string]*string) *EntRiskQueryResponse {
  s.Headers = v
  return s
}

func (s *EntRiskQueryResponse) SetStatusCode(v int32) *EntRiskQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *EntRiskQueryResponse) SetBody(v *EntRiskQueryResponseBody) *EntRiskQueryResponse {
  s.Body = v
  return s
}

func (s *EntRiskQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

