// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdditionalBandwidthResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAdditionalBandwidthResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAdditionalBandwidthResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAdditionalBandwidthResponseBody) *EnableAdditionalBandwidthResponse
  GetBody() *EnableAdditionalBandwidthResponseBody 
}

type EnableAdditionalBandwidthResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAdditionalBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAdditionalBandwidthResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAdditionalBandwidthResponse) GoString() string {
  return s.String()
}

func (s *EnableAdditionalBandwidthResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAdditionalBandwidthResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAdditionalBandwidthResponse) GetBody() *EnableAdditionalBandwidthResponseBody  {
  return s.Body
}

func (s *EnableAdditionalBandwidthResponse) SetHeaders(v map[string]*string) *EnableAdditionalBandwidthResponse {
  s.Headers = v
  return s
}

func (s *EnableAdditionalBandwidthResponse) SetStatusCode(v int32) *EnableAdditionalBandwidthResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAdditionalBandwidthResponse) SetBody(v *EnableAdditionalBandwidthResponseBody) *EnableAdditionalBandwidthResponse {
  s.Body = v
  return s
}

func (s *EnableAdditionalBandwidthResponse) Validate() error {
  return dara.Validate(s)
}

