// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExistNameResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeExistNameResponseBody
	GetResultObject() *bool
}

type DescribeExistNameResponseBody struct {
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

func (s DescribeExistNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExistNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExistNameResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeExistNameResponseBody) SetRequestId(v string) *DescribeExistNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExistNameResponseBody) SetResultObject(v bool) *DescribeExistNameResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeExistNameResponseBody) Validate() error {
	return dara.Validate(s)
}
