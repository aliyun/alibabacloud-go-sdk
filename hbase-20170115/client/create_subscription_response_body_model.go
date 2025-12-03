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
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
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
