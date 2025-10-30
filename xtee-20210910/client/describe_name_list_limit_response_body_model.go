// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeNameListLimitResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeNameListLimitResponseBodyResultObject) *DescribeNameListLimitResponseBody
	GetResultObject() *DescribeNameListLimitResponseBodyResultObject
}

type DescribeNameListLimitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeNameListLimitResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeNameListLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListLimitResponseBody) GetResultObject() *DescribeNameListLimitResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListLimitResponseBody) SetRequestId(v string) *DescribeNameListLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListLimitResponseBody) SetResultObject(v *DescribeNameListLimitResponseBodyResultObject) *DescribeNameListLimitResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListLimitResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNameListLimitResponseBodyResultObject struct {
	// Whether it exceeds the maximum quantity
	//
	// example:
	//
	// true
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum number of creatable items
	//
	// example:
	//
	// 100
	MaxTotalItem *int64 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 101
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeNameListLimitResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListLimitResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListLimitResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeNameListLimitResponseBodyResultObject) GetMaxTotalItem() *int64 {
	return s.MaxTotalItem
}

func (s *DescribeNameListLimitResponseBodyResultObject) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeNameListLimitResponseBodyResultObject) SetLimit(v bool) *DescribeNameListLimitResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeNameListLimitResponseBodyResultObject) SetMaxTotalItem(v int64) *DescribeNameListLimitResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeNameListLimitResponseBodyResultObject) SetTotalItem(v int64) *DescribeNameListLimitResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeNameListLimitResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
