// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLProtectTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeACLProtectTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeACLProtectTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeACLProtectTrendResponseBody) *DescribeACLProtectTrendResponse
	GetBody() *DescribeACLProtectTrendResponseBody
}

type DescribeACLProtectTrendResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeACLProtectTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeACLProtectTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLProtectTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeACLProtectTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeACLProtectTrendResponse) GetBody() *DescribeACLProtectTrendResponseBody {
	return s.Body
}

func (s *DescribeACLProtectTrendResponse) SetHeaders(v map[string]*string) *DescribeACLProtectTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeACLProtectTrendResponse) SetStatusCode(v int32) *DescribeACLProtectTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeACLProtectTrendResponse) SetBody(v *DescribeACLProtectTrendResponseBody) *DescribeACLProtectTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeACLProtectTrendResponse) Validate() error {
	return dara.Validate(s)
}
