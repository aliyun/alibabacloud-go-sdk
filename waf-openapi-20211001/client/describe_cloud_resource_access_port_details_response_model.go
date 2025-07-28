// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessPortDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudResourceAccessPortDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudResourceAccessPortDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudResourceAccessPortDetailsResponseBody) *DescribeCloudResourceAccessPortDetailsResponse
	GetBody() *DescribeCloudResourceAccessPortDetailsResponseBody
}

type DescribeCloudResourceAccessPortDetailsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourceAccessPortDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourceAccessPortDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) GetBody() *DescribeCloudResourceAccessPortDetailsResponseBody {
	return s.Body
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) SetHeaders(v map[string]*string) *DescribeCloudResourceAccessPortDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) SetStatusCode(v int32) *DescribeCloudResourceAccessPortDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) SetBody(v *DescribeCloudResourceAccessPortDetailsResponseBody) *DescribeCloudResourceAccessPortDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponse) Validate() error {
	return dara.Validate(s)
}
