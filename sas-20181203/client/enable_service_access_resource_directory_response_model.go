// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceAccessResourceDirectoryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServiceAccessResourceDirectoryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServiceAccessResourceDirectoryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServiceAccessResourceDirectoryResponseBody) *EnableServiceAccessResourceDirectoryResponse
  GetBody() *EnableServiceAccessResourceDirectoryResponseBody 
}

type EnableServiceAccessResourceDirectoryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServiceAccessResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceAccessResourceDirectoryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceAccessResourceDirectoryResponse) GoString() string {
  return s.String()
}

func (s *EnableServiceAccessResourceDirectoryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServiceAccessResourceDirectoryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServiceAccessResourceDirectoryResponse) GetBody() *EnableServiceAccessResourceDirectoryResponseBody  {
  return s.Body
}

func (s *EnableServiceAccessResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableServiceAccessResourceDirectoryResponse {
  s.Headers = v
  return s
}

func (s *EnableServiceAccessResourceDirectoryResponse) SetStatusCode(v int32) *EnableServiceAccessResourceDirectoryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServiceAccessResourceDirectoryResponse) SetBody(v *EnableServiceAccessResourceDirectoryResponseBody) *EnableServiceAccessResourceDirectoryResponse {
  s.Body = v
  return s
}

func (s *EnableServiceAccessResourceDirectoryResponse) Validate() error {
  return dara.Validate(s)
}

