// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRedeployInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EventRedeployInstanceResponseBody
  GetRequestId() *string 
}

type EventRedeployInstanceResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 125B04C7-3D0D-4245-AF96-14E3758E3F06
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EventRedeployInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EventRedeployInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *EventRedeployInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EventRedeployInstanceResponseBody) SetRequestId(v string) *EventRedeployInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EventRedeployInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

