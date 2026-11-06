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
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If it is not provided, the service returns a 400 error "The specified Lang can not be empty.". Explicitly pass the language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of data records to return for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which to start reading. If you leave this parameter empty, the reading starts from the beginning.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center for threat analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action type of the automatic response rule. Valid values:
	//
	// - doPlaybook: Execute a playbook.
	//
	// - changeEventStatus: Update the event status.
	//
	// - changeThreatLevel: Update the event threat level.
	//
	// - addEventTag: Add an event tag.
	//
	// - deleteEventTag: Delete an event tag.
	//
	// - alertWhitelist: Add an alert to the whitelist.
	//
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The name of the automatic response rule.
	//
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// The status of the automatic response rule. Valid values:
	//
	// - 0: Disabled.
	//
	// - 100: Enabled.
	//
	// example:
	//
	// 0
	ResponseRuleStatus *int32 `json:"ResponseRuleStatus,omitempty" xml:"ResponseRuleStatus,omitempty"`
	// The type of the response rule. Valid values:
	//
	// - preset: Preset.
	//
	// - custom: Custom.
	//
	// example:
	//
	// custom
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The trigger type of the automatic response rule. Valid values:
	//
	// - event: Event occurrence.
	//
	// - event_update: Event update.
	//
	// - alert: Alert occurrence.
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
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
