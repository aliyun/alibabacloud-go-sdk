// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSlaCouponApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDamagedIds(v []*string) *SubmitSlaCouponApplyRequest
	GetDamagedIds() []*string
	SetEcIdAccountIds(v []*SubmitSlaCouponApplyRequestEcIdAccountIds) *SubmitSlaCouponApplyRequest
	GetEcIdAccountIds() []*SubmitSlaCouponApplyRequestEcIdAccountIds
	SetMonth(v int32) *SubmitSlaCouponApplyRequest
	GetMonth() *int32
	SetNbid(v string) *SubmitSlaCouponApplyRequest
	GetNbid() *string
}

type SubmitSlaCouponApplyRequest struct {
	DamagedIds     []*string                                    `json:"DamagedIds,omitempty" xml:"DamagedIds,omitempty" type:"Repeated"`
	EcIdAccountIds []*SubmitSlaCouponApplyRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 202603
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s SubmitSlaCouponApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSlaCouponApplyRequest) GoString() string {
	return s.String()
}

func (s *SubmitSlaCouponApplyRequest) GetDamagedIds() []*string {
	return s.DamagedIds
}

func (s *SubmitSlaCouponApplyRequest) GetEcIdAccountIds() []*SubmitSlaCouponApplyRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *SubmitSlaCouponApplyRequest) GetMonth() *int32 {
	return s.Month
}

func (s *SubmitSlaCouponApplyRequest) GetNbid() *string {
	return s.Nbid
}

func (s *SubmitSlaCouponApplyRequest) SetDamagedIds(v []*string) *SubmitSlaCouponApplyRequest {
	s.DamagedIds = v
	return s
}

func (s *SubmitSlaCouponApplyRequest) SetEcIdAccountIds(v []*SubmitSlaCouponApplyRequestEcIdAccountIds) *SubmitSlaCouponApplyRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *SubmitSlaCouponApplyRequest) SetMonth(v int32) *SubmitSlaCouponApplyRequest {
	s.Month = &v
	return s
}

func (s *SubmitSlaCouponApplyRequest) SetNbid(v string) *SubmitSlaCouponApplyRequest {
	s.Nbid = &v
	return s
}

func (s *SubmitSlaCouponApplyRequest) Validate() error {
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

type SubmitSlaCouponApplyRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s SubmitSlaCouponApplyRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s SubmitSlaCouponApplyRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *SubmitSlaCouponApplyRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *SubmitSlaCouponApplyRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *SubmitSlaCouponApplyRequestEcIdAccountIds) SetAccountIds(v []*int64) *SubmitSlaCouponApplyRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *SubmitSlaCouponApplyRequestEcIdAccountIds) SetEcId(v string) *SubmitSlaCouponApplyRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *SubmitSlaCouponApplyRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
