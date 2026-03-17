// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayGlobalRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody struct {
	// The response code. A value of 200 indicates that the task is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message. A value of Successful indicates that the task is successful.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 407E391C-7FB6-5552-AF39-DBB5CB062E3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) SetCode(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolResponseBody) Validate() error {
	return dara.Validate(s)
}
