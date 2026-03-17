// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayAdminPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayAdminPasswordResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayAdminPasswordResponseBody struct {
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
	// EC8487CD-9338-5FCF-896A-B6D1BB1646A7
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

func (s UpdateSmartAccessGatewayAdminPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayAdminPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) SetCode(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayAdminPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayAdminPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
