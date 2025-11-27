// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostgresExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostgresExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostgresExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostgresExtensionsResponseBody) *DescribePostgresExtensionsResponse
	GetBody() *DescribePostgresExtensionsResponseBody
}

type DescribePostgresExtensionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostgresExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostgresExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostgresExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePostgresExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostgresExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostgresExtensionsResponse) GetBody() *DescribePostgresExtensionsResponseBody {
	return s.Body
}

func (s *DescribePostgresExtensionsResponse) SetHeaders(v map[string]*string) *DescribePostgresExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePostgresExtensionsResponse) SetStatusCode(v int32) *DescribePostgresExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostgresExtensionsResponse) SetBody(v *DescribePostgresExtensionsResponseBody) *DescribePostgresExtensionsResponse {
	s.Body = v
	return s
}

func (s *DescribePostgresExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
