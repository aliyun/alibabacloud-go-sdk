// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStreamURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodStreamURLRequest
	GetOwnerId() *int64
	SetUrl(v string) *DescribeVodStreamURLRequest
	GetUrl() *string
}

type DescribeVodStreamURLRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxx/xxx.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVodStreamURLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStreamURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodStreamURLRequest) GetUrl() *string {
	return s.Url
}

func (s *DescribeVodStreamURLRequest) SetOwnerId(v int64) *DescribeVodStreamURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodStreamURLRequest) SetUrl(v string) *DescribeVodStreamURLRequest {
	s.Url = &v
	return s
}

func (s *DescribeVodStreamURLRequest) Validate() error {
	return dara.Validate(s)
}
