// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateRemediationsRequest
	GetAggregatorId() *string
	SetConfigRuleIds(v string) *ListAggregateRemediationsRequest
	GetConfigRuleIds() *string
}

type ListAggregateRemediationsRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ListAggregateRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateRemediationsRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *ListAggregateRemediationsRequest) SetAggregatorId(v string) *ListAggregateRemediationsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRemediationsRequest) SetConfigRuleIds(v string) *ListAggregateRemediationsRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *ListAggregateRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
