// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDetailsResponseBody) *DescribeInstanceDetailsResponse
	GetBody() *DescribeInstanceDetailsResponseBody
}

type DescribeInstanceDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDetailsResponse) GetBody() *DescribeInstanceDetailsResponseBody {
	return s.Body
}

func (s *DescribeInstanceDetailsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetStatusCode(v int32) *DescribeInstanceDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetBody(v *DescribeInstanceDetailsResponseBody) *DescribeInstanceDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDetailsResponse) Validate() error {
	return dara.Validate(s)
}
