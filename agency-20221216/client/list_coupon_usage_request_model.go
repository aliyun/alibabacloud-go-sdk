// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ListCouponUsageRequest
	GetAccount() *string
	SetCouponTemplateId(v int64) *ListCouponUsageRequest
	GetCouponTemplateId() *int64
	SetPage(v int32) *ListCouponUsageRequest
	GetPage() *int32
	SetPageSize(v int32) *ListCouponUsageRequest
	GetPageSize() *int32
	SetStatus(v string) *ListCouponUsageRequest
	GetStatus() *string
	SetT2PartnerUid(v int64) *ListCouponUsageRequest
	GetT2PartnerUid() *int64
	SetUid(v int64) *ListCouponUsageRequest
	GetUid() *int64
}

type ListCouponUsageRequest struct {
	// 阿里云客户账号
	//
	// example:
	//
	// oqevfbveuadcrduzmf@ttirv.net
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// 优惠券模版id
	//
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// 页码</br>
	//
	//  默认值为1 最小值1
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 分页行数 </br>
	//
	//   默认值20 最大值50 最小值1
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 优惠券状态 </br>
	//
	// AVAILABLE 正常 </br>
	//
	// EXHAUSTED 已用完 </br>
	//
	// EXPIRED 已过期 </br>
	//
	// ABANDONED 已作废 </br>
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// T2伙伴uid
	//
	// 如：123456789
	//
	// example:
	//
	// 123456768
	T2PartnerUid *int64 `json:"T2PartnerUid,omitempty" xml:"T2PartnerUid,omitempty"`
	// 阿里云账号uid
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCouponUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCouponUsageRequest) GoString() string {
	return s.String()
}

func (s *ListCouponUsageRequest) GetAccount() *string {
	return s.Account
}

func (s *ListCouponUsageRequest) GetCouponTemplateId() *int64 {
	return s.CouponTemplateId
}

func (s *ListCouponUsageRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListCouponUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCouponUsageRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCouponUsageRequest) GetT2PartnerUid() *int64 {
	return s.T2PartnerUid
}

func (s *ListCouponUsageRequest) GetUid() *int64 {
	return s.Uid
}

func (s *ListCouponUsageRequest) SetAccount(v string) *ListCouponUsageRequest {
	s.Account = &v
	return s
}

func (s *ListCouponUsageRequest) SetCouponTemplateId(v int64) *ListCouponUsageRequest {
	s.CouponTemplateId = &v
	return s
}

func (s *ListCouponUsageRequest) SetPage(v int32) *ListCouponUsageRequest {
	s.Page = &v
	return s
}

func (s *ListCouponUsageRequest) SetPageSize(v int32) *ListCouponUsageRequest {
	s.PageSize = &v
	return s
}

func (s *ListCouponUsageRequest) SetStatus(v string) *ListCouponUsageRequest {
	s.Status = &v
	return s
}

func (s *ListCouponUsageRequest) SetT2PartnerUid(v int64) *ListCouponUsageRequest {
	s.T2PartnerUid = &v
	return s
}

func (s *ListCouponUsageRequest) SetUid(v int64) *ListCouponUsageRequest {
	s.Uid = &v
	return s
}

func (s *ListCouponUsageRequest) Validate() error {
	return dara.Validate(s)
}
