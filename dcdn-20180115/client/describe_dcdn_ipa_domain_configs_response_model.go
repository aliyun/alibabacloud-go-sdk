// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpaDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpaDomainConfigsResponseBody) *DescribeDcdnIpaDomainConfigsResponse
	GetBody() *DescribeDcdnIpaDomainConfigsResponseBody
}

type DescribeDcdnIpaDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpaDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpaDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpaDomainConfigsResponse) GetBody() *DescribeDcdnIpaDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpaDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponse) SetStatusCode(v int32) *DescribeDcdnIpaDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponse) SetBody(v *DescribeDcdnIpaDomainConfigsResponseBody) *DescribeDcdnIpaDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
