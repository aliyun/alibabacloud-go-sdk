// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v int64) *DescribeCouponRequest
	GetCouponId() *int64
	SetCouponNo(v string) *DescribeCouponRequest
	GetCouponNo() *string
	SetCouponTemplateIdList(v []*int64) *DescribeCouponRequest
	GetCouponTemplateIdList() []*int64
	SetCouponType(v string) *DescribeCouponRequest
	GetCouponType() *string
	SetCurrentPage(v int32) *DescribeCouponRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*DescribeCouponRequestEcIdAccountIds) *DescribeCouponRequest
	GetEcIdAccountIds() []*DescribeCouponRequestEcIdAccountIds
	SetEffectiveEndTime(v int64) *DescribeCouponRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *DescribeCouponRequest
	GetEffectiveStartTime() *int64
	SetExpireEndDate(v int64) *DescribeCouponRequest
	GetExpireEndDate() *int64
	SetExpireStartDate(v int64) *DescribeCouponRequest
	GetExpireStartDate() *int64
	SetIncludeShare(v bool) *DescribeCouponRequest
	GetIncludeShare() *bool
	SetMaxResults(v int32) *DescribeCouponRequest
	GetMaxResults() *int32
	SetNbid(v string) *DescribeCouponRequest
	GetNbid() *string
	SetNextToken(v string) *DescribeCouponRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribeCouponRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeCouponRequest
	GetProductCode() *string
	SetStatus(v string) *DescribeCouponRequest
	GetStatus() *string
}

type DescribeCouponRequest struct {
	// The coupon ID.
	//
	// example:
	//
	// 351430260343
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 554863270150
	CouponNo             *string  `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	CouponTemplateIdList []*int64 `json:"CouponTemplateIdList,omitempty" xml:"CouponTemplateIdList,omitempty" type:"Repeated"`
	// The coupon type.
	//
	// example:
	//
	// CERTAIN
	CouponType *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The enterprise and account list. If this parameter is left empty, the current account is queried.
	EcIdAccountIds []*DescribeCouponRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// The effective end time.
	//
	// example:
	//
	// 1708423156000
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The effective start time.
	//
	// example:
	//
	// 1684750028000
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The expiration end time.
	//
	// example:
	//
	// 1708423156000
	ExpireEndDate *int64 `json:"ExpireEndDate,omitempty" xml:"ExpireEndDate,omitempty"`
	// The expiration start time.
	//
	// example:
	//
	// 1684750028000
	ExpireStartDate *int64 `json:"ExpireStartDate,omitempty" xml:"ExpireStartDate,omitempty"`
	IncludeShare    *bool  `json:"IncludeShare,omitempty" xml:"IncludeShare,omitempty"`
	MaxResults      *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The primary campaign information.
	//
	// example:
	//
	// 2684201000001
	Nbid      *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The status.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponRequest) GetCouponId() *int64 {
	return s.CouponId
}

func (s *DescribeCouponRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeCouponRequest) GetCouponTemplateIdList() []*int64 {
	return s.CouponTemplateIdList
}

func (s *DescribeCouponRequest) GetCouponType() *string {
	return s.CouponType
}

func (s *DescribeCouponRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponRequest) GetEcIdAccountIds() []*DescribeCouponRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DescribeCouponRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *DescribeCouponRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *DescribeCouponRequest) GetExpireEndDate() *int64 {
	return s.ExpireEndDate
}

func (s *DescribeCouponRequest) GetExpireStartDate() *int64 {
	return s.ExpireStartDate
}

func (s *DescribeCouponRequest) GetIncludeShare() *bool {
	return s.IncludeShare
}

func (s *DescribeCouponRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCouponRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeCouponRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCouponRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeCouponRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCouponRequest) SetCouponId(v int64) *DescribeCouponRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponRequest) SetCouponNo(v string) *DescribeCouponRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponRequest) SetCouponTemplateIdList(v []*int64) *DescribeCouponRequest {
	s.CouponTemplateIdList = v
	return s
}

func (s *DescribeCouponRequest) SetCouponType(v string) *DescribeCouponRequest {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponRequest) SetCurrentPage(v int32) *DescribeCouponRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponRequest) SetEcIdAccountIds(v []*DescribeCouponRequestEcIdAccountIds) *DescribeCouponRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeCouponRequest) SetEffectiveEndTime(v int64) *DescribeCouponRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *DescribeCouponRequest) SetEffectiveStartTime(v int64) *DescribeCouponRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *DescribeCouponRequest) SetExpireEndDate(v int64) *DescribeCouponRequest {
	s.ExpireEndDate = &v
	return s
}

func (s *DescribeCouponRequest) SetExpireStartDate(v int64) *DescribeCouponRequest {
	s.ExpireStartDate = &v
	return s
}

func (s *DescribeCouponRequest) SetIncludeShare(v bool) *DescribeCouponRequest {
	s.IncludeShare = &v
	return s
}

func (s *DescribeCouponRequest) SetMaxResults(v int32) *DescribeCouponRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCouponRequest) SetNbid(v string) *DescribeCouponRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponRequest) SetNextToken(v string) *DescribeCouponRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCouponRequest) SetPageSize(v int32) *DescribeCouponRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponRequest) SetProductCode(v string) *DescribeCouponRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeCouponRequest) SetStatus(v string) *DescribeCouponRequest {
	s.Status = &v
	return s
}

func (s *DescribeCouponRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCouponRequestEcIdAccountIds struct {
	// The list of accessed accounts. If this parameter is left empty, all accounts under the current entity ID are selected.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The enterprise entity ID.
	//
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeCouponRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeCouponRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeCouponRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DescribeCouponRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeCouponRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeCouponRequestEcIdAccountIds) SetEcId(v string) *DescribeCouponRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DescribeCouponRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
