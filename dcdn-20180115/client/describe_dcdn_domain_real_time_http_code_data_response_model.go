// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeHttpCodeDataResponse
	GetBody() *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) GetBody() *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
