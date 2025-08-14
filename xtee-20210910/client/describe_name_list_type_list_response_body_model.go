// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListTypeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeNameListTypeListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeNameListTypeListResponseBodyResultObject) *DescribeNameListTypeListResponseBody
	GetResultObject() []*DescribeNameListTypeListResponseBodyResultObject
}

type DescribeNameListTypeListResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeNameListTypeListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeNameListTypeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListTypeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListTypeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListTypeListResponseBody) GetResultObject() []*DescribeNameListTypeListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListTypeListResponseBody) SetRequestId(v string) *DescribeNameListTypeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListTypeListResponseBody) SetResultObject(v []*DescribeNameListTypeListResponseBodyResultObject) *DescribeNameListTypeListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListTypeListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNameListTypeListResponseBodyResultObject struct {
	// Match Key.
	//
	// example:
	//
	// accountId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Variable value
	//
	// example:
	//
	// 账号ID
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeNameListTypeListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListTypeListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListTypeListResponseBodyResultObject) GetKey() *string {
	return s.Key
}

func (s *DescribeNameListTypeListResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListTypeListResponseBodyResultObject) SetKey(v string) *DescribeNameListTypeListResponseBodyResultObject {
	s.Key = &v
	return s
}

func (s *DescribeNameListTypeListResponseBodyResultObject) SetValue(v string) *DescribeNameListTypeListResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeNameListTypeListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
