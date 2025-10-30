// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOperatorListByTypeResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeOperatorListByTypeResponseBodyResultObject) *DescribeOperatorListByTypeResponseBody
	GetResultObject() []*DescribeOperatorListByTypeResponseBodyResultObject
}

type DescribeOperatorListByTypeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeOperatorListByTypeResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperatorListByTypeResponseBody) GetResultObject() []*DescribeOperatorListByTypeResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOperatorListByTypeResponseBody) SetRequestId(v string) *DescribeOperatorListByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorListByTypeResponseBody) SetResultObject(v []*DescribeOperatorListByTypeResponseBodyResultObject) *DescribeOperatorListByTypeResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOperatorListByTypeResponseBody) Validate() error {
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

type DescribeOperatorListByTypeResponseBodyResultObject struct {
	// Return value type
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Operator list
	Operators []*DescribeOperatorListByTypeResponseBodyResultObjectOperators `json:"operators,omitempty" xml:"operators,omitempty" type:"Repeated"`
}

func (s DescribeOperatorListByTypeResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListByTypeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListByTypeResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeOperatorListByTypeResponseBodyResultObject) GetOperators() []*DescribeOperatorListByTypeResponseBodyResultObjectOperators {
	return s.Operators
}

func (s *DescribeOperatorListByTypeResponseBodyResultObject) SetFieldType(v string) *DescribeOperatorListByTypeResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeOperatorListByTypeResponseBodyResultObject) SetOperators(v []*DescribeOperatorListByTypeResponseBodyResultObjectOperators) *DescribeOperatorListByTypeResponseBodyResultObject {
	s.Operators = v
	return s
}

func (s *DescribeOperatorListByTypeResponseBodyResultObject) Validate() error {
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

type DescribeOperatorListByTypeResponseBodyResultObjectOperators struct {
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
	// Operator name
	//
	// example:
	//
	// 等于
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeOperatorListByTypeResponseBodyResultObjectOperators) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListByTypeResponseBodyResultObjectOperators) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) GetCode() *string {
	return s.Code
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) GetHasRightVariable() *bool {
	return s.HasRightVariable
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) GetName() *string {
	return s.Name
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) SetCode(v string) *DescribeOperatorListByTypeResponseBodyResultObjectOperators {
	s.Code = &v
	return s
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) SetHasRightVariable(v bool) *DescribeOperatorListByTypeResponseBodyResultObjectOperators {
	s.HasRightVariable = &v
	return s
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) SetName(v string) *DescribeOperatorListByTypeResponseBodyResultObjectOperators {
	s.Name = &v
	return s
}

func (s *DescribeOperatorListByTypeResponseBodyResultObjectOperators) Validate() error {
	return dara.Validate(s)
}
