// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainUsageDataResponseBody) *DescribeDomainUsageDataResponse
	GetBody() *DescribeDomainUsageDataResponseBody
}

type DescribeDomainUsageDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainUsageDataResponse) GetBody() *DescribeDomainUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUsageDataResponse) SetStatusCode(v int32) *DescribeDomainUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainUsageDataResponse) SetBody(v *DescribeDomainUsageDataResponseBody) *DescribeDomainUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
