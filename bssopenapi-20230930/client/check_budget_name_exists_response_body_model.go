// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBudgetNameExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *CheckBudgetNameExistsResponseBody
	GetBudgetName() *string
	SetExists(v bool) *CheckBudgetNameExistsResponseBody
	GetExists() *bool
	SetRequestId(v string) *CheckBudgetNameExistsResponseBody
	GetRequestId() *string
}

type CheckBudgetNameExistsResponseBody struct {
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// true
	Exists *bool `json:"Exists,omitempty" xml:"Exists,omitempty"`
	// example:
	//
	// F1E2D3C4-B5A6-7890-1234-567890ABCDEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckBudgetNameExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckBudgetNameExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBudgetNameExistsResponseBody) GetBudgetName() *string {
	return s.BudgetName
}

func (s *CheckBudgetNameExistsResponseBody) GetExists() *bool {
	return s.Exists
}

func (s *CheckBudgetNameExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckBudgetNameExistsResponseBody) SetBudgetName(v string) *CheckBudgetNameExistsResponseBody {
	s.BudgetName = &v
	return s
}

func (s *CheckBudgetNameExistsResponseBody) SetExists(v bool) *CheckBudgetNameExistsResponseBody {
	s.Exists = &v
	return s
}

func (s *CheckBudgetNameExistsResponseBody) SetRequestId(v string) *CheckBudgetNameExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBudgetNameExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
