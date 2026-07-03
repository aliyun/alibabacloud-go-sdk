// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayQuotaRuleStatusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGatewayQuotaRuleStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayQuotaRuleStatusResponseBody
	GetRequestId() *string
}

type UpdateGatewayQuotaRuleStatusResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200, 404, 500
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The message content.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayQuotaRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) SetCode(v string) *UpdateGatewayQuotaRuleStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) SetMessage(v string) *UpdateGatewayQuotaRuleStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) SetRequestId(v string) *UpdateGatewayQuotaRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
