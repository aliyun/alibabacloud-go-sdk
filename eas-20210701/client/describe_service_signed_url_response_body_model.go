// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceSignedUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceSignedUrlResponseBody
	GetRequestId() *string
	SetSignedUrl(v string) *DescribeServiceSignedUrlResponseBody
	GetSignedUrl() *string
}

type DescribeServiceSignedUrlResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service URL.
	//
	// example:
	//
	// https://foo-115**.console.cn-hangzhou.eas.pai-ml.com?expire=1735202661&signature=ey*******
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s DescribeServiceSignedUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceSignedUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceSignedUrlResponseBody) GetSignedUrl() *string {
	return s.SignedUrl
}

func (s *DescribeServiceSignedUrlResponseBody) SetRequestId(v string) *DescribeServiceSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceSignedUrlResponseBody) SetSignedUrl(v string) *DescribeServiceSignedUrlResponseBody {
	s.SignedUrl = &v
	return s
}

func (s *DescribeServiceSignedUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
