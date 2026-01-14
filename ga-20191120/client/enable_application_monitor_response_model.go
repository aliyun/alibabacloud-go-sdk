// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationMonitorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationMonitorResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationMonitorResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationMonitorResponseBody) *EnableApplicationMonitorResponse
  GetBody() *EnableApplicationMonitorResponseBody 
}

type EnableApplicationMonitorResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationMonitorResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationMonitorResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationMonitorResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationMonitorResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationMonitorResponse) GetBody() *EnableApplicationMonitorResponseBody  {
  return s.Body
}

func (s *EnableApplicationMonitorResponse) SetHeaders(v map[string]*string) *EnableApplicationMonitorResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationMonitorResponse) SetStatusCode(v int32) *EnableApplicationMonitorResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationMonitorResponse) SetBody(v *EnableApplicationMonitorResponseBody) *EnableApplicationMonitorResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationMonitorResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

