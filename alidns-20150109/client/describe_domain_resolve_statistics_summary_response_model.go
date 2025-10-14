// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainResolveStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainResolveStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainResolveStatisticsSummaryResponseBody) *DescribeDomainResolveStatisticsSummaryResponse
	GetBody() *DescribeDomainResolveStatisticsSummaryResponseBody
}

type DescribeDomainResolveStatisticsSummaryResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResolveStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResolveStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) GetBody() *DescribeDomainResolveStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDomainResolveStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeDomainResolveStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) SetBody(v *DescribeDomainResolveStatisticsSummaryResponseBody) *DescribeDomainResolveStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
