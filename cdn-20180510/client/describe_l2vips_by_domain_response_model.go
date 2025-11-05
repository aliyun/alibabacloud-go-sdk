// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL2VipsByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeL2VipsByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeL2VipsByDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeL2VipsByDomainResponseBody) *DescribeL2VipsByDomainResponse
	GetBody() *DescribeL2VipsByDomainResponseBody
}

type DescribeL2VipsByDomainResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeL2VipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeL2VipsByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeL2VipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeL2VipsByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeL2VipsByDomainResponse) GetBody() *DescribeL2VipsByDomainResponseBody {
	return s.Body
}

func (s *DescribeL2VipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeL2VipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeL2VipsByDomainResponse) SetStatusCode(v int32) *DescribeL2VipsByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeL2VipsByDomainResponse) SetBody(v *DescribeL2VipsByDomainResponseBody) *DescribeL2VipsByDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeL2VipsByDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
