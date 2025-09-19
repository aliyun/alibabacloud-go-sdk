// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSensitiveFileByKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageSensitiveFileByKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageSensitiveFileByKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageSensitiveFileByKeyResponseBody) *DescribeImageSensitiveFileByKeyResponse
	GetBody() *DescribeImageSensitiveFileByKeyResponseBody
}

type DescribeImageSensitiveFileByKeyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageSensitiveFileByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageSensitiveFileByKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSensitiveFileByKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSensitiveFileByKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageSensitiveFileByKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageSensitiveFileByKeyResponse) GetBody() *DescribeImageSensitiveFileByKeyResponseBody {
	return s.Body
}

func (s *DescribeImageSensitiveFileByKeyResponse) SetHeaders(v map[string]*string) *DescribeImageSensitiveFileByKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponse) SetStatusCode(v int32) *DescribeImageSensitiveFileByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponse) SetBody(v *DescribeImageSensitiveFileByKeyResponseBody) *DescribeImageSensitiveFileByKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeImageSensitiveFileByKeyResponse) Validate() error {
	return dara.Validate(s)
}
