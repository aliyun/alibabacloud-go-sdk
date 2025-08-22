// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainUsageDataResponseBody) *DescribeDcdnDomainUsageDataResponse
	GetBody() *DescribeDcdnDomainUsageDataResponseBody
}

type DescribeDcdnDomainUsageDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainUsageDataResponse) GetBody() *DescribeDcdnDomainUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponse) SetBody(v *DescribeDcdnDomainUsageDataResponseBody) *DescribeDcdnDomainUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
