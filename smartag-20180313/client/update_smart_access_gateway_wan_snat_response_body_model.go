// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayWanSnatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayWanSnatResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayWanSnatResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayWanSnatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayWanSnatResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayWanSnatResponseBody struct {
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
	// FAD1E4CB-52A5-520B-BE14-A78F491FBD9C
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

func (s UpdateSmartAccessGatewayWanSnatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayWanSnatResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) SetCode(v string) *UpdateSmartAccessGatewayWanSnatResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayWanSnatResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayWanSnatResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayWanSnatResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponseBody) Validate() error {
	return dara.Validate(s)
}
