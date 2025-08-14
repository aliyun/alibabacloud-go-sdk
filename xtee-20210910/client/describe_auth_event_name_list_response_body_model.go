// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthEventNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAuthEventNameListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeAuthEventNameListResponseBody
	GetResultObject() *bool
}

type DescribeAuthEventNameListResponseBody struct {
	// Request ID.
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

func (s DescribeAuthEventNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthEventNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthEventNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthEventNameListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAuthEventNameListResponseBody) SetRequestId(v string) *DescribeAuthEventNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthEventNameListResponseBody) SetResultObject(v bool) *DescribeAuthEventNameListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAuthEventNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
