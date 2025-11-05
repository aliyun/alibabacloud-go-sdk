// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeSrcHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDomainRealTimeSrcHttpCodeDataResponse
	GetBody() *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
}

type DescribeDomainRealTimeSrcHttpCodeDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) GetBody() *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeSrcHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) SetBody(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDomainRealTimeSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
