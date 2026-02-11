// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileDownloadUrlResponseBody) *DescribeFileDownloadUrlResponse
	GetBody() *DescribeFileDownloadUrlResponseBody
}

type DescribeFileDownloadUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileDownloadUrlResponse) GetBody() *DescribeFileDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeFileDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileDownloadUrlResponse) SetStatusCode(v int32) *DescribeFileDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileDownloadUrlResponse) SetBody(v *DescribeFileDownloadUrlResponseBody) *DescribeFileDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeFileDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
