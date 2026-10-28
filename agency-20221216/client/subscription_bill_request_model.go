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
	// The billing month from which to start pushing bills. Format: yyyy-MM.
	//
	// After a successful subscription, the system automatically pushes data from the start billing cycle to the current time. Data for up to the last six months (including the current month) can be pushed. If the specified month exceeds this range, the subscription is invalid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The bill file format. Valid values:
	//
	// - csv
	//
	// - parquet
	//
	// If you subscribe to multiple file types, use separate buckets to avoid file overwriting.
	//
	// This parameter is required.
	//
	// example:
	//
	// csv
	BillFormat *string `json:"BillFormat,omitempty" xml:"BillFormat,omitempty"`
	// The UID of the account that owns the OSS bucket for storing files.
	//
	// If a partner subscribes and needs to push files to the OSS bucket of a subordinate partner account, this account must be a subordinate partner account of the partner account. Grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the subordinate partner account.
	//
	// If a partner subscribes and pushes files to the partner account\\"s own OSS bucket, grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the partner account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5569414254138836
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The name of the OSS bucket for storing bill files.
	//
	// This parameter is required.
	//
	// example:
	//
	// bill-bucket
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The row count threshold for a single bill file. If the number of rows exceeds this threshold, the file is automatically split into multiple files. The naming convention for split files is `uid_billType_billCycle_SquenceNo_fileNo`.
	//
	// All information in the file names is identical except for fileNo. These files belong to the same billing cycle and the same file type but are split into multiple files.
	//
	// example:
	//
	// 100000
	SubscribeSegmentSize *int32 `json:"SubscribeSegmentSize,omitempty" xml:"SubscribeSegmentSize,omitempty"`
	// The subscription type. Valid values:
	//
	// - PartnerBillingItemDetailForBillingPeriod: billing item detail.
	//
	// - PartnerBillingItemDetailMonthly: billing item monthly bill.
	//
	// - PartnerInstanceDetailForBillingPeriod: instance bill detail.
	//
	// - PartnerInstanceDetailMonthly: instance monthly bill.
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
