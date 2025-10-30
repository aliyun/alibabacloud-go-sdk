// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointZoneConnectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableVpcEndpointZoneConnectionResponseBody
  GetRequestId() *string 
}

type EnableVpcEndpointZoneConnectionResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointZoneConnectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableVpcEndpointZoneConnectionResponseBody) SetRequestId(v string) *EnableVpcEndpointZoneConnectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionResponseBody) Validate() error {
  return dara.Validate(s)
}

