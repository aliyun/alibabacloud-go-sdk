// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubscriptionBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubscribeType(v string) *CancelSubscriptionBillRequest
	GetSubscribeType() *string
}

type CancelSubscriptionBillRequest struct {
	// The type of the bill to which you want to cancel the subscription. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	//
	// This parameter is required.
	//
	// example:
	//
	// PartnerBillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s CancelSubscriptionBillRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillRequest) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *CancelSubscriptionBillRequest) SetSubscribeType(v string) *CancelSubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

func (s *CancelSubscriptionBillRequest) Validate() error {
	return dara.Validate(s)
}
