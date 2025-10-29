// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterUrgentResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EffectCasterUrgentResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EffectCasterUrgentResponse
  GetStatusCode() *int32 
  SetBody(v *EffectCasterUrgentResponseBody) *EffectCasterUrgentResponse
  GetBody() *EffectCasterUrgentResponseBody 
}

type EffectCasterUrgentResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EffectCasterUrgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EffectCasterUrgentResponse) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterUrgentResponse) GoString() string {
  return s.String()
}

func (s *EffectCasterUrgentResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EffectCasterUrgentResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EffectCasterUrgentResponse) GetBody() *EffectCasterUrgentResponseBody  {
  return s.Body
}

func (s *EffectCasterUrgentResponse) SetHeaders(v map[string]*string) *EffectCasterUrgentResponse {
  s.Headers = v
  return s
}

func (s *EffectCasterUrgentResponse) SetStatusCode(v int32) *EffectCasterUrgentResponse {
  s.StatusCode = &v
  return s
}

func (s *EffectCasterUrgentResponse) SetBody(v *EffectCasterUrgentResponseBody) *EffectCasterUrgentResponse {
  s.Body = v
  return s
}

func (s *EffectCasterUrgentResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

