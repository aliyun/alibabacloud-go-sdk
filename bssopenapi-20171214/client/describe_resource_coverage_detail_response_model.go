// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceCoverageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceCoverageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceCoverageDetailResponseBody) *DescribeResourceCoverageDetailResponse
	GetBody() *DescribeResourceCoverageDetailResponseBody
}

type DescribeResourceCoverageDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceCoverageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceCoverageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceCoverageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceCoverageDetailResponse) GetBody() *DescribeResourceCoverageDetailResponseBody {
	return s.Body
}

func (s *DescribeResourceCoverageDetailResponse) SetHeaders(v map[string]*string) *DescribeResourceCoverageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceCoverageDetailResponse) SetStatusCode(v int32) *DescribeResourceCoverageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponse) SetBody(v *DescribeResourceCoverageDetailResponseBody) *DescribeResourceCoverageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceCoverageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
