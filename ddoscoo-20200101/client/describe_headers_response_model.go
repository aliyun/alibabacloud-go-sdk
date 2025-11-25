// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHeadersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHeadersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHeadersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHeadersResponseBody) *DescribeHeadersResponse
	GetBody() *DescribeHeadersResponseBody
}

type DescribeHeadersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHeadersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHeadersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHeadersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHeadersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHeadersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHeadersResponse) GetBody() *DescribeHeadersResponseBody {
	return s.Body
}

func (s *DescribeHeadersResponse) SetHeaders(v map[string]*string) *DescribeHeadersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHeadersResponse) SetStatusCode(v int32) *DescribeHeadersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHeadersResponse) SetBody(v *DescribeHeadersResponseBody) *DescribeHeadersResponse {
	s.Body = v
	return s
}

func (s *DescribeHeadersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
