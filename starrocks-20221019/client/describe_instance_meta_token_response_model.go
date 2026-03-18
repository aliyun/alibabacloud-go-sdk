// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMetaTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceMetaTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceMetaTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceMetaTokenResponseBody) *DescribeInstanceMetaTokenResponse
	GetBody() *DescribeInstanceMetaTokenResponseBody
}

type DescribeInstanceMetaTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMetaTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMetaTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMetaTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMetaTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceMetaTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceMetaTokenResponse) GetBody() *DescribeInstanceMetaTokenResponseBody {
	return s.Body
}

func (s *DescribeInstanceMetaTokenResponse) SetHeaders(v map[string]*string) *DescribeInstanceMetaTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMetaTokenResponse) SetStatusCode(v int32) *DescribeInstanceMetaTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMetaTokenResponse) SetBody(v *DescribeInstanceMetaTokenResponseBody) *DescribeInstanceMetaTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceMetaTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
