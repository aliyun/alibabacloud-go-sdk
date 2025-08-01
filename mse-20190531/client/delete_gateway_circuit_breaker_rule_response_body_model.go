// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteGatewayCircuitBreakerRuleResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteGatewayCircuitBreakerRuleResponseBody
	GetRequestId() *string
}

type DeleteGatewayCircuitBreakerRuleResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 52BA6DA6-A702-4362-A32F-DFF79655****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCircuitBreakerRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayCircuitBreakerRuleResponseBody) SetData(v bool) *DeleteGatewayCircuitBreakerRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleResponseBody) SetRequestId(v string) *DeleteGatewayCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
