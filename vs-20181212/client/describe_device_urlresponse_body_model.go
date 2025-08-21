// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int64) *DescribeDeviceURLResponseBody
	GetExpireTime() *int64
	SetRequestId(v string) *DescribeDeviceURLResponseBody
	GetRequestId() *string
	SetUrl(v string) *DescribeDeviceURLResponseBody
	GetUrl() *string
}

type DescribeDeviceURLResponseBody struct {
	// example:
	//
	// 1639130258
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rtmp://demo.aliyundoc.com/live/live001?auth_key=1639130258-0-0-b2b04fe85ece6*****a6b1a42bc7e
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDeviceURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeDeviceURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceURLResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeDeviceURLResponseBody) SetExpireTime(v int64) *DescribeDeviceURLResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDeviceURLResponseBody) SetRequestId(v string) *DescribeDeviceURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceURLResponseBody) SetUrl(v string) *DescribeDeviceURLResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeDeviceURLResponseBody) Validate() error {
	return dara.Validate(s)
}
