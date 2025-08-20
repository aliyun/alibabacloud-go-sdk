// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeSetResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteChangeSetResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteChangeSetResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteChangeSetResponseBody) *ExecuteChangeSetResponse
  GetBody() *ExecuteChangeSetResponseBody 
}

type ExecuteChangeSetResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteChangeSetResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeSetResponse) GoString() string {
  return s.String()
}

func (s *ExecuteChangeSetResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteChangeSetResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteChangeSetResponse) GetBody() *ExecuteChangeSetResponseBody  {
  return s.Body
}

func (s *ExecuteChangeSetResponse) SetHeaders(v map[string]*string) *ExecuteChangeSetResponse {
  s.Headers = v
  return s
}

func (s *ExecuteChangeSetResponse) SetStatusCode(v int32) *ExecuteChangeSetResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteChangeSetResponse) SetBody(v *ExecuteChangeSetResponseBody) *ExecuteChangeSetResponse {
  s.Body = v
  return s
}

func (s *ExecuteChangeSetResponse) Validate() error {
  return dara.Validate(s)
}

