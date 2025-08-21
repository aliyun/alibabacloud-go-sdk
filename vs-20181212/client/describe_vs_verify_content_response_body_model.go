// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeVsVerifyContentResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeVsVerifyContentResponseBody
	GetRequestId() *string
}

type DescribeVsVerifyContentResponseBody struct {
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c3*****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsVerifyContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeVsVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsVerifyContentResponseBody) SetContent(v string) *DescribeVsVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeVsVerifyContentResponseBody) SetRequestId(v string) *DescribeVsVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsVerifyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
