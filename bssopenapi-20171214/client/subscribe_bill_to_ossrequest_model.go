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
	// The initial billing cycle from which bills start to be pushed. After you subscribe to the bills, the system automatically pushes the data that is generated from the initial billing cycle to the current time. If the SubscribeType parameter is set to MonthBill, this parameter is invalid. Historical data is not pushed again. The data generated within the last year can be pushed.
	//
	// example:
	//
	// 2021-03
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The owner ID of the OSS bucket that stores the bills. This parameter is required if you are a bidder or reseller and want to push data to an OSS bucket of a member account. In this case, you must specify this account as the account used to call this operation and grant the AliyunConsumeDump2OSSRole permission to this account. If you are a regular user, you do not need to set this parameter. By default, your account is used to call this operation.
	//
	// example:
	//
	// 12341324
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The path of the OSS bucket.
	//
	// example:
	//
	// testpath
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	// The type of the account whose bills are to be pushed if multi-tier accounts are involved. Valid values:
	//
	// 	- MA: the master account and a non-managed member account in Finance Cloud
	//
	// 	- ACP1: a member account of a virtual network operator (VNO)
	//
	// Default value: MA.
	//
	// example:
	//
	// MA
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
	// The upper limit of the number of lines in a single file. When the bill file exceeds the upper limit, it will be split into multiple files and merged into a compressed package.
	//
	// example:
	//
	// 300000
	RowLimitPerFile *int32 `json:"RowLimitPerFile,omitempty" xml:"RowLimitPerFile,omitempty"`
	// The OSS bucket that stores the bills to which you want to subscribe.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx-bucket
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The type of the bill to which you want to subscribe. Valid values:
	//
	// 	- BillingItemDetailForBillingPeriod: detailed bills of billable items
	//
	// 	- InstanceDetailForBillingPeriod: detailed bills of instances
	//
	// 	- BillingItemDetailMonthly: billable item-based bills summarized by billing cycle
	//
	// 	- InstanceDetailMonthly: instance-based bills summarized by billing cycle
	//
	// 	- SplitItemDetailDaily: split bills summarized by day
	//
	// 	- MonthBill: monthly bills in the PDF format. You can subscribe to the monthly PDF bills only of the master account.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
	// Whether to protect network communications through the SSL (Secure Sockets Layer) encryption protocol. When this parameter is set to true, it means that SSL encryption is enabled to ensure the security and integrity of data transmission.
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
