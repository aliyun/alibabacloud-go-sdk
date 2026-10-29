// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestSGResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteRequestSGResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteRequestSGResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteRequestSGResponseBody) *ExecuteRequestSGResponse
  GetBody() *ExecuteRequestSGResponseBody 
}

type ExecuteRequestSGResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteRequestSGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteRequestSGResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestSGResponse) GoString() string {
  return s.String()
}

func (s *ExecuteRequestSGResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteRequestSGResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteRequestSGResponse) GetBody() *ExecuteRequestSGResponseBody  {
  return s.Body
}

func (s *ExecuteRequestSGResponse) SetHeaders(v map[string]*string) *ExecuteRequestSGResponse {
  s.Headers = v
  return s
}

func (s *ExecuteRequestSGResponse) SetStatusCode(v int32) *ExecuteRequestSGResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteRequestSGResponse) SetBody(v *ExecuteRequestSGResponseBody) *ExecuteRequestSGResponse {
  s.Body = v
  return s
}

func (s *ExecuteRequestSGResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

