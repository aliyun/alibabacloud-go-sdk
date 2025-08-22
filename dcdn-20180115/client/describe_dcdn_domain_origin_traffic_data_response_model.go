// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainOriginTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainOriginTrafficDataResponseBody) *DescribeDcdnDomainOriginTrafficDataResponse
	GetBody() *DescribeDcdnDomainOriginTrafficDataResponseBody
}

type DescribeDcdnDomainOriginTrafficDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainOriginTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) GetBody() *DescribeDcdnDomainOriginTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainOriginTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) SetBody(v *DescribeDcdnDomainOriginTrafficDataResponseBody) *DescribeDcdnDomainOriginTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
