// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundAppInstanceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RefundAppInstanceForPartnerRequest
	GetBizId() *string
	SetRefundReason(v string) *RefundAppInstanceForPartnerRequest
	GetRefundReason() *string
	SetUserId(v string) *RefundAppInstanceForPartnerRequest
	GetUserId() *string
}

type RefundAppInstanceForPartnerRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RefundReason *string `json:"RefundReason,omitempty" xml:"RefundReason,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RefundAppInstanceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundAppInstanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *RefundAppInstanceForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *RefundAppInstanceForPartnerRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *RefundAppInstanceForPartnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *RefundAppInstanceForPartnerRequest) SetBizId(v string) *RefundAppInstanceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *RefundAppInstanceForPartnerRequest) SetRefundReason(v string) *RefundAppInstanceForPartnerRequest {
	s.RefundReason = &v
	return s
}

func (s *RefundAppInstanceForPartnerRequest) SetUserId(v string) *RefundAppInstanceForPartnerRequest {
	s.UserId = &v
	return s
}

func (s *RefundAppInstanceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
