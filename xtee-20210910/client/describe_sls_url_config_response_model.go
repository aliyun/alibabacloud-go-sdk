// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsUrlConfigResponseBody) *DescribeSlsUrlConfigResponse
	GetBody() *DescribeSlsUrlConfigResponseBody
}

type DescribeSlsUrlConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsUrlConfigResponse) GetBody() *DescribeSlsUrlConfigResponseBody {
	return s.Body
}

func (s *DescribeSlsUrlConfigResponse) SetHeaders(v map[string]*string) *DescribeSlsUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsUrlConfigResponse) SetStatusCode(v int32) *DescribeSlsUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsUrlConfigResponse) SetBody(v *DescribeSlsUrlConfigResponseBody) *DescribeSlsUrlConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsUrlConfigResponse) Validate() error {
	return dara.Validate(s)
}
