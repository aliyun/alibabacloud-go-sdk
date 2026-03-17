// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAGDpiMonitorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSmartAGDpiMonitorResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSmartAGDpiMonitorResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSmartAGDpiMonitorResponseBody) *EnableSmartAGDpiMonitorResponse
  GetBody() *EnableSmartAGDpiMonitorResponseBody 
}

type EnableSmartAGDpiMonitorResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSmartAGDpiMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSmartAGDpiMonitorResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAGDpiMonitorResponse) GoString() string {
  return s.String()
}

func (s *EnableSmartAGDpiMonitorResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSmartAGDpiMonitorResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSmartAGDpiMonitorResponse) GetBody() *EnableSmartAGDpiMonitorResponseBody  {
  return s.Body
}

func (s *EnableSmartAGDpiMonitorResponse) SetHeaders(v map[string]*string) *EnableSmartAGDpiMonitorResponse {
  s.Headers = v
  return s
}

func (s *EnableSmartAGDpiMonitorResponse) SetStatusCode(v int32) *EnableSmartAGDpiMonitorResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSmartAGDpiMonitorResponse) SetBody(v *EnableSmartAGDpiMonitorResponseBody) *EnableSmartAGDpiMonitorResponse {
  s.Body = v
  return s
}

func (s *EnableSmartAGDpiMonitorResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

