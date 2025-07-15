// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawVpcPublishedRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WithdrawVpcPublishedRouteEntriesResponseBody
	GetRequestId() *string
}

type WithdrawVpcPublishedRouteEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawVpcPublishedRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawVpcPublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawVpcPublishedRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawVpcPublishedRouteEntriesResponseBody) SetRequestId(v string) *WithdrawVpcPublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
