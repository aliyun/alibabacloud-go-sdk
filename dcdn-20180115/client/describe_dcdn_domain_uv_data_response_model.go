// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainUvDataResponseBody) *DescribeDcdnDomainUvDataResponse
	GetBody() *DescribeDcdnDomainUvDataResponseBody
}

type DescribeDcdnDomainUvDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainUvDataResponse) GetBody() *DescribeDcdnDomainUvDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainUvDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponse) SetBody(v *DescribeDcdnDomainUvDataResponseBody) *DescribeDcdnDomainUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainUvDataResponse) Validate() error {
	return dara.Validate(s)
}
