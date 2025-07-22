// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanDeductableCommodityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIdsShrink(v string) *GetSavingPlanDeductableCommodityShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *GetSavingPlanDeductableCommodityShrinkRequest
	GetNbid() *string
}

type GetSavingPlanDeductableCommodityShrinkRequest struct {
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s GetSavingPlanDeductableCommodityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanDeductableCommodityShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) SetNbid(v string) *GetSavingPlanDeductableCommodityShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
