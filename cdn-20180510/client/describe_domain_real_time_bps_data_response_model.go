// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeBpsDataResponseBody) *DescribeDomainRealTimeBpsDataResponse
	GetBody() *DescribeDomainRealTimeBpsDataResponseBody
}

type DescribeDomainRealTimeBpsDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeBpsDataResponse) GetBody() *DescribeDomainRealTimeBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponse) SetBody(v *DescribeDomainRealTimeBpsDataResponseBody) *DescribeDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
