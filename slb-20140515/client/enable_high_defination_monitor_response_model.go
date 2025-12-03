// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHighDefinationMonitorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHighDefinationMonitorResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHighDefinationMonitorResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHighDefinationMonitorResponseBody) *EnableHighDefinationMonitorResponse
  GetBody() *EnableHighDefinationMonitorResponseBody 
}

type EnableHighDefinationMonitorResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHighDefinationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHighDefinationMonitorResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHighDefinationMonitorResponse) GoString() string {
  return s.String()
}

func (s *EnableHighDefinationMonitorResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHighDefinationMonitorResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHighDefinationMonitorResponse) GetBody() *EnableHighDefinationMonitorResponseBody  {
  return s.Body
}

func (s *EnableHighDefinationMonitorResponse) SetHeaders(v map[string]*string) *EnableHighDefinationMonitorResponse {
  s.Headers = v
  return s
}

func (s *EnableHighDefinationMonitorResponse) SetStatusCode(v int32) *EnableHighDefinationMonitorResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHighDefinationMonitorResponse) SetBody(v *EnableHighDefinationMonitorResponseBody) *EnableHighDefinationMonitorResponse {
  s.Body = v
  return s
}

func (s *EnableHighDefinationMonitorResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

