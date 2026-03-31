// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRulesReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GenerateConfigRulesReportRequest
	GetClientToken() *string
	SetConfigRuleIds(v string) *GenerateConfigRulesReportRequest
	GetConfigRuleIds() *string
}

type GenerateConfigRulesReportRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the rule. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to query the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// example:
	//
	// cr-25d86457e0d900b5****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GenerateConfigRulesReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateConfigRulesReportRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GenerateConfigRulesReportRequest) SetClientToken(v string) *GenerateConfigRulesReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateConfigRulesReportRequest) SetConfigRuleIds(v string) *GenerateConfigRulesReportRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GenerateConfigRulesReportRequest) Validate() error {
	return dara.Validate(s)
}
