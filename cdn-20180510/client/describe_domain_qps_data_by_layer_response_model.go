// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainQpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainQpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainQpsDataByLayerResponseBody) *DescribeDomainQpsDataByLayerResponse
	GetBody() *DescribeDomainQpsDataByLayerResponseBody
}

type DescribeDomainQpsDataByLayerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainQpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainQpsDataByLayerResponse) GetBody() *DescribeDomainQpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDomainQpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponse) SetStatusCode(v int32) *DescribeDomainQpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponse) SetBody(v *DescribeDomainQpsDataByLayerResponseBody) *DescribeDomainQpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponse) Validate() error {
	return dara.Validate(s)
}
