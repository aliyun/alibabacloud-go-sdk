// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRegionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRegionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRegionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRegionDataResponseBody) *DescribeDomainRegionDataResponse
	GetBody() *DescribeDomainRegionDataResponseBody
}

type DescribeDomainRegionDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRegionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRegionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRegionDataResponse) GetBody() *DescribeDomainRegionDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRegionDataResponse) SetStatusCode(v int32) *DescribeDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRegionDataResponse) SetBody(v *DescribeDomainRegionDataResponseBody) *DescribeDomainRegionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRegionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
