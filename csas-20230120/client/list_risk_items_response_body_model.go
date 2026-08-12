// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRiskItemsResponseBody
	GetRequestId() *string
	SetRiskItems(v []*ListRiskItemsResponseBodyRiskItems) *ListRiskItemsResponseBody
	GetRiskItems() []*ListRiskItemsResponseBodyRiskItems
	SetTotalNum(v int32) *ListRiskItemsResponseBody
	GetTotalNum() *int32
}

type ListRiskItemsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of risk events.
	RiskItems []*ListRiskItemsResponseBodyRiskItems `json:"RiskItems,omitempty" xml:"RiskItems,omitempty" type:"Repeated"`
	// The total number of risk events that match the query conditions.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListRiskItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRiskItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRiskItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRiskItemsResponseBody) GetRiskItems() []*ListRiskItemsResponseBodyRiskItems {
	return s.RiskItems
}

func (s *ListRiskItemsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListRiskItemsResponseBody) SetRequestId(v string) *ListRiskItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRiskItemsResponseBody) SetRiskItems(v []*ListRiskItemsResponseBodyRiskItems) *ListRiskItemsResponseBody {
	s.RiskItems = v
	return s
}

func (s *ListRiskItemsResponseBody) SetTotalNum(v int32) *ListRiskItemsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListRiskItemsResponseBody) Validate() error {
	if s.RiskItems != nil {
		for _, item := range s.RiskItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRiskItemsResponseBodyRiskItems struct {
	// The name of the Agent that generated the risk event. An empty string is returned for non-Agent risk scenarios.
	//
	// example:
	//
	// qoder****
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The AI risk analysis conclusion.
	//
	// example:
	//
	// The account logged on from an unusual device, and the behavior is inconsistent with the historical baseline
	AiConclusion *string `json:"AiConclusion,omitempty" xml:"AiConclusion,omitempty"`
	// The risk judgment provided by AI. An empty string is returned if no AI analysis result exists. Valid values:
	//
	// 	- `Risk`: determined as risky.
	//
	// 	- `Ignore`: determined as not risky.
	//
	// example:
	//
	// Risk
	AiRiskConfirm *string `json:"AiRiskConfirm,omitempty" xml:"AiRiskConfirm,omitempty"`
	// The name of the risk detection item.
	//
	// example:
	//
	// Unusual device logon check
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The department to which the user associated with the risk event belongs.
	//
	// example:
	//
	// Department****
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The original organizational structure information of the user associated with the risk event.
	//
	// example:
	//
	// CN=zhang***,OU=Department****
	GroupInfo *string `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
	// The name of the endpoint device associated with the risk event.
	//
	// example:
	//
	// U-2GW2L4M7-****
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The internal IP address of the endpoint associated with the risk event.
	//
	// example:
	//
	// 192.168.XX.XX
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The risk detection report or risk evidence.
	//
	// example:
	//
	// The account logged on from a new device, and the logon location is inconsistent with the usual location
	Report *string `json:"Report,omitempty" xml:"Report,omitempty"`
	// The list of risk analysis policy names that were hit.
	RiskAnalysisPolicyNames []*string `json:"RiskAnalysisPolicyNames,omitempty" xml:"RiskAnalysisPolicyNames,omitempty" type:"Repeated"`
	// The risk category. Valid values:
	//
	// - `data_safe`: data security.
	//
	// - `identify_safe`: identity security.
	//
	// - `device_safe`: device security.
	//
	// - `access_safe`: access security.
	//
	// - `ai_agent_safe`: Agent security.
	//
	// example:
	//
	// identify_safe
	RiskCategory *string `json:"RiskCategory,omitempty" xml:"RiskCategory,omitempty"`
	// The manually confirmed risk conclusion. An empty string is returned if the event has not been confirmed. Valid values:
	//
	// 	- `Risk`: confirmed as risky.
	//
	// 	- `Ignore`: confirmed as not risky.
	//
	// 	- `Invalid`: confirmed as a false positive.
	//
	// example:
	//
	// Risk
	RiskConfirm *string `json:"RiskConfirm,omitempty" xml:"RiskConfirm,omitempty"`
	// The description of the risk event disposition.
	//
	// example:
	//
	// Upon investigation, the logon was not authorized by the user
	RiskConfirmDesc *string `json:"RiskConfirmDesc,omitempty" xml:"RiskConfirmDesc,omitempty"`
	// The risk description.
	//
	// example:
	//
	// The account logged on from an unusual device
	RiskDesc *string `json:"RiskDesc,omitempty" xml:"RiskDesc,omitempty"`
	// The end time of the risky behavior, in the format of `yyyy-MM-dd HH:mm:ss`.
	//
	// example:
	//
	// 2026-05-21 05:21:00
	RiskEndTime *string `json:"RiskEndTime,omitempty" xml:"RiskEndTime,omitempty"`
	// The list of detection feature or detection item identifiers that triggered the risk event. A risk event may hit multiple identifiers. The specific values vary based on the risk scenario and detection rules.
	RiskFeatureIds []*string `json:"RiskFeatureIds,omitempty" xml:"RiskFeatureIds,omitempty" type:"Repeated"`
	// The time when the risk was detected, in the format of `yyyy-MM-dd HH:mm:ss`.
	//
	// example:
	//
	// 2026-05-20 10:30:00
	RiskFoundTime *string `json:"RiskFoundTime,omitempty" xml:"RiskFoundTime,omitempty"`
	// The risk event ID.
	//
	// example:
	//
	// 69ef648034cf53d7bac7a9c9c912****
	RiskId *string `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The risk level. Valid values:
	//
	// - `High`: high risk.
	//
	// - `Medium`: medium risk.
	//
	// - `Low`: low risk.
	//
	// example:
	//
	// High
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The risk scenario. Valid values:
	//
	// - `account_share`: account sharing.
	//
	// - `account_stolen`: account theft.
	//
	// - `device_share`: device sharing.
	//
	// - `remote_logon`: remote logon.
	//
	// - `sensitive_data_leakage`: sensitive data exfiltration.
	//
	// - `lateral_scanning`: lateral scanning.
	//
	// - `ai_skill_malware`: malicious Skill.
	//
	// - `ai_config_check`: AI configuration check.
	//
	// - `openclaw_vulnerability`: OpenClaw vulnerability.
	//
	// example:
	//
	// account_stolen
	RiskScene *string `json:"RiskScene,omitempty" xml:"RiskScene,omitempty"`
	// The start time of the risky behavior, in the format of `yyyy-MM-dd HH:mm:ss`.
	//
	// example:
	//
	// 2026-05-20 05:20:00
	RiskStartTime *string `json:"RiskStartTime,omitempty" xml:"RiskStartTime,omitempty"`
	// The SASE user ID associated with the risk event.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The name of the Agent Skill that generated the risk event. An empty string is returned for non-Agent risk scenarios.
	//
	// example:
	//
	// frontend-design
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The recommended remediation action for the risk event.
	//
	// example:
	//
	// Verify the account user and freeze the account or reset credentials based on the investigation results
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The disposition status of the risk event. Valid values:
	//
	// 	- `Unprocess`: unprocessed.
	//
	// 	- `Processing`: being processed.
	//
	// 	- `Processed`: processed.
	//
	// example:
	//
	// Unprocess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether AI risk analysis results exist. Valid values:
	//
	// 	- `true`: AI risk analysis results exist.
	//
	// 	- `false`: AI risk analysis results do not exist.
	//
	// example:
	//
	// true
	SupportAnalysis *bool `json:"SupportAnalysis,omitempty" xml:"SupportAnalysis,omitempty"`
	// The username associated with the risk event.
	//
	// example:
	//
	// zhang***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListRiskItemsResponseBodyRiskItems) String() string {
	return dara.Prettify(s)
}

func (s ListRiskItemsResponseBodyRiskItems) GoString() string {
	return s.String()
}

func (s *ListRiskItemsResponseBodyRiskItems) GetAgentName() *string {
	return s.AgentName
}

func (s *ListRiskItemsResponseBodyRiskItems) GetAiConclusion() *string {
	return s.AiConclusion
}

func (s *ListRiskItemsResponseBodyRiskItems) GetAiRiskConfirm() *string {
	return s.AiRiskConfirm
}

func (s *ListRiskItemsResponseBodyRiskItems) GetCheckName() *string {
	return s.CheckName
}

func (s *ListRiskItemsResponseBodyRiskItems) GetDepartment() *string {
	return s.Department
}

func (s *ListRiskItemsResponseBodyRiskItems) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ListRiskItemsResponseBodyRiskItems) GetHostname() *string {
	return s.Hostname
}

func (s *ListRiskItemsResponseBodyRiskItems) GetInnerIp() *string {
	return s.InnerIp
}

func (s *ListRiskItemsResponseBodyRiskItems) GetReport() *string {
	return s.Report
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskAnalysisPolicyNames() []*string {
	return s.RiskAnalysisPolicyNames
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskCategory() *string {
	return s.RiskCategory
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskConfirm() *string {
	return s.RiskConfirm
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskConfirmDesc() *string {
	return s.RiskConfirmDesc
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskDesc() *string {
	return s.RiskDesc
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskEndTime() *string {
	return s.RiskEndTime
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskFeatureIds() []*string {
	return s.RiskFeatureIds
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskFoundTime() *string {
	return s.RiskFoundTime
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskId() *string {
	return s.RiskId
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskScene() *string {
	return s.RiskScene
}

func (s *ListRiskItemsResponseBodyRiskItems) GetRiskStartTime() *string {
	return s.RiskStartTime
}

func (s *ListRiskItemsResponseBodyRiskItems) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListRiskItemsResponseBodyRiskItems) GetSkillName() *string {
	return s.SkillName
}

func (s *ListRiskItemsResponseBodyRiskItems) GetSolution() *string {
	return s.Solution
}

func (s *ListRiskItemsResponseBodyRiskItems) GetStatus() *string {
	return s.Status
}

func (s *ListRiskItemsResponseBodyRiskItems) GetSupportAnalysis() *bool {
	return s.SupportAnalysis
}

func (s *ListRiskItemsResponseBodyRiskItems) GetUsername() *string {
	return s.Username
}

func (s *ListRiskItemsResponseBodyRiskItems) SetAgentName(v string) *ListRiskItemsResponseBodyRiskItems {
	s.AgentName = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetAiConclusion(v string) *ListRiskItemsResponseBodyRiskItems {
	s.AiConclusion = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetAiRiskConfirm(v string) *ListRiskItemsResponseBodyRiskItems {
	s.AiRiskConfirm = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetCheckName(v string) *ListRiskItemsResponseBodyRiskItems {
	s.CheckName = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetDepartment(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Department = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetGroupInfo(v string) *ListRiskItemsResponseBodyRiskItems {
	s.GroupInfo = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetHostname(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Hostname = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetInnerIp(v string) *ListRiskItemsResponseBodyRiskItems {
	s.InnerIp = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetReport(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Report = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskAnalysisPolicyNames(v []*string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskAnalysisPolicyNames = v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskCategory(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskCategory = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskConfirm(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskConfirm = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskConfirmDesc(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskConfirmDesc = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskDesc(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskDesc = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskEndTime(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskEndTime = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskFeatureIds(v []*string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskFeatureIds = v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskFoundTime(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskFoundTime = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskId(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskId = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskLevel(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskLevel = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskScene(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskScene = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetRiskStartTime(v string) *ListRiskItemsResponseBodyRiskItems {
	s.RiskStartTime = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetSaseUserId(v string) *ListRiskItemsResponseBodyRiskItems {
	s.SaseUserId = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetSkillName(v string) *ListRiskItemsResponseBodyRiskItems {
	s.SkillName = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetSolution(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Solution = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetStatus(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Status = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetSupportAnalysis(v bool) *ListRiskItemsResponseBodyRiskItems {
	s.SupportAnalysis = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) SetUsername(v string) *ListRiskItemsResponseBodyRiskItems {
	s.Username = &v
	return s
}

func (s *ListRiskItemsResponseBodyRiskItems) Validate() error {
	return dara.Validate(s)
}
