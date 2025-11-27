// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecretsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecretsResponseBody) *DescribeSecretsResponse
	GetBody() *DescribeSecretsResponseBody
}

type DescribeSecretsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecretsResponse) GetBody() *DescribeSecretsResponseBody {
	return s.Body
}

func (s *DescribeSecretsResponse) SetHeaders(v map[string]*string) *DescribeSecretsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecretsResponse) SetStatusCode(v int32) *DescribeSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecretsResponse) SetBody(v *DescribeSecretsResponseBody) *DescribeSecretsResponse {
	s.Body = v
	return s
}

func (s *DescribeSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
