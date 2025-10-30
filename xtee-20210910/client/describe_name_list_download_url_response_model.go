// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListDownloadUrlResponseBody) *DescribeNameListDownloadUrlResponse
	GetBody() *DescribeNameListDownloadUrlResponseBody
}

type DescribeNameListDownloadUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListDownloadUrlResponse) GetBody() *DescribeNameListDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeNameListDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeNameListDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListDownloadUrlResponse) SetStatusCode(v int32) *DescribeNameListDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListDownloadUrlResponse) SetBody(v *DescribeNameListDownloadUrlResponseBody) *DescribeNameListDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
