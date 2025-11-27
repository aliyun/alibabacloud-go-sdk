// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionInfosResponseBody) *DescribeRegionInfosResponse
	GetBody() *DescribeRegionInfosResponseBody
}

type DescribeRegionInfosResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionInfosResponse) GetBody() *DescribeRegionInfosResponseBody {
	return s.Body
}

func (s *DescribeRegionInfosResponse) SetHeaders(v map[string]*string) *DescribeRegionInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionInfosResponse) SetStatusCode(v int32) *DescribeRegionInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionInfosResponse) SetBody(v *DescribeRegionInfosResponseBody) *DescribeRegionInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
