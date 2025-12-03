// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubDomainResponseBody) *DescribeSubDomainResponse
	GetBody() *DescribeSubDomainResponseBody
}

type DescribeSubDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubDomainResponse) GetBody() *DescribeSubDomainResponseBody {
	return s.Body
}

func (s *DescribeSubDomainResponse) SetHeaders(v map[string]*string) *DescribeSubDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubDomainResponse) SetStatusCode(v int32) *DescribeSubDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubDomainResponse) SetBody(v *DescribeSubDomainResponseBody) *DescribeSubDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeSubDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
