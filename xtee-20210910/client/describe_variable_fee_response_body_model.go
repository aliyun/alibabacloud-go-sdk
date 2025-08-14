// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableFeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableFeeResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeVariableFeeResponseBody
	GetResultObject() *bool
}

type DescribeVariableFeeResponseBody struct {
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

func (s DescribeVariableFeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableFeeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableFeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableFeeResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeVariableFeeResponseBody) SetRequestId(v string) *DescribeVariableFeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableFeeResponseBody) SetResultObject(v bool) *DescribeVariableFeeResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeVariableFeeResponseBody) Validate() error {
	return dara.Validate(s)
}
