// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCountNotScannedImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotScannedCnt(v int32) *DescribeCountNotScannedImageResponseBody
	GetNotScannedCnt() *int32
	SetRequestId(v string) *DescribeCountNotScannedImageResponseBody
	GetRequestId() *string
}

type DescribeCountNotScannedImageResponseBody struct {
	// The number of images that are not scanned.
	//
	// example:
	//
	// 28
	NotScannedCnt *int32 `json:"NotScannedCnt,omitempty" xml:"NotScannedCnt,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCountNotScannedImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCountNotScannedImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCountNotScannedImageResponseBody) GetNotScannedCnt() *int32 {
	return s.NotScannedCnt
}

func (s *DescribeCountNotScannedImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCountNotScannedImageResponseBody) SetNotScannedCnt(v int32) *DescribeCountNotScannedImageResponseBody {
	s.NotScannedCnt = &v
	return s
}

func (s *DescribeCountNotScannedImageResponseBody) SetRequestId(v string) *DescribeCountNotScannedImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCountNotScannedImageResponseBody) Validate() error {
	return dara.Validate(s)
}
