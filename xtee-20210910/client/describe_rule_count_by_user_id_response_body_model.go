// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleCountByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleCountByUserIdResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRuleCountByUserIdResponseBodyResultObject) *DescribeRuleCountByUserIdResponseBody
	GetResultObject() *DescribeRuleCountByUserIdResponseBodyResultObject
}

type DescribeRuleCountByUserIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeRuleCountByUserIdResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRuleCountByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleCountByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleCountByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleCountByUserIdResponseBody) GetResultObject() *DescribeRuleCountByUserIdResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleCountByUserIdResponseBody) SetRequestId(v string) *DescribeRuleCountByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleCountByUserIdResponseBody) SetResultObject(v *DescribeRuleCountByUserIdResponseBodyResultObject) *DescribeRuleCountByUserIdResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleCountByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleCountByUserIdResponseBodyResultObject struct {
	// Whether the limit condition is reached. Values: -**true**: Yes-**false**: No
	//
	// example:
	//
	// false
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum number of items
	//
	// example:
	//
	// 100
	MaxTotalItem *int32 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 27
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeRuleCountByUserIdResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleCountByUserIdResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) GetMaxTotalItem() *int32 {
	return s.MaxTotalItem
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) SetLimit(v bool) *DescribeRuleCountByUserIdResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) SetMaxTotalItem(v int32) *DescribeRuleCountByUserIdResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) SetTotalItem(v int32) *DescribeRuleCountByUserIdResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeRuleCountByUserIdResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
