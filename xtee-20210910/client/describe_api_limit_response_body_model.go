// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiLimitResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeApiLimitResponseBodyResultObject) *DescribeApiLimitResponseBody
	GetResultObject() *DescribeApiLimitResponseBodyResultObject
}

type DescribeApiLimitResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeApiLimitResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeApiLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiLimitResponseBody) GetResultObject() *DescribeApiLimitResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeApiLimitResponseBody) SetRequestId(v string) *DescribeApiLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiLimitResponseBody) SetResultObject(v *DescribeApiLimitResponseBodyResultObject) *DescribeApiLimitResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeApiLimitResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiLimitResponseBodyResultObject struct {
	// Whether the maximum number has been exceeded
	//
	// example:
	//
	// true
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum number of items that can be created
	//
	// example:
	//
	// 150
	MaxTotalItem *int64 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 31
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeApiLimitResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLimitResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeApiLimitResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeApiLimitResponseBodyResultObject) GetMaxTotalItem() *int64 {
	return s.MaxTotalItem
}

func (s *DescribeApiLimitResponseBodyResultObject) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeApiLimitResponseBodyResultObject) SetLimit(v bool) *DescribeApiLimitResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeApiLimitResponseBodyResultObject) SetMaxTotalItem(v int64) *DescribeApiLimitResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeApiLimitResponseBodyResultObject) SetTotalItem(v int64) *DescribeApiLimitResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeApiLimitResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
