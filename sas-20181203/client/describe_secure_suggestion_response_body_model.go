// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecureSuggestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalTime(v int64) *DescribeSecureSuggestionResponseBody
	GetCalTime() *int64
	SetRequestId(v string) *DescribeSecureSuggestionResponseBody
	GetRequestId() *string
	SetSuggestions(v []*DescribeSecureSuggestionResponseBodySuggestions) *DescribeSecureSuggestionResponseBody
	GetSuggestions() []*DescribeSecureSuggestionResponseBodySuggestions
	SetTotalCount(v int32) *DescribeSecureSuggestionResponseBody
	GetTotalCount() *int32
}

type DescribeSecureSuggestionResponseBody struct {
	// The timestamp of security score calculation.
	//
	// example:
	//
	// 1755744253000
	CalTime *int64 `json:"CalTime,omitempty" xml:"CalTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 676F80E3-4B3F-43DA-9CBB-5FF79F202AA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The suggestions on how to handle the risks.
	Suggestions []*DescribeSecureSuggestionResponseBodySuggestions `json:"Suggestions,omitempty" xml:"Suggestions,omitempty" type:"Repeated"`
	// The total number of unhandled security risks.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecureSuggestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBody) GetCalTime() *int64 {
	return s.CalTime
}

func (s *DescribeSecureSuggestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecureSuggestionResponseBody) GetSuggestions() []*DescribeSecureSuggestionResponseBodySuggestions {
	return s.Suggestions
}

func (s *DescribeSecureSuggestionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecureSuggestionResponseBody) SetCalTime(v int64) *DescribeSecureSuggestionResponseBody {
	s.CalTime = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetRequestId(v string) *DescribeSecureSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetSuggestions(v []*DescribeSecureSuggestionResponseBodySuggestions) *DescribeSecureSuggestionResponseBody {
	s.Suggestions = v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetTotalCount(v int32) *DescribeSecureSuggestionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) Validate() error {
	if s.Suggestions != nil {
		for _, item := range s.Suggestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecureSuggestionResponseBodySuggestions struct {
	// The details of the suggestion.
	Detail []*DescribeSecureSuggestionResponseBodySuggestionsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The penalty point of a deduction item.
	//
	// example:
	//
	// 40
	Points *int32 `json:"Points,omitempty" xml:"Points,omitempty"`
	// The type of the unhandled risk. Valid values:
	//
	// 	- **SS_REINFORCE**: missing configuration of key features, such as malicious behavior defense
	//
	// 	- **SS_ALARM**: unhandled alerts
	//
	// 	- **SS_VUL**: unfixed vulnerabilities
	//
	// 	- **SS_HC**: baseline risks
	//
	// 	- **SS_AK**: AccessKey pair leaks
	//
	// 	- **SS_CLOUD_HC**: configuration risks of cloud services
	//
	// 	- **OTHER**: others
	//
	// example:
	//
	// SS_ALARM
	SuggestType *string `json:"SuggestType,omitempty" xml:"SuggestType,omitempty"`
}

func (s DescribeSecureSuggestionResponseBodySuggestions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestions) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetDetail() []*DescribeSecureSuggestionResponseBodySuggestionsDetail {
	return s.Detail
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetPoints() *int32 {
	return s.Points
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetSuggestType() *string {
	return s.SuggestType
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetDetail(v []*DescribeSecureSuggestionResponseBodySuggestionsDetail) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Detail = v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetPoints(v int32) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Points = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetSuggestType(v string) *DescribeSecureSuggestionResponseBodySuggestions {
	s.SuggestType = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecureSuggestionResponseBodySuggestionsDetail struct {
	// The description of the suggestion.
	//
	// example:
	//
	// Malicious tampering of Web pages will affect your normal access to web page content, and may also lead to serious economic losses, brand losses, and even political risks. The webpage tamper-proof service can monitor the website directory in real time and restore the tampered files or directories through backup, so as to ensure that the website information of important systems is not tampered with maliciously and prevent the occurrence of horse hanging, black chain, illegal implantation of terrorist threats, pornography and other content.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The sub-type of the unhandled risk. Valid values:
	//
	// 	- **ALARM_HIGH**: Unhandled Urgency Alerts
	//
	// 	- **ALARM_MEDIUM**: Unhandled Warning Alerts
	//
	// 	- **ALARM_LOW**: Unhandled Reminder Alerts
	//
	// 	- **VUL_EMR_UNCHECK**: Unchecked Urgent Vulnerabilities
	//
	// 	- **VUL_EMR_UNFIX**: Unfixed Urgent Vulnerabilities
	//
	// 	- **VUL_WIN**: Unfixed Windows Server Vulnerabilities
	//
	// 	- **VUL_LINUX**: Unfixed Linux Server Vulnerabilities
	//
	// 	- **VUL_CMS**: Unfixed CMS Vulnerabilities
	//
	// 	- **ACCESSKEY_LEAK**: AccessKey Leakage Risks
	//
	// 	- **HC_WARN**: Baseline Risks
	//
	// 	- **HC_WEAK_EXPLOIT_WARN**: There is a risk of weak passwords exposed by the public network.
	//
	// 	- **HC_WEAK_PASSWORD_WARN**: Risk of weak password
	//
	// 	- **HC_HIGH_EXPLOIT_WARN**: There is a high risk of invasion
	//
	// 	- **HC_OTHER_WARN**: Security Configuration risk
	//
	// 	- **HC_DATABASE_WARN**: Database has security risks
	//
	// 	- **CLOUD_HC_SAS_OPEN**: Security protection has not been installed on the server
	//
	// 	- **CLOUD_HC_AEGIS_OFFLINE**: Server protection status is offline
	//
	// 	- **CLOUD_HC_ACCOUNT_DOUBLE_CHECK**: Two-Factor Authentication not Enabled for Primary Account
	//
	// 	- **CLOUD_HC_RDS**: RDS-database security policy failed, security risks
	//
	// 	- **CLOUD_HC_DDOS**: Risks in Anti-DDoS Pro Back-to-Origin Settings
	//
	// 	- **CLOUD_HC_HIGH_LEVEL**: Cloud product configuration has high risk
	//
	// 	- **CLOUD_HC_OTHER_LEVEL**: Cloud product configuration has medium and low risk risks
	//
	// 	- **OTHER_ATTACH**: Attacks
	//
	// 	- **OTHER_DATABASE_ATTACH**: Database has security risks
	//
	// 	- **REINFORCE_BASELINE**: Config Assessment
	//
	// 	- **REINFORCE_SUSPICIOUS**: Antivirus
	//
	// 	- **REINFORCE_ANALYSIS**: Log Analysis
	//
	// 	- **REINFORCE_AK_LEAK**: AccessKey Leaked Intelligence Detection
	//
	// 	- **REINFORCE_WEB_LOCK**: Website tamper-proofing capability not configured
	//
	// 	- **REINFORCE_BRUTE_FORCE**: Anti brute force cracking
	//
	// 	- **REINFORCE_XPRESS_INSTALL**: One-click client installation
	//
	// 	- **REINFORCE_RANSOMWARE**: Enable anti-extortion strategy
	//
	// 	- **REINFORCE_UNI_RANSOMWARE**: Anti-ransomware for Databases
	//
	// 	- **REINFORCE_VIRUS_SCHEDULE_SCAN**: Periodic virus scan policies not configured
	//
	// 	- **REINFORCE_IMAGE_REPO_SCAN**: No container image scan range configured
	//
	// 	- **REINFORCE_IMAGE_SCAN_TASK**: Image security scan
	//
	// 	- **REINFORCE_K8S_LOG_ANALYSIS**: Container K8s threat detection is disabled
	//
	// 	- **REINFORCE_CONTAINER_NETWORK**: Container Visualization
	//
	// example:
	//
	// REINFORCE_WEB_LOCK
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The name of the unhandled risk.
	//
	// example:
	//
	// Website tamper-proofing capability not configured
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetSubType() *string {
	return s.SubType
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetTitle() *string {
	return s.Title
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetDescription(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Description = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetSubType(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.SubType = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetTitle(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Title = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) Validate() error {
	return dara.Validate(s)
}
