// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVipsByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserVipsByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserVipsByDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserVipsByDomainResponseBody) *DescribeUserVipsByDomainResponse
	GetBody() *DescribeUserVipsByDomainResponseBody
}

type DescribeUserVipsByDomainResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserVipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserVipsByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserVipsByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserVipsByDomainResponse) GetBody() *DescribeUserVipsByDomainResponseBody {
	return s.Body
}

func (s *DescribeUserVipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeUserVipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserVipsByDomainResponse) SetStatusCode(v int32) *DescribeUserVipsByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserVipsByDomainResponse) SetBody(v *DescribeUserVipsByDomainResponseBody) *DescribeUserVipsByDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeUserVipsByDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
