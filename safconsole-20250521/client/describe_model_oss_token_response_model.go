// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOssTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelOssTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelOssTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelOssTokenResponseBody) *DescribeModelOssTokenResponse
	GetBody() *DescribeModelOssTokenResponseBody
}

type DescribeModelOssTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelOssTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelOssTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelOssTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelOssTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelOssTokenResponse) GetBody() *DescribeModelOssTokenResponseBody {
	return s.Body
}

func (s *DescribeModelOssTokenResponse) SetHeaders(v map[string]*string) *DescribeModelOssTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelOssTokenResponse) SetStatusCode(v int32) *DescribeModelOssTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelOssTokenResponse) SetBody(v *DescribeModelOssTokenResponseBody) *DescribeModelOssTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeModelOssTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
