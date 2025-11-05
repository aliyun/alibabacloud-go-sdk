// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTrafficDataResponseBody) *DescribeDomainTrafficDataResponse
	GetBody() *DescribeDomainTrafficDataResponseBody
}

type DescribeDomainTrafficDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTrafficDataResponse) GetBody() *DescribeDomainTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTrafficDataResponse) SetBody(v *DescribeDomainTrafficDataResponseBody) *DescribeDomainTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTrafficDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
