// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSubscriptionResponseBody
	GetRequestId() *string
	SetSubscriptionId(v string) *UpdateSubscriptionResponseBody
	GetSubscriptionId() *string
}

type UpdateSubscriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-******
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubscriptionResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetSubscriptionId(v string) *UpdateSubscriptionResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
