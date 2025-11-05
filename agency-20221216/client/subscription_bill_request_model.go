// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *SubscriptionBillRequest
	GetBeginBillingCycle() *string
	SetBillFormat(v string) *SubscriptionBillRequest
	GetBillFormat() *string
	SetBucketOwnerId(v int64) *SubscriptionBillRequest
	GetBucketOwnerId() *int64
	SetSubscribeBucket(v string) *SubscriptionBillRequest
	GetSubscribeBucket() *string
	SetSubscribeSegmentSize(v int32) *SubscriptionBillRequest
	GetSubscribeSegmentSize() *int32
	SetSubscribeType(v string) *SubscriptionBillRequest
	GetSubscribeType() *string
}

type SubscriptionBillRequest struct {
	// The start month from which the bills are pushed. Specify the value in the yyyy-MM format.
	//
	// After the subscription is generated, the system automatically pushes the bill data that is generated from the month that you specified to the current point in time. Data of up to six months can be pushed. The current month is included. If you subscribe to the bills for more than six months, the subscription is invalid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The file format of the bill. Valid values: csv and parquet.
	//
	// If you subscribe to the bills of multiple file formats, we recommend that you store the bills in different OSS buckets to prevent file overwriting.
	//
	// This parameter is required.
	//
	// example:
	//
	// csv
	BillFormat *string `json:"BillFormat,omitempty" xml:"BillFormat,omitempty"`
	// The ID of the user to which the OSS bucket belongs.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of a subordinate partner account, you must set this parameter to the ID of the subordinate partner account and grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the subordinate partner account.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of your own account, your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5569414254138836
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The name of the Object Storage Service (OSS) bucket in which you want to store the bills.
	//
	// This parameter is required.
	//
	// example:
	//
	// bill-bucket
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The maximum rows in a single bill file. If the number of bill rows exceed the upper limit, the bill is automatically split into multiple files. The name of each split file is in the `uid_billType_billCycle_SquenceNo_fileNo` format.
	//
	// Files whose names are the same except for the fileNo field are of the same type and belong to the same billing cycle.
	//
	// example:
	//
	// 100000
	SubscribeSegmentSize *int32 `json:"SubscribeSegmentSize,omitempty" xml:"SubscribeSegmentSize,omitempty"`
	// The type of the bill to which you want to subscribe. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	//
	// This parameter is required.
	//
	// example:
	//
	// PartnerBillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s SubscriptionBillRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionBillRequest) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *SubscriptionBillRequest) GetBillFormat() *string {
	return s.BillFormat
}

func (s *SubscriptionBillRequest) GetBucketOwnerId() *int64 {
	return s.BucketOwnerId
}

func (s *SubscriptionBillRequest) GetSubscribeBucket() *string {
	return s.SubscribeBucket
}

func (s *SubscriptionBillRequest) GetSubscribeSegmentSize() *int32 {
	return s.SubscribeSegmentSize
}

func (s *SubscriptionBillRequest) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *SubscriptionBillRequest) SetBeginBillingCycle(v string) *SubscriptionBillRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *SubscriptionBillRequest) SetBillFormat(v string) *SubscriptionBillRequest {
	s.BillFormat = &v
	return s
}

func (s *SubscriptionBillRequest) SetBucketOwnerId(v int64) *SubscriptionBillRequest {
	s.BucketOwnerId = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeBucket(v string) *SubscriptionBillRequest {
	s.SubscribeBucket = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeSegmentSize(v int32) *SubscriptionBillRequest {
	s.SubscribeSegmentSize = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeType(v string) *SubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

func (s *SubscriptionBillRequest) Validate() error {
	return dara.Validate(s)
}
