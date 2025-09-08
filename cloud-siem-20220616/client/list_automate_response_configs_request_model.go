// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutomateResponseConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *ListAutomateResponseConfigsRequest
	GetActionType() *string
	SetAutoResponseType(v string) *ListAutomateResponseConfigsRequest
	GetAutoResponseType() *string
	SetCurrentPage(v int32) *ListAutomateResponseConfigsRequest
	GetCurrentPage() *int32
	SetId(v int64) *ListAutomateResponseConfigsRequest
	GetId() *int64
	SetPageSize(v int32) *ListAutomateResponseConfigsRequest
	GetPageSize() *int32
	SetPlaybookUuid(v string) *ListAutomateResponseConfigsRequest
	GetPlaybookUuid() *string
	SetRegionId(v string) *ListAutomateResponseConfigsRequest
	GetRegionId() *string
	SetResponseRuleType(v string) *ListAutomateResponseConfigsRequest
	GetResponseRuleType() *string
	SetRoleFor(v int64) *ListAutomateResponseConfigsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAutomateResponseConfigsRequest
	GetRoleType() *int32
	SetRuleName(v string) *ListAutomateResponseConfigsRequest
	GetRuleName() *string
	SetStatus(v int32) *ListAutomateResponseConfigsRequest
	GetStatus() *int32
	SetSubUserId(v int64) *ListAutomateResponseConfigsRequest
	GetSubUserId() *int64
}

type ListAutomateResponseConfigsRequest struct {
	// The type of the handling action. Valid values:
	//
	// 	- doPlaybook: runs a playbook.
	//
	// 	- changeEventStatus: changes the status of an event.
	//
	// 	- changeThreatLevel: changes the risk level of an event.
	//
	// example:
	//
	// doPlaybook
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_aegis_kill_quara_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 100: enabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsRequest) GetActionType() *string {
	return s.ActionType
}

func (s *ListAutomateResponseConfigsRequest) GetAutoResponseType() *string {
	return s.AutoResponseType
}

func (s *ListAutomateResponseConfigsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAutomateResponseConfigsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListAutomateResponseConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutomateResponseConfigsRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListAutomateResponseConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAutomateResponseConfigsRequest) GetResponseRuleType() *string {
	return s.ResponseRuleType
}

func (s *ListAutomateResponseConfigsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAutomateResponseConfigsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAutomateResponseConfigsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutomateResponseConfigsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAutomateResponseConfigsRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListAutomateResponseConfigsRequest) SetActionType(v string) *ListAutomateResponseConfigsRequest {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetAutoResponseType(v string) *ListAutomateResponseConfigsRequest {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetCurrentPage(v int32) *ListAutomateResponseConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetId(v int64) *ListAutomateResponseConfigsRequest {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPageSize(v int32) *ListAutomateResponseConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPlaybookUuid(v string) *ListAutomateResponseConfigsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRegionId(v string) *ListAutomateResponseConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetResponseRuleType(v string) *ListAutomateResponseConfigsRequest {
	s.ResponseRuleType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRoleFor(v int64) *ListAutomateResponseConfigsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRoleType(v int32) *ListAutomateResponseConfigsRequest {
	s.RoleType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRuleName(v string) *ListAutomateResponseConfigsRequest {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetStatus(v int32) *ListAutomateResponseConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetSubUserId(v int64) *ListAutomateResponseConfigsRequest {
	s.SubUserId = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) Validate() error {
	return dara.Validate(s)
}
