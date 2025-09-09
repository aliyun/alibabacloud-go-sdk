// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WithdrawServiceResponseBody
	GetRequestId() *string
}

type WithdrawServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawServiceResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawServiceResponseBody) SetRequestId(v string) *WithdrawServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
