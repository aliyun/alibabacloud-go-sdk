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
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
