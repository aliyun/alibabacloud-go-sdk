// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeDetailDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeDetailDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeDetailDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeDetailDataResponseBody) *DescribeDomainRealTimeDetailDataResponse
	GetBody() *DescribeDomainRealTimeDetailDataResponseBody
}

type DescribeDomainRealTimeDetailDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeDetailDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeDetailDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeDetailDataResponse) GetBody() *DescribeDomainRealTimeDetailDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeDetailDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponse) SetBody(v *DescribeDomainRealTimeDetailDataResponseBody) *DescribeDomainRealTimeDetailDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
