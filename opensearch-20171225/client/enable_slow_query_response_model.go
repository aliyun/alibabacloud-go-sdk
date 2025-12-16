// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSlowQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSlowQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSlowQueryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSlowQueryResponseBody) *EnableSlowQueryResponse
  GetBody() *EnableSlowQueryResponseBody 
}

type EnableSlowQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSlowQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSlowQueryResponse) GoString() string {
  return s.String()
}

func (s *EnableSlowQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSlowQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSlowQueryResponse) GetBody() *EnableSlowQueryResponseBody  {
  return s.Body
}

func (s *EnableSlowQueryResponse) SetHeaders(v map[string]*string) *EnableSlowQueryResponse {
  s.Headers = v
  return s
}

func (s *EnableSlowQueryResponse) SetStatusCode(v int32) *EnableSlowQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSlowQueryResponse) SetBody(v *EnableSlowQueryResponseBody) *EnableSlowQueryResponse {
  s.Body = v
  return s
}

func (s *EnableSlowQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

