// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckProductResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCheckProductResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCheckProductResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCheckProductResponseBody) *EnableCheckProductResponse
  GetBody() *EnableCheckProductResponseBody 
}

type EnableCheckProductResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCheckProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCheckProductResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckProductResponse) GoString() string {
  return s.String()
}

func (s *EnableCheckProductResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCheckProductResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCheckProductResponse) GetBody() *EnableCheckProductResponseBody  {
  return s.Body
}

func (s *EnableCheckProductResponse) SetHeaders(v map[string]*string) *EnableCheckProductResponse {
  s.Headers = v
  return s
}

func (s *EnableCheckProductResponse) SetStatusCode(v int32) *EnableCheckProductResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCheckProductResponse) SetBody(v *EnableCheckProductResponseBody) *EnableCheckProductResponse {
  s.Body = v
  return s
}

func (s *EnableCheckProductResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

