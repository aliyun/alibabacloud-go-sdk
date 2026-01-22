// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *DeleteCostCenterRuleRequest
	GetCostCenterId() *int64
	SetFilterExpression(v *DeleteCostCenterRuleRequestFilterExpression) *DeleteCostCenterRuleRequest
	GetFilterExpression() *DeleteCostCenterRuleRequestFilterExpression
	SetNbid(v string) *DeleteCostCenterRuleRequest
	GetNbid() *string
}

type DeleteCostCenterRuleRequest struct {
	// example:
	//
	// 637127
	CostCenterId     *int64                                       `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpression *DeleteCostCenterRuleRequestFilterExpression `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty" type:"Struct"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DeleteCostCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *DeleteCostCenterRuleRequest) GetFilterExpression() *DeleteCostCenterRuleRequestFilterExpression {
	return s.FilterExpression
}

func (s *DeleteCostCenterRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteCostCenterRuleRequest) SetCostCenterId(v int64) *DeleteCostCenterRuleRequest {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterRuleRequest) SetFilterExpression(v *DeleteCostCenterRuleRequestFilterExpression) *DeleteCostCenterRuleRequest {
	s.FilterExpression = v
	return s
}

func (s *DeleteCostCenterRuleRequest) SetNbid(v string) *DeleteCostCenterRuleRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCostCenterRuleRequest) Validate() error {
	if s.FilterExpression != nil {
		if err := s.FilterExpression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCostCenterRuleRequestFilterExpression struct {
	// example:
	//
	// NARY
	ExpressionType *string                                                  `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	FilterValues   *DeleteCostCenterRuleRequestFilterExpressionFilterValues `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Struct"`
	Operand        interface{}                                              `json:"Operand,omitempty" xml:"Operand,omitempty"`
	Operands       []interface{}                                            `json:"Operands,omitempty" xml:"Operands,omitempty" type:"Repeated"`
	// example:
	//
	// AND
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
}

func (s DeleteCostCenterRuleRequestFilterExpression) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleRequestFilterExpression) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleRequestFilterExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *DeleteCostCenterRuleRequestFilterExpression) GetFilterValues() *DeleteCostCenterRuleRequestFilterExpressionFilterValues {
	return s.FilterValues
}

func (s *DeleteCostCenterRuleRequestFilterExpression) GetOperand() interface{} {
	return s.Operand
}

func (s *DeleteCostCenterRuleRequestFilterExpression) GetOperands() []interface{} {
	return s.Operands
}

func (s *DeleteCostCenterRuleRequestFilterExpression) GetOperatorType() *string {
	return s.OperatorType
}

func (s *DeleteCostCenterRuleRequestFilterExpression) SetExpressionType(v string) *DeleteCostCenterRuleRequestFilterExpression {
	s.ExpressionType = &v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpression) SetFilterValues(v *DeleteCostCenterRuleRequestFilterExpressionFilterValues) *DeleteCostCenterRuleRequestFilterExpression {
	s.FilterValues = v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpression) SetOperand(v interface{}) *DeleteCostCenterRuleRequestFilterExpression {
	s.Operand = v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpression) SetOperands(v []interface{}) *DeleteCostCenterRuleRequestFilterExpression {
	s.Operands = v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpression) SetOperatorType(v string) *DeleteCostCenterRuleRequestFilterExpression {
	s.OperatorType = &v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpression) Validate() error {
	if s.FilterValues != nil {
		if err := s.FilterValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCostCenterRuleRequestFilterExpressionFilterValues struct {
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

func (s DeleteCostCenterRuleRequestFilterExpressionFilterValues) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleRequestFilterExpressionFilterValues) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) GetCode() *string {
	return s.Code
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) GetCodeName() *string {
	return s.CodeName
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) GetSelectType() *string {
	return s.SelectType
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) GetValues() []*string {
	return s.Values
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) SetCode(v string) *DeleteCostCenterRuleRequestFilterExpressionFilterValues {
	s.Code = &v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) SetCodeName(v string) *DeleteCostCenterRuleRequestFilterExpressionFilterValues {
	s.CodeName = &v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) SetSelectType(v string) *DeleteCostCenterRuleRequestFilterExpressionFilterValues {
	s.SelectType = &v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) SetValues(v []*string) *DeleteCostCenterRuleRequestFilterExpressionFilterValues {
	s.Values = v
	return s
}

func (s *DeleteCostCenterRuleRequestFilterExpressionFilterValues) Validate() error {
	return dara.Validate(s)
}
