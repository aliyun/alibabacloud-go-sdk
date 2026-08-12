// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListRiskItemsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListRiskItemsRequest
	GetPageSize() *int32
	SetPolicyName(v string) *ListRiskItemsRequest
	GetPolicyName() *string
	SetRiskCategory(v string) *ListRiskItemsRequest
	GetRiskCategory() *string
	SetRiskId(v string) *ListRiskItemsRequest
	GetRiskId() *string
	SetRiskLevel(v string) *ListRiskItemsRequest
	GetRiskLevel() *string
	SetRiskScene(v string) *ListRiskItemsRequest
	GetRiskScene() *string
	SetStatus(v string) *ListRiskItemsRequest
	GetStatus() *string
	SetStatusList(v []*string) *ListRiskItemsRequest
	GetStatusList() []*string
	SetUsername(v string) *ListRiskItemsRequest
	GetUsername() *string
}

type ListRiskItemsRequest struct {
	// The current page number in a paging query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page in a paging query. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the risk analysis policy. Fuzzy matching is supported.
	//
	// example:
	//
	// Remote logon risk analysis policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The risk category. Valid values:
	//
	// 	- `data_safe`: data security.
	//
	// 	- `identify_safe`: identity security.
	//
	// 	- `device_safe`: device security.
	//
	// 	- `access_safe`: access security.
	//
	// 	- `ai_agent_safe`: Agent security.
	//
	// example:
	//
	// identify_safe
	RiskCategory *string `json:"RiskCategory,omitempty" xml:"RiskCategory,omitempty"`
	// The risk event ID. If specified, the system performs an exact query for the specified risk event.
	//
	// example:
	//
	// 69ef648034cf53d7bac7a9c9c912****
	RiskId *string `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The risk level. Valid values:
	//
	// 	- `High`: high risk.
	//
	// 	- `Medium`: medium risk.
	//
	// 	- `Low`: low risk.
	//
	// example:
	//
	// High
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The risk scenario. Valid values:
	//
	// 	- `account_share`: account sharing.
	//
	// 	- `account_stolen`: account theft.
	//
	// 	- `device_share`: device sharing.
	//
	// 	- `remote_logon`: remote logon.
	//
	// 	- `sensitive_data_leakage`: sensitive data exfiltration.
	//
	// 	- `lateral_scanning`: lateral scanning.
	//
	// 	- `ai_skill_malware`: malicious Skill.
	//
	// 	- `ai_config_check`: AI configuration check.
	//
	// 	- `openclaw_vulnerability`: OpenClaw vulnerability.
	//
	// example:
	//
	// account_stolen
	RiskScene *string `json:"RiskScene,omitempty" xml:"RiskScene,omitempty"`
	// The disposition status of the risk event. This parameter cannot be used together with `StatusList`.
	//
	// example:
	//
	// Unprocess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of disposition statuses of risk events, in Flat serialization format. This parameter cannot be used together with Status.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The username associated with the risk event. Fuzzy matching is supported. Maximum length: 128 characters.
	//
	// example:
	//
	// zhang***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListRiskItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRiskItemsRequest) GoString() string {
	return s.String()
}

func (s *ListRiskItemsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListRiskItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRiskItemsRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListRiskItemsRequest) GetRiskCategory() *string {
	return s.RiskCategory
}

func (s *ListRiskItemsRequest) GetRiskId() *string {
	return s.RiskId
}

func (s *ListRiskItemsRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListRiskItemsRequest) GetRiskScene() *string {
	return s.RiskScene
}

func (s *ListRiskItemsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRiskItemsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListRiskItemsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListRiskItemsRequest) SetCurrentPage(v int32) *ListRiskItemsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListRiskItemsRequest) SetPageSize(v int32) *ListRiskItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRiskItemsRequest) SetPolicyName(v string) *ListRiskItemsRequest {
	s.PolicyName = &v
	return s
}

func (s *ListRiskItemsRequest) SetRiskCategory(v string) *ListRiskItemsRequest {
	s.RiskCategory = &v
	return s
}

func (s *ListRiskItemsRequest) SetRiskId(v string) *ListRiskItemsRequest {
	s.RiskId = &v
	return s
}

func (s *ListRiskItemsRequest) SetRiskLevel(v string) *ListRiskItemsRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListRiskItemsRequest) SetRiskScene(v string) *ListRiskItemsRequest {
	s.RiskScene = &v
	return s
}

func (s *ListRiskItemsRequest) SetStatus(v string) *ListRiskItemsRequest {
	s.Status = &v
	return s
}

func (s *ListRiskItemsRequest) SetStatusList(v []*string) *ListRiskItemsRequest {
	s.StatusList = v
	return s
}

func (s *ListRiskItemsRequest) SetUsername(v string) *ListRiskItemsRequest {
	s.Username = &v
	return s
}

func (s *ListRiskItemsRequest) Validate() error {
	return dara.Validate(s)
}
