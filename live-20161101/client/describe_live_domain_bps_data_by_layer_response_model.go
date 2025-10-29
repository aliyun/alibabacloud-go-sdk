// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainBpsDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainBpsDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainBpsDataByLayerResponseBody) *DescribeLiveDomainBpsDataByLayerResponse
	GetBody() *DescribeLiveDomainBpsDataByLayerResponseBody
}

type DescribeLiveDomainBpsDataByLayerResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainBpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainBpsDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) GetBody() *DescribeLiveDomainBpsDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainBpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) SetStatusCode(v int32) *DescribeLiveDomainBpsDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) SetBody(v *DescribeLiveDomainBpsDataByLayerResponseBody) *DescribeLiveDomainBpsDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
