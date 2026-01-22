// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *QueryCostCenterRuleResponseBody
	GetCostCenterId() *int64
	SetFilterExpression(v *QueryCostCenterRuleResponseBodyFilterExpression) *QueryCostCenterRuleResponseBody
	GetFilterExpression() *QueryCostCenterRuleResponseBodyFilterExpression
	SetGmtCreate(v string) *QueryCostCenterRuleResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *QueryCostCenterRuleResponseBody
	GetGmtModified() *string
	SetId(v int64) *QueryCostCenterRuleResponseBody
	GetId() *int64
	SetIsDeleted(v int32) *QueryCostCenterRuleResponseBody
	GetIsDeleted() *int32
	SetMetadata(v interface{}) *QueryCostCenterRuleResponseBody
	GetMetadata() interface{}
	SetOwnerAccountId(v int64) *QueryCostCenterRuleResponseBody
	GetOwnerAccountId() *int64
	SetRequestId(v string) *QueryCostCenterRuleResponseBody
	GetRequestId() *string
	SetRootCostCenterId(v int64) *QueryCostCenterRuleResponseBody
	GetRootCostCenterId() *int64
	SetStatus(v string) *QueryCostCenterRuleResponseBody
	GetStatus() *string
}

type QueryCostCenterRuleResponseBody struct {
	// example:
	//
	// 597745
	CostCenterId     *int64                                           `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpression *QueryCostCenterRuleResponseBodyFilterExpression `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty" type:"Struct"`
	// example:
	//
	// Tue Nov 12 14:49:43 CST 2024
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// Wed Oct 16 10:15:37 CST 2024
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 32048
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsDeleted *int32 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1977800748053695
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	RootCostCenterId *int64 `json:"RootCostCenterId,omitempty" xml:"RootCostCenterId,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryCostCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleResponseBody) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterRuleResponseBody) GetFilterExpression() *QueryCostCenterRuleResponseBodyFilterExpression {
	return s.FilterExpression
}

func (s *QueryCostCenterRuleResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryCostCenterRuleResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryCostCenterRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *QueryCostCenterRuleResponseBody) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *QueryCostCenterRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryCostCenterRuleResponseBody) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostCenterRuleResponseBody) GetRootCostCenterId() *int64 {
	return s.RootCostCenterId
}

func (s *QueryCostCenterRuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryCostCenterRuleResponseBody) SetCostCenterId(v int64) *QueryCostCenterRuleResponseBody {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetFilterExpression(v *QueryCostCenterRuleResponseBodyFilterExpression) *QueryCostCenterRuleResponseBody {
	s.FilterExpression = v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetGmtCreate(v string) *QueryCostCenterRuleResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetGmtModified(v string) *QueryCostCenterRuleResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetId(v int64) *QueryCostCenterRuleResponseBody {
	s.Id = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetIsDeleted(v int32) *QueryCostCenterRuleResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetMetadata(v interface{}) *QueryCostCenterRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetOwnerAccountId(v int64) *QueryCostCenterRuleResponseBody {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetRequestId(v string) *QueryCostCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetRootCostCenterId(v int64) *QueryCostCenterRuleResponseBody {
	s.RootCostCenterId = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) SetStatus(v string) *QueryCostCenterRuleResponseBody {
	s.Status = &v
	return s
}

func (s *QueryCostCenterRuleResponseBody) Validate() error {
	if s.FilterExpression != nil {
		if err := s.FilterExpression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCostCenterRuleResponseBodyFilterExpression struct {
	// example:
	//
	// NARY
	ExpressionType *string                                                      `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	FilterValues   *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Struct"`
	Operand        interface{}                                                  `json:"Operand,omitempty" xml:"Operand,omitempty"`
	Operands       []interface{}                                                `json:"Operands,omitempty" xml:"Operands,omitempty" type:"Repeated"`
	// example:
	//
	// AND
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
}

func (s QueryCostCenterRuleResponseBodyFilterExpression) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleResponseBodyFilterExpression) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) GetFilterValues() *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues {
	return s.FilterValues
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) GetOperand() interface{} {
	return s.Operand
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) GetOperands() []interface{} {
	return s.Operands
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) GetOperatorType() *string {
	return s.OperatorType
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) SetExpressionType(v string) *QueryCostCenterRuleResponseBodyFilterExpression {
	s.ExpressionType = &v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) SetFilterValues(v *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) *QueryCostCenterRuleResponseBodyFilterExpression {
	s.FilterValues = v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) SetOperand(v interface{}) *QueryCostCenterRuleResponseBodyFilterExpression {
	s.Operand = v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) SetOperands(v []interface{}) *QueryCostCenterRuleResponseBodyFilterExpression {
	s.Operands = v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) SetOperatorType(v string) *QueryCostCenterRuleResponseBodyFilterExpression {
	s.OperatorType = &v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpression) Validate() error {
	if s.FilterValues != nil {
		if err := s.FilterValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCostCenterRuleResponseBodyFilterExpressionFilterValues struct {
	// example:
	//
	// TAG-test-xxx-key
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// example:
	//
	// IN
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) GetCode() *string {
	return s.Code
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) GetCodeName() *string {
	return s.CodeName
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) GetSelectType() *string {
	return s.SelectType
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) GetValues() []*string {
	return s.Values
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) SetCode(v string) *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues {
	s.Code = &v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) SetCodeName(v string) *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues {
	s.CodeName = &v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) SetSelectType(v string) *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues {
	s.SelectType = &v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) SetValues(v []*string) *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues {
	s.Values = v
	return s
}

func (s *QueryCostCenterRuleResponseBodyFilterExpressionFilterValues) Validate() error {
	return dara.Validate(s)
}
