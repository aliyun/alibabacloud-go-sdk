// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcBlacklistResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EmptyAutoCcBlacklistResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EmptyAutoCcBlacklistResponse
  GetStatusCode() *int32 
  SetBody(v *EmptyAutoCcBlacklistResponseBody) *EmptyAutoCcBlacklistResponse
  GetBody() *EmptyAutoCcBlacklistResponseBody 
}

type EmptyAutoCcBlacklistResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EmptyAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmptyAutoCcBlacklistResponse) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcBlacklistResponse) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcBlacklistResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EmptyAutoCcBlacklistResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EmptyAutoCcBlacklistResponse) GetBody() *EmptyAutoCcBlacklistResponseBody  {
  return s.Body
}

func (s *EmptyAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *EmptyAutoCcBlacklistResponse {
  s.Headers = v
  return s
}

func (s *EmptyAutoCcBlacklistResponse) SetStatusCode(v int32) *EmptyAutoCcBlacklistResponse {
  s.StatusCode = &v
  return s
}

func (s *EmptyAutoCcBlacklistResponse) SetBody(v *EmptyAutoCcBlacklistResponseBody) *EmptyAutoCcBlacklistResponse {
  s.Body = v
  return s
}

func (s *EmptyAutoCcBlacklistResponse) Validate() error {
  return dara.Validate(s)
}

