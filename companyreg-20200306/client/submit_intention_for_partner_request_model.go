// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *SubmitIntentionForPartnerRequest
	GetArea() *string
	SetBizType(v string) *SubmitIntentionForPartnerRequest
	GetBizType() *string
	SetChannel(v string) *SubmitIntentionForPartnerRequest
	GetChannel() *string
	SetCommodityType(v string) *SubmitIntentionForPartnerRequest
	GetCommodityType() *string
	SetContactName(v string) *SubmitIntentionForPartnerRequest
	GetContactName() *string
	SetDescription(v string) *SubmitIntentionForPartnerRequest
	GetDescription() *string
	SetExtInfo(v string) *SubmitIntentionForPartnerRequest
	GetExtInfo() *string
	SetGrade(v int32) *SubmitIntentionForPartnerRequest
	GetGrade() *int32
	SetMobile(v string) *SubmitIntentionForPartnerRequest
	GetMobile() *string
	SetUserId(v string) *SubmitIntentionForPartnerRequest
	GetUserId() *string
}

type SubmitIntentionForPartnerRequest struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// example:
	//
	// esp.isp
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// lingjun
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// Server
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	ContactName   *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// ceshi
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {\\"beianServiceNumber\\":\\"9969c666-0935-4c5b-8042-926ff546e39f\\"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// country
	Grade *int32 `json:"Grade,omitempty" xml:"Grade,omitempty"`
	// example:
	//
	// 18704330000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 1212312312312
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SubmitIntentionForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionForPartnerRequest) GoString() string {
	return s.String()
}

func (s *SubmitIntentionForPartnerRequest) GetArea() *string {
	return s.Area
}

func (s *SubmitIntentionForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *SubmitIntentionForPartnerRequest) GetChannel() *string {
	return s.Channel
}

func (s *SubmitIntentionForPartnerRequest) GetCommodityType() *string {
	return s.CommodityType
}

func (s *SubmitIntentionForPartnerRequest) GetContactName() *string {
	return s.ContactName
}

func (s *SubmitIntentionForPartnerRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitIntentionForPartnerRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *SubmitIntentionForPartnerRequest) GetGrade() *int32 {
	return s.Grade
}

func (s *SubmitIntentionForPartnerRequest) GetMobile() *string {
	return s.Mobile
}

func (s *SubmitIntentionForPartnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *SubmitIntentionForPartnerRequest) SetArea(v string) *SubmitIntentionForPartnerRequest {
	s.Area = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetBizType(v string) *SubmitIntentionForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetChannel(v string) *SubmitIntentionForPartnerRequest {
	s.Channel = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetCommodityType(v string) *SubmitIntentionForPartnerRequest {
	s.CommodityType = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetContactName(v string) *SubmitIntentionForPartnerRequest {
	s.ContactName = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetDescription(v string) *SubmitIntentionForPartnerRequest {
	s.Description = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetExtInfo(v string) *SubmitIntentionForPartnerRequest {
	s.ExtInfo = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetGrade(v int32) *SubmitIntentionForPartnerRequest {
	s.Grade = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetMobile(v string) *SubmitIntentionForPartnerRequest {
	s.Mobile = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) SetUserId(v string) *SubmitIntentionForPartnerRequest {
	s.UserId = &v
	return s
}

func (s *SubmitIntentionForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
