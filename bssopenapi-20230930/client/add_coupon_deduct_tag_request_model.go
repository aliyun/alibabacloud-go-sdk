// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCouponDeductTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *AddCouponDeductTagRequest
	GetCouponId() *string
	SetEcIdAccountIds(v []*AddCouponDeductTagRequestEcIdAccountIds) *AddCouponDeductTagRequest
	GetEcIdAccountIds() []*AddCouponDeductTagRequestEcIdAccountIds
	SetNbid(v string) *AddCouponDeductTagRequest
	GetNbid() *string
	SetTags(v []*AddCouponDeductTagRequestTags) *AddCouponDeductTagRequest
	GetTags() []*AddCouponDeductTagRequestTags
}

type AddCouponDeductTagRequest struct {
	CouponId       *string                                    `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*AddCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                    `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	Tags           []*AddCouponDeductTagRequestTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddCouponDeductTagRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *AddCouponDeductTagRequest) GetEcIdAccountIds() []*AddCouponDeductTagRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *AddCouponDeductTagRequest) GetNbid() *string {
	return s.Nbid
}

func (s *AddCouponDeductTagRequest) GetTags() []*AddCouponDeductTagRequestTags {
	return s.Tags
}

func (s *AddCouponDeductTagRequest) SetCouponId(v string) *AddCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *AddCouponDeductTagRequest) SetEcIdAccountIds(v []*AddCouponDeductTagRequestEcIdAccountIds) *AddCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *AddCouponDeductTagRequest) SetNbid(v string) *AddCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

func (s *AddCouponDeductTagRequest) SetTags(v []*AddCouponDeductTagRequestTags) *AddCouponDeductTagRequest {
	s.Tags = v
	return s
}

func (s *AddCouponDeductTagRequest) Validate() error {
	return dara.Validate(s)
}

type AddCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s AddCouponDeductTagRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *AddCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *AddCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type AddCouponDeductTagRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddCouponDeductTagRequestTags) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagRequestTags) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequestTags) GetKey() *string {
	return s.Key
}

func (s *AddCouponDeductTagRequestTags) GetValue() *string {
	return s.Value
}

func (s *AddCouponDeductTagRequestTags) SetKey(v string) *AddCouponDeductTagRequestTags {
	s.Key = &v
	return s
}

func (s *AddCouponDeductTagRequestTags) SetValue(v string) *AddCouponDeductTagRequestTags {
	s.Value = &v
	return s
}

func (s *AddCouponDeductTagRequestTags) Validate() error {
	return dara.Validate(s)
}
