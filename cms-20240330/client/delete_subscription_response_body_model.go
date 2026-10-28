// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSubscriptionResponseBody
	GetRequestId() *string
	SetSubscriptionId(v string) *DeleteSubscriptionResponseBody
	GetSubscriptionId() *string
}

type DeleteSubscriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8A33DBEA-*****-*****-*****-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s DeleteSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubscriptionResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DeleteSubscriptionResponseBody) SetRequestId(v string) *DeleteSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscriptionResponseBody) SetSubscriptionId(v string) *DeleteSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *DeleteSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
