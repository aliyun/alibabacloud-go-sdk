// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageSensitiveFileListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageSensitiveFileListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageSensitiveFileListResponseBody) *DescribeImageSensitiveFileListResponse
	GetBody() *DescribeImageSensitiveFileListResponseBody
}

type DescribeImageSensitiveFileListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageSensitiveFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageSensitiveFileListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageSensitiveFileListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageSensitiveFileListResponse) GetBody() *DescribeImageSensitiveFileListResponseBody {
	return s.Body
}

func (s *DescribeImageSensitiveFileListResponse) SetHeaders(v map[string]*string) *DescribeImageSensitiveFileListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSensitiveFileListResponse) SetStatusCode(v int32) *DescribeImageSensitiveFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageSensitiveFileListResponse) SetBody(v *DescribeImageSensitiveFileListResponseBody) *DescribeImageSensitiveFileListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageSensitiveFileListResponse) Validate() error {
	return dara.Validate(s)
}
