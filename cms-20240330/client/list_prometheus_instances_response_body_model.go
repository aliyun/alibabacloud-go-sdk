// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPrometheusInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrometheusInstancesResponseBody
	GetNextToken() *string
	SetPrometheusInstances(v []*ListPrometheusInstancesResponseBodyPrometheusInstances) *ListPrometheusInstancesResponseBody
	GetPrometheusInstances() []*ListPrometheusInstancesResponseBodyPrometheusInstances
	SetRequestId(v string) *ListPrometheusInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPrometheusInstancesResponseBody
	GetTotalCount() *int32
}

type ListPrometheusInstancesResponseBody struct {
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken           *string                                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PrometheusInstances []*ListPrometheusInstancesResponseBodyPrometheusInstances `json:"prometheusInstances,omitempty" xml:"prometheusInstances,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 66
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPrometheusInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusInstancesResponseBody) GetPrometheusInstances() []*ListPrometheusInstancesResponseBodyPrometheusInstances {
	return s.PrometheusInstances
}

func (s *ListPrometheusInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPrometheusInstancesResponseBody) SetMaxResults(v int32) *ListPrometheusInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetNextToken(v string) *ListPrometheusInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetPrometheusInstances(v []*ListPrometheusInstancesResponseBodyPrometheusInstances) *ListPrometheusInstancesResponseBody {
	s.PrometheusInstances = v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetRequestId(v string) *ListPrometheusInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetTotalCount(v int32) *ListPrometheusInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) Validate() error {
	if s.PrometheusInstances != nil {
		for _, item := range s.PrometheusInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusInstancesResponseBodyPrometheusInstances struct {
	// example:
	//
	// readWrite
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// example:
	//
	// 2025-08-10T02:07:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// remote-write
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// example:
	//
	// POSTPAY_GB
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// arms
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// example:
	//
	// rw-63549e054ff596a4149927961dff
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// example:
	//
	// test-prom-name
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// example:
	//
	// cn-nanjing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// Prometheus
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// Pending2Running
	Status           *string                                                       `json:"status,omitempty" xml:"status,omitempty"`
	SupportAuthTypes []*string                                                     `json:"supportAuthTypes,omitempty" xml:"supportAuthTypes,omitempty" type:"Repeated"`
	Tags             []*ListPrometheusInstancesResponseBodyPrometheusInstancesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 17073812345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// *
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// default-cms-115214006-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPrometheusInstancesResponseBodyPrometheusInstances) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesResponseBodyPrometheusInstances) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetAccessType() *string {
	return s.AccessType
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetProduct() *string {
	return s.Product
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetStatus() *string {
	return s.Status
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetSupportAuthTypes() []*string {
	return s.SupportAuthTypes
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetTags() []*ListPrometheusInstancesResponseBodyPrometheusInstancesTags {
	return s.Tags
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetAccessType(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.AccessType = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetCreateTime(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.CreateTime = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetInstanceType(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.InstanceType = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetPaymentType(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.PaymentType = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetProduct(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.Product = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetPrometheusInstanceId(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.PrometheusInstanceId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetPrometheusInstanceName(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.PrometheusInstanceName = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetRegionId(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetResourceGroupId(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetResourceType(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetStatus(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.Status = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetSupportAuthTypes(v []*string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.SupportAuthTypes = v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetTags(v []*ListPrometheusInstancesResponseBodyPrometheusInstancesTags) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.Tags = v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetUserId(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.UserId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetVersion(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.Version = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) SetWorkspace(v string) *ListPrometheusInstancesResponseBodyPrometheusInstances {
	s.Workspace = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstances) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusInstancesResponseBodyPrometheusInstancesTags struct {
	// example:
	//
	// testKey
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// testValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListPrometheusInstancesResponseBodyPrometheusInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesResponseBodyPrometheusInstancesTags) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstancesTags) SetKey(v string) *ListPrometheusInstancesResponseBodyPrometheusInstancesTags {
	s.Key = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstancesTags) SetValue(v string) *ListPrometheusInstancesResponseBodyPrometheusInstancesTags {
	s.Value = &v
	return s
}

func (s *ListPrometheusInstancesResponseBodyPrometheusInstancesTags) Validate() error {
	return dara.Validate(s)
}
