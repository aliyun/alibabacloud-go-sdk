// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedeemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveOrNot(v bool) *QueryRedeemRequest
	GetEffectiveOrNot() *bool
	SetExpiryTimeEnd(v string) *QueryRedeemRequest
	GetExpiryTimeEnd() *string
	SetExpiryTimeStart(v string) *QueryRedeemRequest
	GetExpiryTimeStart() *string
	SetPageNum(v int32) *QueryRedeemRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRedeemRequest
	GetPageSize() *int32
}

type QueryRedeemRequest struct {
	// Specifies whether the redemption coupon takes effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EffectiveOrNot *bool `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
	// The end time when the redemption coupon expires. The value must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2018-08-01T00:00:00Z
	ExpiryTimeEnd *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	// The start time when the redemption coupon expires. The value must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2018-08-01T00:00:00Z
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRedeemRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemRequest) GoString() string {
	return s.String()
}

func (s *QueryRedeemRequest) GetEffectiveOrNot() *bool {
	return s.EffectiveOrNot
}

func (s *QueryRedeemRequest) GetExpiryTimeEnd() *string {
	return s.ExpiryTimeEnd
}

func (s *QueryRedeemRequest) GetExpiryTimeStart() *string {
	return s.ExpiryTimeStart
}

func (s *QueryRedeemRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRedeemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRedeemRequest) SetEffectiveOrNot(v bool) *QueryRedeemRequest {
	s.EffectiveOrNot = &v
	return s
}

func (s *QueryRedeemRequest) SetExpiryTimeEnd(v string) *QueryRedeemRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryRedeemRequest) SetExpiryTimeStart(v string) *QueryRedeemRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryRedeemRequest) SetPageNum(v int32) *QueryRedeemRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRedeemRequest) SetPageSize(v int32) *QueryRedeemRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRedeemRequest) Validate() error {
	return dara.Validate(s)
}
