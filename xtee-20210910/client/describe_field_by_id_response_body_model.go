// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFieldByIdResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeFieldByIdResponseBody
	GetResultObject() *bool
}

type DescribeFieldByIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeFieldByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFieldByIdResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeFieldByIdResponseBody) SetRequestId(v string) *DescribeFieldByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldByIdResponseBody) SetResultObject(v bool) *DescribeFieldByIdResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeFieldByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
