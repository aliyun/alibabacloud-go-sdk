// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudResourceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudResourceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudResourceStatusResponseBody) *DescribeCloudResourceStatusResponse
	GetBody() *DescribeCloudResourceStatusResponseBody
}

type DescribeCloudResourceStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudResourceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudResourceStatusResponse) GetBody() *DescribeCloudResourceStatusResponseBody {
	return s.Body
}

func (s *DescribeCloudResourceStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudResourceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourceStatusResponse) SetStatusCode(v int32) *DescribeCloudResourceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourceStatusResponse) SetBody(v *DescribeCloudResourceStatusResponseBody) *DescribeCloudResourceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudResourceStatusResponse) Validate() error {
	return dara.Validate(s)
}
