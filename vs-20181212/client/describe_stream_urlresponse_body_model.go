// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int64) *DescribeStreamURLResponseBody
	GetExpireTime() *int64
	SetRequestId(v string) *DescribeStreamURLResponseBody
	GetRequestId() *string
	SetUrl(v string) *DescribeStreamURLResponseBody
	GetUrl() *string
}

type DescribeStreamURLResponseBody struct {
	// example:
	//
	// 1557977029
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rtmp://demo.aliyundoc.com/live/310101*****7542007?auth_key=1639130258-0-0-b2b04fe85ece6*****a6b1a42bc7e
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeStreamURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeStreamURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamURLResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeStreamURLResponseBody) SetExpireTime(v int64) *DescribeStreamURLResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeStreamURLResponseBody) SetRequestId(v string) *DescribeStreamURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamURLResponseBody) SetUrl(v string) *DescribeStreamURLResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeStreamURLResponseBody) Validate() error {
	return dara.Validate(s)
}
