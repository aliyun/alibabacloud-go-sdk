// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainOriginBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainOriginBpsDataResponseBody) *DescribeDcdnDomainOriginBpsDataResponse
	GetBody() *DescribeDcdnDomainOriginBpsDataResponseBody
}

type DescribeDcdnDomainOriginBpsDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainOriginBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) GetBody() *DescribeDcdnDomainOriginBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainOriginBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) SetBody(v *DescribeDcdnDomainOriginBpsDataResponseBody) *DescribeDcdnDomainOriginBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
