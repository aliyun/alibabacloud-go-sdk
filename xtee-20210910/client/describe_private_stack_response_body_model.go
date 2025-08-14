// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePrivateStackResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribePrivateStackResponseBody
	GetResultObject() *bool
}

type DescribePrivateStackResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return Object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribePrivateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateStackResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateStackResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribePrivateStackResponseBody) SetRequestId(v string) *DescribePrivateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateStackResponseBody) SetResultObject(v bool) *DescribePrivateStackResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribePrivateStackResponseBody) Validate() error {
	return dara.Validate(s)
}
