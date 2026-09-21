// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiKeyResponseBody) *DescribeApiKeyResponse
	GetBody() *DescribeApiKeyResponseBody
}

type DescribeApiKeyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiKeyResponse) GetBody() *DescribeApiKeyResponseBody {
	return s.Body
}

func (s *DescribeApiKeyResponse) SetHeaders(v map[string]*string) *DescribeApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiKeyResponse) SetStatusCode(v int32) *DescribeApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiKeyResponse) SetBody(v *DescribeApiKeyResponseBody) *DescribeApiKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
