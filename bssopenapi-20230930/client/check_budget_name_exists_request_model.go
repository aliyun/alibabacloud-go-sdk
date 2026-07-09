// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBudgetNameExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *CheckBudgetNameExistsRequest
	GetBudgetName() *string
	SetNbid(v string) *CheckBudgetNameExistsRequest
	GetNbid() *string
}

type CheckBudgetNameExistsRequest struct {
	// The budget name.
	//
	// This parameter is required.
	//
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// The primary sales channel ID. If this parameter is left empty, the channel ID of the current user is used by default.
	//
	// example:
	//
	// 2688801000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CheckBudgetNameExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBudgetNameExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckBudgetNameExistsRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *CheckBudgetNameExistsRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CheckBudgetNameExistsRequest) SetBudgetName(v string) *CheckBudgetNameExistsRequest {
	s.BudgetName = &v
	return s
}

func (s *CheckBudgetNameExistsRequest) SetNbid(v string) *CheckBudgetNameExistsRequest {
	s.Nbid = &v
	return s
}

func (s *CheckBudgetNameExistsRequest) Validate() error {
	return dara.Validate(s)
}
