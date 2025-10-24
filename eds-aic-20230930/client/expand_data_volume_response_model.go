// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandDataVolumeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpandDataVolumeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpandDataVolumeResponse
  GetStatusCode() *int32 
  SetBody(v *ExpandDataVolumeResponseBody) *ExpandDataVolumeResponse
  GetBody() *ExpandDataVolumeResponseBody 
}

type ExpandDataVolumeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExpandDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandDataVolumeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpandDataVolumeResponse) GoString() string {
  return s.String()
}

func (s *ExpandDataVolumeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpandDataVolumeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpandDataVolumeResponse) GetBody() *ExpandDataVolumeResponseBody  {
  return s.Body
}

func (s *ExpandDataVolumeResponse) SetHeaders(v map[string]*string) *ExpandDataVolumeResponse {
  s.Headers = v
  return s
}

func (s *ExpandDataVolumeResponse) SetStatusCode(v int32) *ExpandDataVolumeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpandDataVolumeResponse) SetBody(v *ExpandDataVolumeResponseBody) *ExpandDataVolumeResponse {
  s.Body = v
  return s
}

func (s *ExpandDataVolumeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

