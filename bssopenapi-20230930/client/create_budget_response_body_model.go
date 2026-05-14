// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *CreateBudgetResponseBody
	GetBudgetName() *string
	SetRequestId(v string) *CreateBudgetResponseBody
	GetRequestId() *string
}

type CreateBudgetResponseBody struct {
	// example:
	//
	// Department_dev_budget
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// 39EDD65E-68C5-1B17-8440-C729C7591D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBudgetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBudgetResponseBody) GetBudgetName() *string {
	return s.BudgetName
}

func (s *CreateBudgetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBudgetResponseBody) SetBudgetName(v string) *CreateBudgetResponseBody {
	s.BudgetName = &v
	return s
}

func (s *CreateBudgetResponseBody) SetRequestId(v string) *CreateBudgetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBudgetResponseBody) Validate() error {
	return dara.Validate(s)
}
