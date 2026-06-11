// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPrometheusViewsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusViewsResponseBody
	GetNextToken() *string
	SetPrometheusViews(v []*ListPrometheusViewsResponseBodyPrometheusViews) *ListPrometheusViewsResponseBody
	GetPrometheusViews() []*ListPrometheusViewsResponseBodyPrometheusViews
	SetRequestId(v string) *ListPrometheusViewsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPrometheusViewsResponseBody
	GetTotalCount() *int32
}

type ListPrometheusViewsResponseBody struct {
	// The maximum number of records returned.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of Prometheus view instances.
	PrometheusViews []*ListPrometheusViewsResponseBodyPrometheusViews `json:"prometheusViews,omitempty" xml:"prometheusViews,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 66
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPrometheusViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusViewsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusViewsResponseBody) GetPrometheusViews() []*ListPrometheusViewsResponseBodyPrometheusViews {
	return s.PrometheusViews
}

func (s *ListPrometheusViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusViewsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPrometheusViewsResponseBody) SetMaxResults(v int32) *ListPrometheusViewsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusViewsResponseBody) SetNextToken(v string) *ListPrometheusViewsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusViewsResponseBody) SetPrometheusViews(v []*ListPrometheusViewsResponseBodyPrometheusViews) *ListPrometheusViewsResponseBody {
	s.PrometheusViews = v
	return s
}

func (s *ListPrometheusViewsResponseBody) SetRequestId(v string) *ListPrometheusViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusViewsResponseBody) SetTotalCount(v int32) *ListPrometheusViewsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrometheusViewsResponseBody) Validate() error {
	if s.PrometheusViews != nil {
		for _, item := range s.PrometheusViews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusViewsResponseBodyPrometheusViews struct {
	// The time when the instance was created. The time is in UTC and in the \\`yyyy-MM-ddTHH:mmZ\\` format.
	//
	// example:
	//
	// 2025-07-12T02:18:36Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The instance type. Valid values are \\`prom-view\\` for a new-version aggregation view and \\`global-view\\` for an old-version aggregation view.
	//
	// example:
	//
	// prom-view
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The billing method. The value is fixed to \\`FREE\\`.
	//
	// example:
	//
	// FREE
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The product to which the Prometheus instance belongs. Valid values: \\`arms\\` and \\`cms\\`.
	//
	// example:
	//
	// cms
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The number of Prometheus instances in the view.
	//
	// example:
	//
	// 2
	PrometheusInstanceCount *int32 `json:"prometheusInstanceCount,omitempty" xml:"prometheusInstanceCount,omitempty"`
	// The Prometheus view ID.
	//
	// example:
	//
	// view-xxx
	PrometheusViewId *string `json:"prometheusViewId,omitempty" xml:"prometheusViewId,omitempty"`
	// The name of the Prometheus view.
	//
	// example:
	//
	// view1
	PrometheusViewName *string `json:"prometheusViewName,omitempty" xml:"prometheusViewName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource type. The value is fixed to \\`PrometheusView\\`.
	//
	// example:
	//
	// PrometheusView
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The status of the backend data storage.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The version.
	//
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace to which the Prometheus instance belongs.
	//
	// example:
	//
	// default-cms-1490404746278495-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPrometheusViewsResponseBodyPrometheusViews) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsResponseBodyPrometheusViews) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetProduct() *string {
	return s.Product
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetPrometheusInstanceCount() *int32 {
	return s.PrometheusInstanceCount
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetPrometheusViewId() *string {
	return s.PrometheusViewId
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetStatus() *string {
	return s.Status
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetUserId() *string {
	return s.UserId
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetCreateTime(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.CreateTime = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetInstanceType(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.InstanceType = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetPaymentType(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.PaymentType = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetProduct(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.Product = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetPrometheusInstanceCount(v int32) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.PrometheusInstanceCount = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetPrometheusViewId(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.PrometheusViewId = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetPrometheusViewName(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.PrometheusViewName = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetRegionId(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetResourceType(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetStatus(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.Status = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetUserId(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.UserId = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetVersion(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.Version = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) SetWorkspace(v string) *ListPrometheusViewsResponseBodyPrometheusViews {
	s.Workspace = &v
	return s
}

func (s *ListPrometheusViewsResponseBodyPrometheusViews) Validate() error {
	return dara.Validate(s)
}
