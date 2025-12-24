// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistinctReleasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDistinctReleasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDistinctReleasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDistinctReleasesResponseBody) *DescribeDistinctReleasesResponse
	GetBody() *DescribeDistinctReleasesResponseBody
}

type DescribeDistinctReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistinctReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDistinctReleasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistinctReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDistinctReleasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDistinctReleasesResponse) GetBody() *DescribeDistinctReleasesResponseBody {
	return s.Body
}

func (s *DescribeDistinctReleasesResponse) SetHeaders(v map[string]*string) *DescribeDistinctReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetStatusCode(v int32) *DescribeDistinctReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetBody(v *DescribeDistinctReleasesResponseBody) *DescribeDistinctReleasesResponse {
	s.Body = v
	return s
}

func (s *DescribeDistinctReleasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
