// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *CreateCostCenterRuleRequest
	GetCostCenterId() *int64
	SetFilterExpression(v *CreateCostCenterRuleRequestFilterExpression) *CreateCostCenterRuleRequest
	GetFilterExpression() *CreateCostCenterRuleRequestFilterExpression
	SetNbid(v string) *CreateCostCenterRuleRequest
	GetNbid() *string
}

type CreateCostCenterRuleRequest struct {
	// The financial unit ID.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// The rule expression.
	FilterExpression *CreateCostCenterRuleRequestFilterExpression `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty" type:"Struct"`
	// The level-1 marketplace ID. If left empty, the marketplace ID of the current user is used by default.
	//
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CreateCostCenterRuleRequest) GetFilterExpression() *CreateCostCenterRuleRequestFilterExpression {
	return s.FilterExpression
}

func (s *CreateCostCenterRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateCostCenterRuleRequest) SetCostCenterId(v int64) *CreateCostCenterRuleRequest {
	s.CostCenterId = &v
	return s
}

func (s *CreateCostCenterRuleRequest) SetFilterExpression(v *CreateCostCenterRuleRequestFilterExpression) *CreateCostCenterRuleRequest {
	s.FilterExpression = v
	return s
}

func (s *CreateCostCenterRuleRequest) SetNbid(v string) *CreateCostCenterRuleRequest {
	s.Nbid = &v
	return s
}

func (s *CreateCostCenterRuleRequest) Validate() error {
	if s.FilterExpression != nil {
		if err := s.FilterExpression.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCostCenterRuleRequestFilterExpression struct {
	// The operation type.
	//
	// example:
	//
	// NARY
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// The sub-condition filter.
	FilterValues *CreateCostCenterRuleRequestFilterExpressionFilterValues `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Struct"`
	// The filter condition (deprecated).
	//
	// example:
	//
	// None
	Operand interface{} `json:"Operand,omitempty" xml:"Operand,omitempty"`
	// The conditional expression.
	Operands []interface{} `json:"Operands,omitempty" xml:"Operands,omitempty" type:"Repeated"`
	// The relational expression.
	//
	// example:
	//
	// AND
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
}

func (s CreateCostCenterRuleRequestFilterExpression) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleRequestFilterExpression) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleRequestFilterExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *CreateCostCenterRuleRequestFilterExpression) GetFilterValues() *CreateCostCenterRuleRequestFilterExpressionFilterValues {
	return s.FilterValues
}

func (s *CreateCostCenterRuleRequestFilterExpression) GetOperand() interface{} {
	return s.Operand
}

func (s *CreateCostCenterRuleRequestFilterExpression) GetOperands() []interface{} {
	return s.Operands
}

func (s *CreateCostCenterRuleRequestFilterExpression) GetOperatorType() *string {
	return s.OperatorType
}

func (s *CreateCostCenterRuleRequestFilterExpression) SetExpressionType(v string) *CreateCostCenterRuleRequestFilterExpression {
	s.ExpressionType = &v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpression) SetFilterValues(v *CreateCostCenterRuleRequestFilterExpressionFilterValues) *CreateCostCenterRuleRequestFilterExpression {
	s.FilterValues = v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpression) SetOperand(v interface{}) *CreateCostCenterRuleRequestFilterExpression {
	s.Operand = v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpression) SetOperands(v []interface{}) *CreateCostCenterRuleRequestFilterExpression {
	s.Operands = v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpression) SetOperatorType(v string) *CreateCostCenterRuleRequestFilterExpression {
	s.OperatorType = &v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpression) Validate() error {
	if s.FilterValues != nil {
		if err := s.FilterValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCostCenterRuleRequestFilterExpressionFilterValues struct {
	// The filter condition key.
	//
	// example:
	//
	// TAG-test-xxx-key
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The filter condition key name (deprecated).
	//
	// example:
	//
	// None
	CodeName *string `json:"CodeName,omitempty" xml:"CodeName,omitempty"`
	// The association between the code and value.
	//
	// example:
	//
	// IN
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The list of filter condition values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateCostCenterRuleRequestFilterExpressionFilterValues) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleRequestFilterExpressionFilterValues) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) GetCode() *string {
	return s.Code
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) GetCodeName() *string {
	return s.CodeName
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) GetSelectType() *string {
	return s.SelectType
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) GetValues() []*string {
	return s.Values
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) SetCode(v string) *CreateCostCenterRuleRequestFilterExpressionFilterValues {
	s.Code = &v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) SetCodeName(v string) *CreateCostCenterRuleRequestFilterExpressionFilterValues {
	s.CodeName = &v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) SetSelectType(v string) *CreateCostCenterRuleRequestFilterExpressionFilterValues {
	s.SelectType = &v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) SetValues(v []*string) *CreateCostCenterRuleRequestFilterExpressionFilterValues {
	s.Values = v
	return s
}

func (s *CreateCostCenterRuleRequestFilterExpressionFilterValues) Validate() error {
	return dara.Validate(s)
}
