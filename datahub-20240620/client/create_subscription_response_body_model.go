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
	SetSuccess(v string) *CreateSubscriptionResponseBody
	GetSuccess() *string
}

type CreateSubscriptionResponseBody struct {
	// example:
	//
	// 2025092710234722c53d0b08e811d8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1764123368886L0S9H
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateSubscriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSubscriptionResponseBody) SetRequestId(v string) *CreateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetSubscriptionId(v string) *CreateSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *CreateSubscriptionResponseBody) SetSuccess(v string) *CreateSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
