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
	SetSuccess(v bool) *DeleteSubscriptionResponseBody
	GetSuccess() *bool
}

type DeleteSubscriptionResponseBody struct {
	// example:
	//
	// 2026031915480122c53d0b00c2d347
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSubscriptionResponseBody) SetRequestId(v string) *DeleteSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscriptionResponseBody) SetSuccess(v bool) *DeleteSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
