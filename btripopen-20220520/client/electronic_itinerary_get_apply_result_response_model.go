// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryGetApplyResultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ElectronicItineraryGetApplyResultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ElectronicItineraryGetApplyResultResponse
  GetStatusCode() *int32 
  SetBody(v *ElectronicItineraryGetApplyResultResponseBody) *ElectronicItineraryGetApplyResultResponse
  GetBody() *ElectronicItineraryGetApplyResultResponseBody 
}

type ElectronicItineraryGetApplyResultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ElectronicItineraryGetApplyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ElectronicItineraryGetApplyResultResponse) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryGetApplyResultResponse) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryGetApplyResultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ElectronicItineraryGetApplyResultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ElectronicItineraryGetApplyResultResponse) GetBody() *ElectronicItineraryGetApplyResultResponseBody  {
  return s.Body
}

func (s *ElectronicItineraryGetApplyResultResponse) SetHeaders(v map[string]*string) *ElectronicItineraryGetApplyResultResponse {
  s.Headers = v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponse) SetStatusCode(v int32) *ElectronicItineraryGetApplyResultResponse {
  s.StatusCode = &v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponse) SetBody(v *ElectronicItineraryGetApplyResultResponseBody) *ElectronicItineraryGetApplyResultResponse {
  s.Body = v
  return s
}

func (s *ElectronicItineraryGetApplyResultResponse) Validate() error {
  return dara.Validate(s)
}

