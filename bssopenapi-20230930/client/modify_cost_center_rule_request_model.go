// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *ModifyCostCenterRuleRequest
	GetCostCenterId() *int64
	SetFilterExpression(v *ModifyCostCenterRuleRequestFilterExpression) *ModifyCostCenterRuleRequest
	GetFilterExpression() *ModifyCostCenterRuleRequestFilterExpression
	SetNbid(v string) *ModifyCostCenterRuleRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *ModifyCostCenterRuleRequest
	GetOwnerAccountId() *int64
}

type ModifyCostCenterRuleRequest struct {
	// example:
	//
	// 485938
	CostCenterId     *int64                                       `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpression *ModifyCostCenterRuleRequestFilterExpression `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty" type:"Struct"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ModifyCostCenterRuleRequest) GetFilterExpression() *ModifyCostCenterRuleRequestFilterExpression {
	return s.FilterExpression
}

func (s *ModifyCostCenterRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ModifyCostCenterRuleRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *ModifyCostCenterRuleRequest) SetCostCenterId(v int64) *ModifyCostCenterRuleRequest {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterRuleRequest) SetFilterExpression(v *ModifyCostCenterRuleRequestFilterExpression) *ModifyCostCenterRuleRequest {
	s.FilterExpression = v
	return s
}

func (s *ModifyCostCenterRuleRequest) SetNbid(v string) *ModifyCostCenterRuleRequest {
	s.Nbid = &v
	return s
}

func (s *ModifyCostCenterRuleRequest) SetOwnerAccountId(v int64) *ModifyCostCenterRuleRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *ModifyCostCenterRuleRequest) Validate() error {
	if s.FilterExpression != nil {
		if err := s.FilterExpression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCostCenterRuleRequestFilterExpression struct {
	// example:
	//
	// NARY
	ExpressionType *string                                                  `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	FilterValues   *ModifyCostCenterRuleRequestFilterExpressionFilterValues `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Struct"`
	Operand        interface{}                                              `json:"Operand,omitempty" xml:"Operand,omitempty"`
	Operands       []interface{}                                            `json:"Operands,omitempty" xml:"Operands,omitempty" type:"Repeated"`
	// example:
	//
	// AND
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
}

func (s ModifyCostCenterRuleRequestFilterExpression) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleRequestFilterExpression) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleRequestFilterExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *ModifyCostCenterRuleRequestFilterExpression) GetFilterValues() *ModifyCostCenterRuleRequestFilterExpressionFilterValues {
	return s.FilterValues
}

func (s *ModifyCostCenterRuleRequestFilterExpression) GetOperand() interface{} {
	return s.Operand
}

func (s *ModifyCostCenterRuleRequestFilterExpression) GetOperands() []interface{} {
	return s.Operands
}

func (s *ModifyCostCenterRuleRequestFilterExpression) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ModifyCostCenterRuleRequestFilterExpression) SetExpressionType(v string) *ModifyCostCenterRuleRequestFilterExpression {
	s.ExpressionType = &v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpression) SetFilterValues(v *ModifyCostCenterRuleRequestFilterExpressionFilterValues) *ModifyCostCenterRuleRequestFilterExpression {
	s.FilterValues = v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpression) SetOperand(v interface{}) *ModifyCostCenterRuleRequestFilterExpression {
	s.Operand = v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpression) SetOperands(v []interface{}) *ModifyCostCenterRuleRequestFilterExpression {
	s.Operands = v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpression) SetOperatorType(v string) *ModifyCostCenterRuleRequestFilterExpression {
	s.OperatorType = &v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpression) Validate() error {
	if s.FilterValues != nil {
		if err := s.FilterValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCostCenterRuleRequestFilterExpressionFilterValues struct {
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

func (s ModifyCostCenterRuleRequestFilterExpressionFilterValues) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleRequestFilterExpressionFilterValues) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) GetCode() *string {
	return s.Code
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) GetCodeName() *string {
	return s.CodeName
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) GetSelectType() *string {
	return s.SelectType
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) GetValues() []*string {
	return s.Values
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) SetCode(v string) *ModifyCostCenterRuleRequestFilterExpressionFilterValues {
	s.Code = &v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) SetCodeName(v string) *ModifyCostCenterRuleRequestFilterExpressionFilterValues {
	s.CodeName = &v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) SetSelectType(v string) *ModifyCostCenterRuleRequestFilterExpressionFilterValues {
	s.SelectType = &v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) SetValues(v []*string) *ModifyCostCenterRuleRequestFilterExpressionFilterValues {
	s.Values = v
	return s
}

func (s *ModifyCostCenterRuleRequestFilterExpressionFilterValues) Validate() error {
	return dara.Validate(s)
}
