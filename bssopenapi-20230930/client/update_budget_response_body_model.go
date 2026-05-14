// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBudgetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *UpdateBudgetResponseBody
	GetBudgetName() *string
	SetRequestId(v string) *UpdateBudgetResponseBody
	GetRequestId() *string
}

type UpdateBudgetResponseBody struct {
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// 03A59CD4-6C6B-1A62-B64C-F1F2AF7830F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBudgetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBudgetResponseBody) GetBudgetName() *string {
	return s.BudgetName
}

func (s *UpdateBudgetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBudgetResponseBody) SetBudgetName(v string) *UpdateBudgetResponseBody {
	s.BudgetName = &v
	return s
}

func (s *UpdateBudgetResponseBody) SetRequestId(v string) *UpdateBudgetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBudgetResponseBody) Validate() error {
	return dara.Validate(s)
}
