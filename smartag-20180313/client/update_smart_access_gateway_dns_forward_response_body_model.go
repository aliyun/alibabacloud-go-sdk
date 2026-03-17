// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayDnsForwardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody
	GetSuccess() *string
}

type UpdateSmartAccessGatewayDnsForwardResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 284045DE-4C2D-463D-9F27-B6898E67D120
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSmartAccessGatewayDnsForwardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayDnsForwardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) SetCode(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) SetHttpStatusCode(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) SetMessage(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) SetRequestId(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) SetSuccess(v string) *UpdateSmartAccessGatewayDnsForwardResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponseBody) Validate() error {
	return dara.Validate(s)
}
