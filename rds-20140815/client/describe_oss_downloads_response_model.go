// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssDownloadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssDownloadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssDownloadsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssDownloadsResponseBody) *DescribeOssDownloadsResponse
	GetBody() *DescribeOssDownloadsResponseBody
}

type DescribeOssDownloadsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssDownloadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssDownloadsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssDownloadsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssDownloadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssDownloadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssDownloadsResponse) GetBody() *DescribeOssDownloadsResponseBody {
	return s.Body
}

func (s *DescribeOssDownloadsResponse) SetHeaders(v map[string]*string) *DescribeOssDownloadsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssDownloadsResponse) SetStatusCode(v int32) *DescribeOssDownloadsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssDownloadsResponse) SetBody(v *DescribeOssDownloadsResponseBody) *DescribeOssDownloadsResponse {
	s.Body = v
	return s
}

func (s *DescribeOssDownloadsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
