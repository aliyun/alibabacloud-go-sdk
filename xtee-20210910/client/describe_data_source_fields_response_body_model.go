// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDataSourceFieldsResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeDataSourceFieldsResponseBody
	GetResultObject() *bool
}

type DescribeDataSourceFieldsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeDataSourceFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourceFieldsResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeDataSourceFieldsResponseBody) SetRequestId(v string) *DescribeDataSourceFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourceFieldsResponseBody) SetResultObject(v bool) *DescribeDataSourceFieldsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeDataSourceFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}
