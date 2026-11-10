// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICallDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAICallDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAICallDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAICallDetailsResponseBody) *DescribeAICallDetailsResponse
	GetBody() *DescribeAICallDetailsResponseBody
}

type DescribeAICallDetailsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAICallDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAICallDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICallDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAICallDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAICallDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAICallDetailsResponse) GetBody() *DescribeAICallDetailsResponseBody {
	return s.Body
}

func (s *DescribeAICallDetailsResponse) SetHeaders(v map[string]*string) *DescribeAICallDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAICallDetailsResponse) SetStatusCode(v int32) *DescribeAICallDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAICallDetailsResponse) SetBody(v *DescribeAICallDetailsResponseBody) *DescribeAICallDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeAICallDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
