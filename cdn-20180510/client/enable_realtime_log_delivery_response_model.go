// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRealtimeLogDeliveryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableRealtimeLogDeliveryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableRealtimeLogDeliveryResponse
  GetStatusCode() *int32 
  SetBody(v *EnableRealtimeLogDeliveryResponseBody) *EnableRealtimeLogDeliveryResponse
  GetBody() *EnableRealtimeLogDeliveryResponseBody 
}

type EnableRealtimeLogDeliveryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableRealtimeLogDeliveryResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableRealtimeLogDeliveryResponse) GoString() string {
  return s.String()
}

func (s *EnableRealtimeLogDeliveryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableRealtimeLogDeliveryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableRealtimeLogDeliveryResponse) GetBody() *EnableRealtimeLogDeliveryResponseBody  {
  return s.Body
}

func (s *EnableRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *EnableRealtimeLogDeliveryResponse {
  s.Headers = v
  return s
}

func (s *EnableRealtimeLogDeliveryResponse) SetStatusCode(v int32) *EnableRealtimeLogDeliveryResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableRealtimeLogDeliveryResponse) SetBody(v *EnableRealtimeLogDeliveryResponseBody) *EnableRealtimeLogDeliveryResponse {
  s.Body = v
  return s
}

func (s *EnableRealtimeLogDeliveryResponse) Validate() error {
  return dara.Validate(s)
}

