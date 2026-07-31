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
	SetRuleIds(v []*int64) *DescribeBaseSystemRulesRequest
	GetRuleIds() []*int64
	SetRuleName(v string) *DescribeBaseSystemRulesRequest
	GetRuleName() *string
	SetRuleStatus(v int32) *DescribeBaseSystemRulesRequest
	GetRuleStatus() *int32
	SetTemplateId(v int64) *DescribeBaseSystemRulesRequest
	GetTemplateId() *int64
}

type DescribeBaseSystemRulesRequest struct {
	// The detection module. Valid values:
	//
	// - **sqli**: SQL injection.
	//
	// - **xss**: cross-site scripting (XSS).
	//
	// - **cmdi**: OS command injection.
	//
	// - **expression_injection**: expression injection.
	//
	// - **java_deserialization**: Java deserialization.
	//
	// - **dot_net_deserialization**: .NET deserialization.
	//
	// - **php_deserialization**: PHP deserialization.
	//
	// - **code_exec**: code execution.
	//
	// - **ssrf**: server-side request forgery (SSRF).
	//
	// - **path_traversal**: path traversal.
	//
	// - **arbitrary_file_uploading**: arbitrary file upload.
	//
	// - **webshell**: webshell.
	//
	// - **rfilei**: remote file inclusion (RFI).
	//
	// - **lfilei**: local file inclusion (LFI).
	//
	// - **protocol_violation**: protocol violation.
	//
	// - **scanner_behavior**: scanner behavior.
	//
	// - **logic_flaw**: business logic bug.
	//
	// - **arbitrary_file_reading**: arbitrary file reading.
	//
	// - **arbitrary_file_download**: arbitrary file download.
	//
	// - **xxe**: XML external entity injection.
	//
	// - **csrf**: cross-site request forgery.
	//
	// - **crlf**: CRLF.
	//
	// - **other**: other.
	//
	// example:
	//
	// sqli
	DetectType *string `json:"DetectType,omitempty" xml:"DetectType,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the returned rule content. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number to return in a paged query. Default value: **1**, which indicates the first page. For more information about paging, see the PageSize parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Default value: 100, which indicates 100 entries per page. For more information about paging, see the PageNumber parameter.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The risk level. Valid values:
	//
	// - **super_strict**: Super strict.
	//
	// - **strict**: Strict.
	//
	// - **medium**: Medium.
	//
	// - **loose**: Loose.
	//
	// example:
	//
	// loose
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The rule action. Valid values:
	//
	// - **block**: Block.
	//
	// - **monitor**: Monitor.
	//
	// example:
	//
	// block
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the system protection rule to query.
	//
	// example:
	//
	// 113089
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The list of system protection rule IDs to query.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The name of the system protection rule.
	//
	// example:
	//
	// systemRuleTest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule status. Valid values:
	//
	// - **1**: Disabled.
	//
	// - **0**: Enabled.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The ID of the protection template.
	//
	// >
	//
	// > - You can specify this parameter to query the system protection rules in a specific Web core protection rule template.
	//
	// > - If this parameter is left empty, the default settings of system protection rules are queried.
	//
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

func (s *DescribeBaseSystemRulesRequest) GetRuleIds() []*int64 {
	return s.RuleIds
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

func (s *DescribeBaseSystemRulesRequest) SetRuleIds(v []*int64) *DescribeBaseSystemRulesRequest {
	s.RuleIds = v
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
