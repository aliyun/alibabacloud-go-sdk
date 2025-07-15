// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLiveRealtimeLogDeliveryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableLiveRealtimeLogDeliveryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableLiveRealtimeLogDeliveryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableLiveRealtimeLogDeliveryResponseBody) *EnableLiveRealtimeLogDeliveryResponse
  GetBody() *EnableLiveRealtimeLogDeliveryResponseBody 
}

type EnableLiveRealtimeLogDeliveryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableLiveRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableLiveRealtimeLogDeliveryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableLiveRealtimeLogDeliveryResponse) GoString() string {
  return s.String()
}

func (s *EnableLiveRealtimeLogDeliveryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableLiveRealtimeLogDeliveryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableLiveRealtimeLogDeliveryResponse) GetBody() *EnableLiveRealtimeLogDeliveryResponseBody  {
  return s.Body
}

func (s *EnableLiveRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *EnableLiveRealtimeLogDeliveryResponse {
  s.Headers = v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryResponse) SetStatusCode(v int32) *EnableLiveRealtimeLogDeliveryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryResponse) SetBody(v *EnableLiveRealtimeLogDeliveryResponseBody) *EnableLiveRealtimeLogDeliveryResponse {
  s.Body = v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryResponse) Validate() error {
  return dara.Validate(s)
}

