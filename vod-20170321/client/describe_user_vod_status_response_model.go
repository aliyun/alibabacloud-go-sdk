// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVodStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserVodStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserVodStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserVodStatusResponseBody) *DescribeUserVodStatusResponse
	GetBody() *DescribeUserVodStatusResponseBody
}

type DescribeUserVodStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserVodStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserVodStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVodStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserVodStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserVodStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserVodStatusResponse) GetBody() *DescribeUserVodStatusResponseBody {
	return s.Body
}

func (s *DescribeUserVodStatusResponse) SetHeaders(v map[string]*string) *DescribeUserVodStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserVodStatusResponse) SetStatusCode(v int32) *DescribeUserVodStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserVodStatusResponse) SetBody(v *DescribeUserVodStatusResponseBody) *DescribeUserVodStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserVodStatusResponse) Validate() error {
	return dara.Validate(s)
}
