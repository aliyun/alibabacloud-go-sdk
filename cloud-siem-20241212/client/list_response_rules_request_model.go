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
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// example:
	//
	// 0
	ResponseRuleStatus *int32 `json:"ResponseRuleStatus,omitempty" xml:"ResponseRuleStatus,omitempty"`
	// example:
	//
	// custom
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
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
