// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSharingWithResourceDirectoryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSharingWithResourceDirectoryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSharingWithResourceDirectoryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSharingWithResourceDirectoryResponseBody) *EnableSharingWithResourceDirectoryResponse
  GetBody() *EnableSharingWithResourceDirectoryResponseBody 
}

type EnableSharingWithResourceDirectoryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSharingWithResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSharingWithResourceDirectoryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSharingWithResourceDirectoryResponse) GoString() string {
  return s.String()
}

func (s *EnableSharingWithResourceDirectoryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSharingWithResourceDirectoryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSharingWithResourceDirectoryResponse) GetBody() *EnableSharingWithResourceDirectoryResponseBody  {
  return s.Body
}

func (s *EnableSharingWithResourceDirectoryResponse) SetHeaders(v map[string]*string) *EnableSharingWithResourceDirectoryResponse {
  s.Headers = v
  return s
}

func (s *EnableSharingWithResourceDirectoryResponse) SetStatusCode(v int32) *EnableSharingWithResourceDirectoryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSharingWithResourceDirectoryResponse) SetBody(v *EnableSharingWithResourceDirectoryResponseBody) *EnableSharingWithResourceDirectoryResponse {
  s.Body = v
  return s
}

func (s *EnableSharingWithResourceDirectoryResponse) Validate() error {
  return dara.Validate(s)
}

