// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v string) *QueryConfigUnified
	GetAggregate() *string
	SetCheckAfterDataComplete(v bool) *QueryConfigUnified
	GetCheckAfterDataComplete() *bool
	SetDimensions(v []map[string]*string) *QueryConfigUnified
	GetDimensions() []map[string]*string
	SetDurationSecs(v int64) *QueryConfigUnified
	GetDurationSecs() *int64
	SetEnableDataCompleteCheck(v bool) *QueryConfigUnified
	GetEnableDataCompleteCheck() *bool
	SetEntityDomain(v string) *QueryConfigUnified
	GetEntityDomain() *string
	SetEntityFields(v []*EntityFields) *QueryConfigUnified
	GetEntityFields() []*EntityFields
	SetEntityFilters(v []*EntityFilters) *QueryConfigUnified
	GetEntityFilters() []*EntityFilters
	SetEntityType(v string) *QueryConfigUnified
	GetEntityType() *string
	SetExpr(v string) *QueryConfigUnified
	GetExpr() *string
	SetFilterList(v []*FilterList) *QueryConfigUnified
	GetFilterList() []*FilterList
	SetFilterValues(v []*PrometheusMetricFilterValue) *QueryConfigUnified
	GetFilterValues() []*PrometheusMetricFilterValue
	SetGroupFieldList(v []*string) *QueryConfigUnified
	GetGroupFieldList() []*string
	SetGroupId(v string) *QueryConfigUnified
	GetGroupId() *string
	SetGroupType(v string) *QueryConfigUnified
	GetGroupType() *string
	SetJoinings(v []*Joinings) *QueryConfigUnified
	GetJoinings() []*Joinings
	SetLabelFilters(v []*LabelFilters) *QueryConfigUnified
	GetLabelFilters() []*LabelFilters
	SetLegacyRaw(v string) *QueryConfigUnified
	GetLegacyRaw() *string
	SetLegacyType(v string) *QueryConfigUnified
	GetLegacyType() *string
	SetLogSet(v string) *QueryConfigUnified
	GetLogSet() *string
	SetMeasureGroupKey(v string) *QueryConfigUnified
	GetMeasureGroupKey() *string
	SetMeasureList(v []*MeasureList) *QueryConfigUnified
	GetMeasureList() []*MeasureList
	SetMetric(v string) *QueryConfigUnified
	GetMetric() *string
	SetMetricGroupId(v string) *QueryConfigUnified
	GetMetricGroupId() *string
	SetMetricId(v string) *QueryConfigUnified
	GetMetricId() *string
	SetMetricIds(v []*string) *QueryConfigUnified
	GetMetricIds() []*string
	SetMetricSet(v string) *QueryConfigUnified
	GetMetricSet() *string
	SetNamespace(v string) *QueryConfigUnified
	GetNamespace() *string
	SetOffsetSecs(v int64) *QueryConfigUnified
	GetOffsetSecs() *int64
	SetParamValues(v []*PrometheusMetricParamValue) *QueryConfigUnified
	GetParamValues() []*PrometheusMetricParamValue
	SetPromQl(v string) *QueryConfigUnified
	GetPromQl() *string
	SetQueries(v []*Queries) *QueryConfigUnified
	GetQueries() []*Queries
	SetRelationType(v string) *QueryConfigUnified
	GetRelationType() *string
	SetServiceIdList(v []*string) *QueryConfigUnified
	GetServiceIdList() []*string
	SetType(v string) *QueryConfigUnified
	GetType() *string
	SetWindowSecs(v int64) *QueryConfigUnified
	GetWindowSecs() *int64
}

type QueryConfigUnified struct {
	// The aggregation function (used when type=UMODEL_METRICSET_QUERY / UMODEL_LOGSET_QUERY).
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Specifies whether to perform alert detection only after data is complete (originally used when type=PROMETHEUS_MULTI_QUERY). This field overlaps with enableDataCompleteCheck. Using this field in write path returns 400.
	CheckAfterDataComplete *bool `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	// The dimension list (used when type=CLOUD_MONITORING_QUERY. Each dimension is a key/value string mapping).
	Dimensions []map[string]*string `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	// The duration in seconds (used when type=PROMETHEUS_MULTI_QUERY).
	DurationSecs *int64 `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	// Indicates whether the data integrity check is enabled (used when type=PROMETHEUS_SINGLE_QUERY / PROMETHEUS_MULTI_QUERY / PROMETHEUS_PREDEFINED_METRIC_QUERY / PROMETHEUS_METRIC_GROUP_QUERY [deprecated]).
	EnableDataCompleteCheck *bool `json:"enableDataCompleteCheck,omitempty" xml:"enableDataCompleteCheck,omitempty"`
	// The entity domain (used when type=UMODEL_METRICSET_QUERY / UMODEL_METRICSET_MULTI_QUERY / UMODEL_LOGSET_QUERY. Works with entityType/entityFilters to locate UModel entities).
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The entity fields to include in the response (used when type=UMODEL_METRICSET_QUERY / UMODEL_METRICSET_MULTI_QUERY / UMODEL_LOGSET_QUERY).
	EntityFields []*EntityFields `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	// The entity filter list (used when type=UMODEL_METRICSET_QUERY / UMODEL_METRICSET_MULTI_QUERY / UMODEL_LOGSET_QUERY).
	EntityFilters []*EntityFilters `json:"entityFilters,omitempty" xml:"entityFilters,omitempty" type:"Repeated"`
	// The entity type (used when type=UMODEL_METRICSET_QUERY / UMODEL_METRICSET_MULTI_QUERY / UMODEL_LOGSET_QUERY).
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The query expression or SPL statement. Recommended when type=PROMETHEUS_SINGLE_QUERY. Optional when type=UMODEL_METRICSET_QUERY for custom SPL. Required when type=UMODEL_LOGSET_QUERY, where an SPL query statement must be provided (the service layer enforces this requirement).
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The APM filter condition list.
	FilterList []*FilterList `json:"filterList,omitempty" xml:"filterList,omitempty" type:"Repeated"`
	// The list of predefined metric filter values (used when type=PROMETHEUS_PREDEFINED_METRIC_QUERY / PROMETHEUS_METRIC_GROUP_QUERY [deprecated]).
	FilterValues []*PrometheusMetricFilterValue `json:"filterValues,omitempty" xml:"filterValues,omitempty" type:"Repeated"`
	// The group field list (used when type=SLS_MULTI_QUERY and groupType=custom).
	GroupFieldList []*string `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	// The resource group ID (used when type=CLOUD_MONITORING_QUERY and relationType=GROUP).
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The grouping policy (used when type=SLS_MULTI_QUERY): none / label / custom.
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The join list (used when type=SLS_MULTI_QUERY. Maximum of 2: joinings[0] corresponds to the set operation between query 0 and query 1. joinings[1] corresponds to the set operation between query 1 and query 2).
	Joinings []*Joinings `json:"joinings,omitempty" xml:"joinings,omitempty" type:"Repeated"`
	// The label filter conditions (used when type=UMODEL_METRICSET_QUERY. For UMODEL_METRICSET_MULTI_QUERY, place labelFilters in each queries[*] entry).
	LabelFilters []*LabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	// The original V1 query JSON string returned as a fallback when type=UNKNOWN_QUERY and read path parsing fails (contains the field values that triggered the failure, such as filter.operator=ABC). The frontend displays this field as read-only when it is not empty.
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// Returned when type=UNKNOWN_QUERY, indicating that this rule cannot be edited through the new API. Submit a ticket to contact the CloudMonitor team.
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The log set name (used when type=UMODEL_LOGSET_QUERY).
	LogSet *string `json:"logSet,omitempty" xml:"logSet,omitempty"`
	// The measure group key (optional when type=APM_MULTI_QUERY, corresponds to V1 alertMetricInput.groupKey).
	MeasureGroupKey *string `json:"measureGroupKey,omitempty" xml:"measureGroupKey,omitempty"`
	// The APM measure configuration list.
	MeasureList []*MeasureList `json:"measureList,omitempty" xml:"measureList,omitempty" type:"Repeated"`
	// The metric name (required when type=UMODEL_METRICSET_QUERY. Required when type=CLOUD_MONITORING_QUERY, used together with namespace to uniquely identify CloudMonitor monitoring metrics).
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// The metric group ID (used when type=PROMETHEUS_PREDEFINED_METRIC_QUERY / PROMETHEUS_METRIC_GROUP_QUERY [deprecated]).
	MetricGroupId *string `json:"metricGroupId,omitempty" xml:"metricGroupId,omitempty"`
	// The predefined metric ID (used when type=PROMETHEUS_PREDEFINED_METRIC_QUERY).
	MetricId *string `json:"metricId,omitempty" xml:"metricId,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The list of predefined metric IDs (originally used when type=PROMETHEUS_METRIC_GROUP_QUERY). This query type is deprecated. Write path returns 400.
	MetricIds []*string `json:"metricIds,omitempty" xml:"metricIds,omitempty" type:"Repeated"`
	// The metric set name (used when type=UMODEL_METRICSET_QUERY).
	MetricSet *string `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	// The CloudMonitor namespace (Alibaba Cloud service name, used when type=CLOUD_MONITORING_QUERY).
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The query time offset in seconds (used when type=UMODEL_METRICSET_QUERY / UMODEL_LOGSET_QUERY). Works with windowSecs to implement an offset query over the range [T - windowSecs - offsetSecs, T - offsetSecs]. Valid range: [0, 86400].
	OffsetSecs *int64 `json:"offsetSecs,omitempty" xml:"offsetSecs,omitempty"`
	// The list of predefined metric parameter values (used when type=PROMETHEUS_PREDEFINED_METRIC_QUERY / PROMETHEUS_METRIC_GROUP_QUERY [deprecated]).
	ParamValues []*PrometheusMetricParamValue `json:"paramValues,omitempty" xml:"paramValues,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- The legacy Prometheus query statement field. Use expr instead. This field is retained for backward compatibility. The backend automatically normalizes it to expr.
	PromQl *string `json:"promQl,omitempty" xml:"promQl,omitempty"`
	// The subquery list (polymorphic by type): when type=SLS_MULTI_QUERY, each entry is a SlsNamedQueryEntry (timeUnit/start/end/window/expr). When type=PROMETHEUS_MULTI_QUERY, each entry is a PrometheusNamedQueryEntry (name/expr). When type=UMODEL_METRICSET_MULTI_QUERY, each entry is a MetricSetNamedQueryEntry.
	Queries []*Queries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	// The resource relation type (used when type=CLOUD_MONITORING_QUERY).
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// The list of service IDs (used when type=APM_MULTI_QUERY).
	ServiceIdList []*string `json:"serviceIdList,omitempty" xml:"serviceIdList,omitempty" type:"Repeated"`
	// The query type. Valid values and associated fields: PROMETHEUS_SINGLE_QUERY (required: expr. Optional: enableDataCompleteCheck). PROMETHEUS_PREDEFINED_METRIC_QUERY (required: metricGroupId, metricId. Optional: paramValues, filterValues, enableDataCompleteCheck). PROMETHEUS_METRIC_GROUP_QUERY ([deprecated] required: metricGroupId, metricIds. Optional: paramValues, filterValues, enableDataCompleteCheck. Write path returns 400). UMODEL_METRICSET_QUERY (required: metricSet, metric, windowSecs, aggregate. Optional: expr, entityDomain/entityType/entityFilters, labelFilters, entityFields, offsetSecs). UMODEL_METRICSET_MULTI_QUERY (required: queries[*]. Optional: entityDomain/entityType/entityFilters, windowSecs, offsetSecs, aggregate). UMODEL_LOGSET_QUERY (required: logSet, expr, windowSecs, aggregate. Optional: entityDomain/entityType/entityFilters, labelFilters, offsetSecs). APM_MULTI_QUERY (required: serviceIdList, measureList. Optional: filterList, measureGroupKey). CLOUD_MONITORING_QUERY (required: namespace, metric, relationType. When relationType=INSTANCE, dimensions is required. When relationType=GROUP, groupId is required. When relationType=USER, leave both empty). UNKNOWN_QUERY (read-only fallback. Do not use in write path). Do not use non-enumerated values (such as CMS_BASIC_QUERY/SLS_QUERY). The backend returns Invalidtype 400.
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The aggregation time window in seconds (used when type=UMODEL_METRICSET_QUERY / UMODEL_LOGSET_QUERY). Valid range: [60, 86400].
	WindowSecs *int64 `json:"windowSecs,omitempty" xml:"windowSecs,omitempty"`
}

func (s QueryConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigUnified) GoString() string {
	return s.String()
}

func (s *QueryConfigUnified) GetAggregate() *string {
	return s.Aggregate
}

func (s *QueryConfigUnified) GetCheckAfterDataComplete() *bool {
	return s.CheckAfterDataComplete
}

func (s *QueryConfigUnified) GetDimensions() []map[string]*string {
	return s.Dimensions
}

func (s *QueryConfigUnified) GetDurationSecs() *int64 {
	return s.DurationSecs
}

func (s *QueryConfigUnified) GetEnableDataCompleteCheck() *bool {
	return s.EnableDataCompleteCheck
}

func (s *QueryConfigUnified) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *QueryConfigUnified) GetEntityFields() []*EntityFields {
	return s.EntityFields
}

func (s *QueryConfigUnified) GetEntityFilters() []*EntityFilters {
	return s.EntityFilters
}

func (s *QueryConfigUnified) GetEntityType() *string {
	return s.EntityType
}

func (s *QueryConfigUnified) GetExpr() *string {
	return s.Expr
}

func (s *QueryConfigUnified) GetFilterList() []*FilterList {
	return s.FilterList
}

func (s *QueryConfigUnified) GetFilterValues() []*PrometheusMetricFilterValue {
	return s.FilterValues
}

func (s *QueryConfigUnified) GetGroupFieldList() []*string {
	return s.GroupFieldList
}

func (s *QueryConfigUnified) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryConfigUnified) GetGroupType() *string {
	return s.GroupType
}

func (s *QueryConfigUnified) GetJoinings() []*Joinings {
	return s.Joinings
}

func (s *QueryConfigUnified) GetLabelFilters() []*LabelFilters {
	return s.LabelFilters
}

func (s *QueryConfigUnified) GetLegacyRaw() *string {
	return s.LegacyRaw
}

func (s *QueryConfigUnified) GetLegacyType() *string {
	return s.LegacyType
}

func (s *QueryConfigUnified) GetLogSet() *string {
	return s.LogSet
}

func (s *QueryConfigUnified) GetMeasureGroupKey() *string {
	return s.MeasureGroupKey
}

func (s *QueryConfigUnified) GetMeasureList() []*MeasureList {
	return s.MeasureList
}

func (s *QueryConfigUnified) GetMetric() *string {
	return s.Metric
}

func (s *QueryConfigUnified) GetMetricGroupId() *string {
	return s.MetricGroupId
}

func (s *QueryConfigUnified) GetMetricId() *string {
	return s.MetricId
}

func (s *QueryConfigUnified) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *QueryConfigUnified) GetMetricSet() *string {
	return s.MetricSet
}

func (s *QueryConfigUnified) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryConfigUnified) GetOffsetSecs() *int64 {
	return s.OffsetSecs
}

func (s *QueryConfigUnified) GetParamValues() []*PrometheusMetricParamValue {
	return s.ParamValues
}

func (s *QueryConfigUnified) GetPromQl() *string {
	return s.PromQl
}

func (s *QueryConfigUnified) GetQueries() []*Queries {
	return s.Queries
}

func (s *QueryConfigUnified) GetRelationType() *string {
	return s.RelationType
}

func (s *QueryConfigUnified) GetServiceIdList() []*string {
	return s.ServiceIdList
}

func (s *QueryConfigUnified) GetType() *string {
	return s.Type
}

func (s *QueryConfigUnified) GetWindowSecs() *int64 {
	return s.WindowSecs
}

func (s *QueryConfigUnified) SetAggregate(v string) *QueryConfigUnified {
	s.Aggregate = &v
	return s
}

func (s *QueryConfigUnified) SetCheckAfterDataComplete(v bool) *QueryConfigUnified {
	s.CheckAfterDataComplete = &v
	return s
}

func (s *QueryConfigUnified) SetDimensions(v []map[string]*string) *QueryConfigUnified {
	s.Dimensions = v
	return s
}

func (s *QueryConfigUnified) SetDurationSecs(v int64) *QueryConfigUnified {
	s.DurationSecs = &v
	return s
}

func (s *QueryConfigUnified) SetEnableDataCompleteCheck(v bool) *QueryConfigUnified {
	s.EnableDataCompleteCheck = &v
	return s
}

func (s *QueryConfigUnified) SetEntityDomain(v string) *QueryConfigUnified {
	s.EntityDomain = &v
	return s
}

func (s *QueryConfigUnified) SetEntityFields(v []*EntityFields) *QueryConfigUnified {
	s.EntityFields = v
	return s
}

func (s *QueryConfigUnified) SetEntityFilters(v []*EntityFilters) *QueryConfigUnified {
	s.EntityFilters = v
	return s
}

func (s *QueryConfigUnified) SetEntityType(v string) *QueryConfigUnified {
	s.EntityType = &v
	return s
}

func (s *QueryConfigUnified) SetExpr(v string) *QueryConfigUnified {
	s.Expr = &v
	return s
}

func (s *QueryConfigUnified) SetFilterList(v []*FilterList) *QueryConfigUnified {
	s.FilterList = v
	return s
}

func (s *QueryConfigUnified) SetFilterValues(v []*PrometheusMetricFilterValue) *QueryConfigUnified {
	s.FilterValues = v
	return s
}

func (s *QueryConfigUnified) SetGroupFieldList(v []*string) *QueryConfigUnified {
	s.GroupFieldList = v
	return s
}

func (s *QueryConfigUnified) SetGroupId(v string) *QueryConfigUnified {
	s.GroupId = &v
	return s
}

func (s *QueryConfigUnified) SetGroupType(v string) *QueryConfigUnified {
	s.GroupType = &v
	return s
}

func (s *QueryConfigUnified) SetJoinings(v []*Joinings) *QueryConfigUnified {
	s.Joinings = v
	return s
}

func (s *QueryConfigUnified) SetLabelFilters(v []*LabelFilters) *QueryConfigUnified {
	s.LabelFilters = v
	return s
}

func (s *QueryConfigUnified) SetLegacyRaw(v string) *QueryConfigUnified {
	s.LegacyRaw = &v
	return s
}

func (s *QueryConfigUnified) SetLegacyType(v string) *QueryConfigUnified {
	s.LegacyType = &v
	return s
}

func (s *QueryConfigUnified) SetLogSet(v string) *QueryConfigUnified {
	s.LogSet = &v
	return s
}

func (s *QueryConfigUnified) SetMeasureGroupKey(v string) *QueryConfigUnified {
	s.MeasureGroupKey = &v
	return s
}

func (s *QueryConfigUnified) SetMeasureList(v []*MeasureList) *QueryConfigUnified {
	s.MeasureList = v
	return s
}

func (s *QueryConfigUnified) SetMetric(v string) *QueryConfigUnified {
	s.Metric = &v
	return s
}

func (s *QueryConfigUnified) SetMetricGroupId(v string) *QueryConfigUnified {
	s.MetricGroupId = &v
	return s
}

func (s *QueryConfigUnified) SetMetricId(v string) *QueryConfigUnified {
	s.MetricId = &v
	return s
}

func (s *QueryConfigUnified) SetMetricIds(v []*string) *QueryConfigUnified {
	s.MetricIds = v
	return s
}

func (s *QueryConfigUnified) SetMetricSet(v string) *QueryConfigUnified {
	s.MetricSet = &v
	return s
}

func (s *QueryConfigUnified) SetNamespace(v string) *QueryConfigUnified {
	s.Namespace = &v
	return s
}

func (s *QueryConfigUnified) SetOffsetSecs(v int64) *QueryConfigUnified {
	s.OffsetSecs = &v
	return s
}

func (s *QueryConfigUnified) SetParamValues(v []*PrometheusMetricParamValue) *QueryConfigUnified {
	s.ParamValues = v
	return s
}

func (s *QueryConfigUnified) SetPromQl(v string) *QueryConfigUnified {
	s.PromQl = &v
	return s
}

func (s *QueryConfigUnified) SetQueries(v []*Queries) *QueryConfigUnified {
	s.Queries = v
	return s
}

func (s *QueryConfigUnified) SetRelationType(v string) *QueryConfigUnified {
	s.RelationType = &v
	return s
}

func (s *QueryConfigUnified) SetServiceIdList(v []*string) *QueryConfigUnified {
	s.ServiceIdList = v
	return s
}

func (s *QueryConfigUnified) SetType(v string) *QueryConfigUnified {
	s.Type = &v
	return s
}

func (s *QueryConfigUnified) SetWindowSecs(v int64) *QueryConfigUnified {
	s.WindowSecs = &v
	return s
}

func (s *QueryConfigUnified) Validate() error {
	if s.EntityFields != nil {
		for _, item := range s.EntityFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntityFilters != nil {
		for _, item := range s.EntityFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterList != nil {
		for _, item := range s.FilterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterValues != nil {
		for _, item := range s.FilterValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Joinings != nil {
		for _, item := range s.Joinings {
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
	if s.MeasureList != nil {
		for _, item := range s.MeasureList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ParamValues != nil {
		for _, item := range s.ParamValues {
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
	return nil
}
