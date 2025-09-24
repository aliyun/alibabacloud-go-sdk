// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *QueryAvailableInstancesRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *QueryAvailableInstancesRequest
	GetCreateTimeStart() *string
	SetEndTimeEnd(v string) *QueryAvailableInstancesRequest
	GetEndTimeEnd() *string
	SetEndTimeStart(v string) *QueryAvailableInstancesRequest
	GetEndTimeStart() *string
	SetInstanceIDs(v string) *QueryAvailableInstancesRequest
	GetInstanceIDs() *string
	SetOwnerId(v int64) *QueryAvailableInstancesRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryAvailableInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryAvailableInstancesRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryAvailableInstancesRequest
	GetProductCode() *string
	SetProductType(v string) *QueryAvailableInstancesRequest
	GetProductType() *string
	SetRegion(v string) *QueryAvailableInstancesRequest
	GetRegion() *string
	SetRenewStatus(v string) *QueryAvailableInstancesRequest
	GetRenewStatus() *string
	SetSubscriptionType(v string) *QueryAvailableInstancesRequest
	GetSubscriptionType() *string
}

type QueryAvailableInstancesRequest struct {
	// The end time when the specified instance is created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-05-23T12:00:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time when the specified instance is created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-05-23T12:00:00Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-05-23T12:00:00Z.
	//
	// example:
	//
	// 2016-05-23T12:00:00Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-05-23T12:00:00Z.
	//
	// example:
	//
	// 2016-05-23T12:00:00Z
	EndTimeStart *string `json:"EndTimeStart,omitempty" xml:"EndTimeStart,omitempty"`
	// The ID of the instance. Separate multiple IDs with commas (,). You can specify a maximum of 100 IDs.
	//
	// example:
	//
	// rm-xxxxxxxxxxxx
	InstanceIDs *string `json:"InstanceIDs,omitempty" xml:"InstanceIDs,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the service. You can query the service code by calling the **QueryProductList*	- operation or viewing **Codes of Alibaba Cloud services**.
	//
	// >This parameter cannot be left empty if the region is specified.
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
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
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

func (s QueryAvailableInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *QueryAvailableInstancesRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *QueryAvailableInstancesRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *QueryAvailableInstancesRequest) GetEndTimeStart() *string {
	return s.EndTimeStart
}

func (s *QueryAvailableInstancesRequest) GetInstanceIDs() *string {
	return s.InstanceIDs
}

func (s *QueryAvailableInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAvailableInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAvailableInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAvailableInstancesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAvailableInstancesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryAvailableInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryAvailableInstancesRequest) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *QueryAvailableInstancesRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeEnd(v string) *QueryAvailableInstancesRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeStart(v string) *QueryAvailableInstancesRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeEnd(v string) *QueryAvailableInstancesRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeStart(v string) *QueryAvailableInstancesRequest {
	s.EndTimeStart = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetInstanceIDs(v string) *QueryAvailableInstancesRequest {
	s.InstanceIDs = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetOwnerId(v int64) *QueryAvailableInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetPageNum(v int32) *QueryAvailableInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetPageSize(v int32) *QueryAvailableInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetProductCode(v string) *QueryAvailableInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetProductType(v string) *QueryAvailableInstancesRequest {
	s.ProductType = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetRegion(v string) *QueryAvailableInstancesRequest {
	s.Region = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetRenewStatus(v string) *QueryAvailableInstancesRequest {
	s.RenewStatus = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetSubscriptionType(v string) *QueryAvailableInstancesRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAvailableInstancesRequest) Validate() error {
	return dara.Validate(s)
}
