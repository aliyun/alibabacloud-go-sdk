// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceVideoQualityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnhanceVideoQualityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnhanceVideoQualityResponse
  GetStatusCode() *int32 
  SetBody(v *EnhanceVideoQualityResponseBody) *EnhanceVideoQualityResponse
  GetBody() *EnhanceVideoQualityResponseBody 
}

type EnhanceVideoQualityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnhanceVideoQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnhanceVideoQualityResponse) String() string {
  return dara.Prettify(s)
}

func (s EnhanceVideoQualityResponse) GoString() string {
  return s.String()
}

func (s *EnhanceVideoQualityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnhanceVideoQualityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnhanceVideoQualityResponse) GetBody() *EnhanceVideoQualityResponseBody  {
  return s.Body
}

func (s *EnhanceVideoQualityResponse) SetHeaders(v map[string]*string) *EnhanceVideoQualityResponse {
  s.Headers = v
  return s
}

func (s *EnhanceVideoQualityResponse) SetStatusCode(v int32) *EnhanceVideoQualityResponse {
  s.StatusCode = &v
  return s
}

func (s *EnhanceVideoQualityResponse) SetBody(v *EnhanceVideoQualityResponseBody) *EnhanceVideoQualityResponse {
  s.Body = v
  return s
}

func (s *EnhanceVideoQualityResponse) Validate() error {
  return dara.Validate(s)
}

