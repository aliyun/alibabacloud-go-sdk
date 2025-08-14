// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiGroupsResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeApiGroupsResponseBody
	GetResultObject() *bool
}

type DescribeApiGroupsResponseBody struct {
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

func (s DescribeApiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiGroupsResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeApiGroupsResponseBody) SetRequestId(v string) *DescribeApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetResultObject(v bool) *DescribeApiGroupsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
