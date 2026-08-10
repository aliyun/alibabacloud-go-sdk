// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDDLResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteDDLResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteDDLResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteDDLResponseBody) *ExecuteDDLResponse
  GetBody() *ExecuteDDLResponseBody 
}

type ExecuteDDLResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteDDLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteDDLResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDDLResponse) GoString() string {
  return s.String()
}

func (s *ExecuteDDLResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteDDLResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteDDLResponse) GetBody() *ExecuteDDLResponseBody  {
  return s.Body
}

func (s *ExecuteDDLResponse) SetHeaders(v map[string]*string) *ExecuteDDLResponse {
  s.Headers = v
  return s
}

func (s *ExecuteDDLResponse) SetStatusCode(v int32) *ExecuteDDLResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteDDLResponse) SetBody(v *ExecuteDDLResponseBody) *ExecuteDDLResponse {
  s.Body = v
  return s
}

func (s *ExecuteDDLResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

