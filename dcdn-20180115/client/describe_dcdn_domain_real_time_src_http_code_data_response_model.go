// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
	GetBody() *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) GetBody() *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
