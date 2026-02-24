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
	// This parameter is required.
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
