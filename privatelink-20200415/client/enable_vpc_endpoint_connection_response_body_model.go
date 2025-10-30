// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointConnectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableVpcEndpointConnectionResponseBody
  GetRequestId() *string 
}

type EnableVpcEndpointConnectionResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcEndpointConnectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointConnectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointConnectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableVpcEndpointConnectionResponseBody) SetRequestId(v string) *EnableVpcEndpointConnectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableVpcEndpointConnectionResponseBody) Validate() error {
  return dara.Validate(s)
}

