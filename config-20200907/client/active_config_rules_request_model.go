// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *ActiveConfigRulesRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *ActiveConfigRulesRequest
	GetConfigRuleIds() *string
}

type ActiveConfigRulesRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ActiveConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ActiveConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *ActiveConfigRulesRequest) SetCompliancePackId(v string) *ActiveConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ActiveConfigRulesRequest) SetConfigRuleIds(v string) *ActiveConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *ActiveConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
