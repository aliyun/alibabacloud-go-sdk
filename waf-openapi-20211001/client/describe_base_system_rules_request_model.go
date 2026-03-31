// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseSystemRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectType(v string) *DescribeBaseSystemRulesRequest
	GetDetectType() *string
	SetInstanceId(v string) *DescribeBaseSystemRulesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeBaseSystemRulesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeBaseSystemRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBaseSystemRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBaseSystemRulesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeBaseSystemRulesRequest
	GetResourceManagerResourceGroupId() *string
	SetRiskLevel(v string) *DescribeBaseSystemRulesRequest
	GetRiskLevel() *string
	SetRuleAction(v string) *DescribeBaseSystemRulesRequest
	GetRuleAction() *string
	SetRuleId(v int64) *DescribeBaseSystemRulesRequest
	GetRuleId() *int64
	SetRuleName(v string) *DescribeBaseSystemRulesRequest
	GetRuleName() *string
	SetRuleStatus(v int32) *DescribeBaseSystemRulesRequest
	GetRuleStatus() *int32
	SetTemplateId(v int64) *DescribeBaseSystemRulesRequest
	GetTemplateId() *int64
}

type DescribeBaseSystemRulesRequest struct {
	// example:
	//
	// sqli
	DetectType *string `json:"DetectType,omitempty" xml:"DetectType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// loose
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// block
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// example:
	//
	// 113089
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// systemRuleTest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// example:
	//
	// 24354
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeBaseSystemRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesRequest) GetDetectType() *string {
	return s.DetectType
}

func (s *DescribeBaseSystemRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBaseSystemRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBaseSystemRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBaseSystemRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBaseSystemRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBaseSystemRulesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeBaseSystemRulesRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeBaseSystemRulesRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *DescribeBaseSystemRulesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeBaseSystemRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBaseSystemRulesRequest) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *DescribeBaseSystemRulesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeBaseSystemRulesRequest) SetDetectType(v string) *DescribeBaseSystemRulesRequest {
	s.DetectType = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetInstanceId(v string) *DescribeBaseSystemRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetLang(v string) *DescribeBaseSystemRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetPageNumber(v int32) *DescribeBaseSystemRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetPageSize(v int32) *DescribeBaseSystemRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRegionId(v string) *DescribeBaseSystemRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeBaseSystemRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRiskLevel(v string) *DescribeBaseSystemRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRuleAction(v string) *DescribeBaseSystemRulesRequest {
	s.RuleAction = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRuleId(v int64) *DescribeBaseSystemRulesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRuleName(v string) *DescribeBaseSystemRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetRuleStatus(v int32) *DescribeBaseSystemRulesRequest {
	s.RuleStatus = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) SetTemplateId(v int64) *DescribeBaseSystemRulesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeBaseSystemRulesRequest) Validate() error {
	return dara.Validate(s)
}
