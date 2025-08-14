// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOperatorListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeOperatorListResponseBodyResultObject) *DescribeOperatorListResponseBody
	GetResultObject() []*DescribeOperatorListResponseBodyResultObject
}

type DescribeOperatorListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeOperatorListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperatorListResponseBody) GetResultObject() []*DescribeOperatorListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOperatorListResponseBody) SetRequestId(v string) *DescribeOperatorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorListResponseBody) SetResultObject(v []*DescribeOperatorListResponseBodyResultObject) *DescribeOperatorListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOperatorListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOperatorListResponseBodyResultObject struct {
	// Return value type
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Operator list
	Operators []*DescribeOperatorListResponseBodyResultObjectOperators `json:"operators,omitempty" xml:"operators,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeOperatorListResponseBodyResultObject) GetOperators() []*DescribeOperatorListResponseBodyResultObjectOperators {
	return s.Operators
}

func (s *DescribeOperatorListResponseBodyResultObject) SetFieldType(v string) *DescribeOperatorListResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObject) SetOperators(v []*DescribeOperatorListResponseBodyResultObjectOperators) *DescribeOperatorListResponseBodyResultObject {
	s.Operators = v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeOperatorListResponseBodyResultObjectOperators struct {
	// Operator code
	//
	// example:
	//
	// equals
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Whether it contains a right variable
	//
	// example:
	//
	// true
	HasRightVariable *bool `json:"hasRightVariable,omitempty" xml:"hasRightVariable,omitempty"`
	// Description
	//
	// example:
	//
	// 等于
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Operator name
	//
	// example:
	//
	// 等于
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Right variable object
	RightVariables []*DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables `json:"rightVariables,omitempty" xml:"rightVariables,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListResponseBodyResultObjectOperators) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListResponseBodyResultObjectOperators) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) GetCode() *string {
	return s.Code
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) GetHasRightVariable() *bool {
	return s.HasRightVariable
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) GetMemo() *string {
	return s.Memo
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) GetName() *string {
	return s.Name
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) GetRightVariables() []*DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables {
	return s.RightVariables
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) SetCode(v string) *DescribeOperatorListResponseBodyResultObjectOperators {
	s.Code = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) SetHasRightVariable(v bool) *DescribeOperatorListResponseBodyResultObjectOperators {
	s.HasRightVariable = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) SetMemo(v string) *DescribeOperatorListResponseBodyResultObjectOperators {
	s.Memo = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) SetName(v string) *DescribeOperatorListResponseBodyResultObjectOperators {
	s.Name = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) SetRightVariables(v []*DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) *DescribeOperatorListResponseBodyResultObjectOperators {
	s.RightVariables = v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperators) Validate() error {
	return dara.Validate(s)
}

type DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables struct {
	// Field name.
	//
	// example:
	//
	// 年龄
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field type.
	//
	// example:
	//
	// INT
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Field value.
	//
	// example:
	//
	// 20
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) SetFieldName(v string) *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables {
	s.FieldName = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) SetFieldType(v string) *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) SetFieldValue(v string) *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables {
	s.FieldValue = &v
	return s
}

func (s *DescribeOperatorListResponseBodyResultObjectOperatorsRightVariables) Validate() error {
	return dara.Validate(s)
}
