// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainRealTimeTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainRealTimeTrafficDataResponseBody) *DescribeLiveDomainRealTimeTrafficDataResponse
	GetBody() *DescribeLiveDomainRealTimeTrafficDataResponseBody
}

type DescribeLiveDomainRealTimeTrafficDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) GetBody() *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetStatusCode(v int32) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetBody(v *DescribeLiveDomainRealTimeTrafficDataResponseBody) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
