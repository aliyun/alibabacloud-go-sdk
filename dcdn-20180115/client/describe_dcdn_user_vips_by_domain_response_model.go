// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserVipsByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserVipsByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserVipsByDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserVipsByDomainResponseBody) *DescribeDcdnUserVipsByDomainResponse
	GetBody() *DescribeDcdnUserVipsByDomainResponseBody
}

type DescribeDcdnUserVipsByDomainResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserVipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserVipsByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserVipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserVipsByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserVipsByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserVipsByDomainResponse) GetBody() *DescribeDcdnUserVipsByDomainResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserVipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserVipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponse) SetStatusCode(v int32) *DescribeDcdnUserVipsByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponse) SetBody(v *DescribeDcdnUserVipsByDomainResponseBody) *DescribeDcdnUserVipsByDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponse) Validate() error {
	return dara.Validate(s)
}
