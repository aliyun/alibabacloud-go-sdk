// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainQpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainQpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainQpsDataByLayerResponseBody) *DescribeDcdnDomainQpsDataByLayerResponse
	GetBody() *DescribeDcdnDomainQpsDataByLayerResponseBody
}

type DescribeDcdnDomainQpsDataByLayerResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainQpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainQpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) GetBody() *DescribeDcdnDomainQpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainQpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) SetStatusCode(v int32) *DescribeDcdnDomainQpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) SetBody(v *DescribeDcdnDomainQpsDataByLayerResponseBody) *DescribeDcdnDomainQpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
