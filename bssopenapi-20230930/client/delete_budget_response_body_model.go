// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DeleteBudgetResponseBody
	GetBudgetName() *string
	SetRequestId(v string) *DeleteBudgetResponseBody
	GetRequestId() *string
}

type DeleteBudgetResponseBody struct {
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// 7EA6C02D-06D0-4213-9C3B-E67910F7D1EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBudgetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBudgetResponseBody) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DeleteBudgetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBudgetResponseBody) SetBudgetName(v string) *DeleteBudgetResponseBody {
	s.BudgetName = &v
	return s
}

func (s *DeleteBudgetResponseBody) SetRequestId(v string) *DeleteBudgetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBudgetResponseBody) Validate() error {
	return dara.Validate(s)
}
