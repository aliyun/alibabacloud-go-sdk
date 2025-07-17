// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileResponseBody) *DescribeFileResponse
	GetBody() *DescribeFileResponseBody
}

type DescribeFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileResponse) GetBody() *DescribeFileResponseBody {
	return s.Body
}

func (s *DescribeFileResponse) SetHeaders(v map[string]*string) *DescribeFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileResponse) SetStatusCode(v int32) *DescribeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileResponse) SetBody(v *DescribeFileResponseBody) *DescribeFileResponse {
	s.Body = v
	return s
}

func (s *DescribeFileResponse) Validate() error {
	return dara.Validate(s)
}
