// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListResponseRulesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListResponseRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResponseRulesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListResponseRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResponseRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListResponseRulesRequest
	GetRegionId() *string
	SetResponseActionType(v string) *ListResponseRulesRequest
	GetResponseActionType() *string
	SetResponseRuleName(v string) *ListResponseRulesRequest
	GetResponseRuleName() *string
	SetResponseRuleStatus(v int32) *ListResponseRulesRequest
	GetResponseRuleStatus() *int32
	SetResponseRuleType(v string) *ListResponseRulesRequest
	GetResponseRuleType() *string
	SetResponseTriggerType(v string) *ListResponseRulesRequest
	GetResponseTriggerType() *string
	SetRoleFor(v int64) *ListResponseRulesRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListResponseRulesRequest
	GetRoleType() *int32
}

type ListResponseRulesRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. If you leave this parameter empty, the first page of results is returned.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the data management center for threat analysis. Select the region where your assets are located. Valid values:
	//
	// - `cn-hangzhou`: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - `ap-southeast-1`: Your assets are in international regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action of the automated response rule. Valid values:
	//
	// - `doPlaybook`: Executes a playbook.
	//
	// - `changeEventStatus`: Updates the status of an event.
	//
	// - `changeThreatLevel`: Updates the threat level of an event.
	//
	// - `addEventTag`: Adds a tag to an event.
	//
	// - `deleteEventTag`: Removes a tag from an event.
	//
	// - `alertWhitelist`: Adds an alert to the allowlist.
	//
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// The status of the automated response rule. Valid values:
	//
	// - `0`: disabled
	//
	// - `100`: enabled
	//
	// example:
	//
	// 0
	ResponseRuleStatus *int32 `json:"ResponseRuleStatus,omitempty" xml:"ResponseRuleStatus,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// - `preset`: A preset rule.
	//
	// - `custom`: A custom rule.
	//
	// example:
	//
	// custom
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The trigger type of the automated response rule. Valid values:
	//
	// - `event`: An event is generated.
	//
	// - `event_update`: An event is updated.
	//
	// - `alert`: An alert is generated.
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// The ID of a member. An administrator can use this parameter to view data as the specified member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - `0`: Displays data from the current Alibaba Cloud account.
	//
	// - `1`: Displays data from all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListResponseRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRulesRequest) GoString() string {
	return s.String()
}

func (s *ListResponseRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListResponseRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResponseRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResponseRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResponseRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResponseRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResponseRulesRequest) GetResponseActionType() *string {
	return s.ResponseActionType
}

func (s *ListResponseRulesRequest) GetResponseRuleName() *string {
	return s.ResponseRuleName
}

func (s *ListResponseRulesRequest) GetResponseRuleStatus() *int32 {
	return s.ResponseRuleStatus
}

func (s *ListResponseRulesRequest) GetResponseRuleType() *string {
	return s.ResponseRuleType
}

func (s *ListResponseRulesRequest) GetResponseTriggerType() *string {
	return s.ResponseTriggerType
}

func (s *ListResponseRulesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListResponseRulesRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListResponseRulesRequest) SetLang(v string) *ListResponseRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListResponseRulesRequest) SetMaxResults(v int32) *ListResponseRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResponseRulesRequest) SetNextToken(v string) *ListResponseRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResponseRulesRequest) SetPageNumber(v int32) *ListResponseRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResponseRulesRequest) SetPageSize(v int32) *ListResponseRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResponseRulesRequest) SetRegionId(v string) *ListResponseRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListResponseRulesRequest) SetResponseActionType(v string) *ListResponseRulesRequest {
	s.ResponseActionType = &v
	return s
}

func (s *ListResponseRulesRequest) SetResponseRuleName(v string) *ListResponseRulesRequest {
	s.ResponseRuleName = &v
	return s
}

func (s *ListResponseRulesRequest) SetResponseRuleStatus(v int32) *ListResponseRulesRequest {
	s.ResponseRuleStatus = &v
	return s
}

func (s *ListResponseRulesRequest) SetResponseRuleType(v string) *ListResponseRulesRequest {
	s.ResponseRuleType = &v
	return s
}

func (s *ListResponseRulesRequest) SetResponseTriggerType(v string) *ListResponseRulesRequest {
	s.ResponseTriggerType = &v
	return s
}

func (s *ListResponseRulesRequest) SetRoleFor(v int64) *ListResponseRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListResponseRulesRequest) SetRoleType(v int32) *ListResponseRulesRequest {
	s.RoleType = &v
	return s
}

func (s *ListResponseRulesRequest) Validate() error {
	return dara.Validate(s)
}
