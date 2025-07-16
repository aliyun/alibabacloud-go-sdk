// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeSrcBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeSrcBpsDataResponseBody) *DescribeDomainRealTimeSrcBpsDataResponse
	GetBody() *DescribeDomainRealTimeSrcBpsDataResponseBody
}

type DescribeDomainRealTimeSrcBpsDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) GetBody() *DescribeDomainRealTimeSrcBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeDomainRealTimeSrcBpsDataResponseBody) *DescribeDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
