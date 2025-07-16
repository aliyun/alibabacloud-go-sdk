// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeVerifyContentResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeVerifyContentResponseBody
	GetRequestId() *string
}

type DescribeVerifyContentResponseBody struct {
	// The verification content.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c32cd91f
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyContentResponseBody) SetContent(v string) *DescribeVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeVerifyContentResponseBody) SetRequestId(v string) *DescribeVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
