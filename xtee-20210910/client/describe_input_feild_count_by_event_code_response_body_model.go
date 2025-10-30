// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInputFeildCountByEventCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInputFeildCountByEventCodeResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeInputFeildCountByEventCodeResponseBodyResultObject) *DescribeInputFeildCountByEventCodeResponseBody
	GetResultObject() *DescribeInputFeildCountByEventCodeResponseBodyResultObject
}

type DescribeInputFeildCountByEventCodeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeInputFeildCountByEventCodeResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeInputFeildCountByEventCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInputFeildCountByEventCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInputFeildCountByEventCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInputFeildCountByEventCodeResponseBody) GetResultObject() *DescribeInputFeildCountByEventCodeResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeInputFeildCountByEventCodeResponseBody) SetRequestId(v string) *DescribeInputFeildCountByEventCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponseBody) SetResultObject(v *DescribeInputFeildCountByEventCodeResponseBodyResultObject) *DescribeInputFeildCountByEventCodeResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInputFeildCountByEventCodeResponseBodyResultObject struct {
	// Whether it exceeds the maximum quantity
	//
	// example:
	//
	// true
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum number of created items
	//
	// example:
	//
	// 100
	MaxTotalItem *int32 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 8
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeInputFeildCountByEventCodeResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeInputFeildCountByEventCodeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) GetMaxTotalItem() *int32 {
	return s.MaxTotalItem
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) SetLimit(v bool) *DescribeInputFeildCountByEventCodeResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) SetMaxTotalItem(v int32) *DescribeInputFeildCountByEventCodeResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) SetTotalItem(v int32) *DescribeInputFeildCountByEventCodeResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
