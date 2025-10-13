// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationResourceStaticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebApplicationResourceStaticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebApplicationResourceStaticsResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationResourceStaticsBody) *DescribeWebApplicationResourceStaticsResponse
	GetBody() *WebApplicationResourceStaticsBody
}

type DescribeWebApplicationResourceStaticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationResourceStaticsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebApplicationResourceStaticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationResourceStaticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationResourceStaticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebApplicationResourceStaticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebApplicationResourceStaticsResponse) GetBody() *WebApplicationResourceStaticsBody {
	return s.Body
}

func (s *DescribeWebApplicationResourceStaticsResponse) SetHeaders(v map[string]*string) *DescribeWebApplicationResourceStaticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebApplicationResourceStaticsResponse) SetStatusCode(v int32) *DescribeWebApplicationResourceStaticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebApplicationResourceStaticsResponse) SetBody(v *WebApplicationResourceStaticsBody) *DescribeWebApplicationResourceStaticsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebApplicationResourceStaticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
