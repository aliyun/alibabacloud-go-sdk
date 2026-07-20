// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListConfigRulesRequest(v *GetSecuritySuggestionListRequestListConfigRulesRequest) *GetSecuritySuggestionListRequest
	GetListConfigRulesRequest() *GetSecuritySuggestionListRequestListConfigRulesRequest
}

type GetSecuritySuggestionListRequest struct {
	ListConfigRulesRequest *GetSecuritySuggestionListRequestListConfigRulesRequest `json:"ListConfigRulesRequest,omitempty" xml:"ListConfigRulesRequest,omitempty" type:"Struct"`
}

func (s GetSecuritySuggestionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListRequest) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListRequest) GetListConfigRulesRequest() *GetSecuritySuggestionListRequestListConfigRulesRequest {
	return s.ListConfigRulesRequest
}

func (s *GetSecuritySuggestionListRequest) SetListConfigRulesRequest(v *GetSecuritySuggestionListRequestListConfigRulesRequest) *GetSecuritySuggestionListRequest {
	s.ListConfigRulesRequest = v
	return s
}

func (s *GetSecuritySuggestionListRequest) Validate() error {
	if s.ListConfigRulesRequest != nil {
		if err := s.ListConfigRulesRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecuritySuggestionListRequestListConfigRulesRequest struct {
	// example:
	//
	// cp-d7b061dbe91500aa179a
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetSecuritySuggestionListRequestListConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListRequestListConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetCompliancePackId(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetComplianceType(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetConfigRuleName(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetConfigRuleState(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetKeyword(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.Keyword = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetPageNumber(v int32) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetPageSize(v int32) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetResourceTypes(v string) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) SetRiskLevel(v int32) *GetSecuritySuggestionListRequestListConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *GetSecuritySuggestionListRequestListConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
