// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPromotionOfferForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *GrantPromotionOfferForPartnerRequest
	GetActivityCode() *string
	SetActivityId(v string) *GrantPromotionOfferForPartnerRequest
	GetActivityId() *string
	SetBelongId(v string) *GrantPromotionOfferForPartnerRequest
	GetBelongId() *string
	SetChannel(v string) *GrantPromotionOfferForPartnerRequest
	GetChannel() *string
	SetEmployeeCode(v string) *GrantPromotionOfferForPartnerRequest
	GetEmployeeCode() *string
	SetRemark(v string) *GrantPromotionOfferForPartnerRequest
	GetRemark() *string
}

type GrantPromotionOfferForPartnerRequest struct {
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// The activity ID.
	//
	// example:
	//
	// 232
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// The channel.
	//
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The employee code.
	//
	// example:
	//
	// WB02409424
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// The operation remarks (audit information).
	//
	// example:
	//
	// hz-maven
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s GrantPromotionOfferForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantPromotionOfferForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GrantPromotionOfferForPartnerRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *GrantPromotionOfferForPartnerRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *GrantPromotionOfferForPartnerRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *GrantPromotionOfferForPartnerRequest) GetChannel() *string {
	return s.Channel
}

func (s *GrantPromotionOfferForPartnerRequest) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *GrantPromotionOfferForPartnerRequest) GetRemark() *string {
	return s.Remark
}

func (s *GrantPromotionOfferForPartnerRequest) SetActivityCode(v string) *GrantPromotionOfferForPartnerRequest {
	s.ActivityCode = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) SetActivityId(v string) *GrantPromotionOfferForPartnerRequest {
	s.ActivityId = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) SetBelongId(v string) *GrantPromotionOfferForPartnerRequest {
	s.BelongId = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) SetChannel(v string) *GrantPromotionOfferForPartnerRequest {
	s.Channel = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) SetEmployeeCode(v string) *GrantPromotionOfferForPartnerRequest {
	s.EmployeeCode = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) SetRemark(v string) *GrantPromotionOfferForPartnerRequest {
	s.Remark = &v
	return s
}

func (s *GrantPromotionOfferForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
