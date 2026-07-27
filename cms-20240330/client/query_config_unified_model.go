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
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// Deprecated
	CheckAfterDataComplete  *bool                          `json:"checkAfterDataComplete,omitempty" xml:"checkAfterDataComplete,omitempty"`
	Dimensions              []map[string]*string           `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	DurationSecs            *int64                         `json:"durationSecs,omitempty" xml:"durationSecs,omitempty"`
	EnableDataCompleteCheck *bool                          `json:"enableDataCompleteCheck,omitempty" xml:"enableDataCompleteCheck,omitempty"`
	EntityDomain            *string                        `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	EntityFields            []*EntityFields                `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	EntityFilters           []*EntityFilters               `json:"entityFilters,omitempty" xml:"entityFilters,omitempty" type:"Repeated"`
	EntityType              *string                        `json:"entityType,omitempty" xml:"entityType,omitempty"`
	Expr                    *string                        `json:"expr,omitempty" xml:"expr,omitempty"`
	FilterList              []*FilterList                  `json:"filterList,omitempty" xml:"filterList,omitempty" type:"Repeated"`
	FilterValues            []*PrometheusMetricFilterValue `json:"filterValues,omitempty" xml:"filterValues,omitempty" type:"Repeated"`
	GroupFieldList          []*string                      `json:"groupFieldList,omitempty" xml:"groupFieldList,omitempty" type:"Repeated"`
	GroupId                 *string                        `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupType               *string                        `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Joinings                []*Joinings                    `json:"joinings,omitempty" xml:"joinings,omitempty" type:"Repeated"`
	LabelFilters            []*LabelFilters                `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	LegacyRaw               *string                        `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	LegacyType              *string                        `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	LogSet                  *string                        `json:"logSet,omitempty" xml:"logSet,omitempty"`
	MeasureGroupKey         *string                        `json:"measureGroupKey,omitempty" xml:"measureGroupKey,omitempty"`
	MeasureList             []*MeasureList                 `json:"measureList,omitempty" xml:"measureList,omitempty" type:"Repeated"`
	Metric                  *string                        `json:"metric,omitempty" xml:"metric,omitempty"`
	MetricGroupId           *string                        `json:"metricGroupId,omitempty" xml:"metricGroupId,omitempty"`
	MetricId                *string                        `json:"metricId,omitempty" xml:"metricId,omitempty"`
	// Deprecated
	MetricIds   []*string                     `json:"metricIds,omitempty" xml:"metricIds,omitempty" type:"Repeated"`
	MetricSet   *string                       `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	Namespace   *string                       `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OffsetSecs  *int64                        `json:"offsetSecs,omitempty" xml:"offsetSecs,omitempty"`
	ParamValues []*PrometheusMetricParamValue `json:"paramValues,omitempty" xml:"paramValues,omitempty" type:"Repeated"`
	// Deprecated
	PromQl        *string    `json:"promQl,omitempty" xml:"promQl,omitempty"`
	Queries       []*Queries `json:"queries,omitempty" xml:"queries,omitempty" type:"Repeated"`
	RelationType  *string    `json:"relationType,omitempty" xml:"relationType,omitempty"`
	ServiceIdList []*string  `json:"serviceIdList,omitempty" xml:"serviceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	WindowSecs *int64  `json:"windowSecs,omitempty" xml:"windowSecs,omitempty"`
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
