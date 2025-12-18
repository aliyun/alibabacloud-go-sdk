// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *UpdateAggregateConfigRuleResponseBody
	GetConfigRuleId() *string
	SetRequestId(v string) *UpdateAggregateConfigRuleResponseBody
	GetRequestId() *string
}

type UpdateAggregateConfigRuleResponseBody struct {
	// The ID of the rule.
	//
	// example:
	//
	// cr-4e3d626622af0080****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleResponseBody) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateAggregateConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggregateConfigRuleResponseBody) SetConfigRuleId(v string) *UpdateAggregateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleResponseBody) SetRequestId(v string) *UpdateAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggregateConfigRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
