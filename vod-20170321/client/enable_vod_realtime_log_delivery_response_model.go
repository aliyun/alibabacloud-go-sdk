// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVodRealtimeLogDeliveryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableVodRealtimeLogDeliveryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableVodRealtimeLogDeliveryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableVodRealtimeLogDeliveryResponseBody) *EnableVodRealtimeLogDeliveryResponse
  GetBody() *EnableVodRealtimeLogDeliveryResponseBody 
}

type EnableVodRealtimeLogDeliveryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVodRealtimeLogDeliveryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryResponse) GoString() string {
  return s.String()
}

func (s *EnableVodRealtimeLogDeliveryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableVodRealtimeLogDeliveryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableVodRealtimeLogDeliveryResponse) GetBody() *EnableVodRealtimeLogDeliveryResponseBody  {
  return s.Body
}

func (s *EnableVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *EnableVodRealtimeLogDeliveryResponse {
  s.Headers = v
  return s
}

func (s *EnableVodRealtimeLogDeliveryResponse) SetStatusCode(v int32) *EnableVodRealtimeLogDeliveryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableVodRealtimeLogDeliveryResponse) SetBody(v *EnableVodRealtimeLogDeliveryResponseBody) *EnableVodRealtimeLogDeliveryResponse {
  s.Body = v
  return s
}

func (s *EnableVodRealtimeLogDeliveryResponse) Validate() error {
  return dara.Validate(s)
}

