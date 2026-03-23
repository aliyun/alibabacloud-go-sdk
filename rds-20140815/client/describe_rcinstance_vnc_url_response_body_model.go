// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceVncUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginUrl(v string) *DescribeRCInstanceVncUrlResponseBody
	GetLoginUrl() *string
	SetRequestId(v string) *DescribeRCInstanceVncUrlResponseBody
	GetRequestId() *string
}

type DescribeRCInstanceVncUrlResponseBody struct {
	LoginUrl  *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCInstanceVncUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceVncUrlResponseBody) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DescribeRCInstanceVncUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceVncUrlResponseBody) SetLoginUrl(v string) *DescribeRCInstanceVncUrlResponseBody {
	s.LoginUrl = &v
	return s
}

func (s *DescribeRCInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeRCInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceVncUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
