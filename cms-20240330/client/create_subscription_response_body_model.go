// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSubscriptionResponseBody
	GetRequestId() *string
	SetSubscriptionId(v string) *CreateSubscriptionResponseBody
	GetSubscriptionId() *string
}

type CreateSubscriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s CreateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubscriptionResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *CreateSubscriptionResponseBody) SetRequestId(v string) *CreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetSubscriptionId(v string) *CreateSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
