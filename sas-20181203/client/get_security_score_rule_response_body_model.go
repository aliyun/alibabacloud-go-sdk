// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityScoreRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableStatus(v bool) *GetSecurityScoreRuleResponseBody
	GetEnableStatus() *bool
	SetRequestId(v string) *GetSecurityScoreRuleResponseBody
	GetRequestId() *string
	SetSecurityScoreCategoryList(v []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) *GetSecurityScoreRuleResponseBody
	GetSecurityScoreCategoryList() []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList
	SetSecurityScoreRuleList(v []*GetSecurityScoreRuleResponseBodySecurityScoreRuleList) *GetSecurityScoreRuleResponseBody
	GetSecurityScoreRuleList() []*GetSecurityScoreRuleResponseBodySecurityScoreRuleList
}

type GetSecurityScoreRuleResponseBody struct {
	// The enabling status of the custom security scoring rule. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of new security score rules.
	SecurityScoreCategoryList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList `json:"SecurityScoreCategoryList,omitempty" xml:"SecurityScoreCategoryList,omitempty" type:"Repeated"`
	// The list of legacy security score rules.
	SecurityScoreRuleList []*GetSecurityScoreRuleResponseBodySecurityScoreRuleList `json:"SecurityScoreRuleList,omitempty" xml:"SecurityScoreRuleList,omitempty" type:"Repeated"`
}

func (s GetSecurityScoreRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBody) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *GetSecurityScoreRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityScoreRuleResponseBody) GetSecurityScoreCategoryList() []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList {
	return s.SecurityScoreCategoryList
}

func (s *GetSecurityScoreRuleResponseBody) GetSecurityScoreRuleList() []*GetSecurityScoreRuleResponseBodySecurityScoreRuleList {
	return s.SecurityScoreRuleList
}

func (s *GetSecurityScoreRuleResponseBody) SetEnableStatus(v bool) *GetSecurityScoreRuleResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBody) SetRequestId(v string) *GetSecurityScoreRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBody) SetSecurityScoreCategoryList(v []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) *GetSecurityScoreRuleResponseBody {
	s.SecurityScoreCategoryList = v
	return s
}

func (s *GetSecurityScoreRuleResponseBody) SetSecurityScoreRuleList(v []*GetSecurityScoreRuleResponseBodySecurityScoreRuleList) *GetSecurityScoreRuleResponseBody {
	s.SecurityScoreRuleList = v
	return s
}

func (s *GetSecurityScoreRuleResponseBody) Validate() error {
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

type GetSecurityScoreRuleResponseBodySecurityScoreCategoryList struct {
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
	// The deduction threshold of the security score rule category.
	//
	// example:
	//
	// 30
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction list of security score rule types.
	SecurityRuleList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList `json:"SecurityRuleList,omitempty" xml:"SecurityRuleList,omitempty" type:"Repeated"`
	// The name of the security score rule category.
	//
	// example:
	//
	// Security Response
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) GetCategory() *string {
	return s.Category
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) GetScore() *int32 {
	return s.Score
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) GetSecurityRuleList() []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList {
	return s.SecurityRuleList
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) GetTitle() *string {
	return s.Title
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) SetCategory(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList {
	s.Category = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) SetScore(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList {
	s.Score = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) SetSecurityRuleList(v []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList {
	s.SecurityRuleList = v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) SetTitle(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList {
	s.Title = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryList) Validate() error {
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

type GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList struct {
	// The type of the security score rule.
	//
	// example:
	//
	// SS_AI_RISK
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The deduction threshold of the security score rule type.
	//
	// example:
	//
	// 10
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction list of security score rule sub-items.
	SecurityScoreItemList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
	// The name of the security score rule type.
	//
	// example:
	//
	// AI Application Risks
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) GetScore() *int32 {
	return s.Score
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) GetSecurityScoreItemList() []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	return s.SecurityScoreItemList
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) GetTitle() *string {
	return s.Title
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) SetRuleType(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList {
	s.RuleType = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) SetScore(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList {
	s.Score = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) SetSecurityScoreItemList(v []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList {
	s.SecurityScoreItemList = v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) SetTitle(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList {
	s.Title = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList) Validate() error {
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

type GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList struct {
	// The deduction value of the individual item.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction threshold of the individual item.
	//
	// example:
	//
	// 10
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
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
	// - **SS_SDDP_DATA_RISK**: Pending data security risks to remediate.
	//
	// - **SS_WAF_API_RISK**: Pending API security risks.
	//
	// - **SS_DDOS_BH_ASSET**: Assets in Black Hole Activated status.
	//
	// - **SS_SAS_AK_LEAK**: Unhandled AccessKey/SecretKey leakage events.
	//
	// - **SS_PRODUCT_CONNECT**: Security products not in Normal connection status.
	//
	// - **SS_KEY_CONFIG**: Key feature configuration.
	//
	// - **SS_PRODUCT_EXPIRE**: Products about to expire.
	//
	// - **SS_AI_RISK**: AI application risks.
	//
	// example:
	//
	// SSI_AI_VUL_RISK
	SubRuleType *string `json:"SubRuleType,omitempty" xml:"SubRuleType,omitempty"`
	// The name of the security score rule sub-item.
	//
	// example:
	//
	// Unhandled application vulnerabilities exist.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetScore() *int32 {
	return s.Score
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetScoreThreshold() *int32 {
	return s.ScoreThreshold
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetSubRuleType() *string {
	return s.SubRuleType
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) GetTitle() *string {
	return s.Title
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetScore(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.Score = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetScoreThreshold(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.ScoreThreshold = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetSubRuleType(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.SubRuleType = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) SetTitle(v string) *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList {
	s.Title = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList) Validate() error {
	return dara.Validate(s)
}

type GetSecurityScoreRuleResponseBodySecurityScoreRuleList struct {
	// The type of the security score rule. Valid values:
	//
	// - SS_REINFORCE: Key feature configuration.
	//
	// - SS_ALARM: Pending alerts.
	//
	// - SS_VUL: Pending vulnerabilities to fix.
	//
	// - SS_HC: Baseline issues.
	//
	// - SS_CLOUD_HC: Cloud platform configuration check item issues.
	//
	// - SS_AK: AccessKey leakage risk exists.
	//
	// example:
	//
	// SS_ALARM
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The deduction value of the security score rule.
	//
	// > The configurable range is 0 to 100 points. The sum of all security score rule deduction thresholds must equal 100 points.
	//
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The list of individual deduction items for the security score rule.
	SecurityScoreItemList []*GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
	// The description of the security score rule.
	//
	// example:
	//
	// Unhandled Alerts
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreRuleList) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) GetScore() *int32 {
	return s.Score
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) GetSecurityScoreItemList() []*GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList {
	return s.SecurityScoreItemList
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) GetTitle() *string {
	return s.Title
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) SetRuleType(v string) *GetSecurityScoreRuleResponseBodySecurityScoreRuleList {
	s.RuleType = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) SetScore(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreRuleList {
	s.Score = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) SetSecurityScoreItemList(v []*GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) *GetSecurityScoreRuleResponseBodySecurityScoreRuleList {
	s.SecurityScoreItemList = v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) SetTitle(v string) *GetSecurityScoreRuleResponseBodySecurityScoreRuleList {
	s.Title = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleList) Validate() error {
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

type GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList struct {
	// The deduction value of the individual item.
	//
	// example:
	//
	// 3
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction threshold of the individual item.
	//
	// > The configurable range is 0 to the deduction threshold of the security score rule.
	//
	// example:
	//
	// 5
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The sub-rule type of the security score deduction item. The relationship between security score categorization types and sub-rule types is as follows:
	//
	// - SS_REINFORCE: Key feature configuration.
	//
	//   - XPRESS_INSTALL: Security Center service authorization is not enabled.
	//
	//   - REINFORCE_SUSPICIOUS: Anti-virus feature is not enabled.
	//
	//   - RANSOMWARE: Anti-ransomware policy is not enabled.
	//
	//   - WEB_LOCK: Web tamper-proofing feature is not enabled.
	//
	//   - VIRUS_SCHEDULE_SCAN: Periodic virus scan policy is not enabled.
	//
	//   - IMAGE_REPO_SCAN: Container image scan scope is not configured.
	//
	//   - IMAGE_SCAN_TASK: One-click container image security risk scan has not been executed.
	//
	// - SS_ALARM: Pending alerts.
	//
	//   - ALARM_SERIOUS: One unhandled high-risk alert event exists.
	//
	//   - ALARM_SUSPICIOUS: One unhandled medium-risk alert event exists.
	//
	//   - ALARM_REMIND: One unhandled low-risk alert event exists.
	//
	// - SS_VUL: Pending vulnerabilities to fix.
	//
	//   - CMS_UNFIX: One unfixed CMS vulnerability exists.
	//
	//   - WIN_UNFIX: One unfixed Windows host vulnerability exists.
	//
	//   - CVE_UNFIX: One unfixed Linux host vulnerability exists.
	//
	//   - ERM_UNFIX: One unfixed emergency vulnerability exists.
	//
	//   - ERM_UNCHECK: One undetected emergency vulnerability exists.
	//
	// - SS_HC: Baseline issues.
	//
	//   - WEAK_EXPLOIT: Weak password risk exposed to the Internet exists.
	//
	//   - WEAK_PASSWORD: Weak password risk exists.
	//
	//   - HC_EXPLOIT: High-risk intrusion vulnerability exists.
	//
	//   - HC_OTHER_WARNING: Security configuration risk exists.
	//
	// - SS_CLOUD_HC: Cloud platform configuration check item issues.
	//
	//   - CSPM_CIEM_NOT_PASS: One failed CIEM check item exists.
	//
	//   - CSPM_RISK_NOT_PASS: One failed security risk check item exists.
	//
	//   - CSPM_COMPLIANCE_NOT_PASS: One failed compliance check item exists.
	//
	// - SS_AK: AccessKey leakage risk exists.
	//
	// example:
	//
	// ALARM_SERIOUS
	SubRuleType *string `json:"SubRuleType,omitempty" xml:"SubRuleType,omitempty"`
	// The description of the sub-rule type for the security score deduction item.
	//
	// example:
	//
	// Unhandled Urgent Alert Event Exists
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) GetScore() *int32 {
	return s.Score
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) GetScoreThreshold() *int32 {
	return s.ScoreThreshold
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) GetSubRuleType() *string {
	return s.SubRuleType
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) GetTitle() *string {
	return s.Title
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) SetScore(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList {
	s.Score = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) SetScoreThreshold(v int32) *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList {
	s.ScoreThreshold = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) SetSubRuleType(v string) *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList {
	s.SubRuleType = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) SetTitle(v string) *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList {
	s.Title = &v
	return s
}

func (s *GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList) Validate() error {
	return dara.Validate(s)
}
