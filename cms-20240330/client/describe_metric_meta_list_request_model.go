// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeywords(v string) *DescribeMetricMetaListRequest
	GetKeywords() *string
	SetLabels(v []*DescribeMetricMetaListRequestLabels) *DescribeMetricMetaListRequest
	GetLabels() []*DescribeMetricMetaListRequestLabels
	SetMetaFormat(v string) *DescribeMetricMetaListRequest
	GetMetaFormat() *string
	SetMetricName(v string) *DescribeMetricMetaListRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricMetaListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeMetricMetaListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMetricMetaListRequest
	GetPageSize() *int32
}

type DescribeMetricMetaListRequest struct {
	// The keyword.
	//
	// example:
	//
	// 实例理论
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// Filters resources by label. The following labels are available:
	//
	// - metricCategory: the metric category description.
	//
	// - alertEnable: specifies whether alerting is required.
	//
	// - alertUnit: the recommended alert unit.
	//
	// - unitFactor: the unit conversion factor.
	//
	// - minAlertPeriod: the minimum alert period.
	//
	// - productCategory: the service type category.
	Labels []*DescribeMetricMetaListRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The metadata source. Valid values:
	//
	// - CMS: CloudMonitor Basic monitoring metrics.
	//
	// - PROM_BASIC: Prometheus CloudMonitor basic monitoring metrics.
	//
	// example:
	//
	// CMS
	MetaFormat *string `json:"metaFormat,omitempty" xml:"metaFormat,omitempty"`
	// The metric name.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The namespace, which is used to distinguish between services.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 2000.
	//
	// example:
	//
	// 2000
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s DescribeMetricMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *DescribeMetricMetaListRequest) GetLabels() []*DescribeMetricMetaListRequestLabels {
	return s.Labels
}

func (s *DescribeMetricMetaListRequest) GetMetaFormat() *string {
	return s.MetaFormat
}

func (s *DescribeMetricMetaListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricMetaListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricMetaListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetricMetaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetricMetaListRequest) SetKeywords(v string) *DescribeMetricMetaListRequest {
	s.Keywords = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetLabels(v []*DescribeMetricMetaListRequestLabels) *DescribeMetricMetaListRequest {
	s.Labels = v
	return s
}

func (s *DescribeMetricMetaListRequest) SetMetaFormat(v string) *DescribeMetricMetaListRequest {
	s.MetaFormat = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetMetricName(v string) *DescribeMetricMetaListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetNamespace(v string) *DescribeMetricMetaListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetPageNumber(v int32) *DescribeMetricMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetPageSize(v int32) *DescribeMetricMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricMetaListRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricMetaListRequestLabels struct {
	// The label name.
	//
	// example:
	//
	// productCategory
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The label value.
	//
	// example:
	//
	// ecs
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeMetricMetaListRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListRequestLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListRequestLabels) GetName() *string {
	return s.Name
}

func (s *DescribeMetricMetaListRequestLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricMetaListRequestLabels) SetName(v string) *DescribeMetricMetaListRequestLabels {
	s.Name = &v
	return s
}

func (s *DescribeMetricMetaListRequestLabels) SetValue(v string) *DescribeMetricMetaListRequestLabels {
	s.Value = &v
	return s
}

func (s *DescribeMetricMetaListRequestLabels) Validate() error {
	return dara.Validate(s)
}
