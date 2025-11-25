// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayEnabledProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisabledDays(v int64) *DescribePostpayEnabledProtectionResponseBody
	GetDisabledDays() *int64
	SetDisabledType(v string) *DescribePostpayEnabledProtectionResponseBody
	GetDisabledType() *string
	SetIsEnabledProtection(v bool) *DescribePostpayEnabledProtectionResponseBody
	GetIsEnabledProtection() *bool
	SetIsOpenButDisabled(v bool) *DescribePostpayEnabledProtectionResponseBody
	GetIsOpenButDisabled() *bool
	SetRequestId(v string) *DescribePostpayEnabledProtectionResponseBody
	GetRequestId() *string
}

type DescribePostpayEnabledProtectionResponseBody struct {
	// example:
	//
	// 6
	DisabledDays *int64 `json:"DisabledDays,omitempty" xml:"DisabledDays,omitempty"`
	// example:
	//
	// nat
	DisabledType *string `json:"DisabledType,omitempty" xml:"DisabledType,omitempty"`
	// example:
	//
	// false
	IsEnabledProtection *bool `json:"IsEnabledProtection,omitempty" xml:"IsEnabledProtection,omitempty"`
	// example:
	//
	// false
	IsOpenButDisabled *bool `json:"IsOpenButDisabled,omitempty" xml:"IsOpenButDisabled,omitempty"`
	// example:
	//
	// 95CA5E2B-E5FB-5838-BC50-6A2C763C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePostpayEnabledProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayEnabledProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayEnabledProtectionResponseBody) GetDisabledDays() *int64 {
	return s.DisabledDays
}

func (s *DescribePostpayEnabledProtectionResponseBody) GetDisabledType() *string {
	return s.DisabledType
}

func (s *DescribePostpayEnabledProtectionResponseBody) GetIsEnabledProtection() *bool {
	return s.IsEnabledProtection
}

func (s *DescribePostpayEnabledProtectionResponseBody) GetIsOpenButDisabled() *bool {
	return s.IsOpenButDisabled
}

func (s *DescribePostpayEnabledProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayEnabledProtectionResponseBody) SetDisabledDays(v int64) *DescribePostpayEnabledProtectionResponseBody {
	s.DisabledDays = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponseBody) SetDisabledType(v string) *DescribePostpayEnabledProtectionResponseBody {
	s.DisabledType = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponseBody) SetIsEnabledProtection(v bool) *DescribePostpayEnabledProtectionResponseBody {
	s.IsEnabledProtection = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponseBody) SetIsOpenButDisabled(v bool) *DescribePostpayEnabledProtectionResponseBody {
	s.IsOpenButDisabled = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponseBody) SetRequestId(v string) *DescribePostpayEnabledProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayEnabledProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
