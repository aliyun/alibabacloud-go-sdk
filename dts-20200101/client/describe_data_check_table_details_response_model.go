// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataCheckTableDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataCheckTableDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataCheckTableDetailsResponseBody) *DescribeDataCheckTableDetailsResponse
	GetBody() *DescribeDataCheckTableDetailsResponseBody
}

type DescribeDataCheckTableDetailsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCheckTableDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataCheckTableDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataCheckTableDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataCheckTableDetailsResponse) GetBody() *DescribeDataCheckTableDetailsResponseBody {
	return s.Body
}

func (s *DescribeDataCheckTableDetailsResponse) SetHeaders(v map[string]*string) *DescribeDataCheckTableDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCheckTableDetailsResponse) SetStatusCode(v int32) *DescribeDataCheckTableDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCheckTableDetailsResponse) SetBody(v *DescribeDataCheckTableDetailsResponseBody) *DescribeDataCheckTableDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataCheckTableDetailsResponse) Validate() error {
	return dara.Validate(s)
}
