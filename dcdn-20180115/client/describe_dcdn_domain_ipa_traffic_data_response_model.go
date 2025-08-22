// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainIpaTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainIpaTrafficDataResponseBody) *DescribeDcdnDomainIpaTrafficDataResponse
	GetBody() *DescribeDcdnDomainIpaTrafficDataResponseBody
}

type DescribeDcdnDomainIpaTrafficDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainIpaTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) GetBody() *DescribeDcdnDomainIpaTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainIpaTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) SetBody(v *DescribeDcdnDomainIpaTrafficDataResponseBody) *DescribeDcdnDomainIpaTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
