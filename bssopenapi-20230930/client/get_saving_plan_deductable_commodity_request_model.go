// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanDeductableCommodityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds) *GetSavingPlanDeductableCommodityRequest
	GetEcIdAccountIds() []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds
	SetNbid(v string) *GetSavingPlanDeductableCommodityRequest
	GetNbid() *string
}

type GetSavingPlanDeductableCommodityRequest struct {
	EcIdAccountIds []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                                  `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s GetSavingPlanDeductableCommodityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityRequest) GetEcIdAccountIds() []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *GetSavingPlanDeductableCommodityRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanDeductableCommodityRequest) SetEcIdAccountIds(v []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds) *GetSavingPlanDeductableCommodityRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequest) SetNbid(v string) *GetSavingPlanDeductableCommodityRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequest) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanDeductableCommodityRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanDeductableCommodityRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanDeductableCommodityRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
