// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionResourceResponseBody) *DescribeRegionResourceResponse
	GetBody() *DescribeRegionResourceResponseBody
}

type DescribeRegionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionResourceResponse) GetBody() *DescribeRegionResourceResponseBody {
	return s.Body
}

func (s *DescribeRegionResourceResponse) SetHeaders(v map[string]*string) *DescribeRegionResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionResourceResponse) SetStatusCode(v int32) *DescribeRegionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionResourceResponse) SetBody(v *DescribeRegionResourceResponseBody) *DescribeRegionResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
