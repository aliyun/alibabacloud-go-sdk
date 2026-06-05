// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyUserDataUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyUserDataUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyUserDataUploadUrlResponseBody) *DescribeComfyUserDataUploadUrlResponse
	GetBody() *DescribeComfyUserDataUploadUrlResponseBody
}

type DescribeComfyUserDataUploadUrlResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyUserDataUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyUserDataUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyUserDataUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyUserDataUploadUrlResponse) GetBody() *DescribeComfyUserDataUploadUrlResponseBody {
	return s.Body
}

func (s *DescribeComfyUserDataUploadUrlResponse) SetHeaders(v map[string]*string) *DescribeComfyUserDataUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponse) SetStatusCode(v int32) *DescribeComfyUserDataUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponse) SetBody(v *DescribeComfyUserDataUploadUrlResponseBody) *DescribeComfyUserDataUploadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyUserDataUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
