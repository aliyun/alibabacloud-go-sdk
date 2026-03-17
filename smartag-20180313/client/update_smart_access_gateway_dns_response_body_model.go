// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayDnsResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayDnsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayDnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmartAccessGatewayDnsResponseBody
	GetSuccess() *bool
}

type UpdateSmartAccessGatewayDnsResponseBody struct {
	// The error code. A value of 200 indicates that the task is successful.
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
	// AE9B6C13-86D3-5E56-8C05-BBE58C2BE671
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

func (s UpdateSmartAccessGatewayDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayDnsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) SetCode(v string) *UpdateSmartAccessGatewayDnsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayDnsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) SetSuccess(v bool) *UpdateSmartAccessGatewayDnsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponseBody) Validate() error {
	return dara.Validate(s)
}
