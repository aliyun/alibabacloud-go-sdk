// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClassDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClassDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClassDetailsResponseBody) *DescribeClassDetailsResponse
	GetBody() *DescribeClassDetailsResponseBody
}

type DescribeClassDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClassDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClassDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClassDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClassDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClassDetailsResponse) GetBody() *DescribeClassDetailsResponseBody {
	return s.Body
}

func (s *DescribeClassDetailsResponse) SetHeaders(v map[string]*string) *DescribeClassDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClassDetailsResponse) SetStatusCode(v int32) *DescribeClassDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClassDetailsResponse) SetBody(v *DescribeClassDetailsResponseBody) *DescribeClassDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeClassDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
