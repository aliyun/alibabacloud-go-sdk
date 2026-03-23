// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EffectApConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EffectApConfigResponse
  GetStatusCode() *int32 
  SetBody(v *EffectApConfigResponseBody) *EffectApConfigResponse
  GetBody() *EffectApConfigResponseBody 
}

type EffectApConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EffectApConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EffectApConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s EffectApConfigResponse) GoString() string {
  return s.String()
}

func (s *EffectApConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EffectApConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EffectApConfigResponse) GetBody() *EffectApConfigResponseBody  {
  return s.Body
}

func (s *EffectApConfigResponse) SetHeaders(v map[string]*string) *EffectApConfigResponse {
  s.Headers = v
  return s
}

func (s *EffectApConfigResponse) SetStatusCode(v int32) *EffectApConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *EffectApConfigResponse) SetBody(v *EffectApConfigResponseBody) *EffectApConfigResponse {
  s.Body = v
  return s
}

func (s *EffectApConfigResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

