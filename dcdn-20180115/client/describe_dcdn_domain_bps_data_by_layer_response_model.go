// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainBpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainBpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainBpsDataByLayerResponseBody) *DescribeDcdnDomainBpsDataByLayerResponse
	GetBody() *DescribeDcdnDomainBpsDataByLayerResponseBody
}

type DescribeDcdnDomainBpsDataByLayerResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainBpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainBpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) GetBody() *DescribeDcdnDomainBpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainBpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) SetStatusCode(v int32) *DescribeDcdnDomainBpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) SetBody(v *DescribeDcdnDomainBpsDataByLayerResponseBody) *DescribeDcdnDomainBpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
