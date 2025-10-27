// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListBySensitiveFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageListBySensitiveFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageListBySensitiveFileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageListBySensitiveFileResponseBody) *DescribeImageListBySensitiveFileResponse
	GetBody() *DescribeImageListBySensitiveFileResponseBody
}

type DescribeImageListBySensitiveFileResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageListBySensitiveFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageListBySensitiveFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListBySensitiveFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageListBySensitiveFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageListBySensitiveFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageListBySensitiveFileResponse) GetBody() *DescribeImageListBySensitiveFileResponseBody {
	return s.Body
}

func (s *DescribeImageListBySensitiveFileResponse) SetHeaders(v map[string]*string) *DescribeImageListBySensitiveFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageListBySensitiveFileResponse) SetStatusCode(v int32) *DescribeImageListBySensitiveFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponse) SetBody(v *DescribeImageListBySensitiveFileResponseBody) *DescribeImageListBySensitiveFileResponse {
	s.Body = v
	return s
}

func (s *DescribeImageListBySensitiveFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
