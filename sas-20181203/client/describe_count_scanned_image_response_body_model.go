// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCountScannedImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCountScannedImageResponseBody
	GetRequestId() *string
	SetScannedCount(v int32) *DescribeCountScannedImageResponseBody
	GetScannedCount() *int32
}

type DescribeCountScannedImageResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3F4236AB-7070-538D-85EB-98EBFE6C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of images that are scanned.
	//
	// example:
	//
	// 11
	ScannedCount *int32 `json:"ScannedCount,omitempty" xml:"ScannedCount,omitempty"`
}

func (s DescribeCountScannedImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCountScannedImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCountScannedImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCountScannedImageResponseBody) GetScannedCount() *int32 {
	return s.ScannedCount
}

func (s *DescribeCountScannedImageResponseBody) SetRequestId(v string) *DescribeCountScannedImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCountScannedImageResponseBody) SetScannedCount(v int32) *DescribeCountScannedImageResponseBody {
	s.ScannedCount = &v
	return s
}

func (s *DescribeCountScannedImageResponseBody) Validate() error {
	return dara.Validate(s)
}
