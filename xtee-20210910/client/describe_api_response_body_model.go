// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeApiResponseBody
	GetResultObject() *bool
}

type DescribeApiResponseBody struct {
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

func (s DescribeApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeApiResponseBody) SetRequestId(v string) *DescribeApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiResponseBody) SetResultObject(v bool) *DescribeApiResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeApiResponseBody) Validate() error {
	return dara.Validate(s)
}
