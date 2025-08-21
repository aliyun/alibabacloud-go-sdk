// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewSourceCountriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainViewSourceCountriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainViewSourceCountriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainViewSourceCountriesResponseBody) *DescribeDomainViewSourceCountriesResponse
	GetBody() *DescribeDomainViewSourceCountriesResponseBody
}

type DescribeDomainViewSourceCountriesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainViewSourceCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainViewSourceCountriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainViewSourceCountriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainViewSourceCountriesResponse) GetBody() *DescribeDomainViewSourceCountriesResponseBody {
	return s.Body
}

func (s *DescribeDomainViewSourceCountriesResponse) SetHeaders(v map[string]*string) *DescribeDomainViewSourceCountriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponse) SetStatusCode(v int32) *DescribeDomainViewSourceCountriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponse) SetBody(v *DescribeDomainViewSourceCountriesResponseBody) *DescribeDomainViewSourceCountriesResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponse) Validate() error {
	return dara.Validate(s)
}
