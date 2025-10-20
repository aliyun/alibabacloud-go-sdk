// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBrandResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBrandResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBrandResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBrandResponseBody) *EnableBrandResponse
  GetBody() *EnableBrandResponseBody 
}

type EnableBrandResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBrandResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBrandResponse) GoString() string {
  return s.String()
}

func (s *EnableBrandResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBrandResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBrandResponse) GetBody() *EnableBrandResponseBody  {
  return s.Body
}

func (s *EnableBrandResponse) SetHeaders(v map[string]*string) *EnableBrandResponse {
  s.Headers = v
  return s
}

func (s *EnableBrandResponse) SetStatusCode(v int32) *EnableBrandResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBrandResponse) SetBody(v *EnableBrandResponseBody) *EnableBrandResponse {
  s.Body = v
  return s
}

func (s *EnableBrandResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

