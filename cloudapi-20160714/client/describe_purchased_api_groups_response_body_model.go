// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePurchasedApiGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePurchasedApiGroupsResponseBody
	GetPageSize() *int32
	SetPurchasedApiGroupAttributes(v *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) *DescribePurchasedApiGroupsResponseBody
	GetPurchasedApiGroupAttributes() *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes
	SetRequestId(v string) *DescribePurchasedApiGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePurchasedApiGroupsResponseBody
	GetTotalCount() *int32
}

type DescribePurchasedApiGroupsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The attributes of the API group.
	PurchasedApiGroupAttributes *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes `json:"PurchasedApiGroupAttributes,omitempty" xml:"PurchasedApiGroupAttributes,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36BBBAD4-1CFB-489F-841A-8CA52EEA787E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurchasedApiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePurchasedApiGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePurchasedApiGroupsResponseBody) GetPurchasedApiGroupAttributes() *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes {
	return s.PurchasedApiGroupAttributes
}

func (s *DescribePurchasedApiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePurchasedApiGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPageNumber(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPageSize(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPurchasedApiGroupAttributes(v *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) *DescribePurchasedApiGroupsResponseBody {
	s.PurchasedApiGroupAttributes = v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetRequestId(v string) *DescribePurchasedApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetTotalCount(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes struct {
	PurchasedApiGroupAttribute []*DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute `json:"PurchasedApiGroupAttribute,omitempty" xml:"PurchasedApiGroupAttribute,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) GetPurchasedApiGroupAttribute() []*DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	return s.PurchasedApiGroupAttribute
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) SetPurchasedApiGroupAttribute(v []*DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes {
	s.PurchasedApiGroupAttribute = v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute struct {
	// The billing method.
	//
	// example:
	//
	// POSTPAID
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// The description of the API group.
	//
	// example:
	//
	// 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the API group expires.
	//
	// example:
	//
	// 2021-12-19T00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 8e91f23f88f94348855b82c9a73209f3
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The maximum number of calls.
	//
	// example:
	//
	// 10
	InvokeTimesMax *int64 `json:"InvokeTimesMax,omitempty" xml:"InvokeTimesMax,omitempty"`
	// The current number of calls.
	//
	// example:
	//
	// 10
	InvokeTimesNow *int64 `json:"InvokeTimesNow,omitempty" xml:"InvokeTimesNow,omitempty"`
	// The time when the API group was purchased.
	//
	// example:
	//
	// 2021-12-19T00:00:00
	PurchasedTime *string `json:"PurchasedTime,omitempty" xml:"PurchasedTime,omitempty"`
	// The ID of the region where the API group is located.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the API group.
	//
	// 	- **NORMAL**: The API group is normal.
	//
	// 	- **DELETE**: The API group is deleted.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetBillingType() *string {
	return s.BillingType
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetInvokeTimesMax() *int64 {
	return s.InvokeTimesMax
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetInvokeTimesNow() *int64 {
	return s.InvokeTimesNow
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetPurchasedTime() *string {
	return s.PurchasedTime
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetBillingType(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.BillingType = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetDescription(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetExpireTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetGroupId(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetGroupName(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetInvokeTimesMax(v int64) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.InvokeTimesMax = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetInvokeTimesNow(v int64) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.InvokeTimesNow = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetPurchasedTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.PurchasedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetRegionId(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetStatus(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.Status = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) Validate() error {
	return dara.Validate(s)
}
