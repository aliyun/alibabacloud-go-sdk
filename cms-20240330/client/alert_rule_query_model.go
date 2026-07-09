// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleQuery interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v string) *AlertRuleQuery
	GetAggregate() *string
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
	SetLogSet(v string) *AlertRuleQuery
	GetLogSet() *string
	SetMarkTags(v []*AlertRuleQueryMarkTags) *AlertRuleQuery
	GetMarkTags() []*AlertRuleQueryMarkTags
	SetMetric(v string) *AlertRuleQuery
	GetMetric() *string
	SetMetricSet(v string) *AlertRuleQuery
	GetMetricSet() *string
	SetNamespace(v string) *AlertRuleQuery
	GetNamespace() *string
	SetOffsetSecs(v int64) *AlertRuleQuery
	GetOffsetSecs() *int64
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
	SetWindowSecs(v int64) *AlertRuleQuery
	GetWindowSecs() *int64
}

type AlertRuleQuery struct {
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// Applicable query type: PROMQL_QUERY.
	//
	// Specifies whether to perform alert detection only after data is complete.
	//
	// example:
	//
	// true
	CheckAfterDataComplete *bool `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// The list of filter dimensions for the resource.
	Dimensions []map[string]*string `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	// The domain to which the resource belongs.
	//
	// example:
	//
	// rum
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Applicable query type: PROMQL_QUERY.
	//
	// The duration for which alert data persists. Unit: seconds.
	//
	// example:
	//
	// 60
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The array of entity field filters.
	EntityFields []*AlertRuleQueryEntityFields `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	// The resource filter used to filter target resources.
	EntityFilter *AlertRuleQueryEntityFilter `json:"entityFilter,omitempty" xml:"entityFilter,omitempty" type:"Struct"`
	// Applicable query type: PROMQL_QUERY.
	//
	// The query expression (PromQL).
	//
	// example:
	//
	// sum(sum(max_over_time(kube_pod_status_phase{phase=~\\"Pending\\",job=\\"_kube-state-metrics\\"}[5m])) by (pod)) > 1000
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The set join operation configuration for the results of subquery 1 (queries[0]) and subquery 2 (queries[1]).
	FirstJoin *AlertRuleSlsQueryJoin `json:"firstJoin,omitempty" xml:"firstJoin,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The list of group field names.
	GroupFieldList []*string `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// The ID of the associated application group. This parameter takes effect only when relationType is set to GROUP.
	//
	// example:
	//
	// 23423
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The group type. Valid values:
	//
	// - none: no grouping.
	//
	// - label: automatic label-based grouping.
	//
	// - custom: custom label-based grouping.
	//
	// example:
	//
	// label
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The array of label filters.
	LabelFilters []*AlertRuleQueryLabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	LogSet       *string                       `json:"logSet,omitempty" xml:"logSet,omitempty"`
	MarkTags     []*AlertRuleQueryMarkTags     `json:"markTags,omitempty" xml:"markTags,omitempty" type:"Repeated"`
	// The metric name.
	//
	// example:
	//
	// memory
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// The collection of monitoring metrics.
	//
	// example:
	//
	// cpu_usage
	MetricSet *string `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// The namespace of the metric.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OffsetSecs *int64  `json:"offsetSecs,omitempty" xml:"offsetSecs,omitempty"`
	// Applicable query types: SLS_MULTI_QUERY and APM_MULTI_QUERY.
	//
	// The list of subqueries.
	//
	// For the SLS_MULTI_QUERY query type, a maximum of three subqueries are supported. The number and order of subqueries must match the sub-datasource config in datasource.dsList.
	Queries []*AlertRuleQueryQueries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// The resource scope of the rule query. Valid values:
	//
	// - USER: all resources under the user UID.
	//
	// - GROUP: application group.
	//
	// - INSTANCE: specified instance list.
	//
	// example:
	//
	// USER
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The set join operation configuration for the results of subquery 2 (queries[2]) and subquery 3 (queries[3]).
	SecondJoin *AlertRuleSlsQueryJoin `json:"secondJoin,omitempty" xml:"secondJoin,omitempty"`
	// The list of service IDs.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
	// The query type.
	//
	// Valid values:
	//
	// - PROMQL_QUERY: PromQL query.
	//
	// - SLS_MULTI_QUERY: SLS query.
	//
	// - APM_MULTI_QUERY: APM query.
	//
	// - CMS_BASIC_QUERY: basic cloud service monitoring query.
	//
	// Different query types use different valid fields in the query object. For more information, see the "Applicable query type" description of each field.
	//
	// The query type must match the data source type. The mappings are as follows:
	//
	// - Prometheus data source (PROMETHEUS_DS): PROMQL_QUERY
	//
	// - APM data source (APM_DS): APM_MULTI_QUERY
	//
	// - SLS data source (SLS_MULTI_DS): SLS_MULTI_QUERY
	//
	// - Basic cloud service monitoring data source (CMS_BASIC_DS): CMS_BASIC_QUERY
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMQL_QUERY
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	WindowSecs *int64  `json:"windowSecs,omitempty" xml:"windowSecs,omitempty"`
}

func (s AlertRuleQuery) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQuery) GoString() string {
	return s.String()
}

func (s *AlertRuleQuery) GetAggregate() *string {
	return s.Aggregate
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

func (s *AlertRuleQuery) GetLogSet() *string {
	return s.LogSet
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

func (s *AlertRuleQuery) GetOffsetSecs() *int64 {
	return s.OffsetSecs
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

func (s *AlertRuleQuery) GetWindowSecs() *int64 {
	return s.WindowSecs
}

func (s *AlertRuleQuery) SetAggregate(v string) *AlertRuleQuery {
	s.Aggregate = &v
	return s
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

func (s *AlertRuleQuery) SetLogSet(v string) *AlertRuleQuery {
	s.LogSet = &v
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

func (s *AlertRuleQuery) SetOffsetSecs(v int64) *AlertRuleQuery {
	s.OffsetSecs = &v
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

func (s *AlertRuleQuery) SetWindowSecs(v int64) *AlertRuleQuery {
	s.WindowSecs = &v
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
	// The entity field name.
	//
	// example:
	//
	// instanceId
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The field value.
	//
	// example:
	//
	// i-abc123
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
	// The resource type domain.
	//
	// example:
	//
	// rum
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The list of filter conditions used to further filter resources.
	Filters []*AlertRuleQueryEntityFilterFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// The resource type.
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
	// The field.
	//
	// example:
	//
	// instanceId
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// =
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The matched value.
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
	// The label name.
	//
	// example:
	//
	// app
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The comparison operator that determines how to match the label value.
	//
	// example:
	//
	// =
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The label value.
	//
	// example:
	//
	// web
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	// The ID of the APM predefined metric.
	//
	// example:
	//
	// appstat.jvm.ThreadNewCount
	ApmAlertMetricId *string `json:"apmAlertMetricId,omitempty" xml:"apmAlertMetricId,omitempty"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// The dimension filter configuration for the APM metric. Must be used together with apmAlertMetricId.
	ApmFilters []*AlertRuleQueryQueriesApmFilters `json:"apmFilters,omitempty" xml:"apmFilters,omitempty" type:"Repeated"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// The list of aggregation dimensions for the query, specifying which metric dimensions to aggregate by.
	ApmGroupBy []*string `json:"apmGroupBy,omitempty" xml:"apmGroupBy,omitempty" type:"Repeated"`
	// Applicable query type: ARMS_MULTI_QUERY.
	//
	// The alert data duration.
	//
	// example:
	//
	// 120
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The relative time offset end time.
	//
	// If start and end are specified, do not specify window.
	//
	// example:
	//
	// 0
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// Applicable query types: APM_MULTI_QUERY, SLS_MULTI_QUERY.
	//
	// The query expression.
	//
	// - For APM_MULTI_QUERY, this field is optional and contains the PromQL generated for predefined metrics (used for data preview).
	//
	// - For SLS_MULTI_QUERY, this field contains the SQL query statement.
	//
	// example:
	//
	// sum by (rpc,acs_arms_service_id,pid,rpcType) (sum_over_time_lorc(arms_app_requests_count_ign_destid_endpoint_parent_ppid_prpc{callKind=~\\"http|rpc|custom_entry|server|consumer\\",pid=\\"gaddp9ap8q@cb005ffdf44b8ac\\",source=\\"apm\\"}[1m]))
	Expr         *string                              `json:"expr,omitempty" xml:"expr,omitempty"`
	LabelFilters []*AlertRuleQueryQueriesLabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	Metric       *string                              `json:"metric,omitempty" xml:"metric,omitempty"`
	MetricSet    *string                              `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	Name         *string                              `json:"name,omitempty" xml:"name,omitempty"`
	PromQl       *string                              `json:"promQl,omitempty" xml:"promQl,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The relative time offset start time for the SLS query.
	//
	// If start and end are specified, do not specify window. Example: start=15, timeUnit=minute indicates 15 minutes ago.
	//
	// example:
	//
	// 15
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The time unit for the start, end, and window parameters: day/hour/minute/second.
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// Applicable query type: SLS_MULTI_QUERY.
	//
	// The time frame query interval. If window is specified, do not specify start or end.
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

func (s *AlertRuleQueryQueries) GetLabelFilters() []*AlertRuleQueryQueriesLabelFilters {
	return s.LabelFilters
}

func (s *AlertRuleQueryQueries) GetMetric() *string {
	return s.Metric
}

func (s *AlertRuleQueryQueries) GetMetricSet() *string {
	return s.MetricSet
}

func (s *AlertRuleQueryQueries) GetName() *string {
	return s.Name
}

func (s *AlertRuleQueryQueries) GetPromQl() *string {
	return s.PromQl
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

func (s *AlertRuleQueryQueries) SetLabelFilters(v []*AlertRuleQueryQueriesLabelFilters) *AlertRuleQueryQueries {
	s.LabelFilters = v
	return s
}

func (s *AlertRuleQueryQueries) SetMetric(v string) *AlertRuleQueryQueries {
	s.Metric = &v
	return s
}

func (s *AlertRuleQueryQueries) SetMetricSet(v string) *AlertRuleQueryQueries {
	s.MetricSet = &v
	return s
}

func (s *AlertRuleQueryQueries) SetName(v string) *AlertRuleQueryQueries {
	s.Name = &v
	return s
}

func (s *AlertRuleQueryQueries) SetPromQl(v string) *AlertRuleQueryQueries {
	s.PromQl = &v
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
	if s.LabelFilters != nil {
		for _, item := range s.LabelFilters {
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
	// The dimension in the APM metric.
	//
	// example:
	//
	// rpcType
	Dim *string `json:"dim,omitempty" xml:"dim,omitempty"`
	// The filter operation type. Valid values:
	//
	// - eq: equal to
	//
	// - neq: not equal to
	//
	// - match: regex match
	//
	// - nmatch: regex not match
	//
	// example:
	//
	// eq
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value corresponding to the filter operation.
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

type AlertRuleQueryQueriesLabelFilters struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleQueryQueriesLabelFilters) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleQueryQueriesLabelFilters) GoString() string {
	return s.String()
}

func (s *AlertRuleQueryQueriesLabelFilters) GetName() *string {
	return s.Name
}

func (s *AlertRuleQueryQueriesLabelFilters) GetOperator() *string {
	return s.Operator
}

func (s *AlertRuleQueryQueriesLabelFilters) GetValue() *string {
	return s.Value
}

func (s *AlertRuleQueryQueriesLabelFilters) SetName(v string) *AlertRuleQueryQueriesLabelFilters {
	s.Name = &v
	return s
}

func (s *AlertRuleQueryQueriesLabelFilters) SetOperator(v string) *AlertRuleQueryQueriesLabelFilters {
	s.Operator = &v
	return s
}

func (s *AlertRuleQueryQueriesLabelFilters) SetValue(v string) *AlertRuleQueryQueriesLabelFilters {
	s.Value = &v
	return s
}

func (s *AlertRuleQueryQueriesLabelFilters) Validate() error {
	return dara.Validate(s)
}
