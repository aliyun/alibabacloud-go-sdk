// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAppInfoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditAppInfoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditAppInfoResponse
  GetStatusCode() *int32 
  SetBody(v *EditAppInfoResponseBody) *EditAppInfoResponse
  GetBody() *EditAppInfoResponseBody 
}

type EditAppInfoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditAppInfoResponse) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoResponse) GoString() string {
  return s.String()
}

func (s *EditAppInfoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditAppInfoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditAppInfoResponse) GetBody() *EditAppInfoResponseBody  {
  return s.Body
}

func (s *EditAppInfoResponse) SetHeaders(v map[string]*string) *EditAppInfoResponse {
  s.Headers = v
  return s
}

func (s *EditAppInfoResponse) SetStatusCode(v int32) *EditAppInfoResponse {
  s.StatusCode = &v
  return s
}

func (s *EditAppInfoResponse) SetBody(v *EditAppInfoResponseBody) *EditAppInfoResponse {
  s.Body = v
  return s
}

func (s *EditAppInfoResponse) Validate() error {
  return dara.Validate(s)
}

