// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricMetaListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMetricMetaListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricMetaListResponseBody
	GetRequestId() *string
	SetResources(v *DescribeMetricMetaListResponseBodyResources) *DescribeMetricMetaListResponseBody
	GetResources() *DescribeMetricMetaListResponseBodyResources
	SetSuccess(v bool) *DescribeMetricMetaListResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *DescribeMetricMetaListResponseBody
	GetTotalCount() *string
}

type DescribeMetricMetaListResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0CCE0AF0-053C-4B13-A583-DC9A85785D49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration of the metrics in the resources.
	Resources *DescribeMetricMetaListResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMetricMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricMetaListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricMetaListResponseBody) GetResources() *DescribeMetricMetaListResponseBodyResources {
	return s.Resources
}

func (s *DescribeMetricMetaListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricMetaListResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeMetricMetaListResponseBody) SetCode(v string) *DescribeMetricMetaListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetMessage(v string) *DescribeMetricMetaListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetRequestId(v string) *DescribeMetricMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetResources(v *DescribeMetricMetaListResponseBodyResources) *DescribeMetricMetaListResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetSuccess(v bool) *DescribeMetricMetaListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetTotalCount(v string) *DescribeMetricMetaListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricMetaListResponseBodyResources struct {
	Resource []*DescribeMetricMetaListResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMetricMetaListResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResources) GetResource() []*DescribeMetricMetaListResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeMetricMetaListResponseBodyResources) SetResource(v []*DescribeMetricMetaListResponseBodyResourcesResource) *DescribeMetricMetaListResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricMetaListResponseBodyResourcesResource struct {
	// The metric description.
	//
	// example:
	//
	// CPUUtilization
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The monitoring dimensions of the resource. Multiple monitoring dimensions are separated with commas (,).
	//
	// example:
	//
	// instanceId
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The tags of the metric, including one or more JSON strings.
	//
	// Format: `[{"name":"tag key","value":"tag value"}]`. The `name` can be repeated. The following tags are available:
	//
	// 	- metricCategory: the category of the metric.
	//
	// 	- alertEnable: indicates whether to report alerts for the metric.
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
	// [{\\"name\\":\\"alertUnit\\",\\"value\\":\\"Bytes\\"},{\\"name\\":\\"minAlertPeriod\\",\\"value\\":\\"60\\"},{\\"name\\":\\"metricCategory\\",\\"value\\":\\"instanceId\\"},{\\"name\\":\\"instanceType\\",\\"value\\":\\"disaster\\"},{\\"name\\":\\"is_alarm\\",\\"value\\":\\"true\\"},{\\"name\\":\\"productCategory\\",\\"value\\":\\"kvstore_old\\"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The metric name.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// acs_kvstore
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The statistical periods of the metric. Multiple statistical periods are separated with commas (,).
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60,300
	Periods *string `json:"Periods,omitempty" xml:"Periods,omitempty"`
	// The statistical method. Multiple statistical methods are separated with commas (,).
	//
	// example:
	//
	// Average,Minimum,Maximum
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeMetricMetaListResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetDescription() *string {
	return s.Description
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetLabels() *string {
	return s.Labels
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetPeriods() *string {
	return s.Periods
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) GetUnit() *string {
	return s.Unit
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetDescription(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetDimensions(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetLabels(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetMetricName(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetNamespace(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetPeriods(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Periods = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetStatistics(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetUnit(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Unit = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
