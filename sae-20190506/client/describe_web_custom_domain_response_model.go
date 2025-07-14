// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *WebCustomDomain) *DescribeWebCustomDomainResponse
	GetBody() *WebCustomDomain
}

type DescribeWebCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebCustomDomain   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCustomDomainResponse) GetBody() *WebCustomDomain {
	return s.Body
}

func (s *DescribeWebCustomDomainResponse) SetHeaders(v map[string]*string) *DescribeWebCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCustomDomainResponse) SetStatusCode(v int32) *DescribeWebCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCustomDomainResponse) SetBody(v *WebCustomDomain) *DescribeWebCustomDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeWebCustomDomainResponse) Validate() error {
	return dara.Validate(s)
}
