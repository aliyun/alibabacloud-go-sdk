// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailByRequestIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventDetailByRequestIdResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeEventDetailByRequestIdResponseBody
	GetResultObject() *bool
}

type DescribeEventDetailByRequestIdResponseBody struct {
	// Request ID
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

func (s DescribeEventDetailByRequestIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailByRequestIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailByRequestIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventDetailByRequestIdResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeEventDetailByRequestIdResponseBody) SetRequestId(v string) *DescribeEventDetailByRequestIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDetailByRequestIdResponseBody) SetResultObject(v bool) *DescribeEventDetailByRequestIdResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeEventDetailByRequestIdResponseBody) Validate() error {
	return dara.Validate(s)
}
