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
	// The request parameters.
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
	// The compliance package ID.
	//
	// example:
	//
	// cp-d7b061dbe91500aa179a
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - **COMPLIANT**: Compliant.
	//
	// - **NON_COMPLIANT**: Non-compliant.
	//
	// - **NOT_APPLICABLE**: Not applicable.
	//
	// - **INSUFFICIENT_DATA**: Insufficient data.
	//
	// - **IGNORED**: Ignored.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The rule name.
	//
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The rule running status. Valid values:
	//
	// - **ACTIVE**: Active.
	//
	// - **DELETING**: Being deleted.
	//
	// - **EVALUATING**: Being evaluated.
	//
	// - **INACTIVE**: Inactive.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The query keyword.
	//
	// Supports fuzzy match on the rule ID, rule name, rule description, and rule template identifier.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// > Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// > Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type evaluated by the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - **1**: High risk.
	//
	// - **2**: Medium risk.
	//
	// - **3**: Low risk.
	//
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
