// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVodRealtimeLogDeliveryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableVodRealtimeLogDeliveryResponseBody
  GetRequestId() *string 
}

type EnableVodRealtimeLogDeliveryResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVodRealtimeLogDeliveryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableVodRealtimeLogDeliveryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *EnableVodRealtimeLogDeliveryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableVodRealtimeLogDeliveryResponseBody) Validate() error {
  return dara.Validate(s)
}

