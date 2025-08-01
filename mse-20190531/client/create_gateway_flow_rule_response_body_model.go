// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateGatewayFlowRuleResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateGatewayFlowRuleResponseBody
	GetRequestId() *string
}

type CreateGatewayFlowRuleResponseBody struct {
	// The ID of the rule.
	//
	// example:
	//
	// 608
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 72FC625E-9629-591B-9C01-3F0BFDAB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayFlowRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateGatewayFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayFlowRuleResponseBody) SetData(v int64) *CreateGatewayFlowRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayFlowRuleResponseBody) SetRequestId(v string) *CreateGatewayFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayFlowRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
