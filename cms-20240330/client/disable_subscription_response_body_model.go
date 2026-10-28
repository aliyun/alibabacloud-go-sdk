// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSubscriptionResponseBody
	GetRequestId() *string
	SetSubscriptionId(v string) *DisableSubscriptionResponseBody
	GetSubscriptionId() *string
}

type DisableSubscriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-******
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s DisableSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSubscriptionResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DisableSubscriptionResponseBody) SetRequestId(v string) *DisableSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetSubscriptionId(v string) *DisableSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *DisableSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
