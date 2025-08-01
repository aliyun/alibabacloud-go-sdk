// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTraceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateTraceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateTraceResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateTraceResponseBody) *EvaluateTraceResponse
  GetBody() *EvaluateTraceResponseBody 
}

type EvaluateTraceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateTraceResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTraceResponse) GoString() string {
  return s.String()
}

func (s *EvaluateTraceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateTraceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateTraceResponse) GetBody() *EvaluateTraceResponseBody  {
  return s.Body
}

func (s *EvaluateTraceResponse) SetHeaders(v map[string]*string) *EvaluateTraceResponse {
  s.Headers = v
  return s
}

func (s *EvaluateTraceResponse) SetStatusCode(v int32) *EvaluateTraceResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateTraceResponse) SetBody(v *EvaluateTraceResponseBody) *EvaluateTraceResponse {
  s.Body = v
  return s
}

func (s *EvaluateTraceResponse) Validate() error {
  return dara.Validate(s)
}

