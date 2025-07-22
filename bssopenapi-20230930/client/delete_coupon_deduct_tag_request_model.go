// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponDeductTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *DeleteCouponDeductTagRequest
	GetCouponId() *string
	SetEcIdAccountIds(v []*DeleteCouponDeductTagRequestEcIdAccountIds) *DeleteCouponDeductTagRequest
	GetEcIdAccountIds() []*DeleteCouponDeductTagRequestEcIdAccountIds
	SetNbid(v string) *DeleteCouponDeductTagRequest
	GetNbid() *string
	SetTagKeys(v []*string) *DeleteCouponDeductTagRequest
	GetTagKeys() []*string
}

type DeleteCouponDeductTagRequest struct {
	CouponId       *string                                       `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*DeleteCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                       `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagKeys        []*string                                     `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s DeleteCouponDeductTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *DeleteCouponDeductTagRequest) GetEcIdAccountIds() []*DeleteCouponDeductTagRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DeleteCouponDeductTagRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteCouponDeductTagRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *DeleteCouponDeductTagRequest) SetCouponId(v string) *DeleteCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetEcIdAccountIds(v []*DeleteCouponDeductTagRequestEcIdAccountIds) *DeleteCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetNbid(v string) *DeleteCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetTagKeys(v []*string) *DeleteCouponDeductTagRequest {
	s.TagKeys = v
	return s
}

func (s *DeleteCouponDeductTagRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DeleteCouponDeductTagRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *DeleteCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *DeleteCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
