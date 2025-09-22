// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandPhoneDataVolumeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpandPhoneDataVolumeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpandPhoneDataVolumeResponse
  GetStatusCode() *int32 
  SetBody(v *ExpandPhoneDataVolumeResponseBody) *ExpandPhoneDataVolumeResponse
  GetBody() *ExpandPhoneDataVolumeResponseBody 
}

type ExpandPhoneDataVolumeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExpandPhoneDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandPhoneDataVolumeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpandPhoneDataVolumeResponse) GoString() string {
  return s.String()
}

func (s *ExpandPhoneDataVolumeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpandPhoneDataVolumeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpandPhoneDataVolumeResponse) GetBody() *ExpandPhoneDataVolumeResponseBody  {
  return s.Body
}

func (s *ExpandPhoneDataVolumeResponse) SetHeaders(v map[string]*string) *ExpandPhoneDataVolumeResponse {
  s.Headers = v
  return s
}

func (s *ExpandPhoneDataVolumeResponse) SetStatusCode(v int32) *ExpandPhoneDataVolumeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpandPhoneDataVolumeResponse) SetBody(v *ExpandPhoneDataVolumeResponseBody) *ExpandPhoneDataVolumeResponse {
  s.Body = v
  return s
}

func (s *ExpandPhoneDataVolumeResponse) Validate() error {
  return dara.Validate(s)
}

