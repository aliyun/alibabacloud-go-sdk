// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleQuery interface {
	dara.Model
	String() string
	GoString() string
	SetCheckAfterDataComplete(v bool) *AlertRuleQuery
	GetCheckAfterDataComplete() *bool
	SetDimensions(v []map[string]*string) *AlertRuleQuery
	GetDimensions() []map[string]*string
	SetDomain(v string) *AlertRuleQuery
	GetDomain() *string
	SetDuration(v int64) *AlertRuleQuery
	GetDuration() *int64
	SetEntityFields(v []*AlertRuleQueryEntityFields) *AlertRuleQuery
	GetEntityFields() []*AlertRuleQueryEntityFields
	SetEntityFilter(v *AlertRuleQueryEntityFilter) *AlertRuleQuery
	GetEntityFilter() *AlertRuleQueryEntityFilter
	SetExpr(v string) *AlertRuleQuery
	GetExpr() *string
	SetFirstJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery
	GetFirstJoin() *AlertRuleSlsQueryJoin
	SetGroupFieldList(v []*string) *AlertRuleQuery
	GetGroupFieldList() []*string
	SetGroupId(v string) *AlertRuleQuery
	GetGroupId() *string
	SetGroupType(v string) *AlertRuleQuery
	GetGroupType() *string
	SetLabelFilters(v []*AlertRuleQueryLabelFilters) *AlertRuleQuery
	GetLabelFilters() []*AlertRuleQueryLabelFilters
	SetMetric(v string) *AlertRuleQuery
	GetMetric() *string
	SetMetricSet(v string) *AlertRuleQuery
	GetMetricSet() *string
	SetNamespace(v string) *AlertRuleQuery
	GetNamespace() *string
	SetQueries(v []*AlertRuleQueryQueries) *AlertRuleQuery
	GetQueries() []*AlertRuleQueryQueries
	SetRelationType(v string) *AlertRuleQuery
	GetRelationType() *string
	SetSecondJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery
	GetSecondJoin() *AlertRuleSlsQueryJoin
	SetServiceIds(v []*string) *AlertRuleQuery
	GetServiceIds() []*string
	SetType(v string) *AlertRuleQuery
	GetType() *string
}

type AlertRuleQuery struct {
	CheckAfterDataComplete *bool                         `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	Dimensions             []map[string]*string          `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	Domain                 *string                       `json:"domain,omitempty" xml:"domain,omitempty"`
	Duration               *int64                        `json:"duration,omitempty" xml:"duration,omitempty"`
	EntityFields           []*AlertRuleQueryEntityFields `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	EntityFilter           *AlertRuleQueryEntityFilter   `json:"entityFilter,omitempty" xml:"entityFilter,omitempty" type:"Struct"`
	Expr                   *string                       `json:"expr,omitempty" xml:"expr,omitempty"`
	FirstJoin              *AlertRuleSlsQueryJoin        `json:"firstJoin,omitempty" xml:"firstJoin,omitempty"`
	GroupFieldList         []*string                     `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	GroupId                *string                       `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupType              *string                       `json:"groupType,omitempty" xml:"groupType,omitempty"`
	LabelFilters           []*AlertRuleQueryLabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	Metric                 *string                       `json:"metric,omitempty" xml:"metric,omitempty"`
	MetricSet              *string                       `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	Namespace              *string                       `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Queries                []*AlertRuleQueryQueries      `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	RelationType           *string                       `json:"relationType,omitempty" xml:"relationType,omitempty"`
	SecondJoin             *AlertRuleSlsQueryJoin        `json:"secondJoin,omitempty" xml:"secondJoin,omitempty"`
	ServiceIds             []*string                     `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
	// 查询类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleQuery) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQuery) GoString() string {
	return s.String()
}

func (s *AlertRuleQuery) GetCheckAfterDataComplete() *bool {
	return s.CheckAfterDataComplete
}

func (s *AlertRuleQuery) GetDimensions() []map[string]*string {
	return s.Dimensions
}

func (s *AlertRuleQuery) GetDomain() *string {
	return s.Domain
}

func (s *AlertRuleQuery) GetDuration() *int64 {
	return s.Duration
}

func (s *AlertRuleQuery) GetEntityFields() []*AlertRuleQueryEntityFields {
	return s.EntityFields
}

func (s *AlertRuleQuery) GetEntityFilter() *AlertRuleQueryEntityFilter {
	return s.EntityFilter
}

func (s *AlertRuleQuery) GetExpr() *string {
	return s.Expr
}

func (s *AlertRuleQuery) GetFirstJoin() *AlertRuleSlsQueryJoin {
	return s.FirstJoin
}

func (s *AlertRuleQuery) GetGroupFieldList() []*string {
	return s.GroupFieldList
}

func (s *AlertRuleQuery) GetGroupId() *string {
	return s.GroupId
}

func (s *AlertRuleQuery) GetGroupType() *string {
	return s.GroupType
}

func (s *AlertRuleQuery) GetLabelFilters() []*AlertRuleQueryLabelFilters {
	return s.LabelFilters
}

func (s *AlertRuleQuery) GetMetric() *string {
	return s.Metric
}

func (s *AlertRuleQuery) GetMetricSet() *string {
	return s.MetricSet
}

func (s *AlertRuleQuery) GetNamespace() *string {
	return s.Namespace
}

func (s *AlertRuleQuery) GetQueries() []*AlertRuleQueryQueries {
	return s.Queries
}

func (s *AlertRuleQuery) GetRelationType() *string {
	return s.RelationType
}

func (s *AlertRuleQuery) GetSecondJoin() *AlertRuleSlsQueryJoin {
	return s.SecondJoin
}

func (s *AlertRuleQuery) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *AlertRuleQuery) GetType() *string {
	return s.Type
}

func (s *AlertRuleQuery) SetCheckAfterDataComplete(v bool) *AlertRuleQuery {
	s.CheckAfterDataComplete = &v
	return s
}

func (s *AlertRuleQuery) SetDimensions(v []map[string]*string) *AlertRuleQuery {
	s.Dimensions = v
	return s
}

func (s *AlertRuleQuery) SetDomain(v string) *AlertRuleQuery {
	s.Domain = &v
	return s
}

func (s *AlertRuleQuery) SetDuration(v int64) *AlertRuleQuery {
	s.Duration = &v
	return s
}

func (s *AlertRuleQuery) SetEntityFields(v []*AlertRuleQueryEntityFields) *AlertRuleQuery {
	s.EntityFields = v
	return s
}

func (s *AlertRuleQuery) SetEntityFilter(v *AlertRuleQueryEntityFilter) *AlertRuleQuery {
	s.EntityFilter = v
	return s
}

func (s *AlertRuleQuery) SetExpr(v string) *AlertRuleQuery {
	s.Expr = &v
	return s
}

func (s *AlertRuleQuery) SetFirstJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery {
	s.FirstJoin = v
	return s
}

func (s *AlertRuleQuery) SetGroupFieldList(v []*string) *AlertRuleQuery {
	s.GroupFieldList = v
	return s
}

func (s *AlertRuleQuery) SetGroupId(v string) *AlertRuleQuery {
	s.GroupId = &v
	return s
}

func (s *AlertRuleQuery) SetGroupType(v string) *AlertRuleQuery {
	s.GroupType = &v
	return s
}

func (s *AlertRuleQuery) SetLabelFilters(v []*AlertRuleQueryLabelFilters) *AlertRuleQuery {
	s.LabelFilters = v
	return s
}

func (s *AlertRuleQuery) SetMetric(v string) *AlertRuleQuery {
	s.Metric = &v
	return s
}

func (s *AlertRuleQuery) SetMetricSet(v string) *AlertRuleQuery {
	s.MetricSet = &v
	return s
}

func (s *AlertRuleQuery) SetNamespace(v string) *AlertRuleQuery {
	s.Namespace = &v
	return s
}

func (s *AlertRuleQuery) SetQueries(v []*AlertRuleQueryQueries) *AlertRuleQuery {
	s.Queries = v
	return s
}

func (s *AlertRuleQuery) SetRelationType(v string) *AlertRuleQuery {
	s.RelationType = &v
	return s
}

func (s *AlertRuleQuery) SetSecondJoin(v *AlertRuleSlsQueryJoin) *AlertRuleQuery {
	s.SecondJoin = v
	return s
}

func (s *AlertRuleQuery) SetServiceIds(v []*string) *AlertRuleQuery {
	s.ServiceIds = v
	return s
}

func (s *AlertRuleQuery) SetType(v string) *AlertRuleQuery {
	s.Type = &v
	return s
}

func (s *AlertRuleQuery) Validate() error {
	if s.EntityFields != nil {
		for _, item := range s.EntityFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntityFilter != nil {
		if err := s.EntityFilter.Validate(); err != nil {
			return err
		}
	}
	if s.FirstJoin != nil {
		if err := s.FirstJoin.Validate(); err != nil {
			return err
		}
	}
	if s.LabelFilters != nil {
		for _, item := range s.LabelFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Queries != nil {
		for _, item := range s.Queries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecondJoin != nil {
		if err := s.SecondJoin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AlertRuleQueryEntityFields struct {
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryEntityFields) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryEntityFields) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryEntityFields) GetField() *string {
	return s.Field
}

func (s *AlertRuleQueryEntityFields) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryEntityFields) SetField(v string) *AlertRuleQueryEntityFields {
	s.Field = &v
	return s
}

func (s *AlertRuleQueryEntityFields) SetValue(v string) *AlertRuleQueryEntityFields {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryEntityFields) Validate() error {
	return dara.Validate(s)
}

type AlertRuleQueryEntityFilter struct {
	Domain  *string                              `json:"domain,omitempty" xml:"domain,omitempty"`
	Filters []*AlertRuleQueryEntityFilterFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	Type    *string                              `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleQueryEntityFilter) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryEntityFilter) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryEntityFilter) GetDomain() *string {
	return s.Domain
}

func (s *AlertRuleQueryEntityFilter) GetFilters() []*AlertRuleQueryEntityFilterFilters {
	return s.Filters
}

func (s *AlertRuleQueryEntityFilter) GetType() *string {
	return s.Type
}

func (s *AlertRuleQueryEntityFilter) SetDomain(v string) *AlertRuleQueryEntityFilter {
	s.Domain = &v
	return s
}

func (s *AlertRuleQueryEntityFilter) SetFilters(v []*AlertRuleQueryEntityFilterFilters) *AlertRuleQueryEntityFilter {
	s.Filters = v
	return s
}

func (s *AlertRuleQueryEntityFilter) SetType(v string) *AlertRuleQueryEntityFilter {
	s.Type = &v
	return s
}

func (s *AlertRuleQueryEntityFilter) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleQueryEntityFilterFilters struct {
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryEntityFilterFilters) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryEntityFilterFilters) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryEntityFilterFilters) GetField() *string {
	return s.Field
}

func (s *AlertRuleQueryEntityFilterFilters) GetOperator() *string {
	return s.Operator
}

func (s *AlertRuleQueryEntityFilterFilters) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryEntityFilterFilters) SetField(v string) *AlertRuleQueryEntityFilterFilters {
	s.Field = &v
	return s
}

func (s *AlertRuleQueryEntityFilterFilters) SetOperator(v string) *AlertRuleQueryEntityFilterFilters {
	s.Operator = &v
	return s
}

func (s *AlertRuleQueryEntityFilterFilters) SetValue(v string) *AlertRuleQueryEntityFilterFilters {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryEntityFilterFilters) Validate() error {
	return dara.Validate(s)
}

type AlertRuleQueryLabelFilters struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryLabelFilters) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryLabelFilters) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryLabelFilters) GetName() *string {
	return s.Name
}

func (s *AlertRuleQueryLabelFilters) GetOperator() *string {
	return s.Operator
}

func (s *AlertRuleQueryLabelFilters) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryLabelFilters) SetName(v string) *AlertRuleQueryLabelFilters {
	s.Name = &v
	return s
}

func (s *AlertRuleQueryLabelFilters) SetOperator(v string) *AlertRuleQueryLabelFilters {
	s.Operator = &v
	return s
}

func (s *AlertRuleQueryLabelFilters) SetValue(v string) *AlertRuleQueryLabelFilters {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryLabelFilters) Validate() error {
	return dara.Validate(s)
}

type AlertRuleQueryQueries struct {
	ApmAlertMetricId *string                            `json:"apmAlertMetricId,omitempty" xml:"apmAlertMetricId,omitempty"`
	ApmFilters       []*AlertRuleQueryQueriesApmFilters `json:"apmFilters,omitempty" xml:"apmFilters,omitempty" type:"Repeated"`
	ApmGroupBy       []*string                          `json:"apmGroupBy,omitempty" xml:"apmGroupBy,omitempty" type:"Repeated"`
	Duration         *int64                             `json:"duration,omitempty" xml:"duration,omitempty"`
	// 时间偏移结束时间(相对)，如果指定了start、end，则不指定window。
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// 查询表达式
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// sls查询的时间偏移开始时间(相对)，如果指定了start、end，则不指定window。  例如：start=15， timeUnit=minute，表示15分钟前
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// start和end、window的时间单位： day/hour/minute/second
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 整点时间查询区间。  如果指定了window则不指定start、end
	Window *int64 `json:"window,omitempty" xml:"window,omitempty"`
}

func (s AlertRuleQueryQueries) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryQueries) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryQueries) GetApmAlertMetricId() *string {
	return s.ApmAlertMetricId
}

func (s *AlertRuleQueryQueries) GetApmFilters() []*AlertRuleQueryQueriesApmFilters {
	return s.ApmFilters
}

func (s *AlertRuleQueryQueries) GetApmGroupBy() []*string {
	return s.ApmGroupBy
}

func (s *AlertRuleQueryQueries) GetDuration() *int64 {
	return s.Duration
}

func (s *AlertRuleQueryQueries) GetEnd() *int64 {
	return s.End
}

func (s *AlertRuleQueryQueries) GetExpr() *string {
	return s.Expr
}

func (s *AlertRuleQueryQueries) GetStart() *int64 {
	return s.Start
}

func (s *AlertRuleQueryQueries) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *AlertRuleQueryQueries) GetWindow() *int64 {
	return s.Window
}

func (s *AlertRuleQueryQueries) SetApmAlertMetricId(v string) *AlertRuleQueryQueries {
	s.ApmAlertMetricId = &v
	return s
}

func (s *AlertRuleQueryQueries) SetApmFilters(v []*AlertRuleQueryQueriesApmFilters) *AlertRuleQueryQueries {
	s.ApmFilters = v
	return s
}

func (s *AlertRuleQueryQueries) SetApmGroupBy(v []*string) *AlertRuleQueryQueries {
	s.ApmGroupBy = v
	return s
}

func (s *AlertRuleQueryQueries) SetDuration(v int64) *AlertRuleQueryQueries {
	s.Duration = &v
	return s
}

func (s *AlertRuleQueryQueries) SetEnd(v int64) *AlertRuleQueryQueries {
	s.End = &v
	return s
}

func (s *AlertRuleQueryQueries) SetExpr(v string) *AlertRuleQueryQueries {
	s.Expr = &v
	return s
}

func (s *AlertRuleQueryQueries) SetStart(v int64) *AlertRuleQueryQueries {
	s.Start = &v
	return s
}

func (s *AlertRuleQueryQueries) SetTimeUnit(v string) *AlertRuleQueryQueries {
	s.TimeUnit = &v
	return s
}

func (s *AlertRuleQueryQueries) SetWindow(v int64) *AlertRuleQueryQueries {
	s.Window = &v
	return s
}

func (s *AlertRuleQueryQueries) Validate() error {
	if s.ApmFilters != nil {
		for _, item := range s.ApmFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleQueryQueriesApmFilters struct {
	Dim   *string `json:"dim,omitempty" xml:"dim,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryQueriesApmFilters) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryQueriesApmFilters) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryQueriesApmFilters) GetDim() *string {
	return s.Dim
}

func (s *AlertRuleQueryQueriesApmFilters) GetType() *string {
	return s.Type
}

func (s *AlertRuleQueryQueriesApmFilters) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryQueriesApmFilters) SetDim(v string) *AlertRuleQueryQueriesApmFilters {
	s.Dim = &v
	return s
}

func (s *AlertRuleQueryQueriesApmFilters) SetType(v string) *AlertRuleQueryQueriesApmFilters {
	s.Type = &v
	return s
}

func (s *AlertRuleQueryQueriesApmFilters) SetValue(v string) *AlertRuleQueryQueriesApmFilters {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryQueriesApmFilters) Validate() error {
	return dara.Validate(s)
}
