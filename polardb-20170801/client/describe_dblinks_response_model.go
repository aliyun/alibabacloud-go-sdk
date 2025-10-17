// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLinksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBLinksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBLinksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBLinksResponseBody) *DescribeDBLinksResponse
	GetBody() *DescribeDBLinksResponseBody
}

type DescribeDBLinksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBLinksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBLinksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBLinksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBLinksResponse) GetBody() *DescribeDBLinksResponseBody {
	return s.Body
}

func (s *DescribeDBLinksResponse) SetHeaders(v map[string]*string) *DescribeDBLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBLinksResponse) SetStatusCode(v int32) *DescribeDBLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBLinksResponse) SetBody(v *DescribeDBLinksResponseBody) *DescribeDBLinksResponse {
	s.Body = v
	return s
}

func (s *DescribeDBLinksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
