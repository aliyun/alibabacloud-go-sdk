// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanUserDeductRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds) *GetSavingPlanUserDeductRuleRequest
	GetEcIdAccountIds() []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds
	SetNbid(v string) *GetSavingPlanUserDeductRuleRequest
	GetNbid() *string
	SetPageSize(v int32) *GetSavingPlanUserDeductRuleRequest
	GetPageSize() *int32
	SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleRequest
	GetSpnInstanceCode() *string
}

type GetSavingPlanUserDeductRuleRequest struct {
	CurrentPage     *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds  []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                             `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode *string                                             `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanUserDeductRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSavingPlanUserDeductRuleRequest) GetEcIdAccountIds() []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *GetSavingPlanUserDeductRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanUserDeductRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSavingPlanUserDeductRuleRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *GetSavingPlanUserDeductRuleRequest) SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetEcIdAccountIds(v []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds) *GetSavingPlanUserDeductRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetNbid(v string) *GetSavingPlanUserDeductRuleRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetPageSize(v int32) *GetSavingPlanUserDeductRuleRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanUserDeductRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanUserDeductRuleRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
