// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGatewaySecurityGroupRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteGatewaySecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewaySecurityGroupRuleResponseBody
	GetRequestId() *string
}

type DeleteGatewaySecurityGroupRuleResponseBody struct {
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
	// 8F94B3CC-F4BA-511E-8367-ECBBE486E595
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetCode(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetMessage(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetRequestId(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
