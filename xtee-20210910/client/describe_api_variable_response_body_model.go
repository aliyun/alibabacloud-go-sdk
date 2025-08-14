// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiVariableResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeApiVariableResponseBody
	GetResultObject() *bool
}

type DescribeApiVariableResponseBody struct {
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

func (s DescribeApiVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiVariableResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeApiVariableResponseBody) SetRequestId(v string) *DescribeApiVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiVariableResponseBody) SetResultObject(v bool) *DescribeApiVariableResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeApiVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
