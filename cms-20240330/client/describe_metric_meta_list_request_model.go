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
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// The labels used to filter resources. The following labels are supported:
	//
	// - `metricCategory`: The metric category.
	//
	// - `alertEnable`: Indicates whether to enable alerts.
	//
	// - `alertUnit`: The recommended unit for alerts.
	//
	// - `unitFactor`: The unit conversion factor.
	//
	// - `minAlertPeriod`: The minimum alert period.
	//
	// - `productCategory`: The product category.
	Labels []*DescribeMetricMetaListRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The source of the metadata. Valid values: `CMS` for CloudMonitor metrics and `PROM_BASIC` for basic Prometheus metrics.
	//
	// example:
	//
	// CMS
	MetaFormat *string `json:"metaFormat,omitempty" xml:"metaFormat,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The namespace of the product.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: `2000`.
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
	// The key of the label.
	//
	// example:
	//
	// productCategory
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the label.
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
