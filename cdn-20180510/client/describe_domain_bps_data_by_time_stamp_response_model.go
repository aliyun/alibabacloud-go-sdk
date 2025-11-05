// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByTimeStampResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainBpsDataByTimeStampResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainBpsDataByTimeStampResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainBpsDataByTimeStampResponseBody) *DescribeDomainBpsDataByTimeStampResponse
	GetBody() *DescribeDomainBpsDataByTimeStampResponseBody
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainBpsDataByTimeStampResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainBpsDataByTimeStampResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainBpsDataByTimeStampResponse) GetBody() *DescribeDomainBpsDataByTimeStampResponseBody {
	return s.Body
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataByTimeStampResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetStatusCode(v int32) *DescribeDomainBpsDataByTimeStampResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetBody(v *DescribeDomainBpsDataByTimeStampResponseBody) *DescribeDomainBpsDataByTimeStampResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
