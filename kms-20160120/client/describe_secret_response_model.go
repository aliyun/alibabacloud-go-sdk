// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecretResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecretResponseBody) *DescribeSecretResponse
	GetBody() *DescribeSecretResponseBody
}

type DescribeSecretResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecretResponse) GetBody() *DescribeSecretResponseBody {
	return s.Body
}

func (s *DescribeSecretResponse) SetHeaders(v map[string]*string) *DescribeSecretResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecretResponse) SetStatusCode(v int32) *DescribeSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecretResponse) SetBody(v *DescribeSecretResponseBody) *DescribeSecretResponse {
	s.Body = v
	return s
}

func (s *DescribeSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
