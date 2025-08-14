// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventLogDetailResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeEventLogDetailResponseBody
	GetResultObject() *bool
}

type DescribeEventLogDetailResponseBody struct {
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

func (s DescribeEventLogDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventLogDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventLogDetailResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeEventLogDetailResponseBody) SetRequestId(v string) *DescribeEventLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventLogDetailResponseBody) SetResultObject(v bool) *DescribeEventLogDetailResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeEventLogDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
