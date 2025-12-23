// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceCenterResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableResourceCenterResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableResourceCenterResponse
  GetStatusCode() *int32 
  SetBody(v *EnableResourceCenterResponseBody) *EnableResourceCenterResponse
  GetBody() *EnableResourceCenterResponseBody 
}

type EnableResourceCenterResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceCenterResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceCenterResponse) GoString() string {
  return s.String()
}

func (s *EnableResourceCenterResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableResourceCenterResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableResourceCenterResponse) GetBody() *EnableResourceCenterResponseBody  {
  return s.Body
}

func (s *EnableResourceCenterResponse) SetHeaders(v map[string]*string) *EnableResourceCenterResponse {
  s.Headers = v
  return s
}

func (s *EnableResourceCenterResponse) SetStatusCode(v int32) *EnableResourceCenterResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableResourceCenterResponse) SetBody(v *EnableResourceCenterResponseBody) *EnableResourceCenterResponse {
  s.Body = v
  return s
}

func (s *EnableResourceCenterResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

