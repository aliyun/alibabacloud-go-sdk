// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStreamURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutProtocol(v string) *DescribeVodStreamURLResponseBody
	GetOutProtocol() *string
	SetPort(v int64) *DescribeVodStreamURLResponseBody
	GetPort() *int64
	SetRequestId(v string) *DescribeVodStreamURLResponseBody
	GetRequestId() *string
	SetUrl(v string) *DescribeVodStreamURLResponseBody
	GetUrl() *string
}

type DescribeVodStreamURLResponseBody struct {
	// example:
	//
	// rtsp
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	// example:
	//
	// 8080
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rtsp://domain/live/stream?sign=xxxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVodStreamURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStreamURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLResponseBody) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *DescribeVodStreamURLResponseBody) GetPort() *int64 {
	return s.Port
}

func (s *DescribeVodStreamURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodStreamURLResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeVodStreamURLResponseBody) SetOutProtocol(v string) *DescribeVodStreamURLResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetPort(v int64) *DescribeVodStreamURLResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetRequestId(v string) *DescribeVodStreamURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) SetUrl(v string) *DescribeVodStreamURLResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeVodStreamURLResponseBody) Validate() error {
	return dara.Validate(s)
}
