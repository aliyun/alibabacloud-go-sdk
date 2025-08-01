// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceTrafficPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetCode() *int32
	SetData(v *GatewayService) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetData() *GatewayService
	SetHttpStatusCode(v int32) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayServiceTrafficPolicyResponseBody
	GetSuccess() *bool
}

type UpdateGatewayServiceTrafficPolicyResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// {\\"GatewayUniqueId\\": \\"gw-2b8ebd75dc554c37a4279ba9917379f2\\", \\"Id\\": 417, \\"GatewayTrafficPolicy\\": {\\"LoadBalancerSettings\\": {\\"LoadbalancerType\\": \\"ROUND_ROBIN\\"}, \\"TlsSetting\\": {\\"TlsMode\\": \\"DISABLE\\"}}}
	Data *GatewayService `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 11A61389-F896-5231-A4FB-074D9E2E0055
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayServiceTrafficPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceTrafficPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetData() *GatewayService {
	return s.Data
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetCode(v int32) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetData(v *GatewayService) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetMessage(v string) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetRequestId(v string) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) SetSuccess(v bool) *UpdateGatewayServiceTrafficPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
