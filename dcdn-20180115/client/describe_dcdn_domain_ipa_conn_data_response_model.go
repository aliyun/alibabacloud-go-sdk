// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaConnDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaConnDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainIpaConnDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainIpaConnDataResponseBody) *DescribeDcdnDomainIpaConnDataResponse
	GetBody() *DescribeDcdnDomainIpaConnDataResponseBody
}

type DescribeDcdnDomainIpaConnDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainIpaConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainIpaConnDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaConnDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaConnDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainIpaConnDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainIpaConnDataResponse) GetBody() *DescribeDcdnDomainIpaConnDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainIpaConnDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaConnDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainIpaConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponse) SetBody(v *DescribeDcdnDomainIpaConnDataResponseBody) *DescribeDcdnDomainIpaConnDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataResponse) Validate() error {
	return dara.Validate(s)
}
