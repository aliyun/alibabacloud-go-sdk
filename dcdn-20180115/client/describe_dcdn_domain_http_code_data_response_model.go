// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainHttpCodeDataResponseBody) *DescribeDcdnDomainHttpCodeDataResponse
	GetBody() *DescribeDcdnDomainHttpCodeDataResponseBody
}

type DescribeDcdnDomainHttpCodeDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) GetBody() *DescribeDcdnDomainHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainHttpCodeDataResponseBody) *DescribeDcdnDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
