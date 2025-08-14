// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsVariableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventsVariableListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeEventsVariableListResponseBody
	GetResultObject() *bool
}

type DescribeEventsVariableListResponseBody struct {
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

func (s DescribeEventsVariableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsVariableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsVariableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsVariableListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeEventsVariableListResponseBody) SetRequestId(v string) *DescribeEventsVariableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsVariableListResponseBody) SetResultObject(v bool) *DescribeEventsVariableListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeEventsVariableListResponseBody) Validate() error {
	return dara.Validate(s)
}
