// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceDirectoryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableResourceDirectoryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableResourceDirectoryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableResourceDirectoryResponseBody) *EnableResourceDirectoryResponse
  GetBody() *EnableResourceDirectoryResponseBody 
}

type EnableResourceDirectoryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceDirectoryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceDirectoryResponse) GoString() string {
  return s.String()
}

func (s *EnableResourceDirectoryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableResourceDirectoryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableResourceDirectoryResponse) GetBody() *EnableResourceDirectoryResponseBody  {
  return s.Body
}

func (s *EnableResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableResourceDirectoryResponse {
  s.Headers = v
  return s
}

func (s *EnableResourceDirectoryResponse) SetStatusCode(v int32) *EnableResourceDirectoryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableResourceDirectoryResponse) SetBody(v *EnableResourceDirectoryResponseBody) *EnableResourceDirectoryResponse {
  s.Body = v
  return s
}

func (s *EnableResourceDirectoryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

