// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleOperatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperators(v []*ListConfigRuleOperatorsResponseBodyOperators) *ListConfigRuleOperatorsResponseBody
	GetOperators() []*ListConfigRuleOperatorsResponseBodyOperators
	SetRequestId(v string) *ListConfigRuleOperatorsResponseBody
	GetRequestId() *string
}

type ListConfigRuleOperatorsResponseBody struct {
	// A list of operators.
	Operators []*ListConfigRuleOperatorsResponseBodyOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A68DD98C-DE65-46AC-B2D2-04A4A9AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRuleOperatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleOperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRuleOperatorsResponseBody) GetOperators() []*ListConfigRuleOperatorsResponseBodyOperators {
	return s.Operators
}

func (s *ListConfigRuleOperatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigRuleOperatorsResponseBody) SetOperators(v []*ListConfigRuleOperatorsResponseBodyOperators) *ListConfigRuleOperatorsResponseBody {
	s.Operators = v
	return s
}

func (s *ListConfigRuleOperatorsResponseBody) SetRequestId(v string) *ListConfigRuleOperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigRuleOperatorsResponseBody) Validate() error {
	if s.Operators != nil {
		for _, item := range s.Operators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigRuleOperatorsResponseBodyOperators struct {
	// The data type that the operator applies to.
	//
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The description of the operator, which can be used to explain why a resource is non-compliant.
	//
	// example:
	//
	// The current value must equal the target value (case insensitive)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// StringEquals
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operator for the rule\\"s input parameter. The available operators vary based on the data type retrieved using SelectPath.
	//
	// - If the data type is String, valid values are:
	//
	//   - StringEquals: equals.
	//
	//   - NotStringEquals: does not equal.
	//
	//   - StringIn: is in.
	//
	//   - NotStringIn: is not in.
	//
	//   - StringContains: contains.
	//
	//   - NotStringContains: does not contain.
	//
	// - If the data type is Number, valid values are:
	//
	//   - Equals: equals.
	//
	//   - NotEquals: does not equal.
	//
	//   - Less: is less than.
	//
	//   - LessOrEquals: is less than or equal to.
	//
	//   - Greater: is greater than.
	//
	//   - GreaterOrEquals: is greater than or equal to.
	//
	// - If the data type is a Base64-encoded string, valid values are:
	//
	//   - Base64Contains: contains.
	//
	//   - NotBase64Contains: does not contain.
	//
	//   - Base64ContainsAll: contains all.
	//
	//   - Base64ExcludeAll: excludes all.
	//
	// - If the data type is Array, valid values are:
	//
	//   - Contains: contains.
	//
	//   - NotContains: does not contain.
	//
	//   - In: is in.
	//
	//   - NotIn: is not in.
	//
	//   - ContainsAll: contains all.
	//
	//   - ExcludeAll: excludes all.
	//
	//   - IsEmpty: is empty.
	//
	// example:
	//
	// StringEquals
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s ListConfigRuleOperatorsResponseBodyOperators) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleOperatorsResponseBodyOperators) GoString() string {
	return s.String()
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) GetDataType() *string {
	return s.DataType
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) GetDescription() *string {
	return s.Description
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) GetName() *string {
	return s.Name
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) GetOperator() *string {
	return s.Operator
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) SetDataType(v string) *ListConfigRuleOperatorsResponseBodyOperators {
	s.DataType = &v
	return s
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) SetDescription(v string) *ListConfigRuleOperatorsResponseBodyOperators {
	s.Description = &v
	return s
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) SetName(v string) *ListConfigRuleOperatorsResponseBodyOperators {
	s.Name = &v
	return s
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) SetOperator(v string) *ListConfigRuleOperatorsResponseBodyOperators {
	s.Operator = &v
	return s
}

func (s *ListConfigRuleOperatorsResponseBodyOperators) Validate() error {
	return dara.Validate(s)
}
