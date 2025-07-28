// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPeriodBudgetBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPeriodBudgetBillDTOS(v []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) *QueryPeriodBudgetBillResponseBody
	GetPeriodBudgetBillDTOS() []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS
}

type QueryPeriodBudgetBillResponseBody struct {
	PeriodBudgetBillDTOS []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS `json:"periodBudgetBillDTOS,omitempty" xml:"periodBudgetBillDTOS,omitempty" type:"Repeated"`
}

func (s QueryPeriodBudgetBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPeriodBudgetBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponseBody) GetPeriodBudgetBillDTOS() []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	return s.PeriodBudgetBillDTOS
}

func (s *QueryPeriodBudgetBillResponseBody) SetPeriodBudgetBillDTOS(v []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) *QueryPeriodBudgetBillResponseBody {
	s.PeriodBudgetBillDTOS = v
	return s
}

func (s *QueryPeriodBudgetBillResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS struct {
	Bill         *float64 `json:"bill,omitempty" xml:"bill,omitempty"`
	Budget       *float64 `json:"budget,omitempty" xml:"budget,omitempty"`
	LastYearBill *float64 `json:"lastYearBill,omitempty" xml:"lastYearBill,omitempty"`
	Month        *string  `json:"month,omitempty" xml:"month,omitempty"`
}

func (s QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) String() string {
	return dara.Prettify(s)
}

func (s QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GetBill() *float64 {
	return s.Bill
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GetBudget() *float64 {
	return s.Budget
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GetLastYearBill() *float64 {
	return s.LastYearBill
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GetMonth() *string {
	return s.Month
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetBill(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Bill = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetBudget(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Budget = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetLastYearBill(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.LastYearBill = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetMonth(v string) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Month = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) Validate() error {
	return dara.Validate(s)
}
