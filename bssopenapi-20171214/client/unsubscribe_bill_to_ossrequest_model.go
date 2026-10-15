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
	// For multi-level accounts, specifies the type of bill to push. Valid values:
	//
	// - MA: Pushes the bill of the main account.
	//
	// - ACP1: Pushes the bill of the Partner Platform V2.0 sub-account.
	//
	// Default value: MA.
	//
	// example:
	//
	// MA
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
	// The subscription type. Valid values:
	//
	// - BillingItemDetailForBillingPeriod: billable item-level consumption details.
	//
	// - InstanceDetailForBillingPeriod: instance-level consumption details.
	//
	// - BillingItemDetailMonthly: billable item-level consumption summary by billing cycle.
	//
	// - InstanceDetailMonthly: instance-level consumption summary by billing cycle.
	//
	// - SplitItemDetailDaily: split bill summarized by day.
	//
	// - MonthBill: monthly bill in PDF format. Only MA subscriptions are supported.
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
