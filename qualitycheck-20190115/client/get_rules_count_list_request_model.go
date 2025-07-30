// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRulesCountListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetRulesCountListRequest
	GetBaseMeAgentId() *int64
	SetBusinessName(v string) *GetRulesCountListRequest
	GetBusinessName() *string
	SetBusinessRange(v int32) *GetRulesCountListRequest
	GetBusinessRange() *int32
	SetCategoryName(v string) *GetRulesCountListRequest
	GetCategoryName() *string
	SetCountTotal(v bool) *GetRulesCountListRequest
	GetCountTotal() *bool
	SetCreateEmpid(v string) *GetRulesCountListRequest
	GetCreateEmpid() *string
	SetCreateUserId(v int64) *GetRulesCountListRequest
	GetCreateUserId() *int64
	SetCurrentPage(v int32) *GetRulesCountListRequest
	GetCurrentPage() *int32
	SetEndTime(v string) *GetRulesCountListRequest
	GetEndTime() *string
	SetLastUpdateEmpid(v string) *GetRulesCountListRequest
	GetLastUpdateEmpid() *string
	SetPageNumber(v int32) *GetRulesCountListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetRulesCountListRequest
	GetPageSize() *int32
	SetRequireInfos(v []*string) *GetRulesCountListRequest
	GetRequireInfos() []*string
	SetRid(v int64) *GetRulesCountListRequest
	GetRid() *int64
	SetRuleIdOrRuleName(v string) *GetRulesCountListRequest
	GetRuleIdOrRuleName() *string
	SetRuleScoreSingleType(v int32) *GetRulesCountListRequest
	GetRuleScoreSingleType() *int32
	SetRuleType(v int32) *GetRulesCountListRequest
	GetRuleType() *int32
	SetSchemeId(v int64) *GetRulesCountListRequest
	GetSchemeId() *int64
	SetSourceType(v int32) *GetRulesCountListRequest
	GetSourceType() *int32
	SetStartTime(v string) *GetRulesCountListRequest
	GetStartTime() *string
	SetStatus(v int32) *GetRulesCountListRequest
	GetStatus() *int32
	SetType(v int32) *GetRulesCountListRequest
	GetType() *int32
	SetTypeName(v string) *GetRulesCountListRequest
	GetTypeName() *string
	SetUpdateEndTime(v string) *GetRulesCountListRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *GetRulesCountListRequest
	GetUpdateStartTime() *string
	SetUpdateUserId(v int64) *GetRulesCountListRequest
	GetUpdateUserId() *int64
}

type GetRulesCountListRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	BusinessName  *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 1
	BusinessRange *int32  `json:"BusinessRange,omitempty" xml:"BusinessRange,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// true
	CountTotal *bool `json:"CountTotal,omitempty" xml:"CountTotal,omitempty"`
	// example:
	//
	// 63
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 63
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2022-10-08 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 63
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
	// 123
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 123
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
	// 123
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 2022-10-07 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// example:
	//
	// 2022-10-08 23:59:59
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" xml:"UpdateEndTime,omitempty"`
	// example:
	//
	// 2022-10-07 00:00:00
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" xml:"UpdateStartTime,omitempty"`
	// example:
	//
	// 63
	UpdateUserId *int64 `json:"UpdateUserId,omitempty" xml:"UpdateUserId,omitempty"`
}

func (s GetRulesCountListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountListRequest) GoString() string {
	return s.String()
}

func (s *GetRulesCountListRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetRulesCountListRequest) GetBusinessName() *string {
	return s.BusinessName
}

func (s *GetRulesCountListRequest) GetBusinessRange() *int32 {
	return s.BusinessRange
}

func (s *GetRulesCountListRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetRulesCountListRequest) GetCountTotal() *bool {
	return s.CountTotal
}

func (s *GetRulesCountListRequest) GetCreateEmpid() *string {
	return s.CreateEmpid
}

func (s *GetRulesCountListRequest) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *GetRulesCountListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetRulesCountListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetRulesCountListRequest) GetLastUpdateEmpid() *string {
	return s.LastUpdateEmpid
}

func (s *GetRulesCountListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRulesCountListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRulesCountListRequest) GetRequireInfos() []*string {
	return s.RequireInfos
}

func (s *GetRulesCountListRequest) GetRid() *int64 {
	return s.Rid
}

func (s *GetRulesCountListRequest) GetRuleIdOrRuleName() *string {
	return s.RuleIdOrRuleName
}

func (s *GetRulesCountListRequest) GetRuleScoreSingleType() *int32 {
	return s.RuleScoreSingleType
}

func (s *GetRulesCountListRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *GetRulesCountListRequest) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetRulesCountListRequest) GetSourceType() *int32 {
	return s.SourceType
}

func (s *GetRulesCountListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetRulesCountListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetRulesCountListRequest) GetType() *int32 {
	return s.Type
}

func (s *GetRulesCountListRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *GetRulesCountListRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *GetRulesCountListRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *GetRulesCountListRequest) GetUpdateUserId() *int64 {
	return s.UpdateUserId
}

func (s *GetRulesCountListRequest) SetBaseMeAgentId(v int64) *GetRulesCountListRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRulesCountListRequest) SetBusinessName(v string) *GetRulesCountListRequest {
	s.BusinessName = &v
	return s
}

func (s *GetRulesCountListRequest) SetBusinessRange(v int32) *GetRulesCountListRequest {
	s.BusinessRange = &v
	return s
}

func (s *GetRulesCountListRequest) SetCategoryName(v string) *GetRulesCountListRequest {
	s.CategoryName = &v
	return s
}

func (s *GetRulesCountListRequest) SetCountTotal(v bool) *GetRulesCountListRequest {
	s.CountTotal = &v
	return s
}

func (s *GetRulesCountListRequest) SetCreateEmpid(v string) *GetRulesCountListRequest {
	s.CreateEmpid = &v
	return s
}

func (s *GetRulesCountListRequest) SetCreateUserId(v int64) *GetRulesCountListRequest {
	s.CreateUserId = &v
	return s
}

func (s *GetRulesCountListRequest) SetCurrentPage(v int32) *GetRulesCountListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetRulesCountListRequest) SetEndTime(v string) *GetRulesCountListRequest {
	s.EndTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetLastUpdateEmpid(v string) *GetRulesCountListRequest {
	s.LastUpdateEmpid = &v
	return s
}

func (s *GetRulesCountListRequest) SetPageNumber(v int32) *GetRulesCountListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRulesCountListRequest) SetPageSize(v int32) *GetRulesCountListRequest {
	s.PageSize = &v
	return s
}

func (s *GetRulesCountListRequest) SetRequireInfos(v []*string) *GetRulesCountListRequest {
	s.RequireInfos = v
	return s
}

func (s *GetRulesCountListRequest) SetRid(v int64) *GetRulesCountListRequest {
	s.Rid = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleIdOrRuleName(v string) *GetRulesCountListRequest {
	s.RuleIdOrRuleName = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleScoreSingleType(v int32) *GetRulesCountListRequest {
	s.RuleScoreSingleType = &v
	return s
}

func (s *GetRulesCountListRequest) SetRuleType(v int32) *GetRulesCountListRequest {
	s.RuleType = &v
	return s
}

func (s *GetRulesCountListRequest) SetSchemeId(v int64) *GetRulesCountListRequest {
	s.SchemeId = &v
	return s
}

func (s *GetRulesCountListRequest) SetSourceType(v int32) *GetRulesCountListRequest {
	s.SourceType = &v
	return s
}

func (s *GetRulesCountListRequest) SetStartTime(v string) *GetRulesCountListRequest {
	s.StartTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetStatus(v int32) *GetRulesCountListRequest {
	s.Status = &v
	return s
}

func (s *GetRulesCountListRequest) SetType(v int32) *GetRulesCountListRequest {
	s.Type = &v
	return s
}

func (s *GetRulesCountListRequest) SetTypeName(v string) *GetRulesCountListRequest {
	s.TypeName = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateEndTime(v string) *GetRulesCountListRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateStartTime(v string) *GetRulesCountListRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *GetRulesCountListRequest) SetUpdateUserId(v int64) *GetRulesCountListRequest {
	s.UpdateUserId = &v
	return s
}

func (s *GetRulesCountListRequest) Validate() error {
	return dara.Validate(s)
}
