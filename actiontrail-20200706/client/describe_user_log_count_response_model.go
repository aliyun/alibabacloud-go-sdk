// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserLogCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserLogCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserLogCountResponseBody) *DescribeUserLogCountResponse
	GetBody() *DescribeUserLogCountResponseBody
}

type DescribeUserLogCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLogCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserLogCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserLogCountResponse) GetBody() *DescribeUserLogCountResponseBody {
	return s.Body
}

func (s *DescribeUserLogCountResponse) SetHeaders(v map[string]*string) *DescribeUserLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogCountResponse) SetStatusCode(v int32) *DescribeUserLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLogCountResponse) SetBody(v *DescribeUserLogCountResponseBody) *DescribeUserLogCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUserLogCountResponse) Validate() error {
	return dara.Validate(s)
}
