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
	// The status of the custom settings of the security score feature.
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the new version of the security score rule.
	SecurityScoreCategoryList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryList `json:"SecurityScoreCategoryList,omitempty" xml:"SecurityScoreCategoryList,omitempty" type:"Repeated"`
	// The information about the old version of the security score rule.
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
	// example:
	//
	// 30
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction items of the security score rule.
	SecurityRuleList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleList `json:"SecurityRuleList,omitempty" xml:"SecurityRuleList,omitempty" type:"Repeated"`
	// The category of the security score rule.
	//
	// example:
	//
	// Security governance
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
	// The threshold of deduction for the security score rule type.
	//
	// example:
	//
	// 10
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The sub-deduction items of the security score rule.
	SecurityScoreItemList []*GetSecurityScoreRuleResponseBodySecurityScoreCategoryListSecurityRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
	// The name of the security score rule type.
	//
	// example:
	//
	// AI application risks
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
	// The deduction score for the item.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The threshold of the deduction score for the item.
	//
	// example:
	//
	// 10
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The type of the sub-deduction item. Valid values:
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
	// SSI_AI_VUL_RISK
	SubRuleType *string `json:"SubRuleType,omitempty" xml:"SubRuleType,omitempty"`
	// The name of the sub-deduction item of the security score rule.
	//
	// example:
	//
	// Unfixed application vulnerabilities
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
	// The deduction module that is supported by the security score feature. Valid values:
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
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The deduction items of the deduction module.
	SecurityScoreItemList []*GetSecurityScoreRuleResponseBodySecurityScoreRuleListSecurityScoreItemList `json:"SecurityScoreItemList,omitempty" xml:"SecurityScoreItemList,omitempty" type:"Repeated"`
	// The description of the deduction module.
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
	// The penalty point of the deduction item.
	//
	// example:
	//
	// 3
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The threshold for the deduction item.
	//
	// >  Valid values: 0 to the deduction threshold of the deduction module.
	//
	// example:
	//
	// 5
	ScoreThreshold *int32 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The deduction item of the deduction module. The following list describes the deduction modules and their deduction items:
	//
	// 	- SS_REINFORCE: issue in key feature configuration
	//
	//     	- XPRESS_INSTALL: Security Center is not authorized.
	//
	//     	- REINFORCE_SUSPICIOUS: The antivirus feature is disabled.
	//
	//     	- RANSOMWARE: The anti-ransomware policy is disabled.
	//
	//     	- WEB_LOCK: The web tamper proofing feature is disabled.
	//
	//     	- VIRUS_SCHEDULE_SCAN: The periodic virus scan policy is disabled.
	//
	//     	- IMAGE_REPO_SCAN: The range of container image scan is not configured.
	//
	//     	- IMAGE_SCAN_TASK: The feature of one-click scan of container images for security risks is not performed.
	//
	// 	- SS_ALARM: unhandled alert.
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
	//     	- ERM_UNFIX: An unfixed urgent vulnerability is detected.
	//
	//     	- ERM_UNCHECK: An undetected urgent vulnerability exists.
	//
	// 	- SS_HC: baseline risks
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
	// The description of the deduction item in a deduction module.
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
