// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssAuthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOssAuthStatusResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeOssAuthStatusResponseBody
	GetResultObject() *string
}

type DescribeOssAuthStatusResponseBody struct {
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
	ResultObject *string `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeOssAuthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssAuthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssAuthStatusResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeOssAuthStatusResponseBody) SetRequestId(v string) *DescribeOssAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssAuthStatusResponseBody) SetResultObject(v string) *DescribeOssAuthStatusResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeOssAuthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
