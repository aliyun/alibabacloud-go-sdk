// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayFlowRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteGatewayFlowRuleResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteGatewayFlowRuleResponseBody
	GetRequestId() *string
}

type DeleteGatewayFlowRuleResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2B74E7F7-DF54-5AB1-B8F2-67391B83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayFlowRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayFlowRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFlowRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayFlowRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayFlowRuleResponseBody) SetData(v bool) *DeleteGatewayFlowRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayFlowRuleResponseBody) SetRequestId(v string) *DeleteGatewayFlowRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayFlowRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
