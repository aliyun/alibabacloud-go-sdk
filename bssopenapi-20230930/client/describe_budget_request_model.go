// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DescribeBudgetRequest
	GetBudgetName() *string
	SetNbid(v string) *DescribeBudgetRequest
	GetNbid() *string
}

type DescribeBudgetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	Nbid       *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DescribeBudgetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetRequest) GoString() string {
	return s.String()
}

func (s *DescribeBudgetRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeBudgetRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeBudgetRequest) SetBudgetName(v string) *DescribeBudgetRequest {
	s.BudgetName = &v
	return s
}

func (s *DescribeBudgetRequest) SetNbid(v string) *DescribeBudgetRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeBudgetRequest) Validate() error {
	return dara.Validate(s)
}
