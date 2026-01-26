// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableMetricResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableMetricResponse
  GetStatusCode() *int32 
  SetBody(v *EnableMetricResponseBody) *EnableMetricResponse
  GetBody() *EnableMetricResponseBody 
}

type EnableMetricResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMetricResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricResponse) GoString() string {
  return s.String()
}

func (s *EnableMetricResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableMetricResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableMetricResponse) GetBody() *EnableMetricResponseBody  {
  return s.Body
}

func (s *EnableMetricResponse) SetHeaders(v map[string]*string) *EnableMetricResponse {
  s.Headers = v
  return s
}

func (s *EnableMetricResponse) SetStatusCode(v int32) *EnableMetricResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableMetricResponse) SetBody(v *EnableMetricResponseBody) *EnableMetricResponse {
  s.Body = v
  return s
}

func (s *EnableMetricResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

