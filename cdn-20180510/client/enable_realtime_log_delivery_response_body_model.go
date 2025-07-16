// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRealtimeLogDeliveryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableRealtimeLogDeliveryResponseBody
  GetRequestId() *string 
}

type EnableRealtimeLogDeliveryResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 9732E117-8A37-49FD-A36F-ABBB87556CA7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRealtimeLogDeliveryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableRealtimeLogDeliveryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableRealtimeLogDeliveryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableRealtimeLogDeliveryResponseBody) SetRequestId(v string) *EnableRealtimeLogDeliveryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableRealtimeLogDeliveryResponseBody) Validate() error {
  return dara.Validate(s)
}

