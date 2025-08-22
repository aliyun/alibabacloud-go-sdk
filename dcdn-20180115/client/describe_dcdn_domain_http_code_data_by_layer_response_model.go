// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainHttpCodeDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainHttpCodeDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) *DescribeDcdnDomainHttpCodeDataByLayerResponse
	GetBody() *DescribeDcdnDomainHttpCodeDataByLayerResponseBody
}

type DescribeDcdnDomainHttpCodeDataByLayerResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainHttpCodeDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) GetBody() *DescribeDcdnDomainHttpCodeDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainHttpCodeDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) SetStatusCode(v int32) *DescribeDcdnDomainHttpCodeDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) SetBody(v *DescribeDcdnDomainHttpCodeDataByLayerResponseBody) *DescribeDcdnDomainHttpCodeDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
