// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseSystemRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBaseSystemRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeBaseSystemRulesResponseBodyRules) *DescribeBaseSystemRulesResponseBody
	GetRules() []*DescribeBaseSystemRulesResponseBodyRules
	SetTotalCount(v int64) *DescribeBaseSystemRulesResponseBody
	GetTotalCount() *int64
}

type DescribeBaseSystemRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 80736FA5-FA87-55F6-AA69-C5477C6FE6D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of system protection rules.
	Rules []*DescribeBaseSystemRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBaseSystemRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBaseSystemRulesResponseBody) GetRules() []*DescribeBaseSystemRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeBaseSystemRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBaseSystemRulesResponseBody) SetRequestId(v string) *DescribeBaseSystemRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) SetRules(v []*DescribeBaseSystemRulesResponseBodyRules) *DescribeBaseSystemRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) SetTotalCount(v int64) *DescribeBaseSystemRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBaseSystemRulesResponseBodyRules struct {
	// The CVE ID of the vulnerability that is associated with the system protection rule.
	//
	// example:
	//
	// CVE-2021-34538
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The description of the system protection rule.
	//
	// example:
	//
	// rule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of attack that the system protection rule detects. Valid values:
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
	// - **logic_flaw**: logic flaw.
	//
	// - **arbitrary_file_reading**: arbitrary file read.
	//
	// - **arbitrary_file_download**: arbitrary file download.
	//
	// - **xxe**: external entity injection.
	//
	// - **csrf**: cross-site request forgery (CSRF).
	//
	// - **crlf**: CRLF injection.
	//
	// - **other**: other.
	//
	// example:
	//
	// sqli
	DetectType *string `json:"DetectType,omitempty" xml:"DetectType,omitempty"`
	// The risk level of the system protection rule. Valid values:
	//
	// - **super_strict**: Very Strict.
	//
	// - **strict**: Strict.
	//
	// - **medium**: Medium.
	//
	// - **loose**: Loose.
	//
	// example:
	//
	// super_strict
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The action of the system protection rule. Valid values:
	//
	// - **block**: Block.
	//
	// - **monitor**: Monitor.
	//
	// example:
	//
	// block
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the system protection rule.
	//
	// example:
	//
	// 113089
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the system protection rule.
	//
	// example:
	//
	// systemRuleTest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the system protection rule. Valid values:
	//
	// - **1**: disabled.
	//
	// - **0**: enabled.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The time when the system protection rule was last updated. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665460629000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeBaseSystemRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetCveId() *string {
	return s.CveId
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetDetectType() *string {
	return s.DetectType
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetCveId(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.CveId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetDescription(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetDetectType(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.DetectType = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRiskLevel(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RiskLevel = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleAction(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleAction = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleId(v int64) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleName(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleStatus(v int32) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetUpdateTime(v int64) *DescribeBaseSystemRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
