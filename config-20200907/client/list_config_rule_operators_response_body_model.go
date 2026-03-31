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
	Operators []*ListConfigRuleOperatorsResponseBodyOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A68DD98C-DE65-46AC-B2D2-04A4A9AB5B99
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
	// example:
	//
	// String
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// StringEquals
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
