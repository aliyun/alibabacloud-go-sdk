// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayOspfRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayOspfRouteResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayOspfRouteResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9CF52A9E-FC3C-5CC8-B5EB-90DC03E9A531
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSmartAccessGatewayOspfRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayOspfRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) SetCode(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayOspfRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayOspfRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayOspfRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
