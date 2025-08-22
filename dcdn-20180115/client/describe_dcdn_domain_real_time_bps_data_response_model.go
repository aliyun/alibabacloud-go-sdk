// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeBpsDataResponseBody) *DescribeDcdnDomainRealTimeBpsDataResponse
	GetBody() *DescribeDcdnDomainRealTimeBpsDataResponseBody
}

type DescribeDcdnDomainRealTimeBpsDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) GetBody() *DescribeDcdnDomainRealTimeBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeBpsDataResponseBody) *DescribeDcdnDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
