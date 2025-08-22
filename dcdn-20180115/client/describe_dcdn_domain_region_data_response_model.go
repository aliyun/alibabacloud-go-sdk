// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRegionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRegionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRegionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRegionDataResponseBody) *DescribeDcdnDomainRegionDataResponse
	GetBody() *DescribeDcdnDomainRegionDataResponseBody
}

type DescribeDcdnDomainRegionDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRegionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRegionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRegionDataResponse) GetBody() *DescribeDcdnDomainRegionDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponse) SetBody(v *DescribeDcdnDomainRegionDataResponseBody) *DescribeDcdnDomainRegionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponse) Validate() error {
	return dara.Validate(s)
}
