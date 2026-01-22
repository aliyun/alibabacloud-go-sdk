// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanShareAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSavingPlanShareAccountsRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*GetSavingPlanShareAccountsRequestEcIdAccountIds) *GetSavingPlanShareAccountsRequest
	GetEcIdAccountIds() []*GetSavingPlanShareAccountsRequestEcIdAccountIds
	SetNbid(v string) *GetSavingPlanShareAccountsRequest
	GetNbid() *string
	SetPageSize(v int32) *GetSavingPlanShareAccountsRequest
	GetPageSize() *int32
	SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsRequest
	GetSpnInstanceCode() *string
}

type GetSavingPlanShareAccountsRequest struct {
	CurrentPage     *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds  []*GetSavingPlanShareAccountsRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                            `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize        *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode *string                                            `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanShareAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSavingPlanShareAccountsRequest) GetEcIdAccountIds() []*GetSavingPlanShareAccountsRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *GetSavingPlanShareAccountsRequest) GetNbid() *string {
	return s.Nbid
}

func (s *GetSavingPlanShareAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSavingPlanShareAccountsRequest) GetSpnInstanceCode() *string {
	return s.SpnInstanceCode
}

func (s *GetSavingPlanShareAccountsRequest) SetCurrentPage(v int32) *GetSavingPlanShareAccountsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetEcIdAccountIds(v []*GetSavingPlanShareAccountsRequestEcIdAccountIds) *GetSavingPlanShareAccountsRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetNbid(v string) *GetSavingPlanShareAccountsRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetPageSize(v int32) *GetSavingPlanShareAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) Validate() error {
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

type GetSavingPlanShareAccountsRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanShareAccountsRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanShareAccountsRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanShareAccountsRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
