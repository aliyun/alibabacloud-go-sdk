// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteManualNodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteManualNodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteManualNodeResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteManualNodeResponseBody) *ExecuteManualNodeResponse
  GetBody() *ExecuteManualNodeResponseBody 
}

type ExecuteManualNodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteManualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteManualNodeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeResponse) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteManualNodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteManualNodeResponse) GetBody() *ExecuteManualNodeResponseBody  {
  return s.Body
}

func (s *ExecuteManualNodeResponse) SetHeaders(v map[string]*string) *ExecuteManualNodeResponse {
  s.Headers = v
  return s
}

func (s *ExecuteManualNodeResponse) SetStatusCode(v int32) *ExecuteManualNodeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteManualNodeResponse) SetBody(v *ExecuteManualNodeResponseBody) *ExecuteManualNodeResponse {
  s.Body = v
  return s
}

func (s *ExecuteManualNodeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

