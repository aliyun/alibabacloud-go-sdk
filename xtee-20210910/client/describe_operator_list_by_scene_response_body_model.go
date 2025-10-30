// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListBySceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOperatorListBySceneResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeOperatorListBySceneResponseBodyResultObject) *DescribeOperatorListBySceneResponseBody
	GetResultObject() []*DescribeOperatorListBySceneResponseBodyResultObject
}

type DescribeOperatorListBySceneResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject []*DescribeOperatorListBySceneResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListBySceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperatorListBySceneResponseBody) GetResultObject() []*DescribeOperatorListBySceneResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOperatorListBySceneResponseBody) SetRequestId(v string) *DescribeOperatorListBySceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBody) SetResultObject(v []*DescribeOperatorListBySceneResponseBodyResultObject) *DescribeOperatorListBySceneResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOperatorListBySceneResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOperatorListBySceneResponseBodyResultObject struct {
	// Return value type
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Operator list
	Operators []*DescribeOperatorListBySceneResponseBodyResultObjectOperators `json:"operators,omitempty" xml:"operators,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListBySceneResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeOperatorListBySceneResponseBodyResultObject) GetOperators() []*DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	return s.Operators
}

func (s *DescribeOperatorListBySceneResponseBodyResultObject) SetFieldType(v string) *DescribeOperatorListBySceneResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObject) SetOperators(v []*DescribeOperatorListBySceneResponseBodyResultObjectOperators) *DescribeOperatorListBySceneResponseBodyResultObject {
	s.Operators = v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObject) Validate() error {
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

type DescribeOperatorListBySceneResponseBodyResultObjectOperators struct {
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
	RightVariables []*DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables `json:"rightVariables,omitempty" xml:"rightVariables,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListBySceneResponseBodyResultObjectOperators) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneResponseBodyResultObjectOperators) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) GetCode() *string {
	return s.Code
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) GetHasRightVariable() *bool {
	return s.HasRightVariable
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) GetMemo() *string {
	return s.Memo
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) GetName() *string {
	return s.Name
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) GetRightVariables() []*DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables {
	return s.RightVariables
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) SetCode(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	s.Code = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) SetHasRightVariable(v bool) *DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	s.HasRightVariable = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) SetMemo(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	s.Memo = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) SetName(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	s.Name = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) SetRightVariables(v []*DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) *DescribeOperatorListBySceneResponseBodyResultObjectOperators {
	s.RightVariables = v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperators) Validate() error {
	if s.RightVariables != nil {
		for _, item := range s.RightVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables struct {
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

func (s DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) SetFieldName(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables {
	s.FieldName = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) SetFieldType(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) SetFieldValue(v string) *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables {
	s.FieldValue = &v
	return s
}

func (s *DescribeOperatorListBySceneResponseBodyResultObjectOperatorsRightVariables) Validate() error {
	return dara.Validate(s)
}
