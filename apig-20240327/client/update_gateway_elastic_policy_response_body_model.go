// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayElasticPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayElasticPolicyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGatewayElasticPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayElasticPolicyResponseBody
	GetRequestId() *string
}

type UpdateGatewayElasticPolicyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5B540EB6-7CF6-5326-A312-E3D68446CE07
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayElasticPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayElasticPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayElasticPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayElasticPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayElasticPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayElasticPolicyResponseBody) SetCode(v string) *UpdateGatewayElasticPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayElasticPolicyResponseBody) SetMessage(v string) *UpdateGatewayElasticPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayElasticPolicyResponseBody) SetRequestId(v string) *UpdateGatewayElasticPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayElasticPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
