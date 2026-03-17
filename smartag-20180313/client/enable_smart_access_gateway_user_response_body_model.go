// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAccessGatewayUserResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSmartAccessGatewayUserResponseBody
  GetRequestId() *string 
}

type EnableSmartAccessGatewayUserResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // F5894299-84A2-48C1-A999-28908B99F45D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmartAccessGatewayUserResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAccessGatewayUserResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSmartAccessGatewayUserResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSmartAccessGatewayUserResponseBody) SetRequestId(v string) *EnableSmartAccessGatewayUserResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSmartAccessGatewayUserResponseBody) Validate() error {
  return dara.Validate(s)
}

