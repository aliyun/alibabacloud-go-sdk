// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetEnableDataCompleteCheck(v bool) *QueryConfigUnified
	GetEnableDataCompleteCheck() *bool
	SetEntityDomain(v string) *QueryConfigUnified
	GetEntityDomain() *string
	SetEntityFields(v []*UmodelEntityField) *QueryConfigUnified
	GetEntityFields() []*UmodelEntityField
	SetEntityFilters(v []*UmodelEntityFilter) *QueryConfigUnified
	GetEntityFilters() []*UmodelEntityFilter
	SetEntityType(v string) *QueryConfigUnified
	GetEntityType() *string
	SetExpr(v string) *QueryConfigUnified
	GetExpr() *string
	SetFilterList(v []*ApmFilterConfig) *QueryConfigUnified
	GetFilterList() []*ApmFilterConfig
	SetLabelFilters(v []*UmodelLabelFilter) *QueryConfigUnified
	GetLabelFilters() []*UmodelLabelFilter
	SetMeasureList(v []*ApmMeasureConfig) *QueryConfigUnified
	GetMeasureList() []*ApmMeasureConfig
	SetMetric(v string) *QueryConfigUnified
	GetMetric() *string
	SetMetricSet(v string) *QueryConfigUnified
	GetMetricSet() *string
	SetPromQl(v string) *QueryConfigUnified
	GetPromQl() *string
	SetServiceIdList(v []*string) *QueryConfigUnified
	GetServiceIdList() []*string
	SetType(v string) *QueryConfigUnified
	GetType() *string
}

type QueryConfigUnified struct {
	// Specifies whether to check for data completeness. A value of `true` enables the check.
	EnableDataCompleteCheck *bool `json:"enableDataCompleteCheck,omitempty" xml:"enableDataCompleteCheck,omitempty"`
	// Specifies the domain of the entity, such as `acs` for Alibaba Cloud services.
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// A list of entity fields to include in the response.
	EntityFields []*UmodelEntityField `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	// A list of filters for selecting specific entities.
	EntityFilters []*UmodelEntityFilter `json:"entityFilters,omitempty" xml:"entityFilters,omitempty" type:"Repeated"`
	// Specifies the type of the entity, such as `EcsInstance`.
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// Specifies the expression to post-process query results.
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// A list of Application Performance Monitoring (APM) filter configurations.
	FilterList []*ApmFilterConfig `json:"filterList,omitempty" xml:"filterList,omitempty" type:"Repeated"`
	// A list of filters that match labels.
	LabelFilters []*UmodelLabelFilter `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	// A list of APM measure configurations.
	MeasureList []*ApmMeasureConfig `json:"measureList,omitempty" xml:"measureList,omitempty" type:"Repeated"`
	// Specifies the name of the metric to query.
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// Specifies the metric set that contains the metric.
	MetricSet *string `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	// Specifies the query string in Prometheus Query Language (PromQL).
	PromQl *string `json:"promQl,omitempty" xml:"promQl,omitempty"`
	// A list of service IDs to query.
	ServiceIdList []*string `json:"serviceIdList,omitempty" xml:"serviceIdList,omitempty" type:"Repeated"`
	// The query type.
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigUnified) GoString() string {
	return s.String()
}

func (s *QueryConfigUnified) GetEnableDataCompleteCheck() *bool {
	return s.EnableDataCompleteCheck
}

func (s *QueryConfigUnified) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *QueryConfigUnified) GetEntityFields() []*UmodelEntityField {
	return s.EntityFields
}

func (s *QueryConfigUnified) GetEntityFilters() []*UmodelEntityFilter {
	return s.EntityFilters
}

func (s *QueryConfigUnified) GetEntityType() *string {
	return s.EntityType
}

func (s *QueryConfigUnified) GetExpr() *string {
	return s.Expr
}

func (s *QueryConfigUnified) GetFilterList() []*ApmFilterConfig {
	return s.FilterList
}

func (s *QueryConfigUnified) GetLabelFilters() []*UmodelLabelFilter {
	return s.LabelFilters
}

func (s *QueryConfigUnified) GetMeasureList() []*ApmMeasureConfig {
	return s.MeasureList
}

func (s *QueryConfigUnified) GetMetric() *string {
	return s.Metric
}

func (s *QueryConfigUnified) GetMetricSet() *string {
	return s.MetricSet
}

func (s *QueryConfigUnified) GetPromQl() *string {
	return s.PromQl
}

func (s *QueryConfigUnified) GetServiceIdList() []*string {
	return s.ServiceIdList
}

func (s *QueryConfigUnified) GetType() *string {
	return s.Type
}

func (s *QueryConfigUnified) SetEnableDataCompleteCheck(v bool) *QueryConfigUnified {
	s.EnableDataCompleteCheck = &v
	return s
}

func (s *QueryConfigUnified) SetEntityDomain(v string) *QueryConfigUnified {
	s.EntityDomain = &v
	return s
}

func (s *QueryConfigUnified) SetEntityFields(v []*UmodelEntityField) *QueryConfigUnified {
	s.EntityFields = v
	return s
}

func (s *QueryConfigUnified) SetEntityFilters(v []*UmodelEntityFilter) *QueryConfigUnified {
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

func (s *QueryConfigUnified) SetFilterList(v []*ApmFilterConfig) *QueryConfigUnified {
	s.FilterList = v
	return s
}

func (s *QueryConfigUnified) SetLabelFilters(v []*UmodelLabelFilter) *QueryConfigUnified {
	s.LabelFilters = v
	return s
}

func (s *QueryConfigUnified) SetMeasureList(v []*ApmMeasureConfig) *QueryConfigUnified {
	s.MeasureList = v
	return s
}

func (s *QueryConfigUnified) SetMetric(v string) *QueryConfigUnified {
	s.Metric = &v
	return s
}

func (s *QueryConfigUnified) SetMetricSet(v string) *QueryConfigUnified {
	s.MetricSet = &v
	return s
}

func (s *QueryConfigUnified) SetPromQl(v string) *QueryConfigUnified {
	s.PromQl = &v
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
	return nil
}
