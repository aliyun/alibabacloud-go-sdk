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
	// 是否启用数据完整性检查
	EnableDataCompleteCheck *bool `json:"enableDataCompleteCheck,omitempty" xml:"enableDataCompleteCheck,omitempty"`
	// 实体所属域
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// 需要附带返回的实体字段
	EntityFields []*UmodelEntityField `json:"entityFields,omitempty" xml:"entityFields,omitempty" type:"Repeated"`
	// 实体过滤列表
	EntityFilters []*UmodelEntityFilter `json:"entityFilters,omitempty" xml:"entityFilters,omitempty" type:"Repeated"`
	// 实体类型
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// APM 过滤条件列表
	FilterList []*ApmFilterConfig `json:"filterList,omitempty" xml:"filterList,omitempty" type:"Repeated"`
	// 标签过滤条件
	LabelFilters []*UmodelLabelFilter `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	// APM 度量配置列表
	MeasureList []*ApmMeasureConfig `json:"measureList,omitempty" xml:"measureList,omitempty" type:"Repeated"`
	// 指标名称（type=UMODEL_METRICSET_QUERY）
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// 指标集名称（type=UMODEL_METRICSET_QUERY）
	MetricSet *string `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	// Prometheus 查询语句（type=PROMETHEUS_SINGLE_QUERY）
	PromQl *string `json:"promQl,omitempty" xml:"promQl,omitempty"`
	// 服务 ID 列表（type=APM_MULTI_QUERY）
	ServiceIdList []*string `json:"serviceIdList,omitempty" xml:"serviceIdList,omitempty" type:"Repeated"`
	// 查询类型
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
