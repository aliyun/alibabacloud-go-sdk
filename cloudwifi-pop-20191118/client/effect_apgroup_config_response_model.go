// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApgroupConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EffectApgroupConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EffectApgroupConfigResponse
  GetStatusCode() *int32 
  SetBody(v *EffectApgroupConfigResponseBody) *EffectApgroupConfigResponse
  GetBody() *EffectApgroupConfigResponseBody 
}

type EffectApgroupConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EffectApgroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EffectApgroupConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s EffectApgroupConfigResponse) GoString() string {
  return s.String()
}

func (s *EffectApgroupConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EffectApgroupConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EffectApgroupConfigResponse) GetBody() *EffectApgroupConfigResponseBody  {
  return s.Body
}

func (s *EffectApgroupConfigResponse) SetHeaders(v map[string]*string) *EffectApgroupConfigResponse {
  s.Headers = v
  return s
}

func (s *EffectApgroupConfigResponse) SetStatusCode(v int32) *EffectApgroupConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *EffectApgroupConfigResponse) SetBody(v *EffectApgroupConfigResponseBody) *EffectApgroupConfigResponse {
  s.Body = v
  return s
}

func (s *EffectApgroupConfigResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

