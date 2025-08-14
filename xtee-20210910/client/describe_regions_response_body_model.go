// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeRegionsResponseBody
	GetResultObject() *bool
}

type DescribeRegionsResponseBody struct {
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

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResultObject(v bool) *DescribeRegionsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
