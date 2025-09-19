// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulDetailsResponseBody) *DescribeVulDetailsResponse
	GetBody() *DescribeVulDetailsResponseBody
}

type DescribeVulDetailsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulDetailsResponse) GetBody() *DescribeVulDetailsResponseBody {
	return s.Body
}

func (s *DescribeVulDetailsResponse) SetHeaders(v map[string]*string) *DescribeVulDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDetailsResponse) SetStatusCode(v int32) *DescribeVulDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulDetailsResponse) SetBody(v *DescribeVulDetailsResponseBody) *DescribeVulDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulDetailsResponse) Validate() error {
	return dara.Validate(s)
}
