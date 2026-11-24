// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawQuotaApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *WithdrawQuotaApplicationResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *WithdrawQuotaApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WithdrawQuotaApplicationResponseBody
	GetSuccess() *bool
}

type WithdrawQuotaApplicationResponseBody struct {
	// The ID of the quota application.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FF8CAF0-29D9-4F11-B6A4-FD2CBCA016D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WithdrawQuotaApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawQuotaApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawQuotaApplicationResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *WithdrawQuotaApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawQuotaApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WithdrawQuotaApplicationResponseBody) SetApplicationId(v string) *WithdrawQuotaApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *WithdrawQuotaApplicationResponseBody) SetRequestId(v string) *WithdrawQuotaApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawQuotaApplicationResponseBody) SetSuccess(v bool) *WithdrawQuotaApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *WithdrawQuotaApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
