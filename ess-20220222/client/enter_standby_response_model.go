// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterStandbyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterStandbyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterStandbyResponse
  GetStatusCode() *int32 
  SetBody(v *EnterStandbyResponseBody) *EnterStandbyResponse
  GetBody() *EnterStandbyResponseBody 
}

type EnterStandbyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterStandbyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterStandbyResponse) GoString() string {
  return s.String()
}

func (s *EnterStandbyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterStandbyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterStandbyResponse) GetBody() *EnterStandbyResponseBody  {
  return s.Body
}

func (s *EnterStandbyResponse) SetHeaders(v map[string]*string) *EnterStandbyResponse {
  s.Headers = v
  return s
}

func (s *EnterStandbyResponse) SetStatusCode(v int32) *EnterStandbyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterStandbyResponse) SetBody(v *EnterStandbyResponseBody) *EnterStandbyResponse {
  s.Body = v
  return s
}

func (s *EnterStandbyResponse) Validate() error {
  return dara.Validate(s)
}

