// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainDetailDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainDetailDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainDetailDataByLayerResponseBody) *DescribeDomainDetailDataByLayerResponse
	GetBody() *DescribeDomainDetailDataByLayerResponseBody
}

type DescribeDomainDetailDataByLayerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDetailDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDetailDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainDetailDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainDetailDataByLayerResponse) GetBody() *DescribeDomainDetailDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDomainDetailDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponse) SetStatusCode(v int32) *DescribeDomainDetailDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponse) SetBody(v *DescribeDomainDetailDataByLayerResponseBody) *DescribeDomainDetailDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
