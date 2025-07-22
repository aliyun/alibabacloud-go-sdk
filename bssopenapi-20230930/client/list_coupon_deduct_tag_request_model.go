// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponDeductTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *ListCouponDeductTagRequest
	GetCouponId() *string
	SetEcIdAccountIds(v []*ListCouponDeductTagRequestEcIdAccountIds) *ListCouponDeductTagRequest
	GetEcIdAccountIds() []*ListCouponDeductTagRequestEcIdAccountIds
	SetNbid(v string) *ListCouponDeductTagRequest
	GetNbid() *string
}

type ListCouponDeductTagRequest struct {
	CouponId       *string                                     `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*ListCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                     `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListCouponDeductTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *ListCouponDeductTagRequest) GetEcIdAccountIds() []*ListCouponDeductTagRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *ListCouponDeductTagRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListCouponDeductTagRequest) SetCouponId(v string) *ListCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *ListCouponDeductTagRequest) SetEcIdAccountIds(v []*ListCouponDeductTagRequestEcIdAccountIds) *ListCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *ListCouponDeductTagRequest) SetNbid(v string) *ListCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

func (s *ListCouponDeductTagRequest) Validate() error {
	return dara.Validate(s)
}

type ListCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s ListCouponDeductTagRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *ListCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *ListCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
