// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayBgpRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayBgpRouteResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayBgpRouteResponseBody struct {
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
	// 7F0B079C-2D0E-4ABF-A970-C079F785A09C
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

func (s UpdateSmartAccessGatewayBgpRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayBgpRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) SetCode(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayBgpRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayBgpRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
