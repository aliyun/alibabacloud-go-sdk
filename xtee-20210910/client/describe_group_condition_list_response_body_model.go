// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupConditionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupConditionListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeGroupConditionListResponseBodyResultObject) *DescribeGroupConditionListResponseBody
	GetResultObject() []*DescribeGroupConditionListResponseBodyResultObject
}

type DescribeGroupConditionListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeGroupConditionListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeGroupConditionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupConditionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupConditionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupConditionListResponseBody) GetResultObject() []*DescribeGroupConditionListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeGroupConditionListResponseBody) SetRequestId(v string) *DescribeGroupConditionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupConditionListResponseBody) SetResultObject(v []*DescribeGroupConditionListResponseBodyResultObject) *DescribeGroupConditionListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeGroupConditionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupConditionListResponseBodyResultObject struct {
	// Field key
	//
	// example:
	//
	// key
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// Field value.
	//
	// example:
	//
	// value
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s DescribeGroupConditionListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupConditionListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeGroupConditionListResponseBodyResultObject) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DescribeGroupConditionListResponseBodyResultObject) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeGroupConditionListResponseBodyResultObject) SetFieldKey(v string) *DescribeGroupConditionListResponseBodyResultObject {
	s.FieldKey = &v
	return s
}

func (s *DescribeGroupConditionListResponseBodyResultObject) SetFieldValue(v string) *DescribeGroupConditionListResponseBodyResultObject {
	s.FieldValue = &v
	return s
}

func (s *DescribeGroupConditionListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
