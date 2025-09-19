// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckFixDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckFixDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckFixDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckFixDetailsResponseBody) *DescribeCheckFixDetailsResponse
	GetBody() *DescribeCheckFixDetailsResponseBody
}

type DescribeCheckFixDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckFixDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckFixDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckFixDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckFixDetailsResponse) GetBody() *DescribeCheckFixDetailsResponseBody {
	return s.Body
}

func (s *DescribeCheckFixDetailsResponse) SetHeaders(v map[string]*string) *DescribeCheckFixDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckFixDetailsResponse) SetStatusCode(v int32) *DescribeCheckFixDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckFixDetailsResponse) SetBody(v *DescribeCheckFixDetailsResponseBody) *DescribeCheckFixDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckFixDetailsResponse) Validate() error {
	return dara.Validate(s)
}
