// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckResourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCheckResourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCheckResourceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCheckResourceResponseBody) *EnableCheckResourceResponse
  GetBody() *EnableCheckResourceResponseBody 
}

type EnableCheckResourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCheckResourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckResourceResponse) GoString() string {
  return s.String()
}

func (s *EnableCheckResourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCheckResourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCheckResourceResponse) GetBody() *EnableCheckResourceResponseBody  {
  return s.Body
}

func (s *EnableCheckResourceResponse) SetHeaders(v map[string]*string) *EnableCheckResourceResponse {
  s.Headers = v
  return s
}

func (s *EnableCheckResourceResponse) SetStatusCode(v int32) *EnableCheckResourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCheckResourceResponse) SetBody(v *EnableCheckResourceResponseBody) *EnableCheckResourceResponse {
  s.Body = v
  return s
}

func (s *EnableCheckResourceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

