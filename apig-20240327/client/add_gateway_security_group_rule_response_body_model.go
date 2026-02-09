// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddGatewaySecurityGroupRuleResponseBody
	GetCode() *string
	SetMessage(v string) *AddGatewaySecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewaySecurityGroupRuleResponseBody
	GetRequestId() *string
}

type AddGatewaySecurityGroupRuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A6E90D5-A711-54F4-A489-E33C2021EDDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddGatewaySecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddGatewaySecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewaySecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetCode(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetMessage(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetRequestId(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
