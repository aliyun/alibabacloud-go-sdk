// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainRealTimeBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainRealTimeBpsDataResponseBody) *DescribeLiveDomainRealTimeBpsDataResponse
	GetBody() *DescribeLiveDomainRealTimeBpsDataResponseBody
}

type DescribeLiveDomainRealTimeBpsDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) GetBody() *DescribeLiveDomainRealTimeBpsDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetBody(v *DescribeLiveDomainRealTimeBpsDataResponseBody) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
