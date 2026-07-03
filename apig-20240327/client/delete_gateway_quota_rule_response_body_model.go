// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayQuotaRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGatewayQuotaRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteGatewayQuotaRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayQuotaRuleResponseBody
	GetRequestId() *string
}

type DeleteGatewayQuotaRuleResponseBody struct {
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

func (s DeleteGatewayQuotaRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayQuotaRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayQuotaRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGatewayQuotaRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayQuotaRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayQuotaRuleResponseBody) SetCode(v string) *DeleteGatewayQuotaRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayQuotaRuleResponseBody) SetMessage(v string) *DeleteGatewayQuotaRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayQuotaRuleResponseBody) SetRequestId(v string) *DeleteGatewayQuotaRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayQuotaRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
