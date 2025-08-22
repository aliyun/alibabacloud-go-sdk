// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpaDomainCidrResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpaDomainCidrResponseBody) *DescribeDcdnIpaDomainCidrResponse
	GetBody() *DescribeDcdnIpaDomainCidrResponseBody
}

type DescribeDcdnIpaDomainCidrResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpaDomainCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpaDomainCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainCidrResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpaDomainCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpaDomainCidrResponse) GetBody() *DescribeDcdnIpaDomainCidrResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpaDomainCidrResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainCidrResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaDomainCidrResponse) SetStatusCode(v int32) *DescribeDcdnIpaDomainCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpaDomainCidrResponse) SetBody(v *DescribeDcdnIpaDomainCidrResponseBody) *DescribeDcdnIpaDomainCidrResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpaDomainCidrResponse) Validate() error {
	return dara.Validate(s)
}
