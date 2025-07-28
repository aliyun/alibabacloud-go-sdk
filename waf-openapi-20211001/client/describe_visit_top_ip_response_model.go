// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitTopIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVisitTopIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVisitTopIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVisitTopIpResponseBody) *DescribeVisitTopIpResponse
	GetBody() *DescribeVisitTopIpResponseBody
}

type DescribeVisitTopIpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVisitTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVisitTopIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVisitTopIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVisitTopIpResponse) GetBody() *DescribeVisitTopIpResponseBody {
	return s.Body
}

func (s *DescribeVisitTopIpResponse) SetHeaders(v map[string]*string) *DescribeVisitTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitTopIpResponse) SetStatusCode(v int32) *DescribeVisitTopIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitTopIpResponse) SetBody(v *DescribeVisitTopIpResponseBody) *DescribeVisitTopIpResponse {
	s.Body = v
	return s
}

func (s *DescribeVisitTopIpResponse) Validate() error {
	return dara.Validate(s)
}
