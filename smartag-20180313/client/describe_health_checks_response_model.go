// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthChecksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthChecksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthChecksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthChecksResponseBody) *DescribeHealthChecksResponse
	GetBody() *DescribeHealthChecksResponseBody
}

type DescribeHealthChecksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthChecksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthChecksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthChecksResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthChecksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthChecksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthChecksResponse) GetBody() *DescribeHealthChecksResponseBody {
	return s.Body
}

func (s *DescribeHealthChecksResponse) SetHeaders(v map[string]*string) *DescribeHealthChecksResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthChecksResponse) SetStatusCode(v int32) *DescribeHealthChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthChecksResponse) SetBody(v *DescribeHealthChecksResponseBody) *DescribeHealthChecksResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthChecksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
