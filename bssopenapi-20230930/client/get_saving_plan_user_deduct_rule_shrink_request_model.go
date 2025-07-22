// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanUserDeductRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *GetSavingPlanUserDeductRuleShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *GetSavingPlanUserDeductRuleShrinkRequest
	GetNbid() *string
	SetPageSize(v int32) *GetSavingPlanUserDeductRuleShrinkRequest
	GetPageSize() *int32
	SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleShrinkRequest
	GetSpnInstanceCode() *string
}

type GetSavingPlanUserDeductRuleShrinkRequest struct {
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode      *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanUserDeductRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetNbid(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetPageSize(v int32) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
