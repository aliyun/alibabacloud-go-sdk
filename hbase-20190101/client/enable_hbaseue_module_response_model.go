// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueModuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHBaseueModuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHBaseueModuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHBaseueModuleResponseBody) *EnableHBaseueModuleResponse
  GetBody() *EnableHBaseueModuleResponseBody 
}

type EnableHBaseueModuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHBaseueModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHBaseueModuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueModuleResponse) GoString() string {
  return s.String()
}

func (s *EnableHBaseueModuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHBaseueModuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHBaseueModuleResponse) GetBody() *EnableHBaseueModuleResponseBody  {
  return s.Body
}

func (s *EnableHBaseueModuleResponse) SetHeaders(v map[string]*string) *EnableHBaseueModuleResponse {
  s.Headers = v
  return s
}

func (s *EnableHBaseueModuleResponse) SetStatusCode(v int32) *EnableHBaseueModuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHBaseueModuleResponse) SetBody(v *EnableHBaseueModuleResponseBody) *EnableHBaseueModuleResponse {
  s.Body = v
  return s
}

func (s *EnableHBaseueModuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

