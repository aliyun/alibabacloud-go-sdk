// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainTrafficDataResponseBody) *DescribeDcdnDomainTrafficDataResponse
	GetBody() *DescribeDcdnDomainTrafficDataResponseBody
}

type DescribeDcdnDomainTrafficDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainTrafficDataResponse) GetBody() *DescribeDcdnDomainTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponse) SetBody(v *DescribeDcdnDomainTrafficDataResponseBody) *DescribeDcdnDomainTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
