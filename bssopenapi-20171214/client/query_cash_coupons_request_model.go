// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCashCouponsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveOrNot(v bool) *QueryCashCouponsRequest
	GetEffectiveOrNot() *bool
	SetExpiryTimeEnd(v string) *QueryCashCouponsRequest
	GetExpiryTimeEnd() *string
	SetExpiryTimeStart(v string) *QueryCashCouponsRequest
	GetExpiryTimeStart() *string
}

type QueryCashCouponsRequest struct {
	// Specifies whether the voucher is valid. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EffectiveOrNot *bool `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
	// The end of the expiration time range for the voucher. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-08-01T00:00:00Z
	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	// The start of the expiration time range for the voucher. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-09-01T00:00:00Z
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
}

func (s QueryCashCouponsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCashCouponsRequest) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsRequest) GetEffectiveOrNot() *bool {
	return s.EffectiveOrNot
}

func (s *QueryCashCouponsRequest) GetExpiryTimeEnd() *string {
	return s.ExpiryTimeEnd
}

func (s *QueryCashCouponsRequest) GetExpiryTimeStart() *string {
	return s.ExpiryTimeStart
}

func (s *QueryCashCouponsRequest) SetEffectiveOrNot(v bool) *QueryCashCouponsRequest {
	s.EffectiveOrNot = &v
	return s
}

func (s *QueryCashCouponsRequest) SetExpiryTimeEnd(v string) *QueryCashCouponsRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryCashCouponsRequest) SetExpiryTimeStart(v string) *QueryCashCouponsRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryCashCouponsRequest) Validate() error {
	return dara.Validate(s)
}
