// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceAccountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceAccountResponseBody) *DescribeServiceAccountResponse
	GetBody() *DescribeServiceAccountResponseBody
}

type DescribeServiceAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceAccountResponse) GetBody() *DescribeServiceAccountResponseBody {
	return s.Body
}

func (s *DescribeServiceAccountResponse) SetHeaders(v map[string]*string) *DescribeServiceAccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAccountResponse) SetStatusCode(v int32) *DescribeServiceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAccountResponse) SetBody(v *DescribeServiceAccountResponseBody) *DescribeServiceAccountResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
