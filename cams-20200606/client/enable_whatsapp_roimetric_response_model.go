// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWhatsappROIMetricResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWhatsappROIMetricResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWhatsappROIMetricResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWhatsappROIMetricResponseBody) *EnableWhatsappROIMetricResponse
  GetBody() *EnableWhatsappROIMetricResponseBody 
}

type EnableWhatsappROIMetricResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWhatsappROIMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWhatsappROIMetricResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWhatsappROIMetricResponse) GoString() string {
  return s.String()
}

func (s *EnableWhatsappROIMetricResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWhatsappROIMetricResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWhatsappROIMetricResponse) GetBody() *EnableWhatsappROIMetricResponseBody  {
  return s.Body
}

func (s *EnableWhatsappROIMetricResponse) SetHeaders(v map[string]*string) *EnableWhatsappROIMetricResponse {
  s.Headers = v
  return s
}

func (s *EnableWhatsappROIMetricResponse) SetStatusCode(v int32) *EnableWhatsappROIMetricResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWhatsappROIMetricResponse) SetBody(v *EnableWhatsappROIMetricResponseBody) *EnableWhatsappROIMetricResponse {
  s.Body = v
  return s
}

func (s *EnableWhatsappROIMetricResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

