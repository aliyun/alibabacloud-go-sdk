// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvisorChecksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvisorChecksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvisorChecksResponseBody) *DescribeAdvisorChecksResponse
	GetBody() *DescribeAdvisorChecksResponseBody
}

type DescribeAdvisorChecksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorChecksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorChecksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvisorChecksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvisorChecksResponse) GetBody() *DescribeAdvisorChecksResponseBody {
	return s.Body
}

func (s *DescribeAdvisorChecksResponse) SetHeaders(v map[string]*string) *DescribeAdvisorChecksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorChecksResponse) SetStatusCode(v int32) *DescribeAdvisorChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorChecksResponse) SetBody(v *DescribeAdvisorChecksResponseBody) *DescribeAdvisorChecksResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvisorChecksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
