// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeVodVerifyContentResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeVodVerifyContentResponseBody
	GetRequestId() *string
}

type DescribeVodVerifyContentResponseBody struct {
	// The verification content.
	//
	// example:
	//
	// verify_dffeb661*****3a59c31cd91f
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-4688-634BDBE6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeVodVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodVerifyContentResponseBody) SetContent(v string) *DescribeVodVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeVodVerifyContentResponseBody) SetRequestId(v string) *DescribeVodVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodVerifyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
