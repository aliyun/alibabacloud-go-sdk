// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeFileDownloadUrlResponseBody
	GetCode() *int64
	SetRequestId(v string) *DescribeFileDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeFileDownloadUrlResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *DescribeFileDownloadUrlResponseBody
	GetSuccess() *bool
}

type DescribeFileDownloadUrlResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// xxx
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFileDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileDownloadUrlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeFileDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileDownloadUrlResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeFileDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFileDownloadUrlResponseBody) SetCode(v int64) *DescribeFileDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileDownloadUrlResponseBody) SetRequestId(v string) *DescribeFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileDownloadUrlResponseBody) SetResultObject(v string) *DescribeFileDownloadUrlResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeFileDownloadUrlResponseBody) SetSuccess(v bool) *DescribeFileDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFileDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
