// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedVulnTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnprotectedVulnTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnprotectedVulnTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnprotectedVulnTrendResponseBody) *DescribeUnprotectedVulnTrendResponse
	GetBody() *DescribeUnprotectedVulnTrendResponseBody
}

type DescribeUnprotectedVulnTrendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnprotectedVulnTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnprotectedVulnTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedVulnTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedVulnTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnprotectedVulnTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnprotectedVulnTrendResponse) GetBody() *DescribeUnprotectedVulnTrendResponseBody {
	return s.Body
}

func (s *DescribeUnprotectedVulnTrendResponse) SetHeaders(v map[string]*string) *DescribeUnprotectedVulnTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponse) SetStatusCode(v int32) *DescribeUnprotectedVulnTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponse) SetBody(v *DescribeUnprotectedVulnTrendResponseBody) *DescribeUnprotectedVulnTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeUnprotectedVulnTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
