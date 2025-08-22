// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeDetailDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeDetailDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeDetailDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeDetailDataResponseBody) *DescribeDcdnDomainRealTimeDetailDataResponse
	GetBody() *DescribeDcdnDomainRealTimeDetailDataResponseBody
}

type DescribeDcdnDomainRealTimeDetailDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeDetailDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) GetBody() *DescribeDcdnDomainRealTimeDetailDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) SetBody(v *DescribeDcdnDomainRealTimeDetailDataResponseBody) *DescribeDcdnDomainRealTimeDetailDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) Validate() error {
	return dara.Validate(s)
}
