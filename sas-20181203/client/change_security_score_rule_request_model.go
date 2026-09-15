// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSecurityScoreRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalType(v string) *ChangeSecurityScoreRuleRequest
	GetCalType() *string
	SetResetSecurityScoreRule(v bool) *ChangeSecurityScoreRuleRequest
	GetResetSecurityScoreRule() *bool
	SetResourceDirectoryAccountId(v int64) *ChangeSecurityScoreRuleRequest
	GetResourceDirectoryAccountId() *int64
	SetSecurityScoreCategoryList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) *ChangeSecurityScoreRuleRequest
	GetSecurityScoreCategoryList() []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList
	SetSecurityScoreRuleList(v []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList) *ChangeSecurityScoreRuleRequest
	GetSecurityScoreRuleList() []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList
}

type ChangeSecurityScoreRuleRequest struct {
	// Specifies whether to modify the new version or legacy security score rules. If the value is **home_security_score**, the new version security score rules are modified. Otherwise, the legacy security score rules are modified by default.
	//
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// Specifies whether to reset to the system default rules. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	ResetSecurityScoreRule *bool `json:"ResetSecurityScoreRule,omitempty" xml:"ResetSecurityScoreRule,omitempty"`
	// The ID of the member account in the resource directory.
	//
	// > Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The list of new version security score rule deductions.
	SecurityScoreCategoryList []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList `json:"SecurityScoreCategoryList,omitempty" xml:"SecurityScoreCategoryList,omitempty" type:"Repeated"`
	// The list of legacy security score rules.
	SecurityScoreRuleList []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList `json:"SecurityScoreRuleList,omitempty" xml:"SecurityScoreRuleList,omitempty" type:"Repeated"`
}

func (s ChangeSecurityScoreRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequest) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequest) GetCalType() *string {
	return s.CalType
}

func (s *ChangeSecurityScoreRuleRequest) GetResetSecurityScoreRule() *bool {
	return s.ResetSecurityScoreRule
}

func (s *ChangeSecurityScoreRuleRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ChangeSecurityScoreRuleRequest) GetSecurityScoreCategoryList() []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList {
	return s.SecurityScoreCategoryList
}

func (s *ChangeSecurityScoreRuleRequest) GetSecurityScoreRuleList() []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList {
	return s.SecurityScoreRuleList
}

func (s *ChangeSecurityScoreRuleRequest) SetCalType(v string) *ChangeSecurityScoreRuleRequest {
	s.CalType = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) SetResetSecurityScoreRule(v bool) *ChangeSecurityScoreRuleRequest {
	s.ResetSecurityScoreRule = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) SetResourceDirectoryAccountId(v int64) *ChangeSecurityScoreRuleRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) SetSecurityScoreCategoryList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) *ChangeSecurityScoreRuleRequest {
	s.SecurityScoreCategoryList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) SetSecurityScoreRuleList(v []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList) *ChangeSecurityScoreRuleRequest {
	s.SecurityScoreRuleList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) Validate() error {
	if s.SecurityScoreCategoryList != nil {
		for _, item := range s.SecurityScoreCategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityScoreRuleList != nil {
		for _, item := range s.SecurityScoreRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryList struct {
	// The category of the security score rule. Valid values:
	//
	// - **SS_SAS_HANDLE**: Security governance.
	//
	// - **SS_SAS_RESPOND**: Security response.
	//
	// example:
	//
	// SS_SAS_HANDLE
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The deduction threshold for the security score rule category.
	//
	// > The valid range is 0 to 100. The sum of all security score rule category deduction thresholds must equal 100.
	//
	// example:
	//
	// 20
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The list of deductions by security score rule type.
	SecurityRuleList []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList `json:"SecurityRuleList,omitempty" xml:"SecurityRuleList,omitempty" type:"Repeated"`
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) GetCategory() *string {
	return s.Category
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) GetScoreThreshold() *int32 {
	return s.ScoreThreshold
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) GetSecurityRuleList() []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList {
	return s.SecurityRuleList
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) SetCategory(v string) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList {
	s.Category = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) SetScoreThreshold(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList {
	s.ScoreThreshold = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) SetSecurityRuleList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList {
	s.SecurityRuleList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) Validate() error {
	if s.SecurityRuleList != nil {
		for _, item := range s.SecurityRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList struct {
	// The type of the security score rule sub-item. Valid values:
	//
	// - **SS_SAS_WEAK_PW**: Pending weak passwords to fix.
	//
	// - **SS_SAS_ALARM**: Pending Security Center alerts.
	//
	// - **SS_SAS_EMG_VUL**: Pending emergency vulnerabilities to fix.
	//
	// - **SS_SAS_APP_VUL**: Pending application vulnerabilities to fix.
	//
	// - **SS_SAS_SYS_VUL**: Pending system vulnerabilities to fix.
	//
	// - **SS_SAS_CLOUD_HC**: Pending Cloud Security Posture Management (CSPM) risks.
	//
	// - **SS_SDDP_DATA_RISK**: Pending data security risks to address.
	//
	// - **SS_WAF_API_RISK**: Pending API security risks.
	//
	// - **SS_DDOS_BH_ASSET**: Assets in DDoS blackhole filtering status.
	//
	// - **SS_SAS_AK_LEAK**: Unhandled AccessKey/SecretKey leak events.
	//
	// - **SS_PRODUCT_CONNECT**: Security products not properly connected.
	//
	// - **SS_KEY_CONFIG**: Key feature configuration.
	//
	// - **SS_PRODUCT_EXPIRE**: Products about to expire.
	//
	// - **SS_AI_RISK**: AI application risks.
	//
	// example:
	//
	// SS_REINFORCE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The deduction threshold for the security score rule type.
	//
	// > The valid range is 0 to the deduction threshold of the security score rule category.
	//
	// example:
	//
	// 10
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The list of deductions for security score rule sub-items.
	SecurityScoreItemList []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) GetScore() *int32 {
	return s.Score
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) GetSecurityScoreItemList() []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	return s.SecurityScoreItemList
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) SetRuleType(v string) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList {
	s.RuleType = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) SetScore(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList {
	s.Score = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) SetSecurityScoreItemList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList {
	s.SecurityScoreItemList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList) Validate() error {
	if s.SecurityScoreItemList != nil {
		for _, item := range s.SecurityScoreItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList struct {
	// The deduction value for the individual item.
	//
	// example:
	//
	// 2
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction threshold for the individual item.
	//
	// > The valid range is 0 to the deduction threshold of the security score rule type.
	//
	// example:
	//
	// 5
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The security score rule sub-item.
	//
	// example:
	//
	// SSI_KEY_CONFIG
	SubRuleType *string `json:"SubRuleType,omitempty" xml:"SubRuleType,omitempty"`
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetScore() *int32 {
	return s.Score
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetScoreThreshold() *int32 {
	return s.ScoreThreshold
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetSubRuleType() *string {
	return s.SubRuleType
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetScore(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.Score = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetScoreThreshold(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.ScoreThreshold = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetSubRuleType(v string) *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.SubRuleType = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) Validate() error {
	return dara.Validate(s)
}

type ChangeSecurityScoreRuleRequestSecurityScoreRuleList struct {
	// The type of the security score rule. Valid values:
	//
	// - SS_REINFORCE: Key feature configuration.
	//
	// - SS_ALARM: Pending alerts.
	//
	// - SS_VUL: Pending vulnerabilities.
	//
	// - SS_HC: Baseline issues.
	//
	// - SS_CLOUD_HC: Cloud platform configuration check item issues.
	//
	// - SS_AK: AccessKey pair leak risk.
	//
	// example:
	//
	// SS_ALARM
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The deduction value of the security score rule.
	//
	// > The valid range is 0 to 100. The sum of all security score rule deduction thresholds must equal 100.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The list of individual deduction items for the security score rule.
	SecurityScoreItemList []*ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreRuleList) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreRuleList) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) GetScore() *int32 {
	return s.Score
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) GetSecurityScoreItemList() []*ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList {
	return s.SecurityScoreItemList
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) SetRuleType(v string) *ChangeSecurityScoreRuleRequestSecurityScoreRuleList {
	s.RuleType = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) SetScore(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreRuleList {
	s.Score = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) SetSecurityScoreItemList(v []*ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) *ChangeSecurityScoreRuleRequestSecurityScoreRuleList {
	s.SecurityScoreItemList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleList) Validate() error {
	if s.SecurityScoreItemList != nil {
		for _, item := range s.SecurityScoreItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList struct {
	// The deduction value for the individual item.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction threshold for the individual item.
	//
	// > The valid range is 0 to the deduction threshold of the security score rule.
	//
	// example:
	//
	// 10
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The sub-rule type of the individual deduction item. The mapping between security score types and sub-rule types is as follows:
	//
	// - SS_REINFORCE: Key feature configuration.
	//
	//   - XPRESS_INSTALL: Security Center service authorization is not enabled.
	//
	//   - REINFORCE_SUSPICIOUS: The anti-virus feature is not enabled.
	//
	//   - RANSOMWARE: The anti-ransomware policy is not enabled.
	//
	//   - WEB_LOCK: The web tamper-proofing feature is not enabled.
	//
	//   - VIRUS_SCHEDULE_SCAN: The periodic virus scan policy is not enabled.
	//
	//   - IMAGE_REPO_SCAN: The container image scan scope is not configured.
	//
	//   - IMAGE_SCAN_TASK: The one-click container image security risk scan has not been executed.
	//
	// - SS_ALARM: Pending alerts.
	//
	//   - ALARM_SERIOUS: One unhandled high-risk alert event exists.
	//
	//   - ALARM_SUSPICIOUS: One unhandled medium-risk alert event exists.
	//
	//   - ALARM_REMIND: One unhandled low-risk alert event exists.
	//
	// - SS_VUL: Pending vulnerabilities.
	//
	//   - CMS_UNFIX: One unfixed CMS vulnerability exists.
	//
	//   - WIN_UNFIX: One unfixed Windows host vulnerability exists.
	//
	//   - CVE_UNFIX: One unfixed Linux host vulnerability exists.
	//
	//   - ERM_UNFIX: One unfixed emergency vulnerability exists.
	//
	//   - ERM_UNCHECK: One unscanned emergency vulnerability exists.
	//
	// - SS_HC: Baseline issues.
	//
	//   - WEAK_EXPLOIT: A weak password risk exposed to the Internet exists.
	//
	//   - WEAK_PASSWORD: A weak password risk exists.
	//
	//   - HC_EXPLOIT: A high-risk intrusion vulnerability exists.
	//
	//   - HC_OTHER_WARNING: A security configuration risk exists.
	//
	// - SS_CLOUD_HC: Cloud platform configuration check item issues.
	//
	//   - CSPM_CIEM_NOT_PASS: One failed CIEM check item exists.
	//
	//   - CSPM_RISK_NOT_PASS: One failed security risk check item exists.
	//
	//   - CSPM_COMPLIANCE_NOT_PASS: One failed compliance check item exists.
	//
	// - SS_AK: AccessKey pair leak risk. Categorization not applicable.
	//
	// example:
	//
	// ALARM_SERIOUS
	SubRuleType *string `json:"SubRuleType,omitempty" xml:"SubRuleType,omitempty"`
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) GetScore() *int32 {
	return s.Score
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) GetScoreThreshold() *int32 {
	return s.ScoreThreshold
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) GetSubRuleType() *string {
	return s.SubRuleType
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) SetScore(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList {
	s.Score = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) SetScoreThreshold(v int32) *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList {
	s.ScoreThreshold = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) SetSubRuleType(v string) *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList {
	s.SubRuleType = &v
	return s
}

func (s *ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList) Validate() error {
	return dara.Validate(s)
}
