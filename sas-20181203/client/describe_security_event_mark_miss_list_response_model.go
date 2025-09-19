// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventMarkMissListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventMarkMissListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventMarkMissListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventMarkMissListResponseBody) *DescribeSecurityEventMarkMissListResponse
	GetBody() *DescribeSecurityEventMarkMissListResponseBody
}

type DescribeSecurityEventMarkMissListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventMarkMissListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventMarkMissListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventMarkMissListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventMarkMissListResponse) GetBody() *DescribeSecurityEventMarkMissListResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventMarkMissListResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventMarkMissListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponse) SetStatusCode(v int32) *DescribeSecurityEventMarkMissListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponse) SetBody(v *DescribeSecurityEventMarkMissListResponseBody) *DescribeSecurityEventMarkMissListResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponse) Validate() error {
	return dara.Validate(s)
}
