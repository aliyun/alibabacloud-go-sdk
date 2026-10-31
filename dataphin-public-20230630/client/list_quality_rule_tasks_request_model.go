// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRuleTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityRuleTasksRequestListQuery) *ListQualityRuleTasksRequest
	GetListQuery() *ListQualityRuleTasksRequestListQuery
	SetOpTenantId(v int64) *ListQualityRuleTasksRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListQualityRuleTasksRequest
	GetOpUserId() *string
}

type ListQualityRuleTasksRequest struct {
	// The paging query conditions.
	ListQuery *ListQualityRuleTasksRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListQualityRuleTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksRequest) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksRequest) GetListQuery() *ListQualityRuleTasksRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityRuleTasksRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityRuleTasksRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListQualityRuleTasksRequest) SetListQuery(v *ListQualityRuleTasksRequestListQuery) *ListQualityRuleTasksRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityRuleTasksRequest) SetOpTenantId(v int64) *ListQualityRuleTasksRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityRuleTasksRequest) SetOpUserId(v string) *ListQualityRuleTasksRequest {
	s.OpUserId = &v
	return s
}

func (s *ListQualityRuleTasksRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRuleTasksRequestListQuery struct {
	// The business date.
	//
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The rule types. Valid values: CONSISTENT (consistency), EFFECTIVE (effectiveness), TIMELINESE (timeliness), ACCURATE (accuracy), UNIQUENESS (uniqueness), COMPLETENESS (completeness), STABILITY (stability), and CUSTOM (custom).
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// The search keyword. You can search by field name or rule name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule strengths. Valid values: STRONG (strong) and WEAK (weak).
	RuleStrengthList []*string `json:"RuleStrengthList,omitempty" xml:"RuleStrengthList,omitempty" type:"Repeated"`
	// The rule labels. Valid values: DEFAULT (default label), DATA_STANDARD_MANUAL (manually created standard rules), DATA_STANDARD_AUTO (standard quality rules created by automatic creation), PIPELINE (rules created by pipelines), and DATA_MODELING (data modeling).
	RuleTagList []*string `json:"RuleTagList,omitempty" xml:"RuleTagList,omitempty" type:"Repeated"`
	// The task statuses. Valid values: NOT_RUN (not run), WAITING (waiting), RUNNING (running), SUCCESS (successful), FAILED (failed), CANCEL (canceled), TIMEOUT (timed out), and OFFLINE (offline).
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The validation results. Valid values: NOT_RUN (not run), WAITING (waiting), RUNNING (running), PASS (passed), NOT_PASS (not passed), FAILED (failed), OFFLINE (offline and requires restart), CANCEL (canceled), and TIMEOUT (timed out).
	ValidateResultList []*string `json:"ValidateResultList,omitempty" xml:"ValidateResultList,omitempty" type:"Repeated"`
	// The quality watchtask ID.
	//
	// example:
	//
	// 1
	WatchTaskId *int64 `json:"WatchTaskId,omitempty" xml:"WatchTaskId,omitempty"`
}

func (s ListQualityRuleTasksRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksRequestListQuery) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityRuleTasksRequestListQuery) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *ListQualityRuleTasksRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityRuleTasksRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityRuleTasksRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityRuleTasksRequestListQuery) GetRuleStrengthList() []*string {
	return s.RuleStrengthList
}

func (s *ListQualityRuleTasksRequestListQuery) GetRuleTagList() []*string {
	return s.RuleTagList
}

func (s *ListQualityRuleTasksRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListQualityRuleTasksRequestListQuery) GetValidateResultList() []*string {
	return s.ValidateResultList
}

func (s *ListQualityRuleTasksRequestListQuery) GetWatchTaskId() *int64 {
	return s.WatchTaskId
}

func (s *ListQualityRuleTasksRequestListQuery) SetBizDate(v string) *ListQualityRuleTasksRequestListQuery {
	s.BizDate = &v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetCatalogList(v []*string) *ListQualityRuleTasksRequestListQuery {
	s.CatalogList = v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetKeyword(v string) *ListQualityRuleTasksRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetPageNo(v int32) *ListQualityRuleTasksRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetPageSize(v int32) *ListQualityRuleTasksRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetRuleStrengthList(v []*string) *ListQualityRuleTasksRequestListQuery {
	s.RuleStrengthList = v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetRuleTagList(v []*string) *ListQualityRuleTasksRequestListQuery {
	s.RuleTagList = v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetStatusList(v []*string) *ListQualityRuleTasksRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetValidateResultList(v []*string) *ListQualityRuleTasksRequestListQuery {
	s.ValidateResultList = v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) SetWatchTaskId(v int64) *ListQualityRuleTasksRequestListQuery {
	s.WatchTaskId = &v
	return s
}

func (s *ListQualityRuleTasksRequestListQuery) Validate() error {
	return dara.Validate(s)
}
