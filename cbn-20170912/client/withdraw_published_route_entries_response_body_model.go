// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawPublishedRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WithdrawPublishedRouteEntriesResponseBody
	GetRequestId() *string
}

type WithdrawPublishedRouteEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawPublishedRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawPublishedRouteEntriesResponseBody) SetRequestId(v string) *WithdrawPublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
