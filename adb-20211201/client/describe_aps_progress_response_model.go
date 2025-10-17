// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsProgressResponseBody) *DescribeApsProgressResponse
	GetBody() *DescribeApsProgressResponseBody
}

type DescribeApsProgressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsProgressResponse) GetBody() *DescribeApsProgressResponseBody {
	return s.Body
}

func (s *DescribeApsProgressResponse) SetHeaders(v map[string]*string) *DescribeApsProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsProgressResponse) SetStatusCode(v int32) *DescribeApsProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsProgressResponse) SetBody(v *DescribeApsProgressResponseBody) *DescribeApsProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeApsProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
