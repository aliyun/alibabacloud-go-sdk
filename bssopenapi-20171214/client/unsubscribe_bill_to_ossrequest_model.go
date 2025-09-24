// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeBillToOSSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMultAccountRelSubscribe(v string) *UnsubscribeBillToOSSRequest
	GetMultAccountRelSubscribe() *string
	SetSubscribeType(v string) *UnsubscribeBillToOSSRequest
	GetSubscribeType() *string
}

type UnsubscribeBillToOSSRequest struct {
	// The type of accounts whose bills are to be pushed if multi-tier accounts are involved. Valid values:
	//
	// 	- MA: management account.
	//
	// 	- ACP1: member account of a virtual network operator (VNO).
	//
	// Default value: MA.
	//
	// example:
	//
	// MA
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
	// The type of the bill to which you want to subscribe. Valid values:
	//
	// 	- BillingItemDetailForBillingPeriod: bills of billable items
	//
	// 	- InstanceDetailForBillingPeriod: bills of instances
	//
	// 	- BillingItemDetailMonthly: billable item-based bills summarized by billing cycle
	//
	// 	- InstanceDetailMonthly: instance-based bills summarized by billing cycle
	//
	// 	- SplitItemDetailDaily: split bills summarized by day
	//
	// 	- MonthBill: monthly bills in the PDF format You can subscribe to the monthly PDF bills only of the master account.
	//
	// This parameter is required.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s UnsubscribeBillToOSSRequest) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeBillToOSSRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSRequest) GetMultAccountRelSubscribe() *string {
	return s.MultAccountRelSubscribe
}

func (s *UnsubscribeBillToOSSRequest) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *UnsubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *UnsubscribeBillToOSSRequest {
	s.MultAccountRelSubscribe = &v
	return s
}

func (s *UnsubscribeBillToOSSRequest) SetSubscribeType(v string) *UnsubscribeBillToOSSRequest {
	s.SubscribeType = &v
	return s
}

func (s *UnsubscribeBillToOSSRequest) Validate() error {
	return dara.Validate(s)
}
