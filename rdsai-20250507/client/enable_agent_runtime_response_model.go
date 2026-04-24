// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAgentRuntimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAgentRuntimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAgentRuntimeResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAgentRuntimeResponseBody) *EnableAgentRuntimeResponse
  GetBody() *EnableAgentRuntimeResponseBody 
}

type EnableAgentRuntimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAgentRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAgentRuntimeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAgentRuntimeResponse) GoString() string {
  return s.String()
}

func (s *EnableAgentRuntimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAgentRuntimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAgentRuntimeResponse) GetBody() *EnableAgentRuntimeResponseBody  {
  return s.Body
}

func (s *EnableAgentRuntimeResponse) SetHeaders(v map[string]*string) *EnableAgentRuntimeResponse {
  s.Headers = v
  return s
}

func (s *EnableAgentRuntimeResponse) SetStatusCode(v int32) *EnableAgentRuntimeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAgentRuntimeResponse) SetBody(v *EnableAgentRuntimeResponseBody) *EnableAgentRuntimeResponse {
  s.Body = v
  return s
}

func (s *EnableAgentRuntimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

