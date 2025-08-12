// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *DescribeMetricMetaListRequest
	GetLabels() *string
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
	// The tags for filtering metrics. Specify a JSON string.
	//
	// Format: ` [{"name":"tag key","value":"tag value"},{"name":"tag key","value":"tag value"}]  `. The following tags are available:
	//
	// 	- metricCategory: the category of the metric.
	//
	// 	- alertEnable: specifies whether to report alerts for the metric.
	//
	// 	- alertUnit: the unit of the metric in the alerts.
	//
	// 	- unitFactor: the factor for metric unit conversion.
	//
	// 	- minAlertPeriod: the minimum interval at which the alert is reported.
	//
	// 	- productCategory: the category of the service.
	//
	// example:
	//
	// [{"name":"productCategory","value":"kvstore_old"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The metric name. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_kvstore
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMetricMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListRequest) GetLabels() *string {
	return s.Labels
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

func (s *DescribeMetricMetaListRequest) SetLabels(v string) *DescribeMetricMetaListRequest {
	s.Labels = &v
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
	return dara.Validate(s)
}
