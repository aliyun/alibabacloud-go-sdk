// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListRulesV4Request
	GetBaseMeAgentId() *int64
	SetBusinessName(v string) *ListRulesV4Request
	GetBusinessName() *string
	SetBusinessRange(v int32) *ListRulesV4Request
	GetBusinessRange() *int32
	SetCategoryName(v string) *ListRulesV4Request
	GetCategoryName() *string
	SetCountTotal(v bool) *ListRulesV4Request
	GetCountTotal() *bool
	SetCreateEmpid(v string) *ListRulesV4Request
	GetCreateEmpid() *string
	SetCreateUserId(v int64) *ListRulesV4Request
	GetCreateUserId() *int64
	SetCurrentPage(v int32) *ListRulesV4Request
	GetCurrentPage() *int32
	SetEndTime(v string) *ListRulesV4Request
	GetEndTime() *string
	SetLastUpdateEmpid(v string) *ListRulesV4Request
	GetLastUpdateEmpid() *string
	SetPageNumber(v int32) *ListRulesV4Request
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRulesV4Request
	GetPageSize() *int32
	SetRequireInfos(v []*string) *ListRulesV4Request
	GetRequireInfos() []*string
	SetRid(v int64) *ListRulesV4Request
	GetRid() *int64
	SetRuleIdOrRuleName(v string) *ListRulesV4Request
	GetRuleIdOrRuleName() *string
	SetRuleScoreSingleType(v int32) *ListRulesV4Request
	GetRuleScoreSingleType() *int32
	SetRuleType(v int32) *ListRulesV4Request
	GetRuleType() *int32
	SetSchemeId(v int64) *ListRulesV4Request
	GetSchemeId() *int64
	SetSourceType(v int32) *ListRulesV4Request
	GetSourceType() *int32
	SetStartTime(v string) *ListRulesV4Request
	GetStartTime() *string
	SetStatus(v int32) *ListRulesV4Request
	GetStatus() *int32
	SetType(v int32) *ListRulesV4Request
	GetType() *int32
	SetTypeName(v string) *ListRulesV4Request
	GetTypeName() *string
	SetUpdateEndTime(v string) *ListRulesV4Request
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *ListRulesV4Request
	GetUpdateStartTime() *string
	SetUpdateUserId(v int64) *ListRulesV4Request
	GetUpdateUserId() *int64
}

type ListRulesV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	BusinessName  *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	BusinessRange *int32  `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// false
	CountTotal *bool `json:"CountTotal,omitempty" xml:"CountTotal,omitempty"`
	// example:
	//
	// 1
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 1
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2021-11-29 19:11:09
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequireInfos []*string `json:"RequireInfos,omitempty" xml:"RequireInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 895EAD5312634F5AA708E3B3FA79662E
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// xx
	RuleIdOrRuleName *string `json:"RuleIdOrRuleName,omitempty" xml:"RuleIdOrRuleName,omitempty"`
	// example:
	//
	// 1
	RuleScoreSingleType *int32 `json:"RuleScoreSingleType,omitempty" xml:"RuleScoreSingleType,omitempty"`
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 1000000090
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 2021-11-29 18:11:09
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// example:
	//
	// 2021-11-29 18:11:09
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" xml:"UpdateEndTime,omitempty"`
	// example:
	//
	// 2021-11-29 16:11:09
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" xml:"UpdateStartTime,omitempty"`
	// example:
	//
	// 1
	UpdateUserId *int64 `json:"UpdateUserId,omitempty" xml:"UpdateUserId,omitempty"`
}

func (s ListRulesV4Request) String() string {
	return dara.Prettify(s)
}

func (s ListRulesV4Request) GoString() string {
	return s.String()
}

func (s *ListRulesV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListRulesV4Request) GetBusinessName() *string {
	return s.BusinessName
}

func (s *ListRulesV4Request) GetBusinessRange() *int32 {
	return s.BusinessRange
}

func (s *ListRulesV4Request) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListRulesV4Request) GetCountTotal() *bool {
	return s.CountTotal
}

func (s *ListRulesV4Request) GetCreateEmpid() *string {
	return s.CreateEmpid
}

func (s *ListRulesV4Request) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListRulesV4Request) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListRulesV4Request) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRulesV4Request) GetLastUpdateEmpid() *string {
	return s.LastUpdateEmpid
}

func (s *ListRulesV4Request) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRulesV4Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRulesV4Request) GetRequireInfos() []*string {
	return s.RequireInfos
}

func (s *ListRulesV4Request) GetRid() *int64 {
	return s.Rid
}

func (s *ListRulesV4Request) GetRuleIdOrRuleName() *string {
	return s.RuleIdOrRuleName
}

func (s *ListRulesV4Request) GetRuleScoreSingleType() *int32 {
	return s.RuleScoreSingleType
}

func (s *ListRulesV4Request) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListRulesV4Request) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *ListRulesV4Request) GetSourceType() *int32 {
	return s.SourceType
}

func (s *ListRulesV4Request) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRulesV4Request) GetStatus() *int32 {
	return s.Status
}

func (s *ListRulesV4Request) GetType() *int32 {
	return s.Type
}

func (s *ListRulesV4Request) GetTypeName() *string {
	return s.TypeName
}

func (s *ListRulesV4Request) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *ListRulesV4Request) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *ListRulesV4Request) GetUpdateUserId() *int64 {
	return s.UpdateUserId
}

func (s *ListRulesV4Request) SetBaseMeAgentId(v int64) *ListRulesV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListRulesV4Request) SetBusinessName(v string) *ListRulesV4Request {
	s.BusinessName = &v
	return s
}

func (s *ListRulesV4Request) SetBusinessRange(v int32) *ListRulesV4Request {
	s.BusinessRange = &v
	return s
}

func (s *ListRulesV4Request) SetCategoryName(v string) *ListRulesV4Request {
	s.CategoryName = &v
	return s
}

func (s *ListRulesV4Request) SetCountTotal(v bool) *ListRulesV4Request {
	s.CountTotal = &v
	return s
}

func (s *ListRulesV4Request) SetCreateEmpid(v string) *ListRulesV4Request {
	s.CreateEmpid = &v
	return s
}

func (s *ListRulesV4Request) SetCreateUserId(v int64) *ListRulesV4Request {
	s.CreateUserId = &v
	return s
}

func (s *ListRulesV4Request) SetCurrentPage(v int32) *ListRulesV4Request {
	s.CurrentPage = &v
	return s
}

func (s *ListRulesV4Request) SetEndTime(v string) *ListRulesV4Request {
	s.EndTime = &v
	return s
}

func (s *ListRulesV4Request) SetLastUpdateEmpid(v string) *ListRulesV4Request {
	s.LastUpdateEmpid = &v
	return s
}

func (s *ListRulesV4Request) SetPageNumber(v int32) *ListRulesV4Request {
	s.PageNumber = &v
	return s
}

func (s *ListRulesV4Request) SetPageSize(v int32) *ListRulesV4Request {
	s.PageSize = &v
	return s
}

func (s *ListRulesV4Request) SetRequireInfos(v []*string) *ListRulesV4Request {
	s.RequireInfos = v
	return s
}

func (s *ListRulesV4Request) SetRid(v int64) *ListRulesV4Request {
	s.Rid = &v
	return s
}

func (s *ListRulesV4Request) SetRuleIdOrRuleName(v string) *ListRulesV4Request {
	s.RuleIdOrRuleName = &v
	return s
}

func (s *ListRulesV4Request) SetRuleScoreSingleType(v int32) *ListRulesV4Request {
	s.RuleScoreSingleType = &v
	return s
}

func (s *ListRulesV4Request) SetRuleType(v int32) *ListRulesV4Request {
	s.RuleType = &v
	return s
}

func (s *ListRulesV4Request) SetSchemeId(v int64) *ListRulesV4Request {
	s.SchemeId = &v
	return s
}

func (s *ListRulesV4Request) SetSourceType(v int32) *ListRulesV4Request {
	s.SourceType = &v
	return s
}

func (s *ListRulesV4Request) SetStartTime(v string) *ListRulesV4Request {
	s.StartTime = &v
	return s
}

func (s *ListRulesV4Request) SetStatus(v int32) *ListRulesV4Request {
	s.Status = &v
	return s
}

func (s *ListRulesV4Request) SetType(v int32) *ListRulesV4Request {
	s.Type = &v
	return s
}

func (s *ListRulesV4Request) SetTypeName(v string) *ListRulesV4Request {
	s.TypeName = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateEndTime(v string) *ListRulesV4Request {
	s.UpdateEndTime = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateStartTime(v string) *ListRulesV4Request {
	s.UpdateStartTime = &v
	return s
}

func (s *ListRulesV4Request) SetUpdateUserId(v int64) *ListRulesV4Request {
	s.UpdateUserId = &v
	return s
}

func (s *ListRulesV4Request) Validate() error {
	return dara.Validate(s)
}
