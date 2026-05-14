// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DescribeBudgetsRequest
	GetBudgetName() *string
	SetBudgetType(v string) *DescribeBudgetsRequest
	GetBudgetType() *string
	SetExpireStatus(v string) *DescribeBudgetsRequest
	GetExpireStatus() *string
	SetNbid(v string) *DescribeBudgetsRequest
	GetNbid() *string
	SetPageNo(v int32) *DescribeBudgetsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeBudgetsRequest
	GetPageSize() *int32
}

type DescribeBudgetsRequest struct {
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// CONSUME
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// example:
	//
	// NOT_EXPIRED
	ExpireStatus *string `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	Nbid         *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBudgetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeBudgetsRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeBudgetsRequest) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *DescribeBudgetsRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeBudgetsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeBudgetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBudgetsRequest) SetBudgetName(v string) *DescribeBudgetsRequest {
	s.BudgetName = &v
	return s
}

func (s *DescribeBudgetsRequest) SetBudgetType(v string) *DescribeBudgetsRequest {
	s.BudgetType = &v
	return s
}

func (s *DescribeBudgetsRequest) SetExpireStatus(v string) *DescribeBudgetsRequest {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeBudgetsRequest) SetNbid(v string) *DescribeBudgetsRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeBudgetsRequest) SetPageNo(v int32) *DescribeBudgetsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeBudgetsRequest) SetPageSize(v int32) *DescribeBudgetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBudgetsRequest) Validate() error {
	return dara.Validate(s)
}
