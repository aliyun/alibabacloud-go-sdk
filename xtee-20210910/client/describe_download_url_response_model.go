// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadUrlResponseBody) *DescribeDownloadUrlResponse
	GetBody() *DescribeDownloadUrlResponseBody
}

type DescribeDownloadUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadUrlResponse) GetBody() *DescribeDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadUrlResponse) SetStatusCode(v int32) *DescribeDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadUrlResponse) SetBody(v *DescribeDownloadUrlResponseBody) *DescribeDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
