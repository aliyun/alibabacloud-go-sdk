// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVolDingdingMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDingdingUrl(v string) *DescribeVolDingdingMessageResponseBody
	GetDingdingUrl() *string
	SetRequestId(v string) *DescribeVolDingdingMessageResponseBody
	GetRequestId() *string
}

type DescribeVolDingdingMessageResponseBody struct {
	// The QR code address of the DingTalk group.
	//
	// example:
	//
	// https://www.wikihow.com/images_en/thumb/4/48/Get-the-URL-for-Pictures-Step-4-Version-4.jpg/v4-728px-Get-the-URL-for-Pictures-Step-4-Version-4.jpg.webp
	DingdingUrl *string `json:"DingdingUrl,omitempty" xml:"DingdingUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A437E93-47EE-548F-ABCE-13F89AA85585
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVolDingdingMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVolDingdingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageResponseBody) GetDingdingUrl() *string {
	return s.DingdingUrl
}

func (s *DescribeVolDingdingMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVolDingdingMessageResponseBody) SetDingdingUrl(v string) *DescribeVolDingdingMessageResponseBody {
	s.DingdingUrl = &v
	return s
}

func (s *DescribeVolDingdingMessageResponseBody) SetRequestId(v string) *DescribeVolDingdingMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVolDingdingMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
