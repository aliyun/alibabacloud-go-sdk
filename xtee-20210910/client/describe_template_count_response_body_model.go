// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTemplateCountResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeTemplateCountResponseBodyResultObject) *DescribeTemplateCountResponseBody
	GetResultObject() *DescribeTemplateCountResponseBodyResultObject
}

type DescribeTemplateCountResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeTemplateCountResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeTemplateCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateCountResponseBody) GetResultObject() *DescribeTemplateCountResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTemplateCountResponseBody) SetRequestId(v string) *DescribeTemplateCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateCountResponseBody) SetResultObject(v *DescribeTemplateCountResponseBodyResultObject) *DescribeTemplateCountResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTemplateCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplateCountResponseBodyResultObject struct {
	// Template quantity limit.
	//
	// example:
	//
	// false
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum count
	//
	// example:
	//
	// 1000
	MaxTotalItem *int32 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total count.
	//
	// example:
	//
	// 13
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeTemplateCountResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCountResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCountResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeTemplateCountResponseBodyResultObject) GetMaxTotalItem() *int32 {
	return s.MaxTotalItem
}

func (s *DescribeTemplateCountResponseBodyResultObject) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeTemplateCountResponseBodyResultObject) SetLimit(v bool) *DescribeTemplateCountResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeTemplateCountResponseBodyResultObject) SetMaxTotalItem(v int32) *DescribeTemplateCountResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeTemplateCountResponseBodyResultObject) SetTotalItem(v int32) *DescribeTemplateCountResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeTemplateCountResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
