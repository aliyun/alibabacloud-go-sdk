// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceAccessControlResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInstanceAccessControlResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInstanceAccessControlResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInstanceAccessControlResponseBody) *EnableInstanceAccessControlResponse
  GetBody() *EnableInstanceAccessControlResponseBody 
}

type EnableInstanceAccessControlResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInstanceAccessControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstanceAccessControlResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceAccessControlResponse) GoString() string {
  return s.String()
}

func (s *EnableInstanceAccessControlResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInstanceAccessControlResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInstanceAccessControlResponse) GetBody() *EnableInstanceAccessControlResponseBody  {
  return s.Body
}

func (s *EnableInstanceAccessControlResponse) SetHeaders(v map[string]*string) *EnableInstanceAccessControlResponse {
  s.Headers = v
  return s
}

func (s *EnableInstanceAccessControlResponse) SetStatusCode(v int32) *EnableInstanceAccessControlResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInstanceAccessControlResponse) SetBody(v *EnableInstanceAccessControlResponseBody) *EnableInstanceAccessControlResponse {
  s.Body = v
  return s
}

func (s *EnableInstanceAccessControlResponse) Validate() error {
  return dara.Validate(s)
}

