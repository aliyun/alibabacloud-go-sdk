// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckDetailsResponseBody) *DescribeCheckDetailsResponse
	GetBody() *DescribeCheckDetailsResponseBody
}

type DescribeCheckDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckDetailsResponse) GetBody() *DescribeCheckDetailsResponseBody {
	return s.Body
}

func (s *DescribeCheckDetailsResponse) SetHeaders(v map[string]*string) *DescribeCheckDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckDetailsResponse) SetStatusCode(v int32) *DescribeCheckDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckDetailsResponse) SetBody(v *DescribeCheckDetailsResponseBody) *DescribeCheckDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
