// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSubscriptionResponseBody
	GetRequestId() *string
	SetSubscription(v *SubscriptionForView) *GetSubscriptionResponseBody
	GetSubscription() *SubscriptionForView
}

type GetSubscriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 111111-222-333-1111-33333
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The subscription object.
	Subscription *SubscriptionForView `json:"subscription,omitempty" xml:"subscription,omitempty"`
}

func (s GetSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionResponseBody) GetSubscription() *SubscriptionForView {
	return s.Subscription
}

func (s *GetSubscriptionResponseBody) SetRequestId(v string) *GetSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetSubscription(v *SubscriptionForView) *GetSubscriptionResponseBody {
	s.Subscription = v
	return s
}

func (s *GetSubscriptionResponseBody) Validate() error {
	if s.Subscription != nil {
		if err := s.Subscription.Validate(); err != nil {
			return err
		}
	}
	return nil
}
