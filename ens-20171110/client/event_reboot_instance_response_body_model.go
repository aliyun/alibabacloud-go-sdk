// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRebootInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EventRebootInstanceResponseBody
  GetRequestId() *string 
}

type EventRebootInstanceResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 125B04C7-3D0D-4245-AF96-14E3758E3F06
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EventRebootInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EventRebootInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *EventRebootInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EventRebootInstanceResponseBody) SetRequestId(v string) *EventRebootInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EventRebootInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

