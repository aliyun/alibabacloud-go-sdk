// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapUrl(v *DescribeCapResponseBodyCapUrl) *DescribeCapResponseBody
	GetCapUrl() *DescribeCapResponseBodyCapUrl
	SetRequestId(v string) *DescribeCapResponseBody
	GetRequestId() *string
}

type DescribeCapResponseBody struct {
	// The download link to the traffic data that is captured when a DDoS attack event occurs.
	CapUrl *DescribeCapResponseBodyCapUrl `json:"CapUrl,omitempty" xml:"CapUrl,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C728D7E9-9A39-52E0-966B-5C33118BDBB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBody) GetCapUrl() *DescribeCapResponseBodyCapUrl {
	return s.CapUrl
}

func (s *DescribeCapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCapResponseBody) SetCapUrl(v *DescribeCapResponseBodyCapUrl) *DescribeCapResponseBody {
	s.CapUrl = v
	return s
}

func (s *DescribeCapResponseBody) SetRequestId(v string) *DescribeCapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCapResponseBodyCapUrl struct {
	// The download link to the traffic data.
	//
	// example:
	//
	// http://beaver-pack****.oss-cn-hangzhou.aliyuncs.com/ddos-2021****-121.89.XX.XX.cap?Expires=1637****&OSSAccessKeyId=LTAI****************&Signature=******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeCapResponseBodyCapUrl) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapResponseBodyCapUrl) GoString() string {
	return s.String()
}

func (s *DescribeCapResponseBodyCapUrl) GetUrl() *string {
	return s.Url
}

func (s *DescribeCapResponseBodyCapUrl) SetUrl(v string) *DescribeCapResponseBodyCapUrl {
	s.Url = &v
	return s
}

func (s *DescribeCapResponseBodyCapUrl) Validate() error {
	return dara.Validate(s)
}
