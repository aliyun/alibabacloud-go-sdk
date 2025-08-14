// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagsListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeTagsListResponseBody
	GetResultObject() *bool
}

type DescribeTagsListResponseBody struct {
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

func (s DescribeTagsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeTagsListResponseBody) SetRequestId(v string) *DescribeTagsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsListResponseBody) SetResultObject(v bool) *DescribeTagsListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeTagsListResponseBody) Validate() error {
	return dara.Validate(s)
}
