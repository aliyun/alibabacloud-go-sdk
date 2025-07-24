// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceInstanceResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceInstanceResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody
	GetResources() []*ListServiceInstanceResourcesResponseBodyResources
}

type ListServiceInstanceResourcesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*ListServiceInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceResourcesResponseBody) GetResources() []*ListServiceInstanceResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListServiceInstanceResourcesResponseBody) SetMaxResults(v int32) *ListServiceInstanceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetNextToken(v string) *ListServiceInstanceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetRequestId(v string) *ListServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceResourcesResponseBodyResources struct {
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription
	//
	// 	- PayAsYouGo
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// RDS
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The renewal state. Valid values:
	//
	// 	- AutoRenewal
	//
	// 	- ManualRenewal
	//
	// 	- NotRenewal
	//
	// example:
	//
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The renewal period.
	//
	// example:
	//
	// 1
	RenewalPeriod *int32 `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:sag:cn-hangzhou:130920852836****:ccn/ccn-b3qf0q439sq2de****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The state of the resource. Valid values:
	//
	// 	- running
	//
	// 	- waiting
	//
	// 	- terminated
	//
	// >  This parameter is returned only for containers.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServiceInstanceResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetPayType() *string {
	return s.PayType
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetProductType() *string {
	return s.ProductType
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetRenewalPeriod() *int32 {
	return s.RenewalPeriod
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetRenewalPeriodUnit() *string {
	return s.RenewalPeriodUnit
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetResourceARN() *string {
	return s.ResourceARN
}

func (s *ListServiceInstanceResourcesResponseBodyResources) GetStatus() *string {
	return s.Status
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetCreateTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetExpireTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetPayType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductCode(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewStatus = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriod(v int32) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriod = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriodUnit(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetResourceARN(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ResourceARN = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
