// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorTaskLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeMonitorTaskLimitResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeMonitorTaskLimitResponseBodyResultObject) *DescribeMonitorTaskLimitResponseBody
	GetResultObject() *DescribeMonitorTaskLimitResponseBodyResultObject
}

type DescribeMonitorTaskLimitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeMonitorTaskLimitResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeMonitorTaskLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorTaskLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorTaskLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorTaskLimitResponseBody) GetResultObject() *DescribeMonitorTaskLimitResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeMonitorTaskLimitResponseBody) SetRequestId(v string) *DescribeMonitorTaskLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorTaskLimitResponseBody) SetResultObject(v *DescribeMonitorTaskLimitResponseBodyResultObject) *DescribeMonitorTaskLimitResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeMonitorTaskLimitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorTaskLimitResponseBodyResultObject struct {
	// Whether the maximum limit has been reached
	//
	// example:
	//
	// false
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum number of items
	//
	// example:
	//
	// 1000
	MaxTotalItem *int32 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 5
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeMonitorTaskLimitResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorTaskLimitResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) GetMaxTotalItem() *int32 {
	return s.MaxTotalItem
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) SetLimit(v bool) *DescribeMonitorTaskLimitResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) SetMaxTotalItem(v int32) *DescribeMonitorTaskLimitResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) SetTotalItem(v int32) *DescribeMonitorTaskLimitResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeMonitorTaskLimitResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
