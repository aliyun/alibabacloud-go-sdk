// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateRegionResourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateRegionResourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateRegionResourceResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateRegionResourceResponseBody) *EvaluateRegionResourceResponse
  GetBody() *EvaluateRegionResourceResponseBody 
}

type EvaluateRegionResourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateRegionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateRegionResourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateRegionResourceResponse) GoString() string {
  return s.String()
}

func (s *EvaluateRegionResourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateRegionResourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateRegionResourceResponse) GetBody() *EvaluateRegionResourceResponseBody  {
  return s.Body
}

func (s *EvaluateRegionResourceResponse) SetHeaders(v map[string]*string) *EvaluateRegionResourceResponse {
  s.Headers = v
  return s
}

func (s *EvaluateRegionResourceResponse) SetStatusCode(v int32) *EvaluateRegionResourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateRegionResourceResponse) SetBody(v *EvaluateRegionResourceResponseBody) *EvaluateRegionResourceResponse {
  s.Body = v
  return s
}

func (s *EvaluateRegionResourceResponse) Validate() error {
  return dara.Validate(s)
}

