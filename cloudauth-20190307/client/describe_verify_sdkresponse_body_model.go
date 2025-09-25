// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySDKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifySDKResponseBody
	GetRequestId() *string
	SetSdkUrl(v string) *DescribeVerifySDKResponseBody
	GetSdkUrl() *string
}

type DescribeVerifySDKResponseBody struct {
	// The ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SDK download URL. When not empty, it indicates that the generation is complete.
	//
	// example:
	//
	// https://www.xxx.com
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeVerifySDKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySDKResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifySDKResponseBody) GetSdkUrl() *string {
	return s.SdkUrl
}

func (s *DescribeVerifySDKResponseBody) SetRequestId(v string) *DescribeVerifySDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySDKResponseBody) SetSdkUrl(v string) *DescribeVerifySDKResponseBody {
	s.SdkUrl = &v
	return s
}

func (s *DescribeVerifySDKResponseBody) Validate() error {
	return dara.Validate(s)
}
