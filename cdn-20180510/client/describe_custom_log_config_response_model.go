// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomLogConfigResponseBody) *DescribeCustomLogConfigResponse
	GetBody() *DescribeCustomLogConfigResponseBody
}

type DescribeCustomLogConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomLogConfigResponse) GetBody() *DescribeCustomLogConfigResponseBody {
	return s.Body
}

func (s *DescribeCustomLogConfigResponse) SetHeaders(v map[string]*string) *DescribeCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLogConfigResponse) SetStatusCode(v int32) *DescribeCustomLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomLogConfigResponse) SetBody(v *DescribeCustomLogConfigResponseBody) *DescribeCustomLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
