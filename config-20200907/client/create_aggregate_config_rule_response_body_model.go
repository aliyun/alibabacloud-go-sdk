// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *CreateAggregateConfigRuleResponseBody
	GetConfigRuleId() *string
	SetRequestId(v string) *CreateAggregateConfigRuleResponseBody
	GetRequestId() *string
}

type CreateAggregateConfigRuleResponseBody struct {
	// The rule ID.
	//
	// example:
	//
	// cr-4e3d626622af0080****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5895065A-196C-4254-8AD8-14EFC31EEF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleResponseBody) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateAggregateConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregateConfigRuleResponseBody) SetConfigRuleId(v string) *CreateAggregateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateConfigRuleResponseBody) SetRequestId(v string) *CreateAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregateConfigRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
