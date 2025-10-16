// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionInfoResponseBody) *DescribeRegionInfoResponse
	GetBody() *DescribeRegionInfoResponseBody
}

type DescribeRegionInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionInfoResponse) GetBody() *DescribeRegionInfoResponseBody {
	return s.Body
}

func (s *DescribeRegionInfoResponse) SetHeaders(v map[string]*string) *DescribeRegionInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionInfoResponse) SetStatusCode(v int32) *DescribeRegionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionInfoResponse) SetBody(v *DescribeRegionInfoResponseBody) *DescribeRegionInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
