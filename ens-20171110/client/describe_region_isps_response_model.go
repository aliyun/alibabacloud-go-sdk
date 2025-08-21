// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionIspsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionIspsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionIspsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionIspsResponseBody) *DescribeRegionIspsResponse
	GetBody() *DescribeRegionIspsResponseBody
}

type DescribeRegionIspsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionIspsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionIspsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionIspsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionIspsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionIspsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionIspsResponse) GetBody() *DescribeRegionIspsResponseBody {
	return s.Body
}

func (s *DescribeRegionIspsResponse) SetHeaders(v map[string]*string) *DescribeRegionIspsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionIspsResponse) SetStatusCode(v int32) *DescribeRegionIspsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionIspsResponse) SetBody(v *DescribeRegionIspsResponseBody) *DescribeRegionIspsResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionIspsResponse) Validate() error {
	return dara.Validate(s)
}
