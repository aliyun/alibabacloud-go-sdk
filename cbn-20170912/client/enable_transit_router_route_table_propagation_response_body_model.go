// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTransitRouterRouteTablePropagationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableTransitRouterRouteTablePropagationResponseBody
  GetRequestId() *string 
}

type EnableTransitRouterRouteTablePropagationResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 59CF8BF9-DE61-421E-B903-D56AF46A303C
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableTransitRouterRouteTablePropagationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableTransitRouterRouteTablePropagationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableTransitRouterRouteTablePropagationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableTransitRouterRouteTablePropagationResponseBody) SetRequestId(v string) *EnableTransitRouterRouteTablePropagationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableTransitRouterRouteTablePropagationResponseBody) Validate() error {
  return dara.Validate(s)
}

