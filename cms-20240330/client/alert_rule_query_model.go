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
	SetMarkTags(v []*AlertRuleQueryMarkTags) *AlertRuleQuery
	GetMarkTags() []*AlertRuleQueryMarkTags
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
	// Applicable query type: PROMQL_QUERY.
	//
	// Whether to perform alert evaluation only after data completeness is ensured.
	//
	// example:
	//
	// true
	CheckAfterDataComplete *bool `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// List of filtering dimensions for the resource.
	Dimensions []map[string]*string `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	// 资源所属的领域。
	//
	// example:
	//
	// rum
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Applicable query type: PROMQL_QUERY.
	//
	// Duration of alert data, in seconds.
	//
	// example:
	//
	// 60
	Duration     *int64                        `json:"duration,omitempty" xml:"duration,omitempty"`
	EntityFields []*AlertRuleQueryEntityFields `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	// 资源过滤器，用于筛选目标资源。
	EntityFilter *AlertRuleQueryEntityFilter `json:"entityFilter,omitempty" xml:"entityFilter,omitempty" type:"Struct"`
	// Applicable query type: PROMQL_QUERY.
	//
	// Query expression (PromQL).
	//
	// example:
	//
	// sum(sum(max_over_time(kube_pod_status_phase{phase=~\\"Pending\\",job=\\"_kube-state-metrics\\"}[5m])) by (pod)) > 1000
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Configuration for the set join operation between the results of subquery 1 (queries[0]) and subquery 2 (queries[1]).
	FirstJoin *AlertRuleSlsQueryJoin `json:"firstJoin,omitempty" xml:"firstJoin,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// List of grouping field names.
	GroupFieldList []*string `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// Associated application group ID, valid only when relationType = GROUP.
	//
	// example:
	//
	// 23423
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Grouping type, with the following possible values:
	//
	// - none: No grouping.
	//
	// - label: Automatic label grouping.
	//
	// - custom: Custom label grouping.
	//
	// example:
	//
	// label
	GroupType    *string                       `json:"groupType,omitempty" xml:"groupType,omitempty"`
	LabelFilters []*AlertRuleQueryLabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	MarkTags     []*AlertRuleQueryMarkTags     `json:"markTags,omitempty" xml:"markTags,omitempty" type:"Repeated"`
	// 指标名。
	//
	// example:
	//
	// memory
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// 监控指标集合。
	//
	// example:
	//
	// cpu_usage
	MetricSet *string `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// Namespace of the metric.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Applicable query types: SLS_MULTI_QUERY, APM_MULTI_QUERY.
	//
	// List of subqueries.
	//
	// For the SLS_MULTI_QUERY type, the list can contain up to three subqueries, and the number and order of subqueries must match the sub-datasource configurations in datasource.dsList.
	Queries []*AlertRuleQueryQueries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// Resource scope for the rule query, with the following allowed values:
	//
	// - USER: All resources under the user\\"s UID.
	//
	// - GROUP: Application group.
	//
	// - INSTANCE: Specified list of instances.
	//
	// example:
	//
	// USER
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Configuration for the set join operation between the results of subquery 2 (queries[2]) and subquery 3 (queries[3]).
	SecondJoin *AlertRuleSlsQueryJoin `json:"secondJoin,omitempty" xml:"secondJoin,omitempty"`
	// Service ID list.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
	// Query type.
	//
	// Valid values:
	//
	// - PROMQL_QUERY: PromQL query
	//
	// - SLS_MULTI_QUERY: SLS query
	//
	// - APM_MULTI_QUERY: APM query
	//
	// - CMS_BASIC_QUERY: Basic CloudMonitor query
	//
	// The valid fields within the query object vary depending on the query type. Refer to the "Applicable query type" description in each field\\"s documentation for details.
	//
	// The query type must match the data source type, with the following correspondences:
	//
	// - Prometheus data source (PROMETHEUS_DS): PROMQL_QUERY
	//
	// - APM data source (APM_DS): APM_MULTI_QUERY
	//
	// - SLS data source (SLS_MULTI_DS): SLS_MULTI_QUERY
	//
	// - Basic CloudMonitor data source (CMS_BASIC_DS): CMS_BASIC_QUERY.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMQL_QUERY
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

func (s *AlertRuleQuery) GetMarkTags() []*AlertRuleQueryMarkTags {
	return s.MarkTags
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

func (s *AlertRuleQuery) SetMarkTags(v []*AlertRuleQueryMarkTags) *AlertRuleQuery {
	s.MarkTags = v
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
	if s.MarkTags != nil {
		for _, item := range s.MarkTags {
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
	// 资源类型域。
	//
	// example:
	//
	// rum
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 过滤条件列表，用于进一步筛选资源。
	Filters []*AlertRuleQueryEntityFilterFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// 资源类型。
	//
	// example:
	//
	// apm
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// 字段
	//
	// example:
	//
	// instanceId
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// 比较运算符。
	//
	// example:
	//
	// =
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 匹配的值。
	//
	// example:
	//
	// wait_throw
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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

type AlertRuleQueryMarkTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryMarkTags) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryMarkTags) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryMarkTags) GetKey() *string {
	return s.Key
}

func (s *AlertRuleQueryMarkTags) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryMarkTags) SetKey(v string) *AlertRuleQueryMarkTags {
	s.Key = &v
	return s
}

func (s *AlertRuleQueryMarkTags) SetValue(v string) *AlertRuleQueryMarkTags {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryMarkTags) Validate() error {
	return dara.Validate(s)
}

type AlertRuleQueryQueries struct {
	// Applicable query type: APM_MULTI_QUERY.
	//
	// ID of the APM predefined metric.
	//
	// example:
	//
	// appstat.jvm.ThreadNewCount
	ApmAlertMetricId *string `json:"apmAlertMetricId,omitempty" xml:"apmAlertMetricId,omitempty"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// Dimension filter configuration for APM metrics. Must be used in conjunction with apmAlertMetricId.
	ApmFilters []*AlertRuleQueryQueriesApmFilters `json:"apmFilters,omitempty" xml:"apmFilters,omitempty" type:"Repeated"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// List of aggregation dimensions for the query, i.e., the dimensions by which the metric is aggregated.
	ApmGroupBy []*string `json:"apmGroupBy,omitempty" xml:"apmGroupBy,omitempty" type:"Repeated"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// Alert (data) duration.
	//
	// example:
	//
	// 120
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Time offset end time (relative).
	//
	// If start and end are specified, do not specify window.
	//
	// example:
	//
	// 0
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// Applicable query types: APM_MULTI_QUERY, SLS_MULTI_QUERY.
	//
	// Query expression.
	//
	// - For APM_MULTI_QUERY, this field is optional and contains the PromQL generated for predefined metrics (used for data preview).
	//
	// - For SLS_MULTI_QUERY, this field contains the SQL query statement.
	//
	// example:
	//
	// sum by (rpc,acs_arms_service_id,pid,rpcType) (sum_over_time_lorc(arms_app_requests_count_ign_destid_endpoint_parent_ppid_prpc{callKind=~\\"http|rpc|custom_entry|server|consumer\\",pid=\\"gaddp9ap8q@cb005ffdf44b8ac\\",source=\\"apm\\"}[1m]))
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// SLS query time offset start time (relative).
	//
	// If start and end are specified, do not specify window. For example: start=15, timeUnit=minute, which means 15 minutes ago.
	//
	// example:
	//
	// 15
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Time units for the start, end, and window parameters: day/hour/minute/second.
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// Exact-hour time query interval. If window is specified, start and end should not be specified.
	//
	// example:
	//
	// 1
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
	// Dimension in APM metrics.
	//
	// example:
	//
	// rpcType
	Dim *string `json:"dim,omitempty" xml:"dim,omitempty"`
	// Filter operation types:
	//
	// - eq: equals.
	//
	// - neq: not equals.
	//
	// - match: regular expression match.
	//
	// - nmatch: regular expression not match.
	//
	// example:
	//
	// eq
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The corresponding value for the filter operation.
	//
	// example:
	//
	// h3ji7a0y9i@2ac80e27fdfd0a2
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
