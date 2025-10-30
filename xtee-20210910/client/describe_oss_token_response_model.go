// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssTokenResponseBody) *DescribeOssTokenResponse
	GetBody() *DescribeOssTokenResponseBody
}

type DescribeOssTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssTokenResponse) GetBody() *DescribeOssTokenResponseBody {
	return s.Body
}

func (s *DescribeOssTokenResponse) SetHeaders(v map[string]*string) *DescribeOssTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssTokenResponse) SetStatusCode(v int32) *DescribeOssTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssTokenResponse) SetBody(v *DescribeOssTokenResponseBody) *DescribeOssTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeOssTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
