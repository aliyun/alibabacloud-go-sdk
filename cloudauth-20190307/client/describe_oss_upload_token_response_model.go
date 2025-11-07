// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssUploadTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssUploadTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssUploadTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssUploadTokenResponseBody) *DescribeOssUploadTokenResponse
	GetBody() *DescribeOssUploadTokenResponseBody
}

type DescribeOssUploadTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssUploadTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssUploadTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssUploadTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssUploadTokenResponse) GetBody() *DescribeOssUploadTokenResponseBody {
	return s.Body
}

func (s *DescribeOssUploadTokenResponse) SetHeaders(v map[string]*string) *DescribeOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetStatusCode(v int32) *DescribeOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssUploadTokenResponse) SetBody(v *DescribeOssUploadTokenResponseBody) *DescribeOssUploadTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeOssUploadTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
