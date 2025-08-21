// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStreamURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodStreamURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodStreamURLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodStreamURLResponseBody) *DescribeVodStreamURLResponse
	GetBody() *DescribeVodStreamURLResponseBody
}

type DescribeVodStreamURLResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodStreamURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodStreamURLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStreamURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodStreamURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodStreamURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodStreamURLResponse) GetBody() *DescribeVodStreamURLResponseBody {
	return s.Body
}

func (s *DescribeVodStreamURLResponse) SetHeaders(v map[string]*string) *DescribeVodStreamURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodStreamURLResponse) SetStatusCode(v int32) *DescribeVodStreamURLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodStreamURLResponse) SetBody(v *DescribeVodStreamURLResponseBody) *DescribeVodStreamURLResponse {
	s.Body = v
	return s
}

func (s *DescribeVodStreamURLResponse) Validate() error {
	return dara.Validate(s)
}
