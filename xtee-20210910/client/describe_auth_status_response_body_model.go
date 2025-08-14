// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAuthStatusResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeAuthStatusResponseBody
	GetResultObject() *bool
}

type DescribeAuthStatusResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeAuthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthStatusResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAuthStatusResponseBody) SetRequestId(v string) *DescribeAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthStatusResponseBody) SetResultObject(v bool) *DescribeAuthStatusResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAuthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
