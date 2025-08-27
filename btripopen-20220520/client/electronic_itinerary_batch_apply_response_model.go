// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectronicItineraryBatchApplyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ElectronicItineraryBatchApplyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ElectronicItineraryBatchApplyResponse
  GetStatusCode() *int32 
  SetBody(v *ElectronicItineraryBatchApplyResponseBody) *ElectronicItineraryBatchApplyResponse
  GetBody() *ElectronicItineraryBatchApplyResponseBody 
}

type ElectronicItineraryBatchApplyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ElectronicItineraryBatchApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ElectronicItineraryBatchApplyResponse) String() string {
  return dara.Prettify(s)
}

func (s ElectronicItineraryBatchApplyResponse) GoString() string {
  return s.String()
}

func (s *ElectronicItineraryBatchApplyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ElectronicItineraryBatchApplyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ElectronicItineraryBatchApplyResponse) GetBody() *ElectronicItineraryBatchApplyResponseBody  {
  return s.Body
}

func (s *ElectronicItineraryBatchApplyResponse) SetHeaders(v map[string]*string) *ElectronicItineraryBatchApplyResponse {
  s.Headers = v
  return s
}

func (s *ElectronicItineraryBatchApplyResponse) SetStatusCode(v int32) *ElectronicItineraryBatchApplyResponse {
  s.StatusCode = &v
  return s
}

func (s *ElectronicItineraryBatchApplyResponse) SetBody(v *ElectronicItineraryBatchApplyResponseBody) *ElectronicItineraryBatchApplyResponse {
  s.Body = v
  return s
}

func (s *ElectronicItineraryBatchApplyResponse) Validate() error {
  return dara.Validate(s)
}

