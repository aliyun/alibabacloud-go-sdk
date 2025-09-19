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
	SetSecurityScoreCategoryList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) *ChangeSecurityScoreRuleRequest
	GetSecurityScoreCategoryList() []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList
	SetSecurityScoreRuleList(v []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList) *ChangeSecurityScoreRuleRequest
	GetSecurityScoreRuleList() []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList
}

type ChangeSecurityScoreRuleRequest struct {
	// The old or new version of the security score rule. If you set this parameter to **home_security_score**, the new version of the security score rule is changed. Otherwise, the old version of the security score rule is changed by default.
	//
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// Specifies whether to reset to the system default rule. Valid values:
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// false
	ResetSecurityScoreRule *bool `json:"ResetSecurityScoreRule,omitempty" xml:"ResetSecurityScoreRule,omitempty"`
	// The information about the new version of the security score rule.
	SecurityScoreCategoryList []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList `json:"SecurityScoreCategoryList,omitempty" xml:"SecurityScoreCategoryList,omitempty" type:"Repeated"`
	// The information about the old version of the security score rule.
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

func (s *ChangeSecurityScoreRuleRequest) SetSecurityScoreCategoryList(v []*ChangeSecurityScoreRuleRequestSecurityScoreCategoryList) *ChangeSecurityScoreRuleRequest {
	s.SecurityScoreCategoryList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) SetSecurityScoreRuleList(v []*ChangeSecurityScoreRuleRequestSecurityScoreRuleList) *ChangeSecurityScoreRuleRequest {
	s.SecurityScoreRuleList = v
	return s
}

func (s *ChangeSecurityScoreRuleRequest) Validate() error {
	return dara.Validate(s)
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryList struct {
	// The category of the security score rule. Valid values:
	//
	// 	- **SS_SAS_HANDLE**: security governance.
	//
	// 	- **SS_SAS_RESPOND**: security response.
	//
	// example:
	//
	// SS_SAS_HANDLE
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The threshold of deduction for the security score rule type.
	//
	// >  Valid values: 0 to 100. The sum of the deduction thresholds for all deduction modules must be equal to 100.
	//
	// example:
	//
	// 20
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The deduction items of the security score rule.
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
	return dara.Validate(s)
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleList struct {
	// The deducted module that is supported by the security score feature. The type of the sub-deduction item. Valid values:
	//
	// 	- **SS_SAS_WEAK_PW**: unhandled weak password risk.
	//
	// 	- **SS_SAS_ALARM**: unhandled alert in Security Center.
	//
	// 	- **SS_SAS_EMG_VUL**: unfixed urgent vulnerability.
	//
	// 	- **SS_SAS_APP_VUL**: unfixed application vulnerability.
	//
	// 	- **SS_SAS_SYS_VUL**: unfixed system vulnerability.
	//
	// 	- **SS_SAS_CLOUD_HC**: unhandled cloud security posture management (CSPM) risk.
	//
	// 	- **SS_SDDP_DATA_RISK**: unhandled data security risk.
	//
	// 	- **SS_WAF_API_RISK**: unhandled API security risk.
	//
	// 	- **SS_DDOS_BH_ASSET**: asset on which blackhole filtering is triggered.
	//
	// 	- **SS_SAS_AK_LEAK**: unhandled AK/SK leak event.
	//
	// 	- **SS_PRODUCT_CONNECT**: security service not integrated.
	//
	// 	- **SS_KEY_CONFIG**: key feature configuration.
	//
	// 	- **SS_PRODUCT_EXPIRE**: service that is about to expire.
	//
	// 	- **SS_AI_RISK**: AI application risk.
	//
	// example:
	//
	// SS_REINFORCE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The threshold of deduction for the security score rule type.
	//
	// >  Valid values: 0 to the deduction threshold of the deduction module.
	//
	// example:
	//
	// 10
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The sub-deduction items of the security score rule.
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
	return dara.Validate(s)
}

type ChangeSecurityScoreRuleRequestSecurityScoreCategoryListSecurityRuleListSecurityScoreItemList struct {
	// The penalty point of the deduction item.
	//
	// example:
	//
	// 2
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The threshold for the deduction item.
	//
	// >  Valid values: 0 to the deduction threshold of the deduction module.
	//
	// example:
	//
	// 5
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The sub-deduction item of the security score rule.
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
	// 	- SS_REINFORCE: issue in key feature configuration
	//
	// 	- SS_ALARM: unhandled alert
	//
	// 	- SS_VUL: unfixed vulnerability
	//
	// 	- SS_HC: baseline risk
	//
	// 	- SS_CLOUD_HC: risk item of configuration assessment
	//
	// 	- SS_AK: risk of AccessKey pair leaks
	//
	// example:
	//
	// SS_ALARM
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The deduction threshold of the deduction module.
	//
	// >  Valid values: 0 to 100. The sum of the deduction thresholds for all deduction modules must be equal to 100.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction items of the deduction module.
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
	return dara.Validate(s)
}

type ChangeSecurityScoreRuleRequestSecurityScoreRuleListSecurityScoreItemList struct {
	// The penalty point of the deduction item.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The threshold for the deduction item.
	//
	// >  Valid values: 0 to the deduction threshold of the deduction module.
	//
	// example:
	//
	// 10
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The deduction item of the deduction module. The following list describes the deduction modules and their deduction items:
	//
	// 	- SS_REINFORCE: issue in key feature configuration
	//
	//     	- XPRESS_INSTALL: Security Center is not authorized.
	//
	//     	- REINFORCE_SUSPICIOUS: The anti-virus feature is disabled.
	//
	//     	- RANSOMWARE: The anti-ransomware policy is disabled.
	//
	//     	- WEB_LOCK: The web tamper proofing feature is disabled.
	//
	//     	- VIRUS_SCHEDULE_SCAN: The periodic virus scan policy is disabled.
	//
	//     	- IMAGE_REPO_SCAN: The container image scan range is not configured.
	//
	//     	- IMAGE_SCAN_TASK: The feature of one-click scan of container images for security risks is not performed.
	//
	// 	- SS_ALARM: unhandled alert
	//
	//     	- ALARM_SERIOUS: An unhandled high-risk alert event is detected.
	//
	//     	- ALARM_SUSPICIOUS: An unhandled medium-risk alarm event is detected.
	//
	//     	- ALARM_REMIND: An unhandled low-risk alarm event is detected.
	//
	// 	- SS_VUL: unfixed vulnerability
	//
	//     	- CMS_UNFIX: An unfixed Web-CMS vulnerability is detected.
	//
	//     	- WIN_UNFIX: An unfixed Windows host vulnerability is detected.
	//
	//     	- CVE_UNFIX: An unfixed Linux host vulnerability is detected.
	//
	//     	- ERM_UNFIX: An unfixed emergency vulnerability is detected.
	//
	//     	- ERM_UNCHECK: An undetected emergency vulnerability exists.
	//
	// 	- SS_HC: baseline risk
	//
	//     	- WEAK_EXPLOIT: Weak passwords are exposed to the Internet.
	//
	//     	- WEAK_PASSWORD: Weak passwords exist.
	//
	//     	- HC_EXPLOIT: The data source may be hacked.
	//
	//     	- HC_OTHER_WARNING: Security configuration risks exist.
	//
	// 	- SS_CLOUD_HC: Cloud platform configuration check item problem.
	//
	//     	- CSPM_CIEM_NOT_PASS: A CIEM check item failed the check.
	//
	//     	- CSPM_RISK_NOT_PASS: A security risk check item failed the check.
	//
	//     	- CSPM_COMPLIANCE_NOT_PASS: A compliance check item failed the check.
	//
	// 	- SS_AK: risk of AccessKey pair leaks
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
