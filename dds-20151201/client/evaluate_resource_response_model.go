// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateResourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateResourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateResourceResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateResourceResponseBody) *EvaluateResourceResponse
  GetBody() *EvaluateResourceResponseBody 
}

type EvaluateResourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateResourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateResourceResponse) GoString() string {
  return s.String()
}

func (s *EvaluateResourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateResourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateResourceResponse) GetBody() *EvaluateResourceResponseBody  {
  return s.Body
}

func (s *EvaluateResourceResponse) SetHeaders(v map[string]*string) *EvaluateResourceResponse {
  s.Headers = v
  return s
}

func (s *EvaluateResourceResponse) SetStatusCode(v int32) *EvaluateResourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateResourceResponse) SetBody(v *EvaluateResourceResponseBody) *EvaluateResourceResponse {
  s.Body = v
  return s
}

func (s *EvaluateResourceResponse) Validate() error {
  return dara.Validate(s)
}

