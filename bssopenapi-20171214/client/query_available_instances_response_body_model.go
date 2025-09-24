// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAvailableInstancesResponseBody
	GetCode() *string
	SetData(v *QueryAvailableInstancesResponseBodyData) *QueryAvailableInstancesResponseBody
	GetData() *QueryAvailableInstancesResponseBodyData
	SetMessage(v string) *QueryAvailableInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAvailableInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAvailableInstancesResponseBody
	GetSuccess() *bool
}

type QueryAvailableInstancesResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryAvailableInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7C15585-8349-4C62-BEE4-5A391841B9BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAvailableInstancesResponseBody) GetData() *QueryAvailableInstancesResponseBodyData {
	return s.Data
}

func (s *QueryAvailableInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAvailableInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvailableInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAvailableInstancesResponseBody) SetCode(v string) *QueryAvailableInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetData(v *QueryAvailableInstancesResponseBodyData) *QueryAvailableInstancesResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetMessage(v string) *QueryAvailableInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetRequestId(v string) *QueryAvailableInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetSuccess(v bool) *QueryAvailableInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAvailableInstancesResponseBodyData struct {
	// The instances returned.
	InstanceList []*QueryAvailableInstancesResponseBodyDataInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAvailableInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBodyData) GetInstanceList() []*QueryAvailableInstancesResponseBodyDataInstanceList {
	return s.InstanceList
}

func (s *QueryAvailableInstancesResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAvailableInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAvailableInstancesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryAvailableInstancesResponseBodyData) SetInstanceList(v []*QueryAvailableInstancesResponseBodyDataInstanceList) *QueryAvailableInstancesResponseBodyData {
	s.InstanceList = v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetPageNum(v int32) *QueryAvailableInstancesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetPageSize(v int32) *QueryAvailableInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetTotalCount(v int32) *QueryAvailableInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryAvailableInstancesResponseBodyDataInstanceList struct {
	// The time when the specified instance was created.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the instance was expired.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the specified instance was expected to be released.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	ExpectedReleaseTime *string `json:"ExpectedReleaseTime,omitempty" xml:"ExpectedReleaseTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1049056
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The ID of the instance owner.
	//
	// example:
	//
	// 325352345
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the instance was released.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The renewal status of the specified instance. Valid values:
	//
	// 	- AutoRenewal: The instance is automatically renewed.
	//
	// 	- ManualRenewal: The instance is manually renewed.
	//
	// 	- NotRenewal: The instance is not renewed.
	//
	// example:
	//
	// ManualRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The number of auto-renewal cycles.
	//
	// example:
	//
	// 1
	RenewalDuration *int32 `json:"RenewalDuration,omitempty" xml:"RenewalDuration,omitempty"`
	// The unit of the auto-renewal cycle. Valid values:
	//
	// 	- M: month
	//
	// 	- Y: year
	//
	// example:
	//
	// M
	RenewalDurationUnit *string `json:"RenewalDurationUnit,omitempty" xml:"RenewalDurationUnit,omitempty"`
	// The seller.
	//
	// example:
	//
	// 123123123
	Seller *string `json:"Seller,omitempty" xml:"Seller,omitempty"`
	// The ID of the seller.
	//
	// example:
	//
	// 123123123
	SellerId *int64 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Creating: The instance is being created. WaitForExpire: The instance is about to expire. Normal: The instance can properly run. Expired: The instance is expired.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the specified instance was suspended.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The sub-status of the specified instance.
	//
	// example:
	//
	// Normal: The pay-as-you-go module can properly run. WaitForLimit: The pay-as-you-go module is about to be limited due to overdue payments. BandwidthLimited: The pay-as-you-go module is limited due to overdue payments.
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryAvailableInstancesResponseBodyDataInstanceList) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableInstancesResponseBodyDataInstanceList) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetExpectedReleaseTime() *string {
	return s.ExpectedReleaseTime
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetInstanceID() *string {
	return s.InstanceID
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetProductType() *string {
	return s.ProductType
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetRegion() *string {
	return s.Region
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetRenewalDuration() *int32 {
	return s.RenewalDuration
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetRenewalDurationUnit() *string {
	return s.RenewalDurationUnit
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetSeller() *string {
	return s.Seller
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetSellerId() *int64 {
	return s.SellerId
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetStatus() *string {
	return s.Status
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetStopTime() *string {
	return s.StopTime
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetSubStatus() *string {
	return s.SubStatus
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetCreateTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.CreateTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetEndTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.EndTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetExpectedReleaseTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ExpectedReleaseTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetInstanceID(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.InstanceID = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetOwnerId(v int64) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetProductCode(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ProductCode = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetProductType(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ProductType = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRegion(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Region = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetReleaseTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ReleaseTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewStatus = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewalDuration(v int32) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewalDuration = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewalDurationUnit(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewalDurationUnit = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSeller(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Seller = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSellerId(v int64) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SellerId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Status = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetStopTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.StopTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSubStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SubStatus = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSubscriptionType(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) Validate() error {
	return dara.Validate(s)
}
