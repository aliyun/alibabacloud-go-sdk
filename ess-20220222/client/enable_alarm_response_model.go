// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlarmResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAlarmResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAlarmResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAlarmResponseBody) *EnableAlarmResponse
  GetBody() *EnableAlarmResponseBody 
}

type EnableAlarmResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAlarmResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAlarmResponse) GoString() string {
  return s.String()
}

func (s *EnableAlarmResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAlarmResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAlarmResponse) GetBody() *EnableAlarmResponseBody  {
  return s.Body
}

func (s *EnableAlarmResponse) SetHeaders(v map[string]*string) *EnableAlarmResponse {
  s.Headers = v
  return s
}

func (s *EnableAlarmResponse) SetStatusCode(v int32) *EnableAlarmResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAlarmResponse) SetBody(v *EnableAlarmResponseBody) *EnableAlarmResponse {
  s.Body = v
  return s
}

func (s *EnableAlarmResponse) Validate() error {
  return dara.Validate(s)
}

