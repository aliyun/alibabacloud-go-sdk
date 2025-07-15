// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterVideoResourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EffectCasterVideoResourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EffectCasterVideoResourceResponse
  GetStatusCode() *int32 
  SetBody(v *EffectCasterVideoResourceResponseBody) *EffectCasterVideoResourceResponse
  GetBody() *EffectCasterVideoResourceResponseBody 
}

type EffectCasterVideoResourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EffectCasterVideoResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EffectCasterVideoResourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterVideoResourceResponse) GoString() string {
  return s.String()
}

func (s *EffectCasterVideoResourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EffectCasterVideoResourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EffectCasterVideoResourceResponse) GetBody() *EffectCasterVideoResourceResponseBody  {
  return s.Body
}

func (s *EffectCasterVideoResourceResponse) SetHeaders(v map[string]*string) *EffectCasterVideoResourceResponse {
  s.Headers = v
  return s
}

func (s *EffectCasterVideoResourceResponse) SetStatusCode(v int32) *EffectCasterVideoResourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EffectCasterVideoResourceResponse) SetBody(v *EffectCasterVideoResourceResponseBody) *EffectCasterVideoResourceResponse {
  s.Body = v
  return s
}

func (s *EffectCasterVideoResourceResponse) Validate() error {
  return dara.Validate(s)
}

