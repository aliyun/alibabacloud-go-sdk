// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) *DescribeDcdnDomainRealTimeSrcBpsDataResponse
	GetBody() *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) GetBody() *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) *DescribeDcdnDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
