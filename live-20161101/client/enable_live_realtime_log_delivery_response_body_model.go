// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLiveRealtimeLogDeliveryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableLiveRealtimeLogDeliveryResponseBody
  GetRequestId() *string 
}

type EnableLiveRealtimeLogDeliveryResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 9732E117-8A37-49FD-A36F-ABBB87556CA7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLiveRealtimeLogDeliveryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableLiveRealtimeLogDeliveryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableLiveRealtimeLogDeliveryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableLiveRealtimeLogDeliveryResponseBody) SetRequestId(v string) *EnableLiveRealtimeLogDeliveryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableLiveRealtimeLogDeliveryResponseBody) Validate() error {
  return dara.Validate(s)
}

