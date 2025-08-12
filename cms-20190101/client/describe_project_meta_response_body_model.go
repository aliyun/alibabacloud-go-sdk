// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeProjectMetaResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeProjectMetaResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeProjectMetaResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeProjectMetaResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeProjectMetaResponseBody
	GetRequestId() *string
	SetResources(v *DescribeProjectMetaResponseBodyResources) *DescribeProjectMetaResponseBody
	GetResources() *DescribeProjectMetaResponseBodyResources
	SetSuccess(v bool) *DescribeProjectMetaResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeProjectMetaResponseBody
	GetTotal() *string
}

type DescribeProjectMetaResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C2061B2-3B1B-43BF-A4A4-C53426F479C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the cloud service.
	Resources *DescribeProjectMetaResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
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
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeProjectMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProjectMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProjectMetaResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeProjectMetaResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeProjectMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectMetaResponseBody) GetResources() *DescribeProjectMetaResponseBodyResources {
	return s.Resources
}

func (s *DescribeProjectMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectMetaResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeProjectMetaResponseBody) SetCode(v string) *DescribeProjectMetaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetMessage(v string) *DescribeProjectMetaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageNumber(v string) *DescribeProjectMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageSize(v string) *DescribeProjectMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetRequestId(v string) *DescribeProjectMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetResources(v *DescribeProjectMetaResponseBodyResources) *DescribeProjectMetaResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetSuccess(v bool) *DescribeProjectMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetTotal(v string) *DescribeProjectMetaResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectMetaResponseBodyResources struct {
	Resource []*DescribeProjectMetaResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeProjectMetaResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResources) GetResource() []*DescribeProjectMetaResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeProjectMetaResponseBodyResources) SetResource(v []*DescribeProjectMetaResponseBodyResourcesResource) *DescribeProjectMetaResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeProjectMetaResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectMetaResponseBodyResourcesResource struct {
	// The description.
	//
	// example:
	//
	// CDN
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tags. Tags are used to filter services.
	//
	// Tags are returned in the following format: `[{"name":"Tag key","value":"Tag value"}, {"name":"Tag key","value":"Tag value"}]`. The following tags are commonly used:
	//
	// 	- alertUnit: the unit of the metric value in alerts. If the unit is small, the original metric value may be too large. In this case, you can use the `alertUnit` tag to specify an appropriate unit. This tag is used in CloudMonitor.
	//
	// 	- minAlertPeriod: the minimum time interval to report a new alert. The interval at which monitoring data is reported. The value is usually 1 minute.
	//
	// 	- metricCategory: the service specification. Example: kvstore_sharding. Some Alibaba Cloud services have multiple specifications that are defined in the same namespace. This parameter is used to identify the specifications.
	//
	// 	- is_alarm: indicates whether an alert rule can be configured. We recommend that you do not use the special tags in the CloudMonitor console.
	//
	// example:
	//
	// [{"groupFlag":true}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The namespace of the cloud service. Format: `acs_Service name abbreviation`. For more information about namespaces, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_cdn
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeProjectMetaResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetDescription() *string {
	return s.Description
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetLabels() *string {
	return s.Labels
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetDescription(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetLabels(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetNamespace(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
