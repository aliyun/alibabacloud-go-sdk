// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanShareAccountsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSavingPlanShareAccountsShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *GetSavingPlanShareAccountsShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *GetSavingPlanShareAccountsShrinkRequest
	GetNbid() *string
	SetPageSize(v int32) *GetSavingPlanShareAccountsShrinkRequest
	GetPageSize() *int32
	SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsShrinkRequest
	GetSpnInstanceCode() *string
}

type GetSavingPlanShareAccountsShrinkRequest struct {
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode      *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanShareAccountsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSavingPlanShareAccountsShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *GetSavingPlanShareAccountsShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanShareAccountsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSavingPlanShareAccountsShrinkRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetCurrentPage(v int32) *GetSavingPlanShareAccountsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetNbid(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetPageSize(v int32) *GetSavingPlanShareAccountsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
