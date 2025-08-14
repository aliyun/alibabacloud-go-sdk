// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApisResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeApisResponseBody
	GetResultObject() *bool
}

type DescribeApisResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeApisResponseBody) SetRequestId(v string) *DescribeApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisResponseBody) SetResultObject(v bool) *DescribeApisResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeApisResponseBody) Validate() error {
	return dara.Validate(s)
}
