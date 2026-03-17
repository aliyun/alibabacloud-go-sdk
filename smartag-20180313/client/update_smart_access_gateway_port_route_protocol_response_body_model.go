// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayPortRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayPortRouteProtocolResponseBody struct {
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
	// A7DB1241-17CB-5457-83A9-638162A514C6
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

func (s UpdateSmartAccessGatewayPortRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayPortRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) SetCode(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayPortRouteProtocolResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolResponseBody) Validate() error {
	return dara.Validate(s)
}
