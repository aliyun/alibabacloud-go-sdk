// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateConfigRulesReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GenerateAggregateConfigRulesReportRequest
	GetAggregatorId() *string
	SetClientToken(v string) *GenerateAggregateConfigRulesReportRequest
	GetClientToken() *string
	SetConfigRuleIds(v string) *GenerateAggregateConfigRulesReportRequest
	GetConfigRuleIds() *string
}

type GenerateAggregateConfigRulesReportRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule ID. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-25d86457e0d900b5****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GenerateAggregateConfigRulesReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GenerateAggregateConfigRulesReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateAggregateConfigRulesReportRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GenerateAggregateConfigRulesReportRequest) SetAggregatorId(v string) *GenerateAggregateConfigRulesReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportRequest) SetClientToken(v string) *GenerateAggregateConfigRulesReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportRequest) SetConfigRuleIds(v string) *GenerateAggregateConfigRulesReportRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportRequest) Validate() error {
	return dara.Validate(s)
}
