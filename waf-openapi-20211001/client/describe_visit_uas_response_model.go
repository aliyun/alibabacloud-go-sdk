// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitUasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVisitUasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVisitUasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVisitUasResponseBody) *DescribeVisitUasResponse
	GetBody() *DescribeVisitUasResponseBody
}

type DescribeVisitUasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVisitUasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVisitUasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitUasResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVisitUasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVisitUasResponse) GetBody() *DescribeVisitUasResponseBody {
	return s.Body
}

func (s *DescribeVisitUasResponse) SetHeaders(v map[string]*string) *DescribeVisitUasResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitUasResponse) SetStatusCode(v int32) *DescribeVisitUasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitUasResponse) SetBody(v *DescribeVisitUasResponseBody) *DescribeVisitUasResponse {
	s.Body = v
	return s
}

func (s *DescribeVisitUasResponse) Validate() error {
	return dara.Validate(s)
}
