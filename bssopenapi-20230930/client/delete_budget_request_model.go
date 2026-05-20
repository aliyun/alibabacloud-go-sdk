// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DeleteBudgetRequest
	GetBudgetName() *string
	SetNbid(v string) *DeleteBudgetRequest
	GetNbid() *string
}

type DeleteBudgetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DeleteBudgetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetRequest) GoString() string {
	return s.String()
}

func (s *DeleteBudgetRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DeleteBudgetRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteBudgetRequest) SetBudgetName(v string) *DeleteBudgetRequest {
	s.BudgetName = &v
	return s
}

func (s *DeleteBudgetRequest) SetNbid(v string) *DeleteBudgetRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteBudgetRequest) Validate() error {
	return dara.Validate(s)
}
