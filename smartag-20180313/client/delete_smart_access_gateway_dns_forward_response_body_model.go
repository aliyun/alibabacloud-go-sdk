// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayDnsForwardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody
	GetSuccess() *string
}

type DeleteSmartAccessGatewayDnsForwardResponseBody struct {
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
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0BAAF2B9-88B8-4574-BDBE-102A90EE3FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSmartAccessGatewayDnsForwardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayDnsForwardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) SetCode(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) SetHttpStatusCode(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) SetMessage(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) SetRequestId(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) SetSuccess(v string) *DeleteSmartAccessGatewayDnsForwardResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponseBody) Validate() error {
	return dara.Validate(s)
}
