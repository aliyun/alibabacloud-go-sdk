// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyUserDataDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyUserDataDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyUserDataDownloadUrlResponseBody) *DescribeComfyUserDataDownloadUrlResponse
	GetBody() *DescribeComfyUserDataDownloadUrlResponseBody
}

type DescribeComfyUserDataDownloadUrlResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyUserDataDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyUserDataDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyUserDataDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyUserDataDownloadUrlResponse) GetBody() *DescribeComfyUserDataDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeComfyUserDataDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeComfyUserDataDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponse) SetStatusCode(v int32) *DescribeComfyUserDataDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponse) SetBody(v *DescribeComfyUserDataDownloadUrlResponseBody) *DescribeComfyUserDataDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
