// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHostAvailabilityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHostAvailabilityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHostAvailabilityResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHostAvailabilityResponseBody) *EnableHostAvailabilityResponse
  GetBody() *EnableHostAvailabilityResponseBody 
}

type EnableHostAvailabilityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHostAvailabilityResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHostAvailabilityResponse) GoString() string {
  return s.String()
}

func (s *EnableHostAvailabilityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHostAvailabilityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHostAvailabilityResponse) GetBody() *EnableHostAvailabilityResponseBody  {
  return s.Body
}

func (s *EnableHostAvailabilityResponse) SetHeaders(v map[string]*string) *EnableHostAvailabilityResponse {
  s.Headers = v
  return s
}

func (s *EnableHostAvailabilityResponse) SetStatusCode(v int32) *EnableHostAvailabilityResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHostAvailabilityResponse) SetBody(v *EnableHostAvailabilityResponseBody) *EnableHostAvailabilityResponse {
  s.Body = v
  return s
}

func (s *EnableHostAvailabilityResponse) Validate() error {
  return dara.Validate(s)
}

