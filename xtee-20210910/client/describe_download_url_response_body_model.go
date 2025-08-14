// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeDownloadUrlResponseBody
	GetResultObject() *string
}

type DescribeDownloadUrlResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *string `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadUrlResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeDownloadUrlResponseBody) SetRequestId(v string) *DescribeDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadUrlResponseBody) SetResultObject(v string) *DescribeDownloadUrlResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
