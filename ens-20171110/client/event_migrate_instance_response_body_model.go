// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventMigrateInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EventMigrateInstanceResponseBody
  GetRequestId() *string 
}

type EventMigrateInstanceResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 125B04C7-3D0D-4245-AF96-14E3758E3F06
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EventMigrateInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EventMigrateInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *EventMigrateInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EventMigrateInstanceResponseBody) SetRequestId(v string) *EventMigrateInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EventMigrateInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

