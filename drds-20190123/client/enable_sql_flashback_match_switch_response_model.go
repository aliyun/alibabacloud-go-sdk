// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlFlashbackMatchSwitchResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSqlFlashbackMatchSwitchResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSqlFlashbackMatchSwitchResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSqlFlashbackMatchSwitchResponseBody) *EnableSqlFlashbackMatchSwitchResponse
  GetBody() *EnableSqlFlashbackMatchSwitchResponseBody 
}

type EnableSqlFlashbackMatchSwitchResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSqlFlashbackMatchSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchResponse) GoString() string {
  return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSqlFlashbackMatchSwitchResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSqlFlashbackMatchSwitchResponse) GetBody() *EnableSqlFlashbackMatchSwitchResponseBody  {
  return s.Body
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetHeaders(v map[string]*string) *EnableSqlFlashbackMatchSwitchResponse {
  s.Headers = v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetStatusCode(v int32) *EnableSqlFlashbackMatchSwitchResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) SetBody(v *EnableSqlFlashbackMatchSwitchResponseBody) *EnableSqlFlashbackMatchSwitchResponse {
  s.Body = v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchResponse) Validate() error {
  return dara.Validate(s)
}

