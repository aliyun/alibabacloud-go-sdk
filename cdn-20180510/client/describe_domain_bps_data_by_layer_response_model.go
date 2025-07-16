// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainBpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainBpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainBpsDataByLayerResponseBody) *DescribeDomainBpsDataByLayerResponse
	GetBody() *DescribeDomainBpsDataByLayerResponseBody
}

type DescribeDomainBpsDataByLayerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainBpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainBpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainBpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainBpsDataByLayerResponse) GetBody() *DescribeDomainBpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDomainBpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponse) SetStatusCode(v int32) *DescribeDomainBpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponse) SetBody(v *DescribeDomainBpsDataByLayerResponseBody) *DescribeDomainBpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
