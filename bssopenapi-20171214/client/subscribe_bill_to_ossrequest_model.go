// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeBillToOSSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *SubscribeBillToOSSRequest
	GetBeginBillingCycle() *string
	SetBucketOwnerId(v int64) *SubscribeBillToOSSRequest
	GetBucketOwnerId() *int64
	SetBucketPath(v string) *SubscribeBillToOSSRequest
	GetBucketPath() *string
	SetMultAccountRelSubscribe(v string) *SubscribeBillToOSSRequest
	GetMultAccountRelSubscribe() *string
	SetRowLimitPerFile(v int32) *SubscribeBillToOSSRequest
	GetRowLimitPerFile() *int32
	SetSubscribeBucket(v string) *SubscribeBillToOSSRequest
	GetSubscribeBucket() *string
	SetSubscribeType(v string) *SubscribeBillToOSSRequest
	GetSubscribeType() *string
	SetUsingSsl(v string) *SubscribeBillToOSSRequest
	GetUsingSsl() *string
}

type SubscribeBillToOSSRequest struct {
	// The start billing cycle for pushing. After a successful subscription, the system automatically pushes data from the start billing cycle to the current time. This parameter is invalid for monthly bill PDF subscriptions, and historical data is not re-pushed. Data within the last year can be pushed.
	//
	// example:
	//
	// 2021-03
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The UID of the OSS owner that stores the files. Specify this parameter if you are a Bid/Reseller subscriber and need to push files to a sub-account\\"s OSS. The specified account must be a sub-account of the calling account, and the AliyunConsumeDump2OSSRole permission must be granted to this account. Regular users do not need to specify this parameter. The default value is the calling account.
	//
	// example:
	//
	// 12341324
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The specified directory under the bucket to which the subscription is delivered.
	//
	// example:
	//
	// testpath
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	// For multi-level accounts, specifies the type of bills to push. Valid values:
	//
	// - MA: Pushes bills for the primary account and non-trusteeship linked accounts in the financial cloud.
	//
	// - ACP1: Pushes bills for Partner Platform V2.0 sub-accounts.
	//
	// Default value: MA.
	//
	// example:
	//
	// MA
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
	// The maximum number of rows per file. If a bill file exceeds this limit, it is split into multiple files and compressed into a single archive.
	//
	// example:
	//
	// 300000
	RowLimitPerFile *int32 `json:"RowLimitPerFile,omitempty" xml:"RowLimitPerFile,omitempty"`
	// The name of the OSS bucket that stores the subscribed bill files.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx-bucket
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The type of the subscribed bill. Valid values:
	//
	// - BillingItemDetailForBillingPeriod: billing item details.
	//
	// - InstanceDetailForBillingPeriod: instance bill details.
	//
	// - BillingItemDetailMonthly: billing item summary by billing cycle.
	//
	// - InstanceDetailMonthly: instance bill summary by billing cycle.
	//
	// - SplitItemDetailDaily: split bill daily summary.
	//
	// - MonthBill: monthly bill PDF. Only MA subscription is supported.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
	// Specifies whether to use the Secure Sockets Layer (SSL) encryption protocol to protect network communication. Valid values:
	//
	// - true: Enables SSL encryption to ensure the security and integrity of data transmission.
	//
	// - false (default): Does not enable SSL encryption.
	//
	// example:
	//
	// true
	UsingSsl *string `json:"UsingSsl,omitempty" xml:"UsingSsl,omitempty"`
}

func (s SubscribeBillToOSSRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeBillToOSSRequest) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSRequest) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *SubscribeBillToOSSRequest) GetBucketOwnerId() *int64 {
	return s.BucketOwnerId
}

func (s *SubscribeBillToOSSRequest) GetBucketPath() *string {
	return s.BucketPath
}

func (s *SubscribeBillToOSSRequest) GetMultAccountRelSubscribe() *string {
	return s.MultAccountRelSubscribe
}

func (s *SubscribeBillToOSSRequest) GetRowLimitPerFile() *int32 {
	return s.RowLimitPerFile
}

func (s *SubscribeBillToOSSRequest) GetSubscribeBucket() *string {
	return s.SubscribeBucket
}

func (s *SubscribeBillToOSSRequest) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *SubscribeBillToOSSRequest) GetUsingSsl() *string {
	return s.UsingSsl
}

func (s *SubscribeBillToOSSRequest) SetBeginBillingCycle(v string) *SubscribeBillToOSSRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetBucketOwnerId(v int64) *SubscribeBillToOSSRequest {
	s.BucketOwnerId = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetBucketPath(v string) *SubscribeBillToOSSRequest {
	s.BucketPath = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *SubscribeBillToOSSRequest {
	s.MultAccountRelSubscribe = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetRowLimitPerFile(v int32) *SubscribeBillToOSSRequest {
	s.RowLimitPerFile = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetSubscribeBucket(v string) *SubscribeBillToOSSRequest {
	s.SubscribeBucket = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetSubscribeType(v string) *SubscribeBillToOSSRequest {
	s.SubscribeType = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetUsingSsl(v string) *SubscribeBillToOSSRequest {
	s.UsingSsl = &v
	return s
}

func (s *SubscribeBillToOSSRequest) Validate() error {
	return dara.Validate(s)
}
