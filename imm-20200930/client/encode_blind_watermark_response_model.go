// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncodeBlindWatermarkResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EncodeBlindWatermarkResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EncodeBlindWatermarkResponse
  GetStatusCode() *int32 
  SetBody(v *EncodeBlindWatermarkResponseBody) *EncodeBlindWatermarkResponse
  GetBody() *EncodeBlindWatermarkResponseBody 
}

type EncodeBlindWatermarkResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EncodeBlindWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncodeBlindWatermarkResponse) String() string {
  return dara.Prettify(s)
}

func (s EncodeBlindWatermarkResponse) GoString() string {
  return s.String()
}

func (s *EncodeBlindWatermarkResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EncodeBlindWatermarkResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EncodeBlindWatermarkResponse) GetBody() *EncodeBlindWatermarkResponseBody  {
  return s.Body
}

func (s *EncodeBlindWatermarkResponse) SetHeaders(v map[string]*string) *EncodeBlindWatermarkResponse {
  s.Headers = v
  return s
}

func (s *EncodeBlindWatermarkResponse) SetStatusCode(v int32) *EncodeBlindWatermarkResponse {
  s.StatusCode = &v
  return s
}

func (s *EncodeBlindWatermarkResponse) SetBody(v *EncodeBlindWatermarkResponseBody) *EncodeBlindWatermarkResponse {
  s.Body = v
  return s
}

func (s *EncodeBlindWatermarkResponse) Validate() error {
  return dara.Validate(s)
}

