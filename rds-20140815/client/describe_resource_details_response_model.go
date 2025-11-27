// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceDetailsResponseBody) *DescribeResourceDetailsResponse
	GetBody() *DescribeResourceDetailsResponseBody
}

type DescribeResourceDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceDetailsResponse) GetBody() *DescribeResourceDetailsResponseBody {
	return s.Body
}

func (s *DescribeResourceDetailsResponse) SetHeaders(v map[string]*string) *DescribeResourceDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceDetailsResponse) SetStatusCode(v int32) *DescribeResourceDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceDetailsResponse) SetBody(v *DescribeResourceDetailsResponseBody) *DescribeResourceDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
