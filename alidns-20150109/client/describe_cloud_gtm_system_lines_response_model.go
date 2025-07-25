// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmSystemLinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmSystemLinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmSystemLinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmSystemLinesResponseBody) *DescribeCloudGtmSystemLinesResponse
	GetBody() *DescribeCloudGtmSystemLinesResponseBody
}

type DescribeCloudGtmSystemLinesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmSystemLinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmSystemLinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSystemLinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSystemLinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmSystemLinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmSystemLinesResponse) GetBody() *DescribeCloudGtmSystemLinesResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmSystemLinesResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmSystemLinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponse) SetStatusCode(v int32) *DescribeCloudGtmSystemLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponse) SetBody(v *DescribeCloudGtmSystemLinesResponseBody) *DescribeCloudGtmSystemLinesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponse) Validate() error {
	return dara.Validate(s)
}
