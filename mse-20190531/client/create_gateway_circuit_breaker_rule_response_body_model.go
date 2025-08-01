// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayCircuitBreakerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateGatewayCircuitBreakerRuleResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateGatewayCircuitBreakerRuleResponseBody
	GetRequestId() *string
}

type CreateGatewayCircuitBreakerRuleResponseBody struct {
	// example:
	//
	// 28
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayCircuitBreakerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayCircuitBreakerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayCircuitBreakerRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateGatewayCircuitBreakerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayCircuitBreakerRuleResponseBody) SetData(v int64) *CreateGatewayCircuitBreakerRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleResponseBody) SetRequestId(v string) *CreateGatewayCircuitBreakerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayCircuitBreakerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
