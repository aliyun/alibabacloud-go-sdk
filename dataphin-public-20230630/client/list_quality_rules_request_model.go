// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityRulesRequestListQuery) *ListQualityRulesRequest
	GetListQuery() *ListQualityRulesRequestListQuery
	SetOpTenantId(v int64) *ListQualityRulesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListQualityRulesRequest
	GetOpUserId() *string
}

type ListQualityRulesRequest struct {
	// The paged query conditions, including paging parameters.
	ListQuery *ListQualityRulesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListQualityRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityRulesRequest) GetListQuery() *ListQualityRulesRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityRulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityRulesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListQualityRulesRequest) SetListQuery(v *ListQualityRulesRequestListQuery) *ListQualityRulesRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityRulesRequest) SetOpTenantId(v int64) *ListQualityRulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityRulesRequest) SetOpUserId(v string) *ListQualityRulesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListQualityRulesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRulesRequestListQuery struct {
	// The rule type. Valid values: CONSISTENT (Consistency), EFFECTIVE (Effectiveness), TIMELINESE (Timeliness), ACCURATE (Accuracy), UNIQUENESS (Uniqueness), COMPLETENESS (Completeness), STABILITY (Stability), and CUSTOM (Custom).
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// The search keyword. Supports searching by rule name and validation object.
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
	// The rule strength. Valid values: STRONG (Strong) and WEAK (Weak).
	RuleStrengthList []*string `json:"RuleStrengthList,omitempty" xml:"RuleStrengthList,omitempty" type:"Repeated"`
	// The effective status of the rule. Valid values: ENABLE (Enabled) and DISABLE (Disabled).
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The rule template.
	TemplateIdList []*int64 `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
	// The task status. Valid values: NOT_RUN (Not run), WAITING (Waiting), RUNNING (Running), SUCCESS (Succeeded), FAILED (Failed), CANCEL (Canceled), TIMEOUT (Timed out), and OFFLINE (Offline).
	TestRunTaskStatusList []*string `json:"TestRunTaskStatusList,omitempty" xml:"TestRunTaskStatusList,omitempty" type:"Repeated"`
	// The validation result of the rule test run. Valid values: NOT_RUN (Not run), WAITING (Waiting), RUNNING (Running), PASS (Passed), NOT_PASS (Not passed), FAILED (Failed), OFFLINE (Offline, requires restart), CANCEL (Canceled), and TIMEOUT (Timed out).
	TestRunTaskValidateResultList []*string `json:"TestRunTaskValidateResultList,omitempty" xml:"TestRunTaskValidateResultList,omitempty" type:"Repeated"`
	// The monitoring ID.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s ListQualityRulesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityRulesRequestListQuery) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *ListQualityRulesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityRulesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityRulesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityRulesRequestListQuery) GetRuleStrengthList() []*string {
	return s.RuleStrengthList
}

func (s *ListQualityRulesRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListQualityRulesRequestListQuery) GetTemplateIdList() []*int64 {
	return s.TemplateIdList
}

func (s *ListQualityRulesRequestListQuery) GetTestRunTaskStatusList() []*string {
	return s.TestRunTaskStatusList
}

func (s *ListQualityRulesRequestListQuery) GetTestRunTaskValidateResultList() []*string {
	return s.TestRunTaskValidateResultList
}

func (s *ListQualityRulesRequestListQuery) GetWatchId() *int64 {
	return s.WatchId
}

func (s *ListQualityRulesRequestListQuery) SetCatalogList(v []*string) *ListQualityRulesRequestListQuery {
	s.CatalogList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetKeyword(v string) *ListQualityRulesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetPageNo(v int32) *ListQualityRulesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetPageSize(v int32) *ListQualityRulesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetRuleStrengthList(v []*string) *ListQualityRulesRequestListQuery {
	s.RuleStrengthList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetStatusList(v []*string) *ListQualityRulesRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetTemplateIdList(v []*int64) *ListQualityRulesRequestListQuery {
	s.TemplateIdList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetTestRunTaskStatusList(v []*string) *ListQualityRulesRequestListQuery {
	s.TestRunTaskStatusList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetTestRunTaskValidateResultList(v []*string) *ListQualityRulesRequestListQuery {
	s.TestRunTaskValidateResultList = v
	return s
}

func (s *ListQualityRulesRequestListQuery) SetWatchId(v int64) *ListQualityRulesRequestListQuery {
	s.WatchId = &v
	return s
}

func (s *ListQualityRulesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
