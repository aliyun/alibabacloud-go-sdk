// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateVideoQualityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateVideoQualityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateVideoQualityResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateVideoQualityResponseBody) *EvaluateVideoQualityResponse
  GetBody() *EvaluateVideoQualityResponseBody 
}

type EvaluateVideoQualityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateVideoQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateVideoQualityResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityResponse) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateVideoQualityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateVideoQualityResponse) GetBody() *EvaluateVideoQualityResponseBody  {
  return s.Body
}

func (s *EvaluateVideoQualityResponse) SetHeaders(v map[string]*string) *EvaluateVideoQualityResponse {
  s.Headers = v
  return s
}

func (s *EvaluateVideoQualityResponse) SetStatusCode(v int32) *EvaluateVideoQualityResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateVideoQualityResponse) SetBody(v *EvaluateVideoQualityResponseBody) *EvaluateVideoQualityResponse {
  s.Body = v
  return s
}

func (s *EvaluateVideoQualityResponse) Validate() error {
  return dara.Validate(s)
}

