// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPeriodBudgetBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectIds(v string) *QueryPeriodBudgetBillRequest
	GetObjectIds() *string
	SetObjectType(v string) *QueryPeriodBudgetBillRequest
	GetObjectType() *string
	SetPeriod(v string) *QueryPeriodBudgetBillRequest
	GetPeriod() *string
}

type QueryPeriodBudgetBillRequest struct {
	ObjectIds  *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Period     *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryPeriodBudgetBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPeriodBudgetBillRequest) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillRequest) GetObjectIds() *string {
	return s.ObjectIds
}

func (s *QueryPeriodBudgetBillRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *QueryPeriodBudgetBillRequest) GetPeriod() *string {
	return s.Period
}

func (s *QueryPeriodBudgetBillRequest) SetObjectIds(v string) *QueryPeriodBudgetBillRequest {
	s.ObjectIds = &v
	return s
}

func (s *QueryPeriodBudgetBillRequest) SetObjectType(v string) *QueryPeriodBudgetBillRequest {
	s.ObjectType = &v
	return s
}

func (s *QueryPeriodBudgetBillRequest) SetPeriod(v string) *QueryPeriodBudgetBillRequest {
	s.Period = &v
	return s
}

func (s *QueryPeriodBudgetBillRequest) Validate() error {
	return dara.Validate(s)
}
