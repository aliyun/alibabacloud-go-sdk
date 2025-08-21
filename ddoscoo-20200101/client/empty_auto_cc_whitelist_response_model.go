// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcWhitelistResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EmptyAutoCcWhitelistResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EmptyAutoCcWhitelistResponse
  GetStatusCode() *int32 
  SetBody(v *EmptyAutoCcWhitelistResponseBody) *EmptyAutoCcWhitelistResponse
  GetBody() *EmptyAutoCcWhitelistResponseBody 
}

type EmptyAutoCcWhitelistResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EmptyAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmptyAutoCcWhitelistResponse) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcWhitelistResponse) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcWhitelistResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EmptyAutoCcWhitelistResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EmptyAutoCcWhitelistResponse) GetBody() *EmptyAutoCcWhitelistResponseBody  {
  return s.Body
}

func (s *EmptyAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *EmptyAutoCcWhitelistResponse {
  s.Headers = v
  return s
}

func (s *EmptyAutoCcWhitelistResponse) SetStatusCode(v int32) *EmptyAutoCcWhitelistResponse {
  s.StatusCode = &v
  return s
}

func (s *EmptyAutoCcWhitelistResponse) SetBody(v *EmptyAutoCcWhitelistResponseBody) *EmptyAutoCcWhitelistResponse {
  s.Body = v
  return s
}

func (s *EmptyAutoCcWhitelistResponse) Validate() error {
  return dara.Validate(s)
}

