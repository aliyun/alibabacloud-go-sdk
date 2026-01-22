// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponItemListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v int64) *DescribeCouponItemListRequest
	GetCouponId() *int64
	SetCurrentPage(v int32) *DescribeCouponItemListRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*DescribeCouponItemListRequestEcIdAccountIds) *DescribeCouponItemListRequest
	GetEcIdAccountIds() []*DescribeCouponItemListRequestEcIdAccountIds
	SetName(v string) *DescribeCouponItemListRequest
	GetName() *string
	SetNbid(v string) *DescribeCouponItemListRequest
	GetNbid() *string
	SetPageSize(v int32) *DescribeCouponItemListRequest
	GetPageSize() *int32
}

type DescribeCouponItemListRequest struct {
	// example:
	//
	// 59104570
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 1
	CurrentPage    *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*DescribeCouponItemListRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Name           *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCouponItemListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListRequest) GetCouponId() *int64 {
	return s.CouponId
}

func (s *DescribeCouponItemListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponItemListRequest) GetEcIdAccountIds() []*DescribeCouponItemListRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DescribeCouponItemListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCouponItemListRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeCouponItemListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponItemListRequest) SetCouponId(v int64) *DescribeCouponItemListRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetCurrentPage(v int32) *DescribeCouponItemListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetEcIdAccountIds(v []*DescribeCouponItemListRequestEcIdAccountIds) *DescribeCouponItemListRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeCouponItemListRequest) SetName(v string) *DescribeCouponItemListRequest {
	s.Name = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetNbid(v string) *DescribeCouponItemListRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetPageSize(v int32) *DescribeCouponItemListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponItemListRequest) Validate() error {
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

type DescribeCouponItemListRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeCouponItemListRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeCouponItemListRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) SetEcId(v string) *DescribeCouponItemListRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
