// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataCorrectResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteDataCorrectResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteDataCorrectResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteDataCorrectResponseBody) *ExecuteDataCorrectResponse
  GetBody() *ExecuteDataCorrectResponseBody 
}

type ExecuteDataCorrectResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteDataCorrectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteDataCorrectResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataCorrectResponse) GoString() string {
  return s.String()
}

func (s *ExecuteDataCorrectResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteDataCorrectResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteDataCorrectResponse) GetBody() *ExecuteDataCorrectResponseBody  {
  return s.Body
}

func (s *ExecuteDataCorrectResponse) SetHeaders(v map[string]*string) *ExecuteDataCorrectResponse {
  s.Headers = v
  return s
}

func (s *ExecuteDataCorrectResponse) SetStatusCode(v int32) *ExecuteDataCorrectResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteDataCorrectResponse) SetBody(v *ExecuteDataCorrectResponseBody) *ExecuteDataCorrectResponse {
  s.Body = v
  return s
}

func (s *ExecuteDataCorrectResponse) Validate() error {
  return dara.Validate(s)
}

